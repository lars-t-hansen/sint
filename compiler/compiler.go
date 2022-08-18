// Simple compiler from standard sexpr form.
//
// Built-in syntax (macros) have reserved names, which is not very "Scheme"
// but is just fine in practice.  There are no user-defined macros here.

// This takes an expression encoded as a list and both macro-expands and compiles.
// But this is not the only way to
// encode it: instead, the code reader could return a proper syntax-checked AST.
// In full Scheme that's hard because lexical bindings can shadow macros and there
// can also be local macros; compilation and macro expansion are intertwined.
// For a simple system it's more viable.

package compiler

import (
	"math/big"
	. "sint"
)

type Compiler struct {
	s *Scheme

	// REMEMBER to update isKeyword when this list is extended
	andSym    *Symbol
	beginSym  *Symbol
	caseSym   *Symbol
	condSym   *Symbol
	defineSym *Symbol
	elseSym   *Symbol
	ifSym     *Symbol
	lambdaSym *Symbol
	letSym    *Symbol
	letrecSym *Symbol
	orSym     *Symbol
	quoteSym  *Symbol
	setSym    *Symbol
}

type CEnv struct {
	link  *CEnv
	names []*Symbol
}

func lookup(env *CEnv, s *Symbol) (int, int, bool) {
	levels := 0
	for env != nil {
		for offset := 0; offset < len(env.names); offset++ {
			if env.names[offset] == s {
				return levels, offset, true
			}
		}
		env = env.link
	}
	return 0, 0, false
}

func (c *Compiler) compileToplevel(v Val) Code {
	length, exprIsList := c.checkProperList(v)
	if exprIsList && length >= 3 && car(v) == c.defineSym {
		return c.compileToplevelDefinition(v)
	}
	return c.compileExpr(v, nil)
}

func (c *Compiler) compileToplevelDefinition(v Val) Code {
	nameOrSignature := cadr(v)
	if s, ok := nameOrSignature.(*Symbol); ok {
		return &Setglobal{
			Name: s,
			Rhs:  c.compileExpr(caddr(v), nil),
		}
	}
	if len, improper, ok := c.checkPossiblyImproperList(cdadr(v)); ok && len > 0 || improper {
		if s, ok := car(cadr(v)).(*Symbol); ok {
			// check that remaining elements are distinct symbols
			// create a lambda expression and compile it
			var body Code
			lam := &Lambda{Fixed: len, Rest: improper, Body: body}
			return &Setglobal{
				Name: s,
				Rhs:  c.compileExpr(lam, nil),
			}
		}
	}
	panic("Invalid top-level definition")
}

func (c *Compiler) compileExpr(v Val, env *CEnv) Code {
	switch e := v.(type) {
	case *big.Int:
		return &Quote{Value: e}
	case *big.Float:
		return &Quote{Value: e}
	case *True:
		return &Quote{Value: e}
	case *False:
		return &Quote{Value: e}
	case *Unspecified:
		return &Quote{Value: e}
	case *Undefined:
		return &Quote{Value: e}
		/*case string:
		return &Quote{e}*/
	case *Symbol:
		return c.compileRef(e, env)
	case *Cons:
		len, exprIsList := c.checkProperList(e)
		if !exprIsList {
			panic("Improper list used as expression")
		}
		if len == 0 {
			panic("Unquoted empty list used as expression")
		}
		if kwd, ok := e.Car.(*Symbol); ok {
			if kwd == c.ifSym {
				return c.compileIf(e, len, env)
			}
			if kwd == c.setSym {
				return c.compileSet(e, len, env)
			}
			// More macros here!
		}
		return c.compileCall(e, len, env)
	default:
		panic("Bad expression")
	}
}

func (c *Compiler) compileCall(l Val, _ int, env *CEnv) Code {
	var exprs []Code
	for l != c.s.NullVal {
		exprs = append(exprs, c.compileExpr(car(l), env))
		l = cdr(l)
	}
	return &Call{Exprs: exprs}
}

func (c *Compiler) compileIf(l *Cons, len int, env *CEnv) Code {
	if len != 3 && len != 4 {
		panic("if: Illegal form")
	}
	test := cadr(l)
	consequent := caddr(l)
	var alternate Val = c.s.UnspecifiedVal
	if len == 4 {
		alternate = cadddr(l)
	}
	return &If{
		Test:       c.compileExpr(test, env),
		Consequent: c.compileExpr(consequent, env),
		Alternate:  c.compileExpr(alternate, env),
	}
}

func (c *Compiler) compileRef(s *Symbol, env *CEnv) Code {
	if levels, offset, ok := lookup(env, s); ok {
		return &Lexical{Levels: levels, Offset: offset}
	}
	if c.isKeyword(s) {
		panic("Keyword used as variable reference: " + s.Name)
	}
	return &Global{Name: s}
}

func (c *Compiler) compileSet(l *Cons, len int, env *CEnv) Code {
	if len != 3 {
		panic("set!: Illegal form")
	}
	place := cadr(l)
	expr := caddr(l)
	placeName, nameIsSymbol := place.(*Symbol)
	if !nameIsSymbol {
		panic("set!: Illegal variable name: " + place.String())
	}
	rhs := c.compileExpr(expr, env)
	if levels, offset, ok := lookup(env, placeName); ok {
		return &Setlex{Levels: levels, Offset: offset, Rhs: rhs}
	}
	if c.isKeyword(placeName) {
		panic("Keyword used as variable name: " + placeName.Name)
	}
	return &Setglobal{Name: placeName, Rhs: rhs}
}

func car(v Val) Val {
	if c, ok := v.(*Cons); ok {
		return c.Car
	}
	panic("CAR: failed" + v.String())
}

func cdr(v Val) Val {
	if c, ok := v.(*Cons); ok {
		return c.Cdr
	}
	panic("CDR: failed" + v.String())
}

func cadr(v Val) Val {
	return car(cdr(v))
}

func caddr(v Val) Val {
	return car(cdr(cdr(v)))
}

func cadddr(v Val) Val {
	return car(cdr(cdr(cdr(v))))
}

func (c *Compiler) checkProperList(v Val) (int, bool) {
	// Check that v is a proper list, and return its length if it is
	// FIXME: This needs to deal with circularity
	len := 0
	for {
		cell, ok := v.(*Cons)
		if !ok {
			break
		}
		v = cell.Cdr
		len++
	}
	if v == c.s.NullVal {
		return len, true
	}
	return -1, false
}

func (c *Compiler) isKeyword(s *Symbol) bool {
	// FIXME: Stupid to do a linear search here.  Alternative is a map, or that
	// environment lookup will find a denotation that is 'keyword', it all amounts
	// to the same thing
	if s == c.andSym || s == c.beginSym || s == c.caseSym || s == c.condSym || s == c.defineSym ||
		s == c.elseSym || s == c.ifSym || s == c.lambdaSym || s == c.letSym || s == c.letrecSym ||
		s == c.orSym || s == c.quoteSym || s == c.setSym {
		return true
	}
	return false
}
