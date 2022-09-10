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

func initNumbersPrimitives(ctx *Scheme) {
	addPrimitive(ctx, "sint:inexact-float?", 1, false, primInexactFloatp)
	addPrimitive(ctx, "sint:exact-integer?", 1, false, primExactIntegerp)
	addPrimitive(ctx, "finite?", 1, false, primFinitep)
	addPrimitive(ctx, "infinite?", 1, false, primInfinitep)
	addPrimitive(ctx, "+", 0, true, primAdd)
	addPrimitive(ctx, "-", 1, true, primSub)
	addPrimitive(ctx, "*", 0, true, primMul)
	addPrimitive(ctx, "/", 1, true, primDiv)
	addPrimitive(ctx, "<", 2, true, primLess)
	addPrimitive(ctx, "<=", 2, true, primLessOrEqual)
	addPrimitive(ctx, "=", 2, true, primEqual)
	addPrimitive(ctx, ">", 2, true, primGreater)
	addPrimitive(ctx, ">=", 2, true, primGreaterOrEqual)
	addPrimitive(ctx, "number->string", 1, false, primNumber2String)
	addPrimitive(ctx, "string->number", 1, true, primString2Number)
	addPrimitive(ctx, "inexact", 1, false, primInexact)
	addPrimitive(ctx, "exact", 1, false, primExact)
	addPrimitive(ctx, "abs", 1, false, primAbs)
	addPrimitive(ctx, "floor", 1, false, primFloor)
	addPrimitive(ctx, "ceiling", 1, false, primCeiling)
	addPrimitive(ctx, "truncate", 1, false, primTruncate)
	addPrimitive(ctx, "round", 1, false, primRound)
	addPrimitive(ctx, "bitwise-and", 0, true, primBitwiseAnd)
	addPrimitive(ctx, "bitwise-or", 0, true, primBitwiseOr)
	addPrimitive(ctx, "bitwise-xor", 0, true, primBitwiseXor)
	addPrimitive(ctx, "bitwise-and-not", 2, false, primBitwiseAndNot)
	addPrimitive(ctx, "bitwise-not", 1, false, primBitwiseNot)
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

func primFinitep(ctx *Scheme, args []Val) (Val, int) {
	v := args[0]
	if _, ok := v.(*big.Int); ok {
		return ctx.TrueVal, 1
	}
	if fv, ok := v.(*big.Float); ok {
		if fv.IsInf() {
			return ctx.FalseVal, 1
		}
		return ctx.TrueVal, 1
	}
	return ctx.Error("finite?: not a number: " + v.String())
}

func primInfinitep(ctx *Scheme, args []Val) (Val, int) {
	v := args[0]
	if _, ok := v.(*big.Int); ok {
		return ctx.TrueVal, 1
	}
	if fv, ok := v.(*big.Float); ok {
		if fv.IsInf() {
			return ctx.TrueVal, 1
		}
		return ctx.FalseVal, 1
	}
	return ctx.Error("infinite?: not a number: " + v.String())
}

func primAdd(ctx *Scheme, args []Val) (Val, int) {
	if len(args) == 0 {
		return ctx.Zero, 1
	}
	if len(args) == 1 {
		return checkNumber(ctx, args[0], "+")
	}
	r, nres := add2(ctx, args[0], args[1])
	if nres == EvalUnwind {
		return r, nres
	}
	for _, v := range args[2:] {
		r, nres = add2(ctx, r, v)
		if nres == EvalUnwind {
			return r, nres
		}
	}
	return r, 1
}

func primSub(ctx *Scheme, args []Val) (Val, int) {
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
			return ctx.Error("'-': Not a number: " + args[0].String())
		}
	}
	r, nres := sub2(ctx, args[0], args[1])
	if nres == EvalUnwind {
		return r, nres
	}
	for _, v := range args[2:] {
		r, nres = sub2(ctx, r, v)
		if nres == EvalUnwind {
			return r, nres
		}
	}
	return r, 1
}

func primMul(ctx *Scheme, args []Val) (Val, int) {
	if len(args) == 0 {
		return big.NewInt(1), 1
	}
	if len(args) == 1 {
		return checkNumber(ctx, args[0], "*")
	}
	r, nres := mul2(ctx, args[0], args[1])
	if nres == EvalUnwind {
		return r, nres
	}
	for _, v := range args[2:] {
		r, nres = mul2(ctx, r, v)
		if nres == EvalUnwind {
			return r, nres
		}
	}
	return r, 1
}

