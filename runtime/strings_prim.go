package runtime

import (
	"math/big"
	. "sint/core"
)

// R7RS 6.7, Strings
// TODO: Lots.  See README.md
func initStringPrimitives(c *Scheme) {
	addPrimitive(c, "string?", 1, false, primStringp)
	addPrimitive(c, "string-length", 1, false, primStringLength)
}

func primStringp(ctx *Scheme, args []Val) Val {
	if _, ok := args[0].(*Str); ok {
		return ctx.TrueVal
	}
	return ctx.FalseVal
}

func primStringLength(ctx *Scheme, args []Val) Val {
	v0 := args[0]
	if s, ok := v0.(*Str); ok {
		return big.NewInt(int64(len(s.Value)))
	}
	panic("string-length: Not a string: " + v0.String())
}
