package runtime

import (
	. "sint/core"
)

// R7RS 6.5, Symbols
// TODO: string->symbol
func initSymbolPrimitives(c *Scheme) {
	addPrimitive(c, "symbol?", 1, false, primSymbolp)
	addPrimitive(c, "symbol->string", 0, false, primSymbol2String)
	addPrimitive(c, "gensym", 0, false, primGensym)

}

func primSymbolp(ctx *Scheme, args []Val) Val {
	if _, ok := args[0].(*Symbol); ok {
		return ctx.TrueVal
	}
	return ctx.FalseVal
}

func primSymbol2String(c *Scheme, args []Val) Val {
	v := args[0]
	if s, ok := v.(*Symbol); ok {
		return &Str{Value: s.Name}
	}
	panic("string->symbol: Not a symbol: " + v.String())
}

func primGensym(c *Scheme, _ []Val) Val {
	return c.Gensym("S")
}