func primDiv(ctx *Scheme, args []Val) (Val, int) {
	if len(args) == 1 {
		var fv big.Float
		switch v := args[0].(type) {
		case *big.Int:
			fv.SetInt(v)
		case *big.Float:
			fv = *v
		default:
			return ctx.Error("'-': Not a number: " + args[0].String())
		}
		return div2(ctx, big.NewFloat(1), &fv)
	}
	r, nres := div2(ctx, args[0], args[1])
	if nres == EvalUnwind {
		return r, nres
	}
	for _, v := range args[2:] {
		r, nres = div2(ctx, r, v)
		if nres == EvalUnwind {
			return r, nres
		}
	}
	return r, 1
}

func primLess(ctx *Scheme, args []Val) (Val, int) {
	for i := 1; i < len(args); i++ {
		res, err := cmp2(ctx, args[i-1], args[i], "<")
		if err != nil {
			return ctx.SignalWrappedError(err)
		}
		if res != -1 {
			return ctx.FalseVal, 1
		}
	}
	return ctx.TrueVal, 1
}

func primLessOrEqual(ctx *Scheme, args []Val) (Val, int) {
	for i := 1; i < len(args); i++ {
		res, err := cmp2(ctx, args[i-1], args[i], "<=")
		if err != nil {
			return ctx.SignalWrappedError(err)
		}
		if res == 1 {
			return ctx.FalseVal, 1
		}
	}
	return ctx.TrueVal, 1
}

func primEqual(ctx *Scheme, args []Val) (Val, int) {
	for i := 1; i < len(args); i++ {
		res, err := cmp2(ctx, args[i-1], args[i], "=")
		if err != nil {
			return ctx.SignalWrappedError(err)
		}
		if res != 0 {
			return ctx.FalseVal, 1
		}
	}
	return ctx.TrueVal, 1
}

func primGreater(ctx *Scheme, args []Val) (Val, int) {
	for i := 1; i < len(args); i++ {
		res, err := cmp2(ctx, args[i-1], args[i], ">")
		if err != nil {
			return ctx.SignalWrappedError(err)
		}
		if res != 1 {
			return ctx.FalseVal, 1
		}
	}
	return ctx.TrueVal, 1
}

func primGreaterOrEqual(ctx *Scheme, args []Val) (Val, int) {
	for i := 1; i < len(args); i++ {
		res, err := cmp2(ctx, args[i-1], args[i], ">=")
		if err != nil {
			return ctx.SignalWrappedError(err)
		}
		if res == -1 {
			return ctx.FalseVal, 1
		}
	}
	return ctx.TrueVal, 1
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
	return ctx.Error("number->string: Not a number: " + v.String())
}

func primString2Number(ctx *Scheme, args []Val) (Val, int) {
	var s *Str
	if v, ok := args[0].(*Str); ok {
		s = v
	} else {
		return ctx.Error("string->number: bad string: " + args[0].String())
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
		return ctx.Error("string->number: bad radix: " + strconv.Itoa(radix))
	}
	var iv big.Int
	if _, ok := iv.SetString(s.Value, radix); ok {
		return &iv, 1
	}
	// This is too aggressive.  The string could have been rejected because
	// there are bad digits for the base.
	if radix != 10 {
		return ctx.Error("string->number: bad radix for inexact number: " + strconv.Itoa(radix))
	}
	var fv big.Float
	if _, ok := fv.SetString(s.Value); ok {
		return &fv, 1
	}
	return ctx.Error("string->number: bad number: " + s.Value)
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
	return ctx.Error("inexact: Not a number: " + v.String())
}

func primExact(ctx *Scheme, args []Val) (Val, int) {
	v := args[0]
	if _, ok := v.(*big.Int); ok {
		return v, 1
	}
	if fv, ok := v.(*big.Float); ok {
		iv, _ := fv.Int(nil)
		if iv == nil {
			return ctx.Error("exact: Infinity can't be converted to exact: " + v.String())
		}
		return iv, 1
	}
	return ctx.Error("exact: Not a number: " + v.String())
}

