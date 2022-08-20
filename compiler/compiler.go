// Simple compiler from standard sexpr form to internal Code form.
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
	doSym     *Symbol
	elseSym   *Symbol
	ifSym     *Symbol
	lambdaSym *Symbol
	letSym    *Symbol
	letrecSym *Symbol
	orSym     *Symbol
	quoteSym  *Symbol
	setSym    *Symbol
	arrowSym  *Symbol
}

// A Compiler currently has no interesting mutable state, it can be reused for
// multiple compilations.

func NewCompiler(s *Scheme) *Compiler {
	c := &Compiler{
		s:         s,
		keywords:  make(map[*Symbol]bool),
		andSym:    s.Intern("and"),
		beginSym:  s.Intern("begin"),
		caseSym:   s.Intern("case"),
		condSym:   s.Intern("cond"),
		defineSym: s.Intern("define"),
		doSym:     s.Intern("do"),
		elseSym:   s.Intern("else"),
		ifSym:     s.Intern("if"),
		lambdaSym: s.Intern("lambda"),
		letSym:    s.Intern("let"),
		letrecSym: s.Intern("letrec"),
		orSym:     s.Intern("or"),
		quoteSym:  s.Intern("quote"),
		setSym:    s.Intern("set!"),
		arrowSym:  s.Intern("=>"),
	}
	c.keywords[c.andSym] = true
	c.keywords[c.beginSym] = true
	c.keywords[c.caseSym] = true
	c.keywords[c.condSym] = true
	c.keywords[c.defineSym] = true
	c.keywords[c.doSym] = true
	c.keywords[c.elseSym] = true
	c.keywords[c.ifSym] = true
	c.keywords[c.lambdaSym] = true
	c.keywords[c.letSym] = true
	c.keywords[c.letrecSym] = true
	c.keywords[c.orSym] = true
	c.keywords[c.quoteSym] = true
	c.keywords[c.setSym] = true
	// arrowSym is not reserved, possibly elseSym should not be either
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
		levels++
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
		bodyList := cddr(v)
		var body Code
		if cdr(bodyList) != c.s.NullVal {
			body = cons(c.beginSym, bodyList)
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
	//os.Stdout.WriteString(v.String() + "\n")
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
		llen, exprIsList := c.checkProperList(e)
		if !exprIsList {
			panic("Improper list used as expression")
		}
		if llen == 0 {
			panic("Unquoted empty list used as expression")
		}
		if kwd, ok := e.Car.(*Symbol); ok {
			if kwd == c.andSym {
				return c.compileAnd(e, llen, env)
			}
			if kwd == c.beginSym {
				return c.compileBegin(e, llen, env)
			}
			// TODO: case
			// TODO: cond
			// TODO: do
			if kwd == c.lambdaSym {
				return c.compileLambda(e, llen, env)
			}
			if kwd == c.letSym {
				return c.compileLet(e, llen, env)
			}
			if kwd == c.letrecSym {
				return c.compileLetrec(e, llen, env)
			}
			if kwd == c.orSym {
				return c.compileOr(e, llen, env)
			}
			if kwd == c.quoteSym {
				return c.compileQuote(e, llen, env)
			}
			if kwd == c.ifSym {
				return c.compileIf(e, llen, env)
			}
			if kwd == c.setSym {
				return c.compileSet(e, llen, env)
			}
			// Fall through to generic "call" case
		}
		return c.compileCall(e, llen, env)
	default:
		panic("Bad expression")
	}
}

func (c *Compiler) compileExprList(l Val, env *cenv) []Code {
	var exprs []Code
	for l != c.s.NullVal {
		exprs = append(exprs, c.compileExpr(car(l), env))
		l = cdr(l)
	}
	return exprs
}

func (c *Compiler) compileExprSlice(es []Val, env *cenv) []Code {
	var exprs []Code
	for _, e := range es {
		exprs = append(exprs, c.compileExpr(e, env))
	}
	return exprs
}

