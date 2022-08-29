// Equivalence primitive procedures.
//
// R7RS 6.2, Numbers, also see numbers.sch
// TODO: /
// TODO: quotient
// TODO: remainder
// TODO: (other numerics as required)

package runtime

import (
	"fmt"
	"math/big"
	. "sint/core"
	"strconv"
)

func initNumbersPrimitives(c *Scheme) {
	addPrimitive(c, "sint:inexact-float?", 1, false, primInexactFloatp)
	addPrimitive(c, "sint:exact-integer?", 1, false, primExactIntegerp)
	addPrimitive(c, "+", 0, true, primAdd)
	addPrimitive(c, "-", 1, true, primSub)
	addPrimitive(c, "*", 0, true, primMul)
	addPrimitive(c, "<", 2, true, primLess)
	addPrimitive(c, "<=", 2, true, primLessOrEqual)
	addPrimitive(c, "=", 2, true, primEqual)
	addPrimitive(c, ">", 2, true, primGreater)
	addPrimitive(c, ">=", 2, true, primGreaterOrEqual)
	addPrimitive(c, "number->string", 1, false, primNumber2String)
	addPrimitive(c, "string->number", 1, true, primString2Number)
	addPrimitive(c, "inexact", 1, false, primInexact)
	addPrimitive(c, "exact", 1, false, primExact)
}

func primInexactFloatp(ctx *Scheme, args []Val) (Val, int) {
	if _, ok := args[0].(*big.Float); ok {
		return ctx.TrueVal, 1
	}
	return ctx.FalseVal, 1
}

func primExactIntegerp(ctx *Scheme, args []Val) (Val, int) {
	if _, ok := args[0].(*big.Int); ok {
		return ctx.TrueVal, 1
	}
	return ctx.FalseVal, 1
}

func primAdd(c *Scheme, args []Val) (Val, int) {
	if len(args) == 0 {
		return c.Zero, 1
	}
	if len(args) == 1 {
		return checkNumber(args[0], "+"), 1
	}
	r := add2(args[0], args[1])
	for _, v := range args[2:] {
		r = add2(r, v)
	}
	return r, 1
}

func primSub(_ *Scheme, args []Val) (Val, int) {
	if len(args) == 1 {
		switch v := args[0].(type) {
		case *big.Int:
			var r big.Int
			r.Neg(v)
			return &r, 1
		case *big.Float:
			var r big.Float
			r.Neg(v)
			return &r, 1
		default:
			panic("'-': Not a number: " + args[0].String())
		}
	}
	r := sub2(args[0], args[1])
	for _, v := range args[2:] {
		r = sub2(r, v)
	}
	return r, 1
}

func primMul(c *Scheme, args []Val) (Val, int) {
	if len(args) == 0 {
		return big.NewInt(1), 1
	}
	if len(args) == 1 {
		return checkNumber(args[0], "*"), 1
	}
	r := mul2(args[0], args[1])
	for _, v := range args[2:] {
		r = mul2(r, v)
	}
	return r, 1
}

func primLess(c *Scheme, args []Val) (Val, int) {
	for i := 1; i < len(args); i++ {
		if cmp2(args[i-1], args[i], "<") != -1 {
			return c.FalseVal, 1
		}
	}
	return c.TrueVal, 1
}

func primLessOrEqual(c *Scheme, args []Val) (Val, int) {
	for i := 1; i < len(args); i++ {
		if cmp2(args[i-1], args[i], "<=") == 1 {
			return c.FalseVal, 1
		}
	}
	return c.TrueVal, 1
}

func primEqual(c *Scheme, args []Val) (Val, int) {
	for i := 1; i < len(args); i++ {
		if cmp2(args[i-1], args[i], "=") != 0 {
			return c.FalseVal, 1
		}
	}
	return c.TrueVal, 1
}

func primGreater(c *Scheme, args []Val) (Val, int) {
	for i := 1; i < len(args); i++ {
		if cmp2(args[i-1], args[i], ">") != 1 {
			return c.FalseVal, 1
		}
	}
	return c.TrueVal, 1
}

func primGreaterOrEqual(c *Scheme, args []Val) (Val, int) {
	for i := 1; i < len(args); i++ {
		if cmp2(args[i-1], args[i], ">=") != -1 {
			return c.FalseVal, 1
		}
	}
	return c.TrueVal, 1
}

