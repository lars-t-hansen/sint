package runtime

import (
	. "sint/core"
)

// R7RS 6.4, Pairs and lists, also see pairs.sch
func initPairPrimitives(c *Scheme) {
	addPrimitive(c, "null?", 1, false, primNullp)
	addPrimitive(c, "pair?", 1, false, primPairp)
	addPrimitive(c, "cons", 2, false, primCons)
	addPrimitive(c, "car", 1, false, primCar)
	addPrimitive(c, "cdr", 1, false, primCdr)
	addPrimitive(c, "set-car!", 2, false, primSetcar)
	addPrimitive(c, "set-cdr!", 2, false, primSetcdr)
}

func primNullp(ctx *Scheme, args []Val) Val {
	if args[0] == ctx.NullVal {
		return ctx.TrueVal
	}
	return ctx.FalseVal
}

func primPairp(ctx *Scheme, args []Val) Val {
	if _, ok := args[0].(*Cons); ok {
		return ctx.TrueVal
	}
	return ctx.FalseVal
}

func checkCons(v Val, fn string) *Cons {
	if c, ok := v.(*Cons); ok {
		return c
	}
	panic(fn + ": Not a pair: " + v.String())
}

func primCons(_ *Scheme, args []Val) Val {
	return &Cons{Car: args[0], Cdr: args[1]}
}

func primCar(_ *Scheme, args []Val) Val {
	return checkCons(args[0], "car").Car
}

func primCdr(_ *Scheme, args []Val) Val {
	return checkCons(args[0], "cdr").Cdr
}

func primSetcar(ctx *Scheme, args []Val) Val {
	checkCons(args[0], "set-car!").Car = args[1]
	return ctx.UnspecifiedVal
}

func primSetcdr(ctx *Scheme, args []Val) Val {
	checkCons(args[0], "set-cdr!").Cdr = args[1]
	return ctx.UnspecifiedVal
}
