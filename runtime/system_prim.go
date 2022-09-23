package runtime

import (
	"math/big"
	"os"
	. "sint/core"
)

func initSystemPrimitives(c *Scheme) {
	addPrimitive(c, "exit", 0, true, primExit)
}

func primExit(ctx *Scheme, a0, _ Val, rest []Val) (Val, int) {
	code := 0
	if a0 != ctx.UndefinedVal {
		v := a0
		if _, ok := v.(*True); ok {
			// nothing
		} else if _, ok := v.(*False); ok {
			code = 1
		} else if n, ok := v.(*big.Int); ok {
			if n.IsInt64() {
				code = int(n.Int64())
			} else {
				code = 1
			}
		} else {
			code = 1
		}
	}
	os.Exit(code)
	return ctx.UnspecifiedVal, 1
}
