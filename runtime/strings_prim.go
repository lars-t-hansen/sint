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
	addPrimitive(c, "string=?", 2, true, primStringEq)
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

func primStringEq(ctx *Scheme, args []Val) Val {
	s0, ok := args[0].(*Str)
	if !ok {
		panic("string=?: not a string: " + args[0].String())
	}
	// Not sure if we ought to check the types of all the arguments here even
	// if the equality test has already failed.
	for _, v := range args[1:] {
		s1, ok := v.(*Str)
		if !ok {
			panic("string=?: not a string: " + v.String())
		}
		if s0.Value != s1.Value {
			return ctx.FalseVal
		}
	}
	return ctx.TrueVal
}
