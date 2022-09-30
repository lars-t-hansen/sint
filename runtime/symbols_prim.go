// Symbols primitive procedures.
//
// R7RS 6.5, Symbols.  Also see symbols.sch.

package runtime

import (
	. "sint/core"
)

func initSymbolPrimitives(ctx *Scheme) {
	addPrimitive(ctx, "symbol?", 1, false, primSymbolp)
	addPrimitive(ctx, "symbol-has-value?", 1, false, primSymbolHasValue)
	addPrimitive(ctx, "symbol-value", 1, false, primSymbolValue)
	addPrimitive(ctx, "symbol->string", 1, false, primSymbol2String)
	addPrimitive(ctx, "string->symbol", 1, false, primString2Symbol)
	addPrimitive(ctx, "gensym", 0, false, primGensym)
	addPrimitive(ctx, "filter-global-variables", 1, false, primFilterGlobals)
}

func primSymbolp(ctx *Scheme, a0, _ Val, _ []Val) (Val, int) {
	if _, ok := a0.(*Symbol); ok {
		return ctx.TrueVal, 1
	}
	return ctx.FalseVal, 1
}

func primSymbolHasValue(ctx *Scheme, a0, _ Val, _ []Val) (Val, int) {
	sym, symErr := ctx.CheckSymbol(a0, "symbol-has-value?")
	if symErr != nil {
		return ctx.SignalWrappedError(symErr)
	}
	if sym.Value != ctx.UndefinedVal {
		return ctx.TrueVal, 1
	}
	return ctx.FalseVal, 1
}

func primSymbolValue(ctx *Scheme, a0, _ Val, _ []Val) (Val, int) {
	sym, symErr := ctx.CheckSymbol(a0, "symbol-value")
	if symErr != nil {
		return ctx.SignalWrappedError(symErr)
	}
	if sym.Value != ctx.UndefinedVal {
		return sym.Value, 1
	}
	return ctx.Error("symbol-value: Symbol has no value", a0)
}

func primSymbol2String(ctx *Scheme, a0, _ Val, _ []Val) (Val, int) {
	sym, symErr := ctx.CheckSymbol(a0, "symbol->string")
	if symErr != nil {
		return ctx.SignalWrappedError(symErr)
	}
	return &Str{Value: sym.Name}, 1
}

func primString2Symbol(ctx *Scheme, a0, _ Val, _ []Val) (Val, int) {
	s, sErr := ctx.CheckString(a0, "string->symbol")
	if sErr != nil {
		return ctx.SignalWrappedError(sErr)
	}
	return ctx.Intern(s), 1
}

func primGensym(ctx *Scheme, _, _ Val, _ []Val) (Val, int) {
	return ctx.Gensym("S"), 1
}

func primFilterGlobals(ctx *Scheme, a0, _ Val, _ []Val) (Val, int) {
	pattern := ""
	if s, ok := a0.(*Str); ok {
		pattern = s.Value
	} else if s, ok := a0.(*Symbol); ok {
		pattern = s.Name
	} else {
		return ctx.Error("filter-global-variables: Not a string or symbol", a0)
	}
	syms := ctx.FindSymbolsByName(pattern)
	l := ctx.NullVal
	for _, s := range syms {
		l = &Cons{Car: s, Cdr: l}
	}
	return l, 1
}
