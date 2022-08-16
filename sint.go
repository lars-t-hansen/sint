package sint

// Values
//
// type Val union {
//   *Cons,
//   *Symbol,
//   *Procedure,
//   *Null,
//   *bool,			// true or false
//   *Unspecified,
//   *Undefined
//   *big.Int,      // exact integer
//   *big.Float,    // inexact real (rational?)
//   *string		// read-only, which violates the spec, but OK for now
// }

type Val interface{}

type Cons struct {
	car Val
	cdr Val
}

type Symbol struct {
	name  string
	value Val // if the symbol is a global variable, otherwise c.undefined
}

type Procedure struct {
	lambda *Lambda
	env    *LexEnv                  // closed-over lexical environment, nil for global procedures and primitives
	primop func(*Scheme, []Val) Val // nil for non-primitives
}

type Null struct{}
type Unspecified struct{}
type Undefined struct{}

// Code and evaluation.
//
// type Code union {
//   *If,
//   *Begin,
//   *Quote,
//   *Call,
//   *Lambda,
//   *Let,
//   *Letrec,
//   *Lexical,
//   *Setlex,
//   *Global,
//   *Setglobal
// }

type Code interface {
	// Documentation: each expression should carry its source location
}

type Quote struct {
	value Val
}

type If struct {
	test       Code
	consequent Code
	alternate  Code
}

type Begin struct {
	exprs []Code
}

type Call struct {
	exprs []Code
}

type Lambda struct {
	fixed int
	rest  bool
	body  Code
	// Documentation: this should carry the doc string and the source code
	// Documentation: This should carry the names of locals in the rib
}

type Let struct {
	exprs []Code
	body  Code
	// Documentation: This should carry the names of locals in the rib
}

type Letrec struct {
	exprs []Code
	body  Code
	// Documentation: This should carry the names of locals in the rib
}

type Lexical struct {
	levels int
	offset int
	// Documentation: This should carry the name of the variable
}

type Setlex struct {
	levels int
	offset int
	rhs    Code
	// Documentation: This should carry the name of the variable
}

type Global struct {
	name *Symbol
}

type Setglobal struct {
	name *Symbol
	rhs  Code
}

type LexEnv struct {
	slots []Val
	link  *LexEnv
	// Documentation: This should carry the names of locals in the rib
}

type Scheme struct {
	unspecified Val
	undefined   Val
	null        Val
	trueVal     Val
	falseVal    Val
	oblist      map[string]*Symbol
}

func NewScheme() *Scheme {
	t := true
	f := false
	c := &Scheme{
		unspecified: &Unspecified{},
		undefined:   &Undefined{},
		null:        &Null{},
		trueVal:     &t,
		falseVal:    &f,
		oblist:      map[string]*Symbol{},
	}
	c.initPrimitives()
	c.initCompiled()
	return c
}
func (c *Scheme) intern(s string) *Symbol {
	if v, ok := c.oblist[s]; ok {
		return v
	}
	sym := &Symbol{s, c.undefined}
	c.oblist[s] = sym
	return sym
}

func (c *Scheme) eval(expr Code, env *LexEnv) Val {
again:
	switch e := expr.(type) {
	case *Quote:
		return e.value
	case *If:
		if isTruthy(c.eval(e.test, env)) {
			expr = e.consequent
		} else {
			expr = e.alternate
		}
		goto again
	case *Begin:
		if len(e.exprs) == 0 {
			return c.unspecified
		}
		c.evalExprs(e.exprs[:len(e.exprs)-1], env)
		expr = e.exprs[len(e.exprs)-1]
		goto again
	case *Call:
		vals := c.evalExprs(e.exprs, env)
		maybeProc := vals[0]
		args := vals[1:]
		if p, ok := maybeProc.(*Procedure); ok {
			if len(args) < p.lambda.fixed {
				panic("Not enough arguments")
			}
			if len(args) > p.lambda.fixed && !p.lambda.rest {
				panic("Too many arguments")
			}
			if p.lambda.body == nil {
				return p.primop(c, args)
			}
			var newEnv *LexEnv = nil
			if !p.lambda.rest {
				newEnv = &LexEnv{args, env}
			} else {
				newSlots := []Val{}
				for i := 0; i < p.lambda.fixed; i++ {
					newSlots = append(newSlots, args[i])
				}
				var l *Cons
				var last *Cons
				for i := p.lambda.fixed; i < len(args); i++ {
					x := &Cons{args[i], c.null}
					if l == nil {
						l = x
					}
					if last != nil {
						last.cdr = x
					}
					last = x
				}
				if l == nil {
					newSlots = append(newSlots, c.null)
				} else {
					newSlots = append(newSlots, l)
				}
				newEnv = &LexEnv{newSlots, env}
			}
			expr = p.lambda.body
			env = newEnv
			goto again
		} else {
			panic("Not a procedure")
		}
	case *Lambda:
		return &Procedure{e, env, nil}
	case *Let:
		vals := c.evalExprs(e.exprs, env)
		newEnv := &LexEnv{vals, env}
		expr = e.body
		env = newEnv
		goto again
	case *Letrec:
		// TODO: Probably a more efficient way to do this?
		slotvals := []Val{}
		for i := 0; i < len(e.exprs); i++ {
			slotvals = append(slotvals, c.unspecified)
		}
		newEnv := &LexEnv{slotvals, env}
		vals := c.evalExprs(e.exprs, newEnv)
		for i, v := range vals {
			slotvals[i] = v
		}
		expr = e.body
		env = newEnv
		goto again
	case *Lexical:
		rib := env
		for levels := e.levels; levels > 0; levels-- {
			rib = rib.link
		}
		return rib.slots[e.offset]
	case *Setlex:
		rhs := c.eval(e.rhs, env)
		rib := env
		for levels := e.levels; levels > 0; levels-- {
			rib = rib.link
		}
		rib.slots[e.offset] = rhs
		return c.unspecified
	case *Global:
		val := e.name.value
		if val == c.undefined {
			panic("Attempting to read undefined global variable")
		}
		return val
	case *Setglobal:
		rhs := c.eval(e.rhs, env)
		e.name.value = rhs
		return c.unspecified
	default:
		panic("Unknown AST type")
	}
}

func (c *Scheme) evalExprs(es []Code, env *LexEnv) []Val {
	vs := []Val{}
	for e := range es {
		vs = append(vs, c.eval(e, env))
	}
	return vs
}

func isTruthy(v Val) bool {
	if b, ok := v.(bool); ok {
		return b
	}
	return true
}
