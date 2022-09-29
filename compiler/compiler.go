// Simple compiler from standard sexpr form to internal Code form.
//
// Built-in syntactic forms have reserved names, which is not very "Scheme"
// but is just fine in practice, and the compiler is also a macro-expander
// for these forms.  There are no user-defined macros at this time.
//
// TODO: A more sophisticated compiler would take an AST form that also carries
// source location information, or at least a map of such information on the side.

package compiler

import (
	"math/big"
	. "sint/core"
)

type Compiler struct {
	s        *SharedScheme
	keywords map[*Symbol]bool
}

// A Compiler currently has no interesting mutable state, it can be reused for
// multiple compilations.

func NewCompiler(s *SharedScheme) *Compiler {
	c := &Compiler{
		s:        s,
		keywords: make(map[*Symbol]bool),
	}

	// These symbols cannot be used as variable names
	c.keywords[s.AndSym] = true
	c.keywords[s.BeginSym] = true
	c.keywords[s.CaseSym] = true
	c.keywords[s.CondSym] = true
	c.keywords[s.DefineSym] = true
	c.keywords[s.DoSym] = true
	c.keywords[s.GoSym] = true
	c.keywords[s.IfSym] = true
	c.keywords[s.LambdaSym] = true
	c.keywords[s.LetSym] = true
	c.keywords[s.LetStarSym] = true
	c.keywords[s.LetValuesSym] = true
	c.keywords[s.LetStarValuesSym] = true
	c.keywords[s.LetrecSym] = true
	c.keywords[s.OrSym] = true
	c.keywords[s.ParameterizeSym] = true
	c.keywords[s.QuoteSym] = true
	c.keywords[s.SetSym] = true
	// arrowSym and elseSym are not reserved

	return c
}

type CompilerError struct {
	msg string
}

func NewCompilerError(msg string) *CompilerError {
	return &CompilerError{msg: msg}
}

func (e *CompilerError) Error() string {
	return e.msg
}

// Returns (nil, *CompilerError) on error, the error holds data explaining
// the problem.  The compiler panics only on internal errors.

func (c *Compiler) CompileToplevel(v Val) (Code, error) {
	length, exprIsList := c.checkProperList(v)
	var compiled Code
	var err error
	if exprIsList && length >= 3 && car(v) == c.s.DefineSym {
		compiled, err = c.compileToplevelDefinition(v)
	} else {
		compiled, err = c.compileExpr(v, &cenv{doc: ""})
	}
	if err != nil {
		return nil, err
	}
	return compiled, nil
}

// Compile-time environment.  There is one of these per lexical rib, and there is always
// an empty one that is outermost, so nil checks for the cenv are never required except
// when walking the environment chain.  The `doc` is a string that can be attached to
// lambda expressions that appear in certain value positions.
//
// Note that this is currently mutable: `doc`` is updated destructively when compiling
// let-like expressions, and `names` is updated during the compilation of let*.  It
// wouldn't be too hard to fix that.

