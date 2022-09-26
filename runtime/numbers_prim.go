// Number primitive procedures.
//
// R7RS 6.2, Numbers, also see numbers.sch

package runtime

import (
	"math/big"
	//lint:ignore ST1001 dot import
	. "sint/core"
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
	addPrimitive(ctx, "quotient", 2, false, primQuotient)
	addPrimitive(ctx, "remainder", 2, false, primRemainder)
	addPrimitive(ctx, "<", 2, true, primLess)
	addPrimitive(ctx, "<=", 2, true, primLessOrEqual)
	addPrimitive(ctx, "=", 2, true, primEqual)
	addPrimitive(ctx, ">", 2, true, primGreater)
	addPrimitive(ctx, ">=", 2, true, primGreaterOrEqual)
	addPrimitive(ctx, "number->string", 1, true, primNumber2String)
	addPrimitive(ctx, "string->number", 1, true, primString2Number)
	addPrimitive(ctx, "inexact", 1, false, primInexact)
	addPrimitive(ctx, "exact", 1, false, primExact)
	addPrimitive(ctx, "abs", 1, false, primAbs)
	addPrimitive(ctx, "sqrt", 1, false, primSqrt)
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

func primInexactFloatp(ctx *Scheme, a0, _ Val, _ []Val) (Val, int) {
	if _, ok := a0.(*big.Float); ok {
		return ctx.TrueVal, 1
	}
	return ctx.FalseVal, 1
}

func primExactIntegerp(ctx *Scheme, a0, _ Val, _ []Val) (Val, int) {
	if _, ok := a0.(*big.Int); ok {
		return ctx.TrueVal, 1
	}
	return ctx.FalseVal, 1
}

func primFinitep(ctx *Scheme, a0, _ Val, _ []Val) (Val, int) {
	if _, ok := a0.(*big.Int); ok {
		return ctx.TrueVal, 1
	}
	if fv, ok := a0.(*big.Float); ok {
		if fv.IsInf() {
			return ctx.FalseVal, 1
		}
		return ctx.TrueVal, 1
	}
	return ctx.Error("finite?: Not a number", a0)
}

func primInfinitep(ctx *Scheme, a0, _ Val, _ []Val) (Val, int) {
	if _, ok := a0.(*big.Int); ok {
		return ctx.FalseVal, 1
	}
	if fv, ok := a0.(*big.Float); ok {
		if fv.IsInf() {
			return ctx.TrueVal, 1
		}
		return ctx.FalseVal, 1
	}
	return ctx.Error("infinite?: Not a number", a0)
}

func primAdd(ctx *Scheme, a0, a1 Val, rest []Val) (Val, int) {
	if a0 == ctx.UndefinedVal {
		return ctx.Zero, 1
	}
	if a1 == ctx.UndefinedVal {
		return checkNumber(ctx, a0, "+")
	}
	r, nres := add2(ctx, a0, a1)
	if nres == EvalUnwind {
		return r, nres
	}
	for _, v := range rest {
		r, nres = add2(ctx, r, v)
		if nres == EvalUnwind {
			return r, nres
		}
	}
	return r, 1
}

func primSub(ctx *Scheme, a0, a1 Val, rest []Val) (Val, int) {
	if a1 == ctx.UndefinedVal {
		switch v := a0.(type) {
		case *big.Int:
			var r big.Int
			r.Neg(v)
			return &r, 1
		case *big.Float:
			var r big.Float
			r.Neg(v)
			return &r, 1
		default:
			return ctx.Error("'-': Not a number", a0)
		}
	}
	r, nres := sub2(ctx, a0, a1)
	if nres == EvalUnwind {
		return r, nres
	}
	for _, v := range rest {
		r, nres = sub2(ctx, r, v)
		if nres == EvalUnwind {
			return r, nres
		}
	}
	return r, 1
}

func primMul(ctx *Scheme, a0, a1 Val, rest []Val) (Val, int) {
	if a0 == ctx.UndefinedVal {
		return big.NewInt(1), 1
	}
	if a1 == ctx.UndefinedVal {
		return checkNumber(ctx, a0, "*")
	}
	r, nres := mul2(ctx, a0, a1)
	if nres == EvalUnwind {
		return r, nres
	}
	for _, v := range rest {
		r, nres = mul2(ctx, r, v)
		if nres == EvalUnwind {
			return r, nres
		}
	}
	return r, 1
}

func primDiv(ctx *Scheme, a0, a1 Val, rest []Val) (Val, int) {
	if a1 == ctx.UndefinedVal {
		var fv big.Float
		switch v := a0.(type) {
		case *big.Int:
			fv.SetInt(v)
		case *big.Float:
			fv = *v
		default:
			return ctx.Error("'-': Not a number", a0)
		}
		return div2(ctx, big.NewFloat(1), &fv)
	}
	r, nres := div2(ctx, a0, a1)
	if nres == EvalUnwind {
		return r, nres
	}
	for _, v := range rest {
		r, nres = div2(ctx, r, v)
		if nres == EvalUnwind {
			return r, nres
		}
	}
	return r, 1
}

func primQuotient(ctx *Scheme, a0, a1 Val, _ []Val) (Val, int) {
	if ia, ib, ok := bothInt(a0, a1); ok {
		var z big.Int
		z.Quo(ia, ib)
		return &z, 1
	}
	return ctx.Error("quotient: numbers must be exact integers", a0, a1)
}

func primRemainder(ctx *Scheme, a0, a1 Val, _ []Val) (Val, int) {
	if ia, ib, ok := bothInt(a0, a1); ok {
		var z big.Int
		z.Rem(ia, ib)
		return &z, 1
	}
	return ctx.Error("remainder: numbers must be exact integers", a0, a1)
}

func primLess(ctx *Scheme, a0, a1 Val, rest []Val) (Val, int) {
	res, err := cmp2(ctx, a0, a1, "<")
	if err != nil {
		return ctx.SignalWrappedError(err)
	}
	if res != -1 {
		return ctx.FalseVal, 1
	}
	prev := a1
	for _, v := range rest {
		res, err := cmp2(ctx, prev, v, "<")
		if err != nil {
			return ctx.SignalWrappedError(err)
		}
		if res != -1 {
			return ctx.FalseVal, 1
		}
		prev = v
	}
	return ctx.TrueVal, 1
}

func primLessOrEqual(ctx *Scheme, a0, a1 Val, rest []Val) (Val, int) {
	res, err := cmp2(ctx, a0, a1, "<=")
	if err != nil {
		return ctx.SignalWrappedError(err)
	}
	if res == 1 {
		return ctx.FalseVal, 1
	}
	prev := a1
	for _, v := range rest {
		res, err := cmp2(ctx, prev, v, "<=")
		if err != nil {
			return ctx.SignalWrappedError(err)
		}
		if res == 1 {
			return ctx.FalseVal, 1
		}
		prev = v
	}
	return ctx.TrueVal, 1
}

func primEqual(ctx *Scheme, a0, a1 Val, rest []Val) (Val, int) {
	res, err := cmp2(ctx, a0, a1, "=")
	if err != nil {
		return ctx.SignalWrappedError(err)
	}
	if res != 0 {
		return ctx.FalseVal, 1
	}
	prev := a1
	for _, v := range rest {
		res, err := cmp2(ctx, prev, v, "=")
		if err != nil {
			return ctx.SignalWrappedError(err)
		}
		if res != 0 {
			return ctx.FalseVal, 1
		}
		prev = v
	}
	return ctx.TrueVal, 1
}

func primGreater(ctx *Scheme, a0, a1 Val, rest []Val) (Val, int) {
	res, err := cmp2(ctx, a0, a1, ">")
	if err != nil {
		return ctx.SignalWrappedError(err)
	}
	if res != 1 {
		return ctx.FalseVal, 1
	}
	prev := a1
	for _, v := range rest {
		res, err := cmp2(ctx, prev, v, ">")
		if err != nil {
			return ctx.SignalWrappedError(err)
		}
		if res != 1 {
			return ctx.FalseVal, 1
		}
		prev = v
	}
	return ctx.TrueVal, 1
}

func primGreaterOrEqual(ctx *Scheme, a0, a1 Val, rest []Val) (Val, int) {
	res, err := cmp2(ctx, a0, a1, ">=")
	if err != nil {
		return ctx.SignalWrappedError(err)
	}
	if res == -1 {
		return ctx.FalseVal, 1
	}
	prev := a1
	for _, v := range rest {
		res, err := cmp2(ctx, prev, v, ">=")
		if err != nil {
			return ctx.SignalWrappedError(err)
		}
		if res == -1 {
			return ctx.FalseVal, 1
		}
		prev = v
	}
	return ctx.TrueVal, 1
}

func parseRadix(v Val) (radix int, radixOk bool) {
	if iv, ok := v.(*big.Int); ok {
		if iv.IsInt64() {
			i := iv.Int64()
			if i == 2 || i == 8 || i == 10 || i == 16 {
				radix = int(i)
				radixOk = true
			}
		}
	}
	return
}

func primNumber2String(ctx *Scheme, a0, a1 Val, _ []Val) (Val, int) {
	radix := 10
	if a1 != ctx.UndefinedVal {
		radixOk := true
		radix, radixOk = parseRadix(a1)
		if !radixOk {
			return ctx.Error("number->string: Bad radix", a1)
		}
	}
	v := NumberToString(a0, radix)
	if v == nil {
		return ctx.Error("number->string: Not a number", a0)
	}
	return v, 1
}

func primString2Number(ctx *Scheme, a0, a1 Val, _ []Val) (Val, int) {
	var s *Str
	if v, ok := a0.(*Str); ok {
		s = v
	} else {
		return ctx.Error("string->number: Not a string", a0)
	}
	radix := -10
	if a1 != ctx.UndefinedVal {
		radixOk := true
		radix, radixOk = parseRadix(a1)
		if !radixOk {
			return ctx.Error("string->number: Bad radix", a1)
		}
	}
	num := StringToNumber(s.Value, radix)
	if num == nil {
		return ctx.Error("string->number: Bad number syntax", s)
	}
	return num, 1
}

func primInexact(ctx *Scheme, a0, _ Val, _ []Val) (Val, int) {
	if iv, ok := a0.(*big.Int); ok {
		var n big.Float
		n.SetInt(iv)
		return &n, 1
	}
	if _, ok := a0.(*big.Float); ok {
		return a0, 1
	}
	return ctx.Error("inexact: Not a number", a0)
}

func primExact(ctx *Scheme, a0, _ Val, _ []Val) (Val, int) {
	if _, ok := a0.(*big.Int); ok {
		return a0, 1
	}
	if fv, ok := a0.(*big.Float); ok {
		iv, _ := fv.Int(nil)
		if iv == nil {
			return ctx.Error("exact: Infinity can't be converted to exact", a0)
		}
		return iv, 1
	}
	return ctx.Error("exact: Not a number", a0)
}

func primAbs(ctx *Scheme, a0, _ Val, _ []Val) (Val, int) {
	if iv, ok := a0.(*big.Int); ok {
		var r big.Int
		return r.Abs(iv), 1
	}
	if fv, ok := a0.(*big.Float); ok {
		var r big.Float
		return r.Abs(fv), 1
	}
	return ctx.Error("abs: Not a number", a0)
}

func primSqrt(ctx *Scheme, a0, _ Val, _ []Val) (Val, int) {
	var r big.Float
	if iv, ok := a0.(*big.Int); ok {
		r.SetInt(iv)
	} else if fv, ok := a0.(*big.Float); ok {
		r = *fv
	} else {
		return ctx.Error("sqrt: Not a number", a0)
	}
	if r.Cmp(big.NewFloat(0.0)) < 0 {
		return ctx.Error("sqrt: Can't take square root of negative number", a0)
	}
	return r.Sqrt(&r), 1
}

const (
	TowardNegativeInfinity = iota
	TowardPositiveInfinity
	None
	Round
)

func primFloor(ctx *Scheme, a0, _ Val, _ []Val) (Val, int) {
	return roundToInteger(ctx, a0, "floor", TowardNegativeInfinity)
}

func primCeiling(ctx *Scheme, a0, a1 Val, _ []Val) (Val, int) {
	return roundToInteger(ctx, a0, "ceiling", TowardPositiveInfinity)
}

func primTruncate(ctx *Scheme, a0, a1 Val, _ []Val) (Val, int) {
	return roundToInteger(ctx, a0, "truncate", None)
}

func primRound(ctx *Scheme, a0, a1 Val, _ []Val) (Val, int) {
	return roundToInteger(ctx, a0, "round", Round)
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
			// FIXME: Issue #2: need to look at the difference between the rounded value and
			// the original value, and round to even
		}
		return iv, 1
	}
	return ctx.Error(name+": Not a number", v)
}

