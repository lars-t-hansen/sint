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
	return ctx.Error("symbol->string: Not a symbol: " + v.String())
}

func primString2Symbol(ctx *Scheme, args []Val) (Val, int) {
	v := args[0]
	if s, ok := v.(*Str); ok {
		return ctx.Intern(s.Value), 1
	}
	return ctx.Error("string->symbol: Not a string: " + v.String())
}

func primGensym(ctx *Scheme, _ []Val) (Val, int) {
	return ctx.Gensym("S"), 1
}
