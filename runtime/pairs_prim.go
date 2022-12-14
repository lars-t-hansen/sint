// Pairs and lists primitive procedures.
//
// R7RS 6.4, Pairs and lists, also see pairs.sch

package runtime

import (
	. "sint/core"
)

func initPairPrimitives(ctx *Scheme) {
	addPrimitive(ctx, "null?", 1, false, primNullp)
	addPrimitive(ctx, "pair?", 1, false, primPairp)
	addPrimitive(ctx, "cons", 2, false, primCons)
	addPrimitive(ctx, "car", 1, false, primCar)
	addPrimitive(ctx, "cdr", 1, false, primCdr)
	addPrimitive(ctx, "set-car!", 2, false, primSetcar)
	addPrimitive(ctx, "set-cdr!", 2, false, primSetcdr)
}

func primNullp(ctx *Scheme, a0, _ Val, _ []Val) (Val, int) {
	if a0 == ctx.NullVal {
		return ctx.TrueVal, 1
	}
	return ctx.FalseVal, 1
}

func primPairp(ctx *Scheme, a0, _ Val, _ []Val) (Val, int) {
	if _, ok := a0.(*Cons); ok {
		return ctx.TrueVal, 1
	}
	return ctx.FalseVal, 1
}

func primCons(_ *Scheme, a0, a1 Val, _ []Val) (Val, int) {
	return &Cons{Car: a0, Cdr: a1}, 1
}

func primCar(ctx *Scheme, a0, _ Val, _ []Val) (Val, int) {
	c, err := ctx.CheckPair(a0, "car")
	if err != nil {
		return ctx.SignalWrappedError(err)
	}
	return c.Car, 1
}

func primCdr(ctx *Scheme, a0, _ Val, _ []Val) (Val, int) {
	c, err := ctx.CheckPair(a0, "cdr")
	if err != nil {
		return ctx.SignalWrappedError(err)
	}
	return c.Cdr, 1
}

func primSetcar(ctx *Scheme, a0, a1 Val, _ []Val) (Val, int) {
	c, err := ctx.CheckPair(a0, "set-car!")
	if err != nil {
		return ctx.SignalWrappedError(err)
	}
	c.Car = a1
	return ctx.UnspecifiedVal, 1
}

func primSetcdr(ctx *Scheme, a0, a1 Val, _ []Val) (Val, int) {
	c, err := ctx.CheckPair(a0, "set-cdr!")
	if err != nil {
		return ctx.SignalWrappedError(err)
	}
	c.Cdr = a1
	return ctx.UnspecifiedVal, 1
}
