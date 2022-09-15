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
		// TODO: Really ought to check that this is a list.  As it is,
		// the system will panic if it is not.
		var xs []Val
		for l := args[1]; l != ctx.NullVal; l = l.(*Cons).Cdr {
			xs = append(xs, l.(*Cons).Car)
		}
		return ctx.Error(s.Value, xs...)
	}
	return ctx.Error("sint:report-error: Not a string", args[0])
}