func primAbs(ctx *Scheme, args []Val) (Val, int) {
	v := args[0]
	if iv, ok := v.(*big.Int); ok {
		var r *big.Int
		return r.Abs(iv), 1
	}
	if fv, ok := v.(*big.Float); ok {
		var r *big.Float
		return r.Abs(fv), 1
	}
	return ctx.Error("abs: Not a number: " + v.String())
}

const (
	TowardNegativeInfinity = iota
	TowardPositiveInfinity
	None
	Round
)

func primFloor(ctx *Scheme, args []Val) (Val, int) {
	return roundToInteger(ctx, args[0], "floor", TowardNegativeInfinity)
}

func primCeiling(ctx *Scheme, args []Val) (Val, int) {
	return roundToInteger(ctx, args[0], "ceiling", TowardPositiveInfinity)
}

func primTruncate(ctx *Scheme, args []Val) (Val, int) {
	return roundToInteger(ctx, args[0], "truncate", None)
}

func primRound(ctx *Scheme, args []Val) (Val, int) {
	return roundToInteger(ctx, args[0], "round", Round)
}

func roundToInteger(ctx *Scheme, v Val, name string, adjust int) (Val, int) {
	if _, ok := v.(*big.Int); ok {
		return v, 1
	}
	if fv, ok := v.(*big.Float); ok {
		iv, acc := fv.Int(nil)
		if iv == nil {
			return fv, 1
		}
		if adjust == TowardNegativeInfinity && acc == big.Above {
			// TODO: Cache the "1"
			iv.Sub(iv, big.NewInt(1))
		} else if adjust == TowardPositiveInfinity && acc == big.Below {
			// TODO: Cache the "1"
			iv.Add(iv, big.NewInt(1))
		} else if adjust == Round && acc != big.Exact {
			// FIXME: need to look at the difference between the rounded value and
			// the original value, and round to even
		}
		return iv, 1
	}
	return ctx.Error(name + ": Not a number: " + v.String())
}

func primBitwiseAnd(ctx *Scheme, args []Val) (Val, int) {
	if len(args) == 0 {
		return ctx.Zero, 1
	}
	if len(args) == 1 {
		return checkNumber(ctx, args[0], "bitwise-and")
	}
	r, nres := and2(ctx, args[0], args[1])
	if nres == EvalUnwind {
		return r, nres
	}
	for _, v := range args[2:] {
		r, nres = and2(ctx, r, v)
		if nres == EvalUnwind {
			return r, nres
		}
	}
	return r, 1
}

func primBitwiseOr(ctx *Scheme, args []Val) (Val, int) {
	if len(args) == 0 {
		return ctx.Zero, 1
	}
	if len(args) == 1 {
		return checkNumber(ctx, args[0], "bitwise-or")
	}
	r, nres := or2(ctx, args[0], args[1])
	if nres == EvalUnwind {
		return r, nres
	}
	for _, v := range args[2:] {
		r, nres = or2(ctx, r, v)
		if nres == EvalUnwind {
			return r, nres
		}
	}
	return r, 1
}

func primBitwiseXor(ctx *Scheme, args []Val) (Val, int) {
	if len(args) == 0 {
		return ctx.Zero, 1
	}
	if len(args) == 1 {
		return checkNumber(ctx, args[0], "bitwise-xor")
	}
	r, nres := xor2(ctx, args[0], args[1])
	if nres == EvalUnwind {
		return r, nres
	}
	for _, v := range args[2:] {
		r, nres = xor2(ctx, r, v)
		if nres == EvalUnwind {
			return r, nres
		}
	}
	return r, 1
}

func primBitwiseAndNot(ctx *Scheme, args []Val) (Val, int) {
	a := args[0]
	b := args[1]
	if ia, ib, ok := bothInt(a, b); ok {
		var z big.Int
		z.AndNot(ia, ib)
		return &z, 1
	}
	return ctx.Error("bitwise-and-not: numbers must be exact integers: " + a.String() + " " + b.String())
}

func primBitwiseNot(ctx *Scheme, args []Val) (Val, int) {
	a := args[0]
	if ia, ok := a.(*big.Int); ok {
		var z big.Int
		z.Not(ia)
		return &z, 1
	}
	return ctx.Error("bitwise-not: not an exact integer: " + a.String())
}

