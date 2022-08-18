// Interpreter core - definitions of values and program nodes; evaluation.

package sint

import (
	"fmt"
	"math/big"
)

// Values
//
// type Val union {
//   *Cons,
//   *Symbol,
//   *Procedure,
//   *Null,			// Singleton
//   *True			// Singleton
//   *False			// Singleton
//   *Unspecified,	// Singleton
//   *Undefined		// Singleton
//   *big.Int,      // Exact integer
//   *big.Float,    // Inexact real (rational?)
//   *string		// String, immutable for now
// }

type Val interface {
	fmt.Stringer
}

type Cons struct {
	car Val
	cdr Val
}

func (c *Cons) String() string {
	return "cons"
}

type Symbol struct {
	name  string
	value Val // if the symbol is a global variable, otherwise c.undefined
}

func (c *Symbol) String() string {
	return "symbol"
}

type Procedure struct {
	lambda *Lambda
	env    *LexEnv                  // closed-over lexical environment, nil for global procedures and primitives
	primop func(*Scheme, []Val) Val // nil for non-primitives
}

func (c *Procedure) String() string {
	return "procedure"
}

type Null struct{}

func (c *Null) String() string {
	return "null"
}

type True struct{}

func (c *True) String() string {
	return "true"
}

type False struct{}

func (c *False) String() string {
	return "false"
}

type Unspecified struct{}

func (c *Unspecified) String() string {
	return "unspecified"
}

type Undefined struct{}

func (c *Undefined) String() string {
	return "undefined"
}

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
//
// TODO: Probably want a let*, at least, to cut down on the number of
// ribs being allocated and the number of eval steps.

type Code interface {
	// Documentation: each expression should carry its source location
	fmt.Stringer
}

type Quote struct {
	value Val
}

func (c *Quote) String() string {
	return "quote"
}

type If struct {
	test       Code
	consequent Code
	alternate  Code
}

func (c *If) String() string {
	return "if"
}

type Begin struct {
	exprs []Code
}

func (c *Begin) String() string {
	return "begin"
}

type Call struct {
	exprs []Code
}

func (c *Call) String() string {
	return "call"
}

type Lambda struct {
	fixed int
	rest  bool
	body  Code
	// Documentation: this should carry the doc string and the source code
	// Documentation: This should carry the names of locals in the rib
}

func (c *Lambda) String() string {
	return "lambda"
}

type Let struct {
	exprs []Code
	body  Code
	// Documentation: This should carry the names of locals in the rib
}

func (c *Let) String() string {
	return "let"
}

type Letrec struct {
	exprs []Code
	body  Code
	// Documentation: This should carry the names of locals in the rib
}

func (c *Letrec) String() string {
	return "letrec"
}

type Lexical struct {
	levels int
	offset int
	// Documentation: This should carry the name of the variable
}

func (c *Lexical) String() string {
	return "lexical"
}

type Setlex struct {
	levels int
	offset int
	rhs    Code
	// Documentation: This should carry the name of the variable
}

func (c *Setlex) String() string {
	return "setlex"
}

type Global struct {
	name *Symbol
}

func (c *Global) String() string {
	return "global"
}

type Setglobal struct {
	name *Symbol
	rhs  Code
}

func (c *Setglobal) String() string {
	return "setglobal"
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
	zero        *big.Int
	fzero       *big.Float
	oblist      map[string]*Symbol
}

func NewScheme() *Scheme {
	c := &Scheme{
		unspecified: &Unspecified{},
		undefined:   &Undefined{},
		null:        &Null{},
		trueVal:     &True{},
		falseVal:    &False{},
		zero:        big.NewInt(0),
		fzero:       big.NewFloat(0),
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
		if c.eval(e.test, env) != c.falseVal {
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
				panic("Not enough arguments") // FIXME msg
			}
			if len(args) > p.lambda.fixed && !p.lambda.rest {
				panic("Too many arguments") // FIXME msg
			}
			if p.lambda.body == nil {
				return p.primop(c, args)
			}
			var newEnv *LexEnv = nil
			// args (really the underlying vals) is freshly allocated,
			// so it's OK to use that storage here.
			if !p.lambda.rest {
				newEnv = &LexEnv{args, env}
			} else {
				// TODO: I think we can do better than this.  Since the storage
				// is fresh, we can store the rest argument in the slot after the
				// slice, if it exists, in which case we avoid copying the
				// array in the append() below.  If there is no extra slot then there's
				// at least a chance that the append() will use capacity that is there.
				newSlots := args[:p.lambda.fixed]
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
			panic("Not a procedure") // FIXME msg
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
		// TODO: Probably there's a more efficient way to do this?  Note we need
		// fresh storage, so at a minimum we need to copy out of a master slice of
		// undefined values.
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
			panic("Undefined global variable '" + e.name.name + "'")
		}
		return val
	case *Setglobal:
		rhs := c.eval(e.rhs, env)
		e.name.value = rhs
		return c.unspecified
	default:
		panic("Bad expression: " + expr.String())
	}
}

func (c *Scheme) evalExprs(es []Code, env *LexEnv) []Val {
	vs := []Val{}
	for _, e := range es {
		vs = append(vs, c.eval(e, env))
	}
	return vs
}
