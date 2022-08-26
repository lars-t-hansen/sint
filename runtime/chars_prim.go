// Character primitive procedures.
//
// R7RS 6.6, Characters
// TODO: char=?
// TODO: char>?
// TODO: char>=?
// TODO: char<?
// TODO: char<=?
// TODO: (and probably many others)

package runtime

import (
	"math/big"
	. "sint/core"
)

func initCharPrimitives(c *Scheme) {
	addPrimitive(c, "char?", 1, false, primCharp)
	addPrimitive(c, "char->integer", 1, false, primChar2Int)
	addPrimitive(c, "integer->char", 1, false, primInt2Char)

}

func primCharp(ctx *Scheme, args []Val) Val {
	if _, ok := args[0].(*Char); ok {
		return ctx.TrueVal
	}
	return ctx.FalseVal
}

func primChar2Int(c *Scheme, args []Val) Val {
	if ch, ok := args[0].(*Char); ok {
		return big.NewInt(int64(ch.Value))
	}
	panic("char->integer: Not a character: " + args[0].String())
}

func primInt2Char(c *Scheme, args []Val) Val {
	if n, ok := args[0].(*big.Int); ok {
		if !n.IsInt64() {
			panic("char->integer: Integer outside character range: " + args[0].String())
		}
		k := n.Int64()
		// TODO: Is this right?
		if k < 0 || k > 0xDFFF {
			panic("char->integer: Integer outside character range: " + args[0].String())
		}
		return &Char{Value: rune(n.Int64())}
	}
	panic("char->integer: Not an exact integer: " + args[0].String())
}