func (c *Compiler) compileAnd(l Val, llen int, env *cenv) Code {
	// (and expr ...)
	if llen == 1 {
		return &Quote{Value: c.s.TrueVal}
	}
	if llen == 2 {
		return c.compileExpr(cadr(l), env)
	}
	return &If{
		Test:       cadr(l),
		Consequent: c.compileExpr(cons(c.andSym, cddr(l)), env),
		Alternate:  &Quote{Value: c.s.FalseVal},
	}
}

func (c *Compiler) compileBegin(l Val, llen int, env *cenv) Code {
	// (begin expr ...)
	if llen == 1 {
		return &Quote{Value: c.s.UnspecifiedVal}
	}
	// Optimization: Single-expression BEGIN becomes just the expression
	if llen == 2 {
		return cadr(l)
	}
	return &Begin{Exprs: c.compileExprList(cdr(l), env)}
}

func (c *Compiler) compileCall(l Val, _ int, env *cenv) Code {
	// (expr expr ...)
	return &Call{Exprs: c.compileExprList(l, env)}
}

func (c *Compiler) compileIf(l *Cons, llen int, env *cenv) Code {
	// (if expr expr)
	// (if expr expr expr)
	if llen != 3 && llen != 4 {
		panic("if: Illegal form")
	}
	test := cadr(l)
	consequent := caddr(l)
	var alternate Val = c.s.UnspecifiedVal
	if llen == 4 {
		alternate = cadddr(l)
	}
	return &If{
		Test:       c.compileExpr(test, env),
		Consequent: c.compileExpr(consequent, env),
		Alternate:  c.compileExpr(alternate, env),
	}
}

func (c *Compiler) compileLambda(l Val, llen int, env *cenv) Code {
	if llen < 3 {
		panic("lambda: Illegal form")
	}
	fixed, rest, formals, ok := c.checkLambdaSignature(cadr(l))
	if !ok {
		panic("lambda: Illegal form")
	}
	bodyExpr := cddr(l)
	if llen > 3 {
		bodyExpr = cons(c.beginSym, bodyExpr)
	}
	newEnv := &cenv{link: env, names: formals}
	compiledBodyExpr := c.compileExpr(bodyExpr, newEnv)
	return &Lambda{Fixed: fixed, Rest: rest, Body: compiledBodyExpr}
}

func (c *Compiler) compileLet(l Val, llen int, env *cenv) Code {
	// (let ((id expr) ...) expr expr ...)
	return c.compileLetOrLetrec(l, llen, env, false)
}

func (c *Compiler) compileLetrec(l Val, llen int, env *cenv) Code {
	// (letrec ((id expr) ...) expr expr ...)
	return c.compileLetOrLetrec(l, llen, env, true)
}

func (c *Compiler) compileLetOrLetrec(l Val, llen int, env *cenv, isLetrec bool) Code {
	name := "let"
	if isLetrec {
		name = "letrec"
	}
	if llen < 3 {
		panic(name + ": Illegal form")
	}
	names, inits, bindingsAreOk := c.checkLetBindings(cadr(l))
	if !bindingsAreOk {
		panic(name + ": Illegal form")
	}
	var bodyExpr Val
	if llen > 3 {
		bodyExpr = cons(c.beginSym, cddr(l))
	} else {
		bodyExpr = car(cddr(l))
	}
	// Optimization: Don't introduce a rib if there are no bindings
	if len(names) == 0 {
		return c.compileExpr(bodyExpr, env)
	}
	var compiledInits []Code
	newEnv := &cenv{link: env, names: names}
	if isLetrec {
		compiledInits = c.compileExprSlice(inits, newEnv)
	} else {
		compiledInits = c.compileExprSlice(inits, env)
	}
	compiledBody := c.compileExpr(bodyExpr, newEnv)
	return &Let{Exprs: compiledInits, Body: compiledBody}
}

