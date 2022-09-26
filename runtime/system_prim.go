package runtime

import (
	"math/big"
	"os"
	. "sint/core"
	"time"
)

func initSystemPrimitives(c *Scheme) {
	addPrimitive(c, "exit", 0, true, primExit)
	addPrimitive(c, "current-jiffy", 0, false, primCurrentJiffy)
	addPrimitive(c, "jiffies-per-second", 0, false, primJiffiesPerSecond)
}

func primExit(ctx *Scheme, a0, _ Val, _ []Val) (Val, int) {
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

func primCurrentJiffy(ctx *Scheme, _, _ Val, _ []Val) (Val, int) {
	return big.NewInt(time.Now().UnixMicro()), 1
}

func primJiffiesPerSecond(ctx *Scheme, _, _ Val, _ []Val) (Val, int) {
	return big.NewInt(1000000), 1
}
