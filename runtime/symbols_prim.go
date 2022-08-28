// Symbols primitive procedures.
//
// R7RS 6.5, Symbols.  Also see symbols.sch.

package runtime

import (
	. "sint/core"
)

func initSymbolPrimitives(c *Scheme) {
	addPrimitive(c, "symbol?", 1, false, primSymbolp)
	addPrimitive(c, "symbol->string", 1, false, primSymbol2String)
	addPrimitive(c, "string->symbol", 1, false, primString2Symbol)
	addPrimitive(c, "gensym", 0, false, primGensym)
}

func primSymbolp(ctx *Scheme, args []Val) (Val, int) {
	if _, ok := args[0].(*Symbol); ok {
		return ctx.TrueVal, 1
	}
	return ctx.FalseVal, 1
}

func primSymbol2String(c *Scheme, args []Val) (Val, int) {
	v := args[0]
	if s, ok := v.(*Symbol); ok {
		return &Str{Value: s.Name}, 1
	}
	panic("symbol->string: Not a symbol: " + v.String())
}

func primString2Symbol(c *Scheme, args []Val) (Val, int) {
	v := args[0]
	if s, ok := v.(*Str); ok {
		return c.Intern(s.Value), 1
	}
	panic("string->symbol: Not a string: " + v.String())
}

func primGensym(c *Scheme, _ []Val) (Val, int) {
	return c.Gensym("S"), 1
}
