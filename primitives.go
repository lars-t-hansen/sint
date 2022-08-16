package sint

import "math/big"

func (c *Scheme) addPrimitive(name string, fixed int, rest bool, primop func(*Scheme, []Val) Val) {
	sym := c.intern(name)
	sym.value = &Procedure{lambda: &Lambda{fixed: fixed, rest: rest, body: nil}, env: nil, primop: primop}
}

func (c *Scheme) initPrimitives() {
	// TODO: These could go in a table, it doesn't have to be code
	c.addPrimitive("cons", 2, false, primCons)
	c.addPrimitive("car", 1, false, primCar)
	c.addPrimitive("cdr", 1, false, primCdr)
	c.addPrimitive("set-car!", 2, false, primSetcar)
	c.addPrimitive("set-cdr!", 2, false, primSetcdr)
	c.addPrimitive("null?", 1, false, primNullp)
	c.addPrimitive("+", 0, true, primAdd)
	c.addPrimitive("-", 1, true, primSub)
	c.addPrimitive("<", 2, true, primLess)
	c.addPrimitive("=", 2, true, primEqual)
}

func checkCons(v Val, fn string) *Cons {
	if c, ok := v.(*Cons); ok {
		return c
	}
	panic(fn + ": Not a pair")
}

func primCons(_ *Scheme, args []Val) Val {
	return &Cons{args[0], args[1]}
}

func primCar(_ *Scheme, args []Val) Val {
	return checkCons(args[0], "CAR").car
}

func primCdr(_ *Scheme, args []Val) Val {
	return checkCons(args[0], "CDR").cdr
}

func primSetcar(ctx *Scheme, args []Val) Val {
	checkCons(args[0], "SET-CAR!").car = args[1]
	return ctx.unspecified
}

func primSetcdr(ctx *Scheme, args []Val) Val {
	checkCons(args[0], "SET-CDR!").cdr = args[1]
	return ctx.unspecified
}

func primNullp(ctx *Scheme, args []Val) Val {
	if args[0] == ctx.null {
		return ctx.trueVal
	}
	return ctx.falseVal
}

func primAdd(_ *Scheme, args []Val) Val {
	if len(args) == 0 {
		return big.NewInt(0)
	}
	if len(args) == 1 {
		assertNumeric(args[0], "unary '+'")
		return args[0]
	}
	intVals, floatVals := checkAndCoerceNumbers(args, "'+'")
	if intVals != nil {
		var r big.Int = *intVals[0]
		for _, v := range intVals[1:] {
			r.Add(&r, v)
		}
		return &r
	} else {
		var r big.Float = *floatVals[0]
		for _, v := range floatVals[1:] {
			r.Add(&r, v)
		}
		return &r
	}
}

func primSub(_ *Scheme, args []Val) Val {
	if len(args) == 1 {
		switch v := args[0].(type) {
		case *big.Int:
			var r big.Int
			r.Neg(v)
			return &r
		case *big.Float:
			var r big.Float
			r.Neg(v)
			return &r
		default:
			panic("Bad operand to unary '-'")
		}
	}
	intVals, floatVals := checkAndCoerceNumbers(args, "'-'")
	if intVals != nil {
		var r big.Int = *intVals[0]
		for _, v := range intVals[1:] {
			r.Sub(&r, v)
		}
		return &r
	} else {
		var r big.Float = *floatVals[0]
		for _, v := range floatVals[1:] {
			r.Sub(&r, v)
		}
		return &r
	}
}

func primLess(_ *Scheme, args []Val) Val {
	return primCompare(args, "'<'", -1)
}

func primEqual(_ *Scheme, args []Val) Val {
	return primCompare(args, "'='", 0)
}

func primCompare(args []Val, name string, expected int) Val {
	intVals, floatVals := checkAndCoerceNumbers(args, name)
	r := true
	if intVals != nil {
		prev := intVals[0]
		for _, v := range intVals[1:] {
			if prev.Cmp(v) != expected {
				r = false
				break
			}
		}
	} else {
		prev := floatVals[0]
		for _, v := range floatVals[1:] {
			if prev.Cmp(v) != expected {
				r = false
				break
			}
		}
	}
	return &r
}

func checkAndCoerceNumbers(vals []Val, irritant string) ([]*big.Int, []*big.Float) {
	// FIXME
}

func assertNumeric(v Val, s string) {
	if _, ok := v.(*big.Int); ok {
		return
	}
	if _, ok := v.(*big.Float); ok {
		return
	}
	panic("Bad operand to " + s)
}
