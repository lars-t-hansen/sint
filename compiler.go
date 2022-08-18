// Simple code reader and compiler, enough for bootstrapping

package sint

import "math/big"

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

func (env *CEnv) lookup(s *Symbol) (int, int, bool) {

}

// TODO: compileToplevel

// This takes an expression encoded as a list and both macro-expands and compiles.
// But this is not the only way to
// encode it: instead, the code reader could return a proper syntax-checked AST.
// In full Scheme that's hard because lexical bindings can shadow macros and there
// can also be local macros; compilation and macro expansion are intertwined.
// For a simple system it's more viable.

func (c *Compiler) compileExpr(v Val, env *CEnv) Code {
	switch e := v.(type) {
	case *big.Int:
		return &Quote{value: e}
	case *big.Float:
		return &Quote{value: e}
	case *True:
		return &Quote{value: e}
	case *False:
		return &Quote{value: e}
	case *Unspecified:
		return &Quote{value: e}
	case *Undefined:
		return &Quote{value: e}
		/*case string:
		return &Quote{e}*/
	case *Symbol:
		return c.compileRef(e, env)
	case *Cons:
		len, exprIsList := c.checkList(e)
		if !exprIsList {
			panic("Improper list used as expression")
		}
		if len == 0 {
			panic("Unquoted empty list used as expression")
		}
		if kwd, ok := e.car.(*Symbol); ok {
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
	for l != c.s.null {
		exprs = append(exprs, c.compileExpr(car(l), env))
		l = cdr(l)
	}
	return &Call{exprs: exprs}
}

func (c *Compiler) compileIf(l *Cons, len int, env *CEnv) Code {
	if len != 3 && len != 4 {
		panic("if: Illegal form")
	}
	test := cadr(l)
	consequent := caddr(l)
	var alternate Val = c.s.unspecified
	if len == 4 {
		alternate = cadddr(l)
	}
	return &If{
		test:       c.compileExpr(test, env),
		consequent: c.compileExpr(consequent, env),
		alternate:  c.compileExpr(alternate, env),
	}
}

func (c *Compiler) compileRef(s *Symbol, env *CEnv) Code {
	if c.isKeyword(s) {
		panic("Keyword used as variable reference: " + s.name)
	}
	if levels, offset, ok := env.lookup(s); ok {
		return &Lexical{levels: levels, offset: offset}
	}
	return &Global{name: s}
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
	if c.isKeyword(placeName) {
		panic("Keyword used as variable name: " + placeName.name)
	}
	rhs := c.compileExpr(expr, env)
	if levels, offset, ok := env.lookup(placeName); ok {
		return &Setlex{levels, offset, rhs}
	}
	return &Setglobal{name: placeName, rhs: rhs}
}

func car(v Val) Val {
	if c, ok := v.(*Cons); ok {
		return c.car
	}
	panic("CAR: failed" + v.String())
}

func cdr(v Val) Val {
	if c, ok := v.(*Cons); ok {
		return c.cdr
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

func (c *Compiler) checkList(v *Cons) (int, bool) {

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
