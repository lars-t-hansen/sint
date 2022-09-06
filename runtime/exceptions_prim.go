package runtime

import (
	. "sint/core"
)

func initExceptionsPrimitives(c *Scheme) {
	addPrimitive(c, "sint:throw-string", 1, false, primThrowString)
}

func primThrowString(ctx *Scheme, args []Val) (Val, int) {
	// This takes one argument, a string
	if s, ok := args[0].(*Str); ok {
		return ctx.Error(s.Value)
	}
	return ctx.Error("sint:throw-string: Not a string: " + args[0].String())
}
