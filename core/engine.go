// Evaluation engine core

package core

import (
	"math/big"
	"strconv"
)

type Scheme struct {
	// Symbol table
	oblist map[string]*Symbol

	// Counter for non-interned symbol name generation
	nextGensym int

	// Singleton values
	UnspecifiedVal Val
	UndefinedVal   Val
	NullVal        Val
	TrueVal        Val
	FalseVal       Val
	EofVal         Val

	// Useful values.  TODO: Flesh this out, and use it in the emitter: Most
	// literal values in programs are 0, 1, and 2, and we could have them
	// all predefined here and could just use them rather than cons them
	// up anew every time.  That said, those are *constant* values and
	// are only consed up when the program is deserialized, not at runtime,
	// so they probably are not all that useful frankly.
	Zero *big.Int

	// Well-known symbols.
	AndSym     *Symbol
	BeginSym   *Symbol
	CaseSym    *Symbol
	CondSym    *Symbol
	DefineSym  *Symbol
	DoSym      *Symbol
	ElseSym    *Symbol
	IfSym      *Symbol
	LambdaSym  *Symbol
	LetSym     *Symbol
	LetrecSym  *Symbol
	OrSym      *Symbol
	QuoteSym   *Symbol
	SetSym     *Symbol
	ArrowSym   *Symbol
	DotSym     *Symbol
	NewlineSym *Symbol
	ReturnSym  *Symbol
	TabSym     *Symbol
	SpaceSym   *Symbol

	// Per-thread state that should be handled differently

	// This is interpreted in the context of the number-of-values flag passed back
	// in the evaluator
	MultiVals []Val
}

func NewScheme() *Scheme {
	s := &Scheme{
		UnspecifiedVal: &Unspecified{},
		UndefinedVal:   &Undefined{},
		NullVal:        &Null{},
		TrueVal:        &True{},
		FalseVal:       &False{},
		EofVal:         &EofObject{},
		Zero:           big.NewInt(0),
		oblist:         map[string]*Symbol{},
		nextGensym:     1000,
	}

	s.AndSym = s.Intern("and")
	s.BeginSym = s.Intern("begin")
	s.CaseSym = s.Intern("case")
	s.CondSym = s.Intern("cond")
	s.DefineSym = s.Intern("define")
	s.DoSym = s.Intern("do")
	s.ElseSym = s.Intern("else")
	s.IfSym = s.Intern("if")
	s.LambdaSym = s.Intern("lambda")
	s.LetSym = s.Intern("let")
	s.LetrecSym = s.Intern("letrec")
	s.OrSym = s.Intern("or")
	s.QuoteSym = s.Intern("quote")
	s.SetSym = s.Intern("set!")
	s.ArrowSym = s.Intern("=>")
	s.DotSym = s.Intern(".")
	s.NewlineSym = s.Intern("newline")
	s.ReturnSym = s.Intern("return")
	s.TabSym = s.Intern("tab")
	s.SpaceSym = s.Intern("space")

	return s
}

func (c *Scheme) Intern(s string) *Symbol {
	if v, ok := c.oblist[s]; ok {
		return v
	}
	sym := &Symbol{Name: s, Value: c.UndefinedVal}
	c.oblist[s] = sym
	return sym
}

func (c *Scheme) Gensym(s string) *Symbol {
	name := ".G" + strconv.Itoa(c.nextGensym) + "." + s
	c.nextGensym++
	return &Symbol{Name: name, Value: c.UndefinedVal}
}

func (c *Scheme) EvalToplevel(expr Code) []Val {
	return c.captureValues(c.eval(expr, nil))
}

func (c *Scheme) Invoke(proc Val, args []Val) []Val {
	newCode, newEnv, prim := c.invokeSetup(proc, args)
	var v Val
	var k int
	if prim != nil {
		v, k = prim(c, args)
	} else {
		v, k = c.eval(newCode, newEnv)
	}
	return c.captureValues(v, k)
}

func (c *Scheme) captureValues(v Val, numVal int) []Val {
	vs := []Val{v}
	if numVal > 1 {
		vs = append(vs, c.MultiVals[:numVal-1]...)
	}
	return vs
}