func primBitwiseAnd(ctx *Scheme, a0, a1 Val, rest []Val) (Val, int) {
	if a0 == ctx.UndefinedVal {
		return ctx.Zero, 1
	}
	if a1 == ctx.UndefinedVal {
		return checkNumber(ctx, a0, "bitwise-and")
	}
	r, nres := and2(ctx, a0, a1)
	if nres == EvalUnwind {
		return r, nres
	}
	for _, v := range rest {
		r, nres = and2(ctx, r, v)
		if nres == EvalUnwind {
			return r, nres
		}
	}
	return r, 1
}

func primBitwiseOr(ctx *Scheme, a0, a1 Val, rest []Val) (Val, int) {
	if a0 == ctx.UndefinedVal {
		return ctx.Zero, 1
	}
	if a1 == ctx.UndefinedVal {
		return checkNumber(ctx, a0, "bitwise-or")
	}
	r, nres := or2(ctx, a0, a1)
	if nres == EvalUnwind {
		return r, nres
	}
	for _, v := range rest {
		r, nres = or2(ctx, r, v)
		if nres == EvalUnwind {
			return r, nres
		}
	}
	return r, 1
}

func primBitwiseXor(ctx *Scheme, a0, a1 Val, rest []Val) (Val, int) {
	if a0 == ctx.UndefinedVal {
		return ctx.Zero, 1
	}
	if a1 == ctx.UndefinedVal {
		return checkNumber(ctx, a0, "bitwise-xor")
	}
	r, nres := xor2(ctx, a0, a1)
	if nres == EvalUnwind {
		return r, nres
	}
	for _, v := range rest {
		r, nres = xor2(ctx, r, v)
		if nres == EvalUnwind {
			return r, nres
		}
	}
	return r, 1
}

