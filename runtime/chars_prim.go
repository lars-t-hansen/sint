// Character primitive procedures.
//
// R7RS 6.6, Characters

package runtime

import (
	"math/big"
	. "sint/core"
)

func initCharPrimitives(ctx *Scheme) {
	addPrimitive(ctx, "char?", 1, false, primCharp)
	addPrimitive(ctx, "char->integer", 1, false, primChar2Int)
	addPrimitive(ctx, "integer->char", 1, false, primInt2Char)
	addPrimitive(ctx, "char=?", 2, false, primCharEq)
	addPrimitive(ctx, "char>?", 2, false, primCharGt)
	addPrimitive(ctx, "char>=?", 2, false, primCharGe)
	addPrimitive(ctx, "char<?", 2, false, primCharLt)
	addPrimitive(ctx, "char<=?", 2, false, primCharLe)
}

func primCharp(ctx *Scheme, a0, _ Val, _ []Val) (Val, int) {
	if _, ok := a0.(*Char); ok {
		return ctx.TrueVal, 1
	}
	return ctx.FalseVal, 1
}

func primChar2Int(ctx *Scheme, a0, _ Val, _ []Val) (Val, int) {
	ch, err := ctx.CheckChar(a0, "char->integer")
	if err != nil {
		return ctx.SignalWrappedError(err)
	}
	return big.NewInt(int64(ch)), 1
}

func primInt2Char(ctx *Scheme, a0, _ Val, _ []Val) (Val, int) {
	// TODO: Is this the right range?  I tend to doubt it.
	n, nErr := ctx.CheckExactIntInRange(a0, "char->integer", 0, 0xDFFF)
	if nErr != nil {
		return ctx.SignalWrappedError(nErr)
	}
	return &Char{Value: rune(n)}, 1
}

func primCharEq(ctx *Scheme, a0, a1 Val, _ []Val) (Val, int) {
	c1, c2, err := checkBothChars(ctx, a0, a1, "char=?")
	if err != nil {
		return ctx.SignalWrappedError(err)
	}
	if c1 == c2 {
		return ctx.TrueVal, 1
	}
	return ctx.FalseVal, 1
}

func primCharGt(ctx *Scheme, a0, a1 Val, _ []Val) (Val, int) {
	c1, c2, err := checkBothChars(ctx, a0, a1, "char>?")
	if err != nil {
		return ctx.SignalWrappedError(err)
	}
	if c1 > c2 {
		return ctx.TrueVal, 1
	}
	return ctx.FalseVal, 1
}

func primCharGe(ctx *Scheme, a0, a1 Val, _ []Val) (Val, int) {
	c1, c2, err := checkBothChars(ctx, a0, a1, "char>=?")
	if err != nil {
		return ctx.SignalWrappedError(err)
	}
	if c1 >= c2 {
		return ctx.TrueVal, 1
	}
	return ctx.FalseVal, 1
}

func primCharLt(ctx *Scheme, a0, a1 Val, _ []Val) (Val, int) {
	c1, c2, err := checkBothChars(ctx, a0, a1, "char<?")
	if err != nil {
		return ctx.SignalWrappedError(err)
	}
	if c1 < c2 {
		return ctx.TrueVal, 1
	}
	return ctx.FalseVal, 1
}

func primCharLe(ctx *Scheme, a0, a1 Val, _ []Val) (Val, int) {
	c1, c2, err := checkBothChars(ctx, a0, a1, "char<=?")
	if err != nil {
		return ctx.SignalWrappedError(err)
	}
	if c1 <= c2 {
		return ctx.TrueVal, 1
	}
	return ctx.FalseVal, 1
}

func checkBothChars(ctx *Scheme, v0 Val, v1 Val, name string) (c0 rune, c1 rune, err *WrappedError) {
	c0, err = ctx.CheckChar(v0, name)
	if err == nil {
		c1, err = ctx.CheckChar(v1, name)
	}
	return
}
