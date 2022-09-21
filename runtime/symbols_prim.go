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

func primSymbolp(ctx *Scheme, args []Val) (Val, int) {
	v0 := args[0]
	if _, ok := v0.(*Symbol); ok {
		return ctx.TrueVal, 1
	}
	return ctx.FalseVal, 1
}

func primSymbolHasValue(ctx *Scheme, args []Val) (Val, int) {
	v0 := args[0]
	if sym, ok := v0.(*Symbol); ok {
		if sym.Value != ctx.UndefinedVal {
			return ctx.TrueVal, 1
		}
		return ctx.FalseVal, 1
	}
	return ctx.Error("symbol-has-value?: Not a symbol", v0)
}

func primSymbolValue(ctx *Scheme, args []Val) (Val, int) {
	v0 := args[0]
	if sym, ok := v0.(*Symbol); ok {
		if sym.Value != ctx.UndefinedVal {
			return sym.Value, 1
		}
		return ctx.Error("symbol-value: Has no value", v0)
	}
	return ctx.Error("symbol-value: Not a symbol", v0)
}

func primSymbol2String(ctx *Scheme, args []Val) (Val, int) {
	v0 := args[0]
	if s, ok := v0.(*Symbol); ok {
		return &Str{Value: s.Name}, 1
	}
	return ctx.Error("symbol->string: Not a symbol", v0)
}

func primString2Symbol(ctx *Scheme, args []Val) (Val, int) {
	v0 := args[0]
	if s, ok := v0.(*Str); ok {
		return ctx.Intern(s.Value), 1
	}
	return ctx.Error("string->symbol: Not a string", v0)
}

func primGensym(ctx *Scheme, _ []Val) (Val, int) {
	return ctx.Gensym("S"), 1
}

func primFilterGlobals(ctx *Scheme, args []Val) (Val, int) {
	v0 := args[0]
	pattern := ""
	if s, ok := v0.(*Str); ok {
		pattern = s.Value
	} else if s, ok := v0.(*Symbol); ok {
		pattern = s.Name
	} else {
		return ctx.Error("filter-global-variables: Not a string", v0)
	}
	syms := ctx.FindSymbolsByName(pattern)
	l := ctx.NullVal
	for _, s := range syms {
		l = &Cons{Car: s, Cdr: l}
	}
	return l, 1
}
