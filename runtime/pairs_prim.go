// Pairs and lists primitive procedures.
//
// R7RS 6.4, Pairs and lists, also see pairs.sch

package runtime

import (
	. "sint/core"
)

func initPairPrimitives(c *Scheme) {
	addPrimitive(c, "null?", 1, false, primNullp)
	addPrimitive(c, "pair?", 1, false, primPairp)
	addPrimitive(c, "cons", 2, false, primCons)
	addPrimitive(c, "car", 1, false, primCar)
	addPrimitive(c, "cdr", 1, false, primCdr)
	addPrimitive(c, "set-car!", 2, false, primSetcar)
	addPrimitive(c, "set-cdr!", 2, false, primSetcdr)
}

func primNullp(ctx *Scheme, args []Val) (Val, int) {
	if args[0] == ctx.NullVal {
		return ctx.TrueVal, 1
	}
	return ctx.FalseVal, 1
}

func primPairp(ctx *Scheme, args []Val) (Val, int) {
	if _, ok := args[0].(*Cons); ok {
		return ctx.TrueVal, 1
	}
	return ctx.FalseVal, 1
}

func primCons(_ *Scheme, args []Val) (Val, int) {
	return &Cons{Car: args[0], Cdr: args[1]}, 1
}

func primCar(ctx *Scheme, args []Val) (Val, int) {
	if c, ok := args[0].(*Cons); ok {
		return c.Car, 1
	}
	return ctx.Error("car: Not a pair: " + args[0].String())
}

func primCdr(ctx *Scheme, args []Val) (Val, int) {
	if c, ok := args[0].(*Cons); ok {
		return c.Cdr, 1
	}
	return ctx.Error("cdr: Not a pair: " + args[0].String())
}

func primSetcar(ctx *Scheme, args []Val) (Val, int) {
	if c, ok := args[0].(*Cons); ok {
		c.Car = args[1]
		return ctx.UnspecifiedVal, 1
	}
	return ctx.Error("set-car!: Not a pair: " + args[0].String())
}

func primSetcdr(ctx *Scheme, args []Val) (Val, int) {
	if c, ok := args[0].(*Cons); ok {
		c.Cdr = args[1]
		return ctx.UnspecifiedVal, 1
	}
	return ctx.Error("set-cdr!: Not a pair: " + args[0].String())
}
