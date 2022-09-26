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

func primSymbolp(ctx *Scheme, a0, _ Val, rest []Val) (Val, int) {
	if _, ok := a0.(*Symbol); ok {
		return ctx.TrueVal, 1
	}
	return ctx.FalseVal, 1
}

func primSymbolHasValue(ctx *Scheme, a0, _ Val, rest []Val) (Val, int) {
	if sym, ok := a0.(*Symbol); ok {
		if sym.Value != ctx.UndefinedVal {
			return ctx.TrueVal, 1
		}
		return ctx.FalseVal, 1
	}
	return ctx.Error("symbol-has-value?: Not a symbol", a0)
}

func primSymbolValue(ctx *Scheme, a0, _ Val, rest []Val) (Val, int) {
	if sym, ok := a0.(*Symbol); ok {
		if sym.Value != ctx.UndefinedVal {
			return sym.Value, 1
		}
		return ctx.Error("symbol-value: Has no value", a0)
	}
	return ctx.Error("symbol-value: Not a symbol", a0)
}

func primSymbol2String(ctx *Scheme, a0, _ Val, rest []Val) (Val, int) {
	if s, ok := a0.(*Symbol); ok {
		return &Str{Value: s.Name}, 1
	}
	return ctx.Error("symbol->string: Not a symbol", a0)
}

func primString2Symbol(ctx *Scheme, a0, _ Val, rest []Val) (Val, int) {
	if s, ok := a0.(*Str); ok {
		return ctx.Intern(s.Value), 1
	}
	return ctx.Error("string->symbol: Not a string", a0)
}

func primGensym(ctx *Scheme, _, _ Val, rest []Val) (Val, int) {
	return ctx.Gensym("S"), 1
}

func primFilterGlobals(ctx *Scheme, a0, _ Val, rest []Val) (Val, int) {
	pattern := ""
	if s, ok := a0.(*Str); ok {
		pattern = s.Value
	} else if s, ok := a0.(*Symbol); ok {
		pattern = s.Name
	} else {
		return ctx.Error("filter-global-variables: Not a string", a0)
	}
	syms := ctx.FindSymbolsByName(pattern)
	l := ctx.NullVal
	for _, s := range syms {
		l = &Cons{Car: s, Cdr: l}
	}
	return l, 1
}
