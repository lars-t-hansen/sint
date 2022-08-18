// Simple compiler from standard sexpr form.
//
// Built-in syntax (macros) have reserved names, which is not very "Scheme"
// but is just fine in practice.  There are no user-defined macros here.
//
// A more sophisticated compiler would take an AST form that also carries source
// location information, or at least a map of such information on the side.

package compiler

import (
	"math/big"
	. "sint/core"
)

type Compiler struct {
	s         *Scheme
	keywords  map[*Symbol]bool
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

func NewCompiler(s *Scheme) *Compiler {
	c := &Compiler{
		s:         s,
		keywords:  make(map[*Symbol]bool),
		andSym:    s.Intern("and"),
		beginSym:  s.Intern("begin"),
		caseSym:   s.Intern("case"),
		condSym:   s.Intern("cond"),
		defineSym: s.Intern("define"),
		elseSym:   s.Intern("else"),
		ifSym:     s.Intern("if"),
		lambdaSym: s.Intern("lambda"),
		letSym:    s.Intern("let"),
		letrecSym: s.Intern("letrec"),
		orSym:     s.Intern("or"),
		quoteSym:  s.Intern("quote"),
		setSym:    s.Intern("set!"),
	}
	c.keywords[c.andSym] = true
	c.keywords[c.beginSym] = true
	c.keywords[c.caseSym] = true
	c.keywords[c.condSym] = true
	c.keywords[c.defineSym] = true
	c.keywords[c.elseSym] = true
	c.keywords[c.ifSym] = true
	c.keywords[c.lambdaSym] = true
	c.keywords[c.letSym] = true
	c.keywords[c.letrecSym] = true
	c.keywords[c.orSym] = true
	c.keywords[c.quoteSym] = true
	c.keywords[c.setSym] = true
	return c
}

func (c *Compiler) CompileToplevel(v Val) Code {
	length, exprIsList := c.checkProperList(v)
	if exprIsList && length >= 3 && car(v) == c.defineSym {
		return c.compileToplevelDefinition(v)
	}
	return c.compileExpr(v, nil)
}

type cenv struct {
	link  *cenv
	names []*Symbol
}

func lookup(env *cenv, s *Symbol) (int, int, bool) {
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

func (c *Compiler) compileToplevelDefinition(v Val) Code {
	nameOrSignature := cadr(v)
	// (define x v)
	if globName, ok := nameOrSignature.(*Symbol); ok {
		return &Setglobal{
			Name: globName,
			Rhs:  c.compileExpr(caddr(v), nil),
		}
	}
	// (define (f arg ... . arg) body ...)
	if fixed, rest, globName, formals, ok := c.checkDefinitionSignature(nameOrSignature); ok {
		bodyList := cdr(cdr(v))
		var body Code
		if cdr(bodyList) != c.s.NullVal {
			body = &Cons{Car: c.beginSym, Cdr: bodyList}
		} else {
			body = car(bodyList)
		}
		lam := &Lambda{
			Fixed: fixed,
			Rest:  rest,
			Body:  c.compileExpr(body, &cenv{link: nil, names: formals})}
		return &Setglobal{
			Name: globName,
			Rhs:  lam,
		}
	}
	panic("Invalid top-level definition")
}

func (c *Compiler) compileExpr(v Val, env *cenv) Code {
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
			// FIXME: More syntax here!

			// Fall through to generic "call" case
		}
		return c.compileCall(e, len, env)
	default:
		panic("Bad expression")
	}
}

func (c *Compiler) compileCall(l Val, _ int, env *cenv) Code {
	// (expr expr ...)
	var exprs []Code
	for l != c.s.NullVal {
		exprs = append(exprs, c.compileExpr(car(l), env))
		l = cdr(l)
	}
	return &Call{Exprs: exprs}
}

func (c *Compiler) compileIf(l *Cons, len int, env *cenv) Code {
	// (if expr expr)
	// (if expr expr expr)
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

func (c *Compiler) compileRef(s *Symbol, env *cenv) Code {
	// ident
	if levels, offset, ok := lookup(env, s); ok {
		return &Lexical{Levels: levels, Offset: offset}
	}
	if c.isKeyword(s) {
		panic("Keyword used as variable reference: " + s.Name)
	}
	return &Global{Name: s}
}

func (c *Compiler) compileSet(l *Cons, len int, env *cenv) Code {
	// (set! ident expr)
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

func (c *Compiler) checkDefinitionSignature(sig Val) (fixed int, rest bool, globName *Symbol, formals []*Symbol, ok bool) {
	// FIXME: This needs to deal with circularity
	var names []*Symbol
	for {
		cell, cellIsCons := sig.(*Cons)
		if !cellIsCons {
			break
		}
		argName, argIsSymbol := cell.Car.(*Symbol)
		if !argIsSymbol {
			ok = false
			return
		}
		names = append(names, argName)
		sig = cell.Cdr
		fixed++
	}
	if sig != c.s.NullVal {
		rest = true
		argName, argIsSymbol := sig.(*Symbol)
		if !argIsSymbol {
			ok = false
			return
		}
		names = append(formals, argName)
	}
	globName = names[0]
	formals = names[1:]
	for i := 0; i < len(formals); i++ {
		for j := i + 1; j < len(formals); j++ {
			if formals[i] == formals[j] {
				ok = false
				return
			}
		}
	}
	ok = true
	return
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
	return c.keywords[s]
}