func (c *Scheme) invokeSetup(proc Val, args []Val) (Code, *lexenv, func(*Scheme, []Val) (Val, int)) {
	if p, ok := proc.(*Procedure); ok {
		if len(args) < p.Lam.Fixed {
			panic("Not enough arguments") // TODO msg
		}
		if len(args) > p.Lam.Fixed && !p.Lam.Rest {
			panic("Too many arguments") // TODO msg
		}
		if p.Lam.Body == nil {
			return nil, nil, p.Primop
		}
		var newEnv *lexenv = nil
		// args (really the underlying vals) is freshly allocated,
		// so it's OK to use that storage here.
		if !p.Lam.Rest {
			newEnv = &lexenv{slots: args, link: p.Env}
		} else {
			// TODO: I think we can do better than this.  Since the storage
			// is fresh, we can store the rest argument in the slot after the
			// slice, if it exists, in which case we avoid copying the
			// array in the append() below.  If there is no extra slot then there's
			// at least a chance that the append() will use capacity that is there.
			newSlots := args[:p.Lam.Fixed]
			var l *Cons
			var last *Cons
			for i := p.Lam.Fixed; i < len(args); i++ {
				x := &Cons{Car: args[i], Cdr: c.NullVal}
				if l == nil {
					l = x
				}
				if last != nil {
					last.Cdr = x
				}
				last = x
			}
			if l == nil {
				newSlots = append(newSlots, c.NullVal)
			} else {
				newSlots = append(newSlots, l)
			}
			newEnv = &lexenv{slots: newSlots, link: p.Env}
		}
		return p.Lam.Body, newEnv, nil
	} else {
		panic("Invoke: Not a procedure" /*+ e.Exprs[0].String() + "\n" + proc.String()*/)
	}
}

type lexenv struct {
	slots []Val
	link  *lexenv
	// TODO: Documentation: This should carry the names of locals in the rib
}

func (c *Scheme) eval(expr Code, env *lexenv) (Val, int) {
again:
	switch instr := expr.(type) {
	case *Quote:
		return instr.Value, 1
	case *If:
		if v, _ := c.eval(instr.Test, env); v != c.FalseVal {
			expr = instr.Consequent
		} else {
			expr = instr.Alternate
		}
		goto again
	case *Begin:
		if len(instr.Exprs) == 0 {
			return c.UnspecifiedVal, 1
		}
		c.evalExprs(instr.Exprs[:len(instr.Exprs)-1], env)
		expr = instr.Exprs[len(instr.Exprs)-1]
		goto again
	case *Call:
		vals := c.evalExprs(instr.Exprs, env)
		maybeProc := vals[0]
		args := vals[1:]
		newCode, newEnv, prim := c.invokeSetup(maybeProc, args)
		if prim != nil {
			return prim(c, args)
		}
		expr = newCode
		env = newEnv
		goto again
	//			panic("Invoke: Not a procedure: " + e.Exprs[0].String() + "\n" + maybeProc.String())
	case *Apply:
		proc, _ := c.eval(instr.Proc, env)
		argList, _ := c.eval(instr.Args, env)
		args := []Val{}
		for {
			if argList == c.NullVal {
				break
			}
			a, ok := argList.(*Cons)
			if !ok {
				panic("sint:apply: Not a list") // TODO: msg
			}
			args = append(args, a.Car)
			argList = a.Cdr
		}
		newCode, newEnv, prim := c.invokeSetup(proc, args)
		if prim != nil {
			return prim(c, args)
		}
		expr = newCode
		env = newEnv
		goto again
	case *Lambda:
		return &Procedure{Lam: instr, Env: env, Primop: nil}, 1
	case *Let:
		vals := c.evalExprs(instr.Exprs, env)
		newEnv := &lexenv{slots: vals, link: env}
		expr = instr.Body
		env = newEnv
		goto again
	case *Letrec:
		// TODO: Probably there's a more efficient way to do this?  Note we need
		// fresh storage, so at a minimum we need to copy out of a master slice of
		// undefined values.
		slotvals := []Val{}
		for i := 0; i < len(instr.Exprs); i++ {
			slotvals = append(slotvals, c.UnspecifiedVal)
		}
		newEnv := &lexenv{slots: slotvals, link: env}
		vals := c.evalExprs(instr.Exprs, newEnv)
		for i, v := range vals {
			slotvals[i] = v
		}
		expr = instr.Body
		env = newEnv
		goto again
	case *Lexical:
		rib := env
		for levels := instr.Levels; levels > 0; levels-- {
			rib = rib.link
		}
		return rib.slots[instr.Offset], 1
	case *Setlex:
		rhs, _ := c.eval(instr.Rhs, env)
		rib := env
		for levels := instr.Levels; levels > 0; levels-- {
			rib = rib.link
		}
		rib.slots[instr.Offset] = rhs
		return c.UnspecifiedVal, 1
	case *Global:
		val := instr.Name.Value
		if val == c.UndefinedVal {
			panic("Undefined global variable '" + instr.Name.Name + "'")
		}
		return val, 1
	case *Setglobal:
		rhs, _ := c.eval(instr.Rhs, env)
		instr.Name.Value = rhs
		return c.UnspecifiedVal, 1
	default:
		panic("Bad expression: " + expr.String())
	}
}

func (c *Scheme) evalExprs(es []Code, env *lexenv) []Val {
	vs := []Val{}
	for _, e := range es {
		r, _ := c.eval(e, env)
		vs = append(vs, r)
	}
	return vs
}