type cenv struct {
	link  *cenv
	doc   string
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

func (c *Compiler) reportError(msg string) (Code, error) {
	return nil, &CompilerError{msg}
}

func (c *Compiler) compileToplevelDefinition(v Val) (Code, error) {
	nameOrSignature := cadr(v)
	// (define x v)
	if globName, ok := nameOrSignature.(*Symbol); ok {
		rhs, err := c.compileExpr(caddr(v), &cenv{doc: globName.Name})
		if err != nil {
			return nil, err
		}
		return &Setglobal{
			Name: globName,
			Rhs:  rhs,
		}, nil
	}
	// (define (f arg ... . arg) body ...)
	if fixed, rest, globName, formals, ok := c.checkDefinitionSignature(nameOrSignature); ok {
		body := c.wrapBodyList(cddr(v))
		bodyc, err := c.compileExpr(body, &cenv{names: formals, doc: globName.Name + " > "})
		if err != nil {
			return nil, err
		}
		lam := &Lambda{
			Fixed: fixed,
			Rest:  rest,
			Body:  bodyc,
			Name:  globName.Name}
		return &Setglobal{
			Name: globName,
			Rhs:  lam,
		}, nil
	}
	return c.reportError("Invalid top-level definition") // TODO: msg
}

// `bodyList`` is a list of expressions.  If its length is 1, return the first element,
// otherwise turn it into a BEGIN.

func (c *Compiler) wrapBodyList(bodyList Val) Val {
	if cdr(bodyList) != c.s.NullVal {
		return cons(c.s.BeginSym, bodyList)
	}
	return car(bodyList)
}

func (c *Compiler) compileExpr(v Val, env *cenv) (Code, error) {
	switch e := v.(type) {
	case *big.Int:
		return &Quote{Value: e}, nil
	case *big.Float:
		return &Quote{Value: e}, nil
	case *Char:
		return &Quote{Value: e}, nil
	case *True:
		return &Quote{Value: e}, nil
	case *False:
		return &Quote{Value: e}, nil
	case *Unspecified:
		return &Quote{Value: e}, nil
	case *Undefined:
		return &Quote{Value: e}, nil
	case *Symbol:
		return c.compileRef(e, env)
	case *Str:
		return &Quote{Value: e}, nil
	case *Regexp:
		return &Quote{Value: e}, nil
	case *Cons:
		llen, exprIsList := c.checkProperList(e)
		if !exprIsList {
			return c.reportError("Improper list used as expression") // TODO: msg
		}
		if llen == 0 {
			return c.reportError("Unquoted empty list used as expression")
		}
		if kwd, ok := e.Car.(*Symbol); ok {
			if kwd == c.s.AndSym {
				return c.compileAnd(e, llen, env)
			}
			if kwd == c.s.BeginSym {
				return c.compileBegin(e, llen, env)
			}
			if kwd == c.s.CaseSym {
				return c.compileCase(e, llen, env)
			}
			if kwd == c.s.CondSym {
				return c.compileCond(e, llen, env)
			}
			if kwd == c.s.DoSym {
				return c.compileDo(e, llen, env)
			}
			if kwd == c.s.GoSym {
				return c.compileGo(e, llen, env)
			}
			if kwd == c.s.LambdaSym {
				return c.compileLambda(e, llen, env)
			}
			if kwd == c.s.LetSym {
				return c.compileLet(e, llen, env)
			}
			if kwd == c.s.LetStarSym {
				return c.compileLetStar(e, llen, env)
			}
			if kwd == c.s.LetValuesSym {
				return c.compileLetValues(e, llen, env)
			}
			if kwd == c.s.LetStarValuesSym {
				return c.compileLetStarValues(e, llen, env)
			}
			if kwd == c.s.LetrecSym {
				return c.compileLetrec(e, llen, env)
			}
			if kwd == c.s.OrSym {
				return c.compileOr(e, llen, env)
			}
			if kwd == c.s.ParameterizeSym {
				return c.compileParameterize(e, llen, env)
			}
			if kwd == c.s.QuoteSym {
				return c.compileQuote(e, llen, env)
			}
			if kwd == c.s.IfSym {
				return c.compileIf(e, llen, env)
			}
			if kwd == c.s.SetSym {
				return c.compileSet(e, llen, env)
			}
			// Fall through to generic "call" case
		}
		return c.compileCall(e, llen, env)
	default:
		return c.reportError("Bad expression: " + v.String())
	}
}

func (c *Compiler) compileExprList(l Val, env *cenv) ([]Code, error) {
	var exprs []Code
	for l != c.s.NullVal {
		ce, err := c.compileExpr(car(l), env)
		if err != nil {
			return nil, err
		}
		exprs = append(exprs, ce)
		l = cdr(l)
	}
	return exprs, nil
}

func (c *Compiler) compileExprSlice(es []Val, env *cenv) ([]Code, error) {
	var exprs []Code
	for _, e := range es {
		ce, err := c.compileExpr(e, env)
		if err != nil {
			return nil, err
		}
		exprs = append(exprs, ce)
	}
	return exprs, nil
}

func (c *Compiler) compileAnd(l Val, llen int, env *cenv) (Code, error) {
	// (and expr ...)
	if llen == 1 {
		return &Quote{Value: c.s.TrueVal}, nil
	}
	if llen == 2 {
		return c.compileExpr(cadr(l), env)
	}
	e1, err1 := c.compileExpr(cadr(l), env)
	if err1 != nil {
		return nil, err1
	}
	e2, err2 := c.compileExpr(cons(c.s.AndSym, cddr(l)), env)
	if err2 != nil {
		return nil, err2
	}
	e3 := &Quote{Value: c.s.FalseVal}
	return &If{Test: e1, Consequent: e2, Alternate: e3}, nil
}

func (c *Compiler) compileBegin(l Val, llen int, env *cenv) (Code, error) {
	// (begin expr ...)
	if llen == 1 {
		return &Quote{Value: c.s.UnspecifiedVal}, nil
	}
	// Optimization: Single-expression BEGIN becomes just the expression
	if llen == 2 {
		return c.compileExpr(cadr(l), env)
	}
	es, err := c.compileExprList(cdr(l), env)
	if err != nil {
		return nil, err
	}
	return &Begin{Exprs: es}, nil
}

func (c *Compiler) compileCall(l Val, _ int, env *cenv) (Code, error) {
	// (expr expr ...)
	es, err := c.compileExprList(l, env)
	if err != nil {
		return nil, err
	}
	return &Call{Exprs: es}, nil
}

func (c *Compiler) compileCase(l Val, llen int, env *cenv) (Code, error) {
	return c.reportError("`case` not implemented")
}

func (c *Compiler) compileCond(l Val, llen int, env *cenv) (Code, error) {
	// ("cond" clause ...)
	// ("cond" clause ... ("else" expr ...))
	// clause ::= (expr expr  ...)
	//          | (expr "=>" expr)
	clauses := cdr(l)
	if clauses == c.s.NullVal {
		return &Quote{Value: c.s.UnspecifiedVal}, nil
	}
	clause := car(clauses)
	rest := cdr(clauses)
	clauseLen, clauseOk := c.checkProperList(clause)
	if !clauseOk || clauseLen == 0 {
		return c.reportError("Bad clause in `cond`:" + clause.String())
	}
	if clauseLen >= 2 && cadr(clause) == c.s.ArrowSym {
		if clauseLen != 3 {
			return c.reportError("Bad clause in `cond`:" + clause.String())
		}
		// expr => fn becomes (let ((v expr)) (if v (fn v) (cond <rest-of-clauses>)))
		v := c.s.Gensym("COND")
		return c.compileExpr(
			c.list(c.s.LetSym, c.list(c.list(v, car(clause))),
				c.list(c.s.IfSym, v,
					c.list(caddr(clause), v),
					cons(c.s.CondSym, rest))),
			env)
	}
	if car(clause) == c.s.ElseSym {
		if clauseLen < 2 {
			return c.reportError("Bad `else` clause in `cond`:" + clause.String())
		}
		if rest != c.s.NullVal {
			return c.reportError("The `else` clause must be last in `cond`:" + clause.String())
		}
		// (else expr expr ...) becomes (begin expr expr ...)
		return c.compileExpr(c.wrapBodyList(cdr(clause)), env)
	}
	if clauseLen == 1 {
		// (expr) becomes (or expr (cond <rest-of-clauses>))
		return c.compileExpr(c.list(c.s.OrSym, car(clause), cons(c.s.CondSym, rest)), env)
	}
	// (expr0 expr1 expr2 ...) becomes (if expr0 (begin expr1 expr2 ...) (cond <rest-of-clauses>))
	return c.compileExpr(c.list(c.s.IfSym, car(clause), c.wrapBodyList(cdr(clause)), cons(c.s.CondSym, rest)), env)
}

func (c *Compiler) compileDo(l Val, llen int, env *cenv) (Code, error) {
	return c.reportError("`do` not implemented")
}

func (c *Compiler) compileGo(l Val, llen int, env *cenv) (Code, error) {
	// (go (expr0 expr1 ...)) becomes (sint:go (let ((v0 expr0) (v1 expr1) ...) (lambda () (v0 v1 ...))))
	if llen != 2 {
		panic("go: invalid syntax: " + l.String())
	}
	callExpr := cadr(l)
	cLen, cOk := c.checkProperList(callExpr)
	if !cOk || cLen == 0 {
		panic("go: invalid call syntax: " + l.String())
	}
	var fstBinding *Cons
	var lastBinding *Cons
	var fstArg *Cons
	var lastArg *Cons
	e := callExpr
	for i := 0; i < cLen; i++ {
		// A temporary name for the value
		v := c.s.Gensym("GO")

		// A new binding for the introduced `let`
		b := c.list(v, car(e))
		bCell := &Cons{Car: b, Cdr: c.s.NullVal}
		e = cdr(e)
		if fstBinding == nil {
			fstBinding = bCell
		} else {
			lastBinding.Cdr = bCell
		}
		lastBinding = bCell

		// Another expr on the call
		aCell := &Cons{Car: v, Cdr: c.s.NullVal}
		if fstArg == nil {
			fstArg = aCell
		} else {
			lastArg.Cdr = aCell
		}
		lastArg = aCell
	}
	if e != c.s.NullVal {
		panic("Inconsistency in implementation of GO")
	}
	lambdaExpr := c.list(c.s.LambdaSym, c.s.NullVal, fstArg)
	letExpr := c.list(c.s.LetSym, fstBinding, lambdaExpr)
	return c.compileExpr(c.list(c.s.Intern("sint:go"), letExpr), env)
}

func (c *Compiler) compileIf(l *Cons, llen int, env *cenv) (Code, error) {
	// (if expr expr)
	// (if expr expr expr)
	if llen != 3 && llen != 4 {
		return c.reportError("if: Illegal form: " + l.String())
	}
	test := cadr(l)
	consequent := caddr(l)
	var alternate Val = c.s.UnspecifiedVal
	if llen == 4 {
		alternate = cadddr(l)
	}
	e1, err1 := c.compileExpr(test, env)
	if err1 != nil {
		return nil, err1
	}
	e2, err2 := c.compileExpr(consequent, env)
	if err2 != nil {
		return nil, err2
	}
	e3, err3 := c.compileExpr(alternate, env)
	if err3 != nil {
		return nil, err3
	}
	return &If{Test: e1, Consequent: e2, Alternate: e3}, nil
}

func (c *Compiler) compileLambda(l Val, llen int, env *cenv) (Code, error) {
	if llen < 3 {
		return c.reportError("lambda: Illegal form: " + l.String())
	}
	fixed, rest, formals, ok := c.checkLambdaSignature(cadr(l))
	if !ok {
		return c.reportError("lambda: Illegal form: " + l.String())
	}
	bodyExpr := c.wrapBodyList(cddr(l))
	newEnv := &cenv{link: env, names: formals, doc: env.doc + " > [lambda]"}
	compiledBodyExpr, err := c.compileExpr(bodyExpr, newEnv)
	if err != nil {
		return nil, err
	}
	return &Lambda{Fixed: fixed, Rest: rest, Body: compiledBodyExpr, Name: env.doc}, nil
}

func (c *Compiler) compileLet(l Val, llen int, env *cenv) (Code, error) {
	// (let ((id expr) ...) expr expr ...)
	// (let id ((id expr) ...) expr expr ...)
	if llen >= 4 {
		if _, ok := cadr(l).(*Symbol); ok {
			return c.compileNamedLet(l, llen, env)
		}
	}
	return c.compileLetOrLetrecOrLetStar(l, llen, env, kLet)
}

func (c *Compiler) compileLetrec(l Val, llen int, env *cenv) (Code, error) {
	// (letrec ((id expr) ...) expr expr ...)
	return c.compileLetOrLetrecOrLetStar(l, llen, env, kLetrec)
}

func (c *Compiler) compileLetStar(l Val, llen int, env *cenv) (Code, error) {
	// (let* ((id expr) ...) expr expr ...)
	return c.compileLetOrLetrecOrLetStar(l, llen, env, kLetStar)
}

func (c *Compiler) compileLetValues(l Val, llen int, env *cenv) (Code, error) {
	// (let-values ((bindings expr) ...) expr expr ...)
	return c.reportError("`let-values` not implemented yet") // TODO
}

func (c *Compiler) compileLetStarValues(l Val, llen int, env *cenv) (Code, error) {
	// (let*-values ((bindings expr) ...) expr expr ...)
	return c.reportError("`let*-values` not implemented yet") // TODO
}

type LetKind int

const (
	kLet LetKind = iota
	kLetrec
	kLetStar
)

var letNames []string = []string{"let", "letrec", "let*"}

func (c *Compiler) compileLetOrLetrecOrLetStar(l Val, llen int, env *cenv, kind LetKind) (Code, error) {
	name := letNames[kind]
	if llen < 3 {
		return c.reportError(name + ": Illegal form: " + l.String())
	}
	names, inits, bindingsAreOk := c.checkLetBindings(cadr(l))
	if !bindingsAreOk {
		return c.reportError(name + ": Illegal form: " + l.String())
	}
	bodyExpr := c.wrapBodyList(cddr(l))
	// Optimization: Don't introduce a rib if there are no bindings
	if len(names) == 0 {
		return c.compileExpr(bodyExpr, env)
	}
	var newEnv *cenv
	var compiledInits []Code
	var err error
	switch kind {
	case kLet:
		savedEnvDoc := env.doc
		for i, init := range inits {
			env.doc = savedEnvDoc + " > " + names[i].Name
			compiledInit, compileErr := c.compileExpr(init, env)
			if compileErr != nil {
				err = compileErr
				break
			}
			compiledInits = append(compiledInits, compiledInit)
		}
		env.doc = savedEnvDoc
		newEnv = &cenv{link: env, names: names, doc: env.doc}
	case kLetrec:
		newEnv = &cenv{link: env, names: names}
		for i, init := range inits {
			newEnv.doc = env.doc + " > " + names[i].Name
			compiledInit, compileErr := c.compileExpr(init, newEnv)
			if compileErr != nil {
				err = compileErr
				break
			}
			compiledInits = append(compiledInits, compiledInit)
		}
		newEnv.doc = env.doc
	case kLetStar:
		newEnv = &cenv{link: env, names: []*Symbol{}}
		for i, init := range inits {
			newEnv.doc = env.doc + " > " + names[i].Name
			compiledInit, compileErr := c.compileExpr(init, newEnv)
			if compileErr != nil {
				err = compileErr
				break
			}
			compiledInits = append(compiledInits, compiledInit)
			newEnv.names = append(newEnv.names, names[i])
		}
		newEnv.doc = env.doc
	default:
		panic("Unexpected LetKind")
	}
	if err != nil {
		return nil, err
	}
	compiledBody, bodyErr := c.compileExpr(bodyExpr, newEnv)
	if bodyErr != nil {
		return nil, bodyErr
	}
	switch kind {
	case kLet:
		return &Let{Exprs: compiledInits, Body: compiledBody}, nil
	case kLetrec:
		return &Letrec{Exprs: compiledInits, Body: compiledBody}, nil
	case kLetStar:
		return &LetStar{Exprs: compiledInits, Body: compiledBody}, nil
	default:
		panic("Unexpected LetKind")
	}
}

func (c *Compiler) compileNamedLet(l Val, llen int, env *cenv) (Code, error) {
	return c.reportError("Named `let` not implemented yet") // TODO
}

func (c *Compiler) compileOr(l Val, llen int, env *cenv) (Code, error) {
	// (or expr ...)
	if llen == 1 {
		return &Quote{Value: c.s.FalseVal}, nil
	}
	if llen == 2 {
		return c.compileExpr(cadr(l), env)
	}
	// Introduce a let binding to avoid repeated evaluation.
	// Optimization: Don't introduce a let if the first operand is a variable.
	first := cadr(l)
	rest := cddr(l)
	useLet := true
	var vname *Symbol
	if s, isSymbol := first.(*Symbol); isSymbol {
		vname = s
		useLet = false
	} else {
		vname = c.s.Gensym("OR")
	}
	e := c.list(c.s.IfSym, vname, vname, cons(c.s.OrSym, rest))
	if useLet {
		e = c.list(c.s.LetSym, c.list(c.list(vname, first)), e)
	}
	return c.compileExpr(e, env)
}

func (c *Compiler) compileParameterize(l Val, llen int, env *cenv) (Code, error) {
	// (parameterize ((p-expr0 v-expr0) ...) expr0 expr1 ...) becomes
	//    (let ((p-name0 p-expr0) ... (v-name0 v-expr0) ...)
	//      (let ((old-name0 (p-name0)) ...)
	//        (sint:dynamic-wind
	//           (lambda ()
	//             (p-name0 v-name0) ...)
	//           (lambda ()
	//              expr0 expr1 ...)
	//           (lambda ()
	//             (p-name0 old-name0) ...))))
	// The p-name*, v-name*, and old-name* are all fresh names.
	// We use sint:dynamic-wind to avoid capturing any dynamic-wind that's lexically bound.
	return c.reportError("`parameterize` not implemented yet") // TODO
}

func (c *Compiler) compileQuote(l Val, llen int, env *cenv) (Code, error) {
	// (quote datum)
	if llen != 2 {
		return c.reportError("quote: Illegal form: " + l.String())
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
	// - a character
	// - a symbol
	// - a proper or improper list, probably without any non-atomic sharing
	//
	// Notably it can't be a procedure or `undefined` or a channel.
	//
	// TODO: Check those things.

	return &Quote{Value: cadr(l)}, nil
}

func (c *Compiler) compileRef(s *Symbol, env *cenv) (Code, error) {
	// ident
	if levels, offset, ok := lookup(env, s); ok {
		return &Lexical{Levels: levels, Offset: offset}, nil
	}
	if c.isKeyword(s) {
		return c.reportError("Keyword used as variable reference: " + s.Name)
	}
	return &Global{Name: s}, nil
}

func (c *Compiler) compileSet(l Val, llen int, env *cenv) (Code, error) {
	// (set! ident expr)
	if llen != 3 {
		return c.reportError("set!: Illegal form: " + l.String())
	}
	place := cadr(l)
	expr := caddr(l)
	placeName, nameIsSymbol := place.(*Symbol)
	if !nameIsSymbol {
		return c.reportError("set!: Illegal variable name: " + place.String())
	}
	rhs, rhsErr := c.compileExpr(expr, env)
	if rhsErr != nil {
		return nil, rhsErr
	}
	if levels, offset, ok := lookup(env, placeName); ok {
		return &Setlex{Levels: levels, Offset: offset, Rhs: rhs}, nil
	}
	if c.isKeyword(placeName) {
		return c.reportError("Keyword used as variable name: " + placeName.Name)
	}
	return &Setglobal{Name: placeName, Rhs: rhs}, nil
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
