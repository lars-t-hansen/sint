// R7RS 6.11 "Exceptions".  Also see exceptions.sch.

package runtime

import (
	. "sint/core"
)

func initExceptionsPrimitives(ctx *Scheme) {
	addPrimitive(ctx, "sint:report-error", 2, false, primReportError)
}

func primReportError(ctx *Scheme, args []Val) (Val, int) {
	if s, ok := args[0].(*Str); ok {
		var xs []Val
		for l := args[1]; l != ctx.NullVal; l = l.(*Cons).Cdr {
			xs = append(xs, l.(*Cons).Car)
		}
		return ctx.Error(s.Value, xs...)
	}
	return ctx.Error("sint:report-error: Not a string", args[0])
}
