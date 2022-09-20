// Symbols primitive procedures.
//
// R7RS 6.5, Symbols.  Also see symbols.sch.

package runtime

import (
	. "sint/core"
)

func initSymbolPrimitives(ctx *Scheme) {
	addPrimitive(ctx, "symbol?", 1, false, primSymbolp)
	addPrimitive(ctx, "symbol->string", 1, false, primSymbol2String)
	addPrimitive(ctx, "string->symbol", 1, false, primString2Symbol)
	addPrimitive(ctx, "gensym", 0, false, primGensym)
	addPrimitive(ctx, "apropos", 1, false, primApropos)
}

func primSymbolp(ctx *Scheme, args []Val) (Val, int) {
	if _, ok := args[0].(*Symbol); ok {
		return ctx.TrueVal, 1
	}
	return ctx.FalseVal, 1
}

func primSymbol2String(ctx *Scheme, args []Val) (Val, int) {
	v := args[0]
	if s, ok := v.(*Symbol); ok {
		return &Str{Value: s.Name}, 1
	}
	return ctx.Error("symbol->string: Not a symbol", v)
}

func primString2Symbol(ctx *Scheme, args []Val) (Val, int) {
	v := args[0]
	if s, ok := v.(*Str); ok {
		return ctx.Intern(s.Value), 1
	}
	return ctx.Error("string->symbol: Not a string", v)
}

func primGensym(ctx *Scheme, _ []Val) (Val, int) {
	return ctx.Gensym("S"), 1
}

func primApropos(ctx *Scheme, args []Val) (Val, int) {
	v := args[0]
	pattern := ""
	if s, ok := v.(*Str); ok {
		pattern = s.Value
	} else if s, ok := v.(*Symbol); ok {
		pattern = s.Name
	} else {
		return ctx.Error("apropos: Not a string", v)
	}
	syms := ctx.FindSymbolsByName(pattern)
	l := ctx.NullVal
	for _, s := range syms {
		l = &Cons{Car: s, Cdr: l}
	}
	return l, 1
}
