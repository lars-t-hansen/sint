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

func primAdd(c *Scheme, args []Val) Val {
	if len(args) == 0 {
		return c.zero
	}
	if len(args) == 1 {
		return checkNumber(args[0], "+")
	}
	r := add2(args[0], args[1])
	for v := range args[2:] {
		r = add2(r, v)
	}
	return r
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
			panic("'-': Not a number")
		}
	}
	r := sub2(args[0], args[1])
	for v := range args[2:] {
		r = sub2(r, v)
	}
	return r
}

func primLess(c *Scheme, args []Val) Val {
	for i := 1; i < len(args); i++ {
		if cmp2(args[i-1], args[i], "<") != -1 {
			return c.falseVal
		}
	}
	return c.trueVal
}

func primEqual(c *Scheme, args []Val) Val {
	for i := 1; i < len(args); i++ {
		if cmp2(args[i-1], args[i], "=") != 0 {
			return c.falseVal
		}
	}
	return c.trueVal
}

func cmp2(a Val, b Val, name string) int {
	if ia, ib, ok := bothInt(a, b); ok {
		return ia.Cmp(ib)
	}
	fa, fb := bothFloat(a, b, name)
	return fa.Cmp(fb)
}

func add2(a Val, b Val) Val {
	if ia, ib, ok := bothInt(a, b); ok {
		var z big.Int
		z.Add(ia, ib)
		return z
	}
	fa, fb := bothFloat(a, b, "+")
	var z big.Float
	z.Add(fa, fb)
	return z
}

func sub2(a Val, b Val) Val {
	if ia, ib, ok := bothInt(a, b); ok {
		var z big.Int
		z.Sub(ia, ib)
		return z
	}
	fa, fb := bothFloat(a, b, "+")
	var z big.Float
	z.Sub(fa, fb)
	return z
}

func bothInt(a Val, b Val) (*big.Int, *big.Int, bool) {
	if ia, ok := a.(*big.Int); ok {
		if ib, ok := a.(*big.Int); ok {
			return ia, ib, true
		}
	}
	return nil, nil, false
}

// Coerce both values to float and return them
func bothFloat(a Val, b Val, name string) (*big.Float, *big.Float) {
	if fa, ok := a.(*big.Float); ok {
		if fb, ok := b.(*big.Float); ok {
			return fa, fb
		}
		if ib, ok := b.(*big.Int); ok {
			var fb big.Float
			fb.SetInt(ib)
			return fa, &fb
		}
		panic("'" + name + "': Not a number") // b
	}
	if ia, ok := a.(*big.Int); ok {
		var fa big.Float
		fa.SetInt(ia)
		if fb, ok := b.(*big.Float); ok {
			return &fa, fb
		}
		if ib, ok := b.(*big.Int); ok {
			var fb big.Float
			fb.SetInt(ib)
			return &fa, &fb
		}
		panic("'" + name + "': Not a number") // b
	}
	panic("'" + name + "': Not a number") // a
}

func checkNumber(v Val, s string) Val {
	if _, ok := v.(*big.Int); ok {
		return v
	}
	if _, ok := v.(*big.Float); ok {
		return v
	}
	panic("Bad operand to " + s)
}
