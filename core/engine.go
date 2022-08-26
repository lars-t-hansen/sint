// Evaluation engine core - definitions of values and program nodes; evaluation.

package core

import (
	"math/big"
	"strconv"
)

// Runtimes and evaluation.
//
// TODO: The Scheme instance should cache a compiler instance, if it has no
// post-init mutable state.

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

	// Useful values
	Zero  *big.Int
	FZero *big.Float

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
		FZero:          big.NewFloat(0),
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

func (c *Scheme) EvalToplevel(expr Code) Val {
	return c.eval(expr, nil)
}

type lexenv struct {
	slots []Val
	link  *lexenv
	// TODO: Documentation: This should carry the names of locals in the rib
}

func (c *Scheme) eval(expr Code, env *lexenv) Val {
again:
	switch e := expr.(type) {
	case *Quote:
		return e.Value
	case *If:
		if c.eval(e.Test, env) != c.FalseVal {
			expr = e.Consequent
		} else {
			expr = e.Alternate
		}
		goto again
	case *Begin:
		if len(e.Exprs) == 0 {
			return c.UnspecifiedVal
		}
		c.evalExprs(e.Exprs[:len(e.Exprs)-1], env)
		expr = e.Exprs[len(e.Exprs)-1]
		goto again
	case *Call:
		// TODO: apply must be supported directly here
		vals := c.evalExprs(e.Exprs, env)
		maybeProc := vals[0]
		args := vals[1:]
		if p, ok := maybeProc.(*Procedure); ok {
			if len(args) < p.Lam.Fixed {
				panic("Not enough arguments") // TODO msg
			}
			if len(args) > p.Lam.Fixed && !p.Lam.Rest {
				panic("Too many arguments") // TODO msg
			}
			if p.Lam.Body == nil {
				return p.Primop(c, args)
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
			expr = p.Lam.Body
			env = newEnv
			goto again
		} else {
			panic("Invoke: Not a procedure: " + e.Exprs[0].String() + "\n" + maybeProc.String())
		}
	case *Apply:
		// FIXME
		panic("Apply not implemented yet")
	case *Lambda:
		return &Procedure{e, env, nil}
	case *Let:
		vals := c.evalExprs(e.Exprs, env)
		newEnv := &lexenv{slots: vals, link: env}
		expr = e.Body
		env = newEnv
		goto again
	case *Letrec:
		// TODO: Probably there's a more efficient way to do this?  Note we need
		// fresh storage, so at a minimum we need to copy out of a master slice of
		// undefined values.
		slotvals := []Val{}
		for i := 0; i < len(e.Exprs); i++ {
			slotvals = append(slotvals, c.UnspecifiedVal)
		}
		newEnv := &lexenv{slots: slotvals, link: env}
		vals := c.evalExprs(e.Exprs, newEnv)
		for i, v := range vals {
			slotvals[i] = v
		}
		expr = e.Body
		env = newEnv
		goto again
	case *Lexical:
		rib := env
		for levels := e.Levels; levels > 0; levels-- {
			rib = rib.link
		}
		return rib.slots[e.Offset]
	case *Setlex:
		rhs := c.eval(e.Rhs, env)
		rib := env
		for levels := e.Levels; levels > 0; levels-- {
			rib = rib.link
		}
		rib.slots[e.Offset] = rhs
		return c.UnspecifiedVal
	case *Global:
		val := e.Name.Value
		if val == c.UndefinedVal {
			panic("Undefined global variable '" + e.Name.Name + "'")
		}
		return val
	case *Setglobal:
		rhs := c.eval(e.Rhs, env)
		e.Name.Value = rhs
		return c.UnspecifiedVal
	default:
		panic("Bad expression: " + expr.String())
	}
}

func (c *Scheme) evalExprs(es []Code, env *lexenv) []Val {
	vs := []Val{}
	for _, e := range es {
		vs = append(vs, c.eval(e, env))
	}
	return vs
}
