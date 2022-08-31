package runtime

import (
	. "sint/core"
)

func initConcurrencyPrimitives(c *Scheme) {
	addPrimitive(c, "sint:go", 1, false, primGo)
}

func primGo(ctx *Scheme, args []Val) (Val, int) {
	ctx.InvokeConcurrent(args[0])
	return ctx.UnspecifiedVal, 1
}
