// Evaluation engine core - definitions of values and program nodes; evaluation.

package core

import (
	"fmt"
	"math/big"
	"strconv"
)

// Values.
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
	Car Val
	Cdr Val
}

func (c *Cons) String() string {
	return "cons"
}

type Symbol struct {
	Name  string
	Value Val // if the symbol is a global variable, otherwise c.undefined
}

func (c *Symbol) String() string {
	return "symbol"
}

type Procedure struct {
	Lam    *Lambda
	Env    *lexenv                  // closed-over lexical environment, nil for global procedures and primitives
	Primop func(*Scheme, []Val) Val // nil for non-primitives
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
	Value Val
}

func (c *Quote) String() string {
	return "quote"
}

type If struct {
	Test       Code
	Consequent Code
	Alternate  Code
}

func (c *If) String() string {
	return "if"
}

type Begin struct {
	Exprs []Code
}

func (c *Begin) String() string {
	return "begin"
}

type Call struct {
	Exprs []Code
}

func (c *Call) String() string {
	return "call"
}

type Lambda struct {
	Fixed int
	Rest  bool
	Body  Code
	// Documentation: this should carry the doc string and the source code
	// Documentation: This should carry the names of locals in the rib
}

func (c *Lambda) String() string {
	return "lambda"
}

type Let struct {
	Exprs []Code
	Body  Code
	// Documentation: This should carry the names of locals in the rib
}

func (c *Let) String() string {
	return "let"
}

type Letrec struct {
	Exprs []Code
	Body  Code
	// Documentation: This should carry the names of locals in the rib
}

func (c *Letrec) String() string {
	return "letrec"
}

type Lexical struct {
	Levels int
	Offset int
	// Documentation: This should carry the name of the variable
}

func (c *Lexical) String() string {
	return "lexical"
}

type Setlex struct {
	Levels int
	Offset int
	Rhs    Code
	// Documentation: This should carry the name of the variable
}

func (c *Setlex) String() string {
	return "setlex"
}

type Global struct {
	Name *Symbol
}

func (c *Global) String() string {
	return "global"
}

type Setglobal struct {
	Name *Symbol
	Rhs  Code
}

func (c *Setglobal) String() string {
	return "setglobal"
}

// Runtimes.

type Scheme struct {
	UnspecifiedVal Val
	UndefinedVal   Val
	NullVal        Val
	TrueVal        Val
	FalseVal       Val
	Zero           *big.Int
	FZero          *big.Float
	oblist         map[string]*Symbol
	nextGensym     int
}

func NewScheme() *Scheme {
	return &Scheme{
		UnspecifiedVal: &Unspecified{},
		UndefinedVal:   &Undefined{},
		NullVal:        &Null{},
		TrueVal:        &True{},
		FalseVal:       &False{},
		Zero:           big.NewInt(0),
		FZero:          big.NewFloat(0),
		oblist:         map[string]*Symbol{},
		nextGensym:     1000,
	}
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
	// Documentation: This should carry the names of locals in the rib
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
				newEnv = &lexenv{args, env}
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
					x := &Cons{args[i], c.NullVal}
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
				newEnv = &lexenv{newSlots, env}
			}
			expr = p.Lam.Body
			env = newEnv
			goto again
		} else {
			panic("Not a procedure") // TODO msg
		}
	case *Lambda:
		return &Procedure{e, env, nil}
	case *Let:
		vals := c.evalExprs(e.Exprs, env)
		newEnv := &lexenv{vals, env}
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
		newEnv := &lexenv{slotvals, env}
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