func (c *Compiler) compileOr(l Val, llen int, env *cenv) Code {
	// (or expr ...)
	if llen == 1 {
		return &Quote{Value: c.s.FalseVal}
	}
	if llen == 2 {
		return c.compileExpr(cadr(l), env)
	}
	// Introduce a let binding to avoid repeated evaluation.
	// Optimization: Don't introduce a let if the first operand is a variable.
	fst := cadr(l)
	useLet := true
	var vname *Symbol
	if s, isSymbol := fst.(*Symbol); isSymbol {
		vname = s
		useLet = false
	} else {
		vname = c.s.Gensym("OR")
	}
	e := cons(c.ifSym, cons(vname, cons(vname, cons(c.orSym, cddr(l)))))
	if useLet {
		e = c.list(c.letSym, c.list(c.list(vname, fst)), e)
	}
	return c.compileExpr(e, env)
}

func (c *Compiler) compileQuote(l Val, llen int, env *cenv) Code {
	// (quote datum)
	if llen != 2 {
		panic("quote: Illegal form")
	}

	// There are probably restrictions on the datum.  It can be:
	//
	// - an exact number
	// - an inexact number
	// - the empty list
	// - unspecified
	// - true
	// - false
	// - a string (TBD)
	// - a character (TBD)
	// - a symbol
	// - a proper or improper list, probably without any non-atomic sharing
	//
	// Notably it can't be a procedure or `undefined` or the eof object?
	//
	// TODO: Check those things.

	return &Quote{Value: cadr(l)}
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

func (c *Compiler) compileSet(l Val, llen int, env *cenv) Code {
	// (set! ident expr)
	if llen != 3 {
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

func (c *Compiler) checkDefinitionSignature(sig Val) (fixed int, rest bool, globName *Symbol, formals []*Symbol, ok bool) {
	f, r, names, k := c.collectNamesFromSignature(sig)
	if !k {
		return
	}
	fixed = f - 1
	rest = r
	globName = names[0]
	formals = names[1:]
	ok = c.namesAreUnique(formals)
	return
}

func (c *Compiler) checkLambdaSignature(sig Val) (fixed int, rest bool, formals []*Symbol, ok bool) {
	f, r, names, k := c.collectNamesFromSignature(sig)
	if !k {
		return
	}
	fixed = f
	rest = r
	formals = names
	ok = c.namesAreUnique(names)
	return
}

func (c *Compiler) checkLetBindings(bindings Val) (names []*Symbol, inits []Val, ok bool) {
	_, bindingsIsList := c.checkProperList(bindings)
	if !bindingsIsList {
		return
	}
	for bindings != c.s.NullVal {
		b := car(bindings)
		bindingLen, bindingIsList := c.checkProperList(b)
		if !bindingIsList || bindingLen != 2 {
			return
		}
		bindingName := car(b)
		bindingExpr := cadr(b)
		nameSym, isSymbol := bindingName.(*Symbol)
		if !isSymbol {
			return
		}
		names = append(names, nameSym)
		inits = append(inits, bindingExpr)
		bindings = cdr(bindings)
	}
	if !c.namesAreUnique(names) {
		return
	}
	ok = true
	return
}

func (c *Compiler) collectNamesFromSignature(sig Val) (fixed int, rest bool, names []*Symbol, ok bool) {
	// TODO: This needs to deal with circularity
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
		names = append(names, argName)
	}
	ok = true
	return
}

func (c *Compiler) checkProperList(v Val) (int, bool) {
	// Check that v is a proper list, and return its length if it is
	// TODO: This needs to deal with circularity
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

func (c *Compiler) namesAreUnique(names []*Symbol) bool {
	for i := 0; i < len(names); i++ {
		for j := i + 1; j < len(names); j++ {
			if names[i] == names[j] {
				return false
			}
		}
	}
	return true
}

func (c *Compiler) list(vs ...Val) Val {
	v := c.s.NullVal
	for i := len(vs) - 1; i >= 0; i-- {
		v = &Cons{Car: vs[i], Cdr: v}
	}
	return v
}

func cons(v Val, vs Val) Val {
	return &Cons{Car: v, Cdr: vs}
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

func cddr(v Val) Val {
	return cdr(cdr(v))
}

func caddr(v Val) Val {
	return car(cdr(cdr(v)))
}

func cadddr(v Val) Val {
	return car(cdr(cdr(cdr(v))))
}