func add2(ctx *Scheme, a Val, b Val) (Val, int) {
	if ia, ib, ok := bothInt(a, b); ok {
		var z big.Int
		z.Add(ia, ib)
		return &z, 1
	}
	fa, fb, err := bothFloat(ctx, a, b, "+")
	if err != nil {
		return ctx.SignalWrappedError(err)
	}
	var z big.Float
	z.Add(fa, fb)
	return &z, 1
}

func sub2(ctx *Scheme, a Val, b Val) (Val, int) {
	if ia, ib, ok := bothInt(a, b); ok {
		var z big.Int
		z.Sub(ia, ib)
		return &z, 1
	}
	fa, fb, err := bothFloat(ctx, a, b, "+")
	if err != nil {
		return ctx.SignalWrappedError(err)
	}
	var z big.Float
	z.Sub(fa, fb)
	return &z, 1
}

func mul2(ctx *Scheme, a Val, b Val) (Val, int) {
	if ia, ib, ok := bothInt(a, b); ok {
		var z big.Int
		z.Mul(ia, ib)
		return &z, 1
	}
	fa, fb, err := bothFloat(ctx, a, b, "*")
	if err != nil {
		return ctx.SignalWrappedError(err)
	}
	var z big.Float
	z.Mul(fa, fb)
	return &z, 1
}

var fzero *big.Float = big.NewFloat(0)

func div2(ctx *Scheme, a Val, b Val) (Val, int) {
	fa, fb, err := bothFloat(ctx, a, b, "/")
	if err != nil {
		return ctx.SignalWrappedError(err)
	}
	if (fa.IsInf() && fb.IsInf()) || (fa.Cmp(fzero) == 0 && fb.Cmp(fzero) == 0) {
		return ctx.Error("/: Result is not a number: (/" + fa.String() + " " + fb.String() + ")")
	}
	var z big.Float
	z.Quo(fa, fb)
	return &z, 1
}

func cmp2(ctx *Scheme, a Val, b Val, name string) (int, *WrappedError) {
	if ia, ib, ok := bothInt(a, b); ok {
		return ia.Cmp(ib), nil
	}
	fa, fb, err := bothFloat(ctx, a, b, name)
	if err != nil {
		return 0, err
	}
	return fa.Cmp(fb), nil
}

func and2(ctx *Scheme, a Val, b Val) (Val, int) {
	if ia, ib, ok := bothInt(a, b); ok {
		var z big.Int
		z.And(ia, ib)
		return &z, 1
	}
	return ctx.Error("bitwise-and: numbers must be exact integers: " + a.String() + " " + b.String())
}

func or2(ctx *Scheme, a Val, b Val) (Val, int) {
	if ia, ib, ok := bothInt(a, b); ok {
		var z big.Int
		z.Or(ia, ib)
		return &z, 1
	}
	return ctx.Error("bitwise-or: numbers must be exact integers: " + a.String() + " " + b.String())
}

func xor2(ctx *Scheme, a Val, b Val) (Val, int) {
	if ia, ib, ok := bothInt(a, b); ok {
		var z big.Int
		z.Xor(ia, ib)
		return &z, 1
	}
	return ctx.Error("bitwise-xor: numbers must be exact integers: " + a.String() + " " + b.String())
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
func bothFloat(ctx *Scheme, a Val, b Val, name string) (*big.Float, *big.Float, *WrappedError) {
	if fa, ok := a.(*big.Float); ok {
		if fb, ok := b.(*big.Float); ok {
			return fa, fb, nil
		}
		if ib, ok := b.(*big.Int); ok {
			var fb big.Float
			fb.SetInt(ib)
			return fa, &fb, nil
		}
		return nil, nil, ctx.WrapError("'" + name + "': Not a number: " + b.String())
	}
	if ia, ok := a.(*big.Int); ok {
		var fa big.Float
		fa.SetInt(ia)
		if fb, ok := b.(*big.Float); ok {
			return &fa, fb, nil
		}
		if ib, ok := b.(*big.Int); ok {
			var fb big.Float
			fb.SetInt(ib)
			return &fa, &fb, nil
		}
		return nil, nil, ctx.WrapError("'" + name + "': Not a number: " + b.String())
	}
	return nil, nil, ctx.WrapError("'" + name + "': Not a number: " + a.String())
}

func checkNumber(ctx *Scheme, v Val, s string) (Val, int) {
	if !isNumber(v) {
		return ctx.Error(s + ": Not a number: " + v.String())
	}
	return v, 1
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
