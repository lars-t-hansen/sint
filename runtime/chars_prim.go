// Character primitive procedures.
//
// R7RS 6.6, Characters

package runtime

import (
	"math/big"
	. "sint/core"
)

func initCharPrimitives(c *Scheme) {
	addPrimitive(c, "char?", 1, false, primCharp)
	addPrimitive(c, "char->integer", 1, false, primChar2Int)
	addPrimitive(c, "integer->char", 1, false, primInt2Char)
	addPrimitive(c, "char=?", 1, false, primCharEq)
	addPrimitive(c, "char>?", 1, false, primCharGt)
	addPrimitive(c, "char>=?", 1, false, primCharGe)
	addPrimitive(c, "char<?", 1, false, primCharLt)
	addPrimitive(c, "char<=?", 1, false, primCharLe)
}

func primCharp(ctx *Scheme, args []Val) (Val, int) {
	if _, ok := args[0].(*Char); ok {
		return ctx.TrueVal, 1
	}
	return ctx.FalseVal, 1
}

func primChar2Int(ctx *Scheme, args []Val) (Val, int) {
	if ch, ok := args[0].(*Char); ok {
		return big.NewInt(int64(ch.Value)), 1
	}
	return ctx.Error("char->integer: Not a character: " + args[0].String())
}

func primInt2Char(ctx *Scheme, args []Val) (Val, int) {
	if n, ok := args[0].(*big.Int); ok {
		if !n.IsInt64() {
			ctx.Error("char->integer: Integer outside character range: " + args[0].String())
		}
		k := n.Int64()
		// TODO: Is this right?
		if k < 0 || k > 0xDFFF {
			return ctx.Error("char->integer: Integer outside character range: " + args[0].String())
		}
		return &Char{Value: rune(n.Int64())}, 1
	}
	return ctx.Error("char->integer: Not an exact integer: " + args[0].String())
}

func primCharEq(c *Scheme, args []Val) (Val, int) {
	c1, c2, err := checkBothChars(c, args[0], args[1], "char=?")
	if err != nil {
		return err, EvalUnwind
	}
	if c1 == c2 {
		return c.TrueVal, 1
	}
	return c.FalseVal, 1
}

func primCharGt(c *Scheme, args []Val) (Val, int) {
	c1, c2, err := checkBothChars(c, args[0], args[1], "char>?")
	if err != nil {
		return err, EvalUnwind
	}
	if c1 > c2 {
		return c.TrueVal, 1
	}
	return c.FalseVal, 1
}

func primCharGe(c *Scheme, args []Val) (Val, int) {
	c1, c2, err := checkBothChars(c, args[0], args[1], "char>=?")
	if err != nil {
		return err, EvalUnwind
	}
	if c1 >= c2 {
		return c.TrueVal, 1
	}
	return c.FalseVal, 1
}

func primCharLt(c *Scheme, args []Val) (Val, int) {
	c1, c2, err := checkBothChars(c, args[0], args[1], "char<?")
	if err != nil {
		return err, EvalUnwind
	}
	if c1 < c2 {
		return c.TrueVal, 1
	}
	return c.FalseVal, 1
}

func primCharLe(c *Scheme, args []Val) (Val, int) {
	c1, c2, err := checkBothChars(c, args[0], args[1], "char<=?")
	if err != nil {
		return err, EvalUnwind
	}
	if c1 <= c2 {
		return c.TrueVal, 1
	}
	return c.FalseVal, 1
}

func checkBothChars(ctx *Scheme, v0 Val, v1 Val, name string) (rune, rune, Val) {
	c0, ok0 := v0.(*Char)
	if !ok0 {
		return 0, 0, ctx.WrapError(name + ": not a character: " + v0.String())
	}
	c1, ok1 := v1.(*Char)
	if !ok1 {
		return 0, 0, ctx.WrapError(name + ": not a character: " + v1.String())
	}
	return c0.Value, c1.Value, nil
}
