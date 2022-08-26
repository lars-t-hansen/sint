package runtime

import (
	. "sint/core"
)

func initExceptionsPrimitives(c *Scheme) {
	addPrimitive(c, "sint:throw-string", 1, false, primThrowString)
}

func primThrowString(c *Scheme, args []Val) Val {
	// This takes one argument, a string
	if s, ok := args[0].(*Str); ok {
		panic(s.Value)
	}
	panic("sint:throw-string: Not a string: " + args[0].String())
}
