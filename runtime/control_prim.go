// Control features primitive procedures.
//
// R7RS 6.10, Control features, also see control.sch

package runtime

import (
	. "sint/core"
)

func initControlPrimitives(c *Scheme) {
	addPrimitive(c, "procedure?", 1, false, primProcedurep)
}

func primProcedurep(ctx *Scheme, args []Val) Val {
	if _, ok := args[0].(*Symbol); ok {
		return ctx.TrueVal
	}
	return ctx.FalseVal
}