func primBitwiseAndNot(ctx *Scheme, a0, a1 Val, _ []Val) (Val, int) {
	if ia, ib, ok := bothInt(a0, a1); ok {
		var z big.Int
		z.AndNot(ia, ib)
		return &z, 1
	}
	return ctx.Error("bitwise-and-not: Numbers must be exact integers", a0, a1)
}

func primBitwiseNot(ctx *Scheme, a0, _ Val, _ []Val) (Val, int) {
	if ia, ok := a0.(*big.Int); ok {
		var z big.Int
		z.Not(ia)
		return &z, 1
	}
	return ctx.Error("bitwise-not: Not an exact integer", a0)
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
		return ctx.Error("/: Result would not be a number", a, b)
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
	return ctx.Error("bitwise-and: numbers must be exact integers", a, b)
}

func or2(ctx *Scheme, a Val, b Val) (Val, int) {
	if ia, ib, ok := bothInt(a, b); ok {
		var z big.Int
		z.Or(ia, ib)
		return &z, 1
	}
	return ctx.Error("bitwise-or: numbers must be exact integers", a, b)
}

func xor2(ctx *Scheme, a Val, b Val) (Val, int) {
	if ia, ib, ok := bothInt(a, b); ok {
		var z big.Int
		z.Xor(ia, ib)
		return &z, 1
	}
	return ctx.Error("bitwise-xor: numbers must be exact integers", a, b)
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
		return nil, nil, ctx.WrapError(name+": Not a number", b)
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
		return nil, nil, ctx.WrapError(name+": Not a number", b)
	}
	return nil, nil, ctx.WrapError(name+": Not a number", a)
}

func checkNumber(ctx *Scheme, v Val, s string) (Val, int) {
	if !isNumber(v) {
		return ctx.Error(s+": Not a number", v)
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