func primNumber2String(ctx *Scheme, args []Val) (Val, int) {
	// FIXME: Radix!
	v := args[0]
	if iv, ok := v.(*big.Int); ok {
		return &Str{Value: fmt.Sprint(iv)}, 1
	}
	if fv, ok := v.(*big.Float); ok {
		return &Str{Value: fmt.Sprint(fv)}, 1
	}
	panic("number->string: Not a number: " + v.String())
}

func primString2Number(ctx *Scheme, args []Val) (Val, int) {
	var s *Str
	if v, ok := args[0].(*Str); ok {
		s = v
	} else {
		panic("string->number: bad string: " + args[0].String())
	}
	radix := 10
	if len(args) > 1 {
		if r, ok := args[1].(*big.Int); ok {
			radix = int(r.Int64())
		}
	}
	switch radix {
	case 2, 8, 10, 16:
		break
	default:
		panic("string->number: bad radix: " + strconv.Itoa(radix))
	}
	var iv big.Int
	if _, ok := iv.SetString(s.Value, radix); ok {
		return &iv, 1
	}
	// This is too aggressive.  The string could have been rejected because
	// there are bad digits for the base.
	if radix != 10 {
		panic("string->number: bad radix for inexact number: " + strconv.Itoa(radix))
	}
	var fv big.Float
	if _, ok := fv.SetString(s.Value); ok {
		return &fv, 1
	}
	panic("string->number: bad number: " + s.Value)
}

func primInexact(ctx *Scheme, args []Val) (Val, int) {
	v := args[0]
	if iv, ok := v.(*big.Int); ok {
		var n big.Float
		n.SetInt(iv)
		return &n, 1
	}
	if _, ok := v.(*big.Float); ok {
		return v, 1
	}
	panic("inexact: Not a number: " + v.String())
}

func primExact(ctx *Scheme, args []Val) (Val, int) {
	v := args[0]
	if _, ok := v.(*big.Int); ok {
		return v, 1
	}
	if fv, ok := v.(*big.Float); ok {
		iv, _ := fv.Int(nil)
		if iv == nil {
			panic("exact: Infinity can't be converted to exact: " + v.String())
		}
		return iv, 1
	}
	panic("exact: Not a number: " + v.String())
}

func add2(a Val, b Val) Val {
	if ia, ib, ok := bothInt(a, b); ok {
		var z big.Int
		z.Add(ia, ib)
		return &z
	}
	fa, fb := bothFloat(a, b, "+")
	var z big.Float
	z.Add(fa, fb)
	return &z
}

func sub2(a Val, b Val) Val {
	if ia, ib, ok := bothInt(a, b); ok {
		var z big.Int
		z.Sub(ia, ib)
		return &z
	}
	fa, fb := bothFloat(a, b, "+")
	var z big.Float
	z.Sub(fa, fb)
	return &z
}

func mul2(a Val, b Val) Val {
	if ia, ib, ok := bothInt(a, b); ok {
		var z big.Int
		z.Mul(ia, ib)
		return &z
	}
	fa, fb := bothFloat(a, b, "*")
	var z big.Float
	z.Mul(fa, fb)
	return &z
}

func cmp2(a Val, b Val, name string) int {
	if ia, ib, ok := bothInt(a, b); ok {
		return ia.Cmp(ib)
	}
	fa, fb := bothFloat(a, b, name)
	return fa.Cmp(fb)
}

func bothInt(a Val, b Val) (*big.Int, *big.Int, bool) {
	if ia, ok := a.(*big.Int); ok {
		if ib, ok := b.(*big.Int); ok {
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
		panic("'" + name + "': Not a number: " + b.String())
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
		panic("'" + name + "': Not a number: " + b.String())
	}
	panic("'" + name + "': Not a number: " + a.String())
}

func checkNumber(v Val, s string) Val {
	if !isNumber(v) {
		panic(s + ": Not a number: " + v.String())
	}
	return v
}

func isNumber(v Val) bool {
	if _, ok := v.(*big.Int); ok {
		return true
	}
	if _, ok := v.(*big.Float); ok {
		return true
	}
	return false
}
