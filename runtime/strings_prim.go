// String primitive procedures.
//
// R7RS 6.7, Strings
// TODO: Lots.  See README.md

package runtime

import (
	"math/big"
	. "sint/core"
	"strings"
	"unicode/utf8"
)

func initStringPrimitives(c *Scheme) {
	addPrimitive(c, "string?", 1, false, primStringp)
	addPrimitive(c, "string-length", 1, false, primStringLength)
	addPrimitive(c, "string-ref", 2, false, primStringRef)
	addPrimitive(c, "sint:string-compare", 2, false, primStringCompare)
	addPrimitive(c, "string-append", 0, true, primStringAppend)
}

func primStringp(ctx *Scheme, args []Val) (Val, int) {
	if _, ok := args[0].(*Str); ok {
		return ctx.TrueVal, 1
	}
	return ctx.FalseVal, 1
}

func primStringLength(ctx *Scheme, args []Val) (Val, int) {
	v0 := args[0]
	if s, ok := v0.(*Str); ok {
		return big.NewInt(int64(len(s.Value))), 1
	}
	panic("string-length: Not a string: " + v0.String())
}

func primStringRef(ctx *Scheme, args []Val) (Val, int) {
	v0 := args[0]
	v1 := args[1]
	if s, ok := v0.(*Str); ok {
		if ix, ok := v1.(*big.Int); ok {
			if ix.IsInt64() && ix.Int64() >= 0 && ix.Int64() < int64(len(s.Value)) {
				ch, size := utf8.DecodeRuneInString(s.Value[int(ix.Int64()):])
				if ch == utf8.RuneError {
					// This can happen when indexing into the middle of a char, for example.
					panic("string-ref: Invalid code point in string at index: " + v1.String())
				}
				ctx.MultiVals = []Val{big.NewInt(int64(size))}
				return &Char{Value: ch}, 2
			}
			panic("string-ref: Out of range: " + v1.String())
		}
		panic("string-ref: Not an exact integer index: " + v1.String())
	}
	panic("string-ref: Not a string: " + v0.String())
}

func primStringCompare(ctx *Scheme, args []Val) (Val, int) {
	v0 := args[0]
	v1 := args[1]
	s0, ok0 := v0.(*Str)
	if !ok0 {
		panic("sint:string-compare: not a string: " + v0.String())
	}
	s1, ok1 := v1.(*Str)
	if !ok1 {
		panic("sint:string-compare: not a string: " + v1.String())
	}
	return big.NewInt(int64(strings.Compare(s0.Value, s1.Value))), 1
}

func primStringAppend(ctx *Scheme, args []Val) (Val, int) {
	s := ""
	for _, v := range args {
		s2, ok := v.(*Str)
		if !ok {
			panic("string-append: Not a string: " + v.String())
		}
		s = s + s2.Value
	}
	return &Str{Value: s}, 1
}
