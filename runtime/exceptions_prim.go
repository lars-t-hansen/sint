// R7RS 6.11 "Exceptions".  Also see exceptions.sch.

package runtime

import (
	. "sint/core"
)

func initExceptionsPrimitives(ctx *Scheme) {
	addPrimitive(ctx, "sint:report-error", 2, false, primReportError)
}

func primReportError(ctx *Scheme, a0, a1 Val, _ []Val) (Val, int) {
	s, sErr := checkString(ctx, a0, "sint:report-error")
	if sErr != nil {
		return ctx.SignalWrappedError(sErr)
	}
	// TODO: Really ought to check that this is a list.  As it is,
	// the system will panic if it is not.  On the other hand, this
	// is in the error handler.
	var xs []Val
	for l := a1; l != ctx.NullVal; l = l.(*Cons).Cdr {
		xs = append(xs, l.(*Cons).Car)
	}
	return ctx.Error(s, xs...)
}
