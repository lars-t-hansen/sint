// String primitive procedures.
//
// R7RS 6.7, Strings.  Also see strings.sch.
//
// Note that strings in sint are Go strings: immutable byte arrays
// holding (mostly) UTF8-encoded Unicode.  String lengths and string
// indices are *byte* lengths and indices.

package runtime

import (
	"math/big"
	. "sint/core"
	"strings"
	"unicode/utf8"
)

func initStringPrimitives(ctx *Scheme) {
	addPrimitive(ctx, "string?", 1, false, primStringp)
	addPrimitive(ctx, "string-length", 1, false, primStringLength)
	addPrimitive(ctx, "string-ref", 2, false, primStringRef)
	addPrimitive(ctx, "sint:string-compare", 2, false, primStringCompare)
	addPrimitive(ctx, "string-append", 0, true, primStringAppend)
	addPrimitive(ctx, "substring", 3, false, primSubstring)
	addPrimitive(ctx, "sint:list->string", 1, false, primList2String)
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
	return ctx.Error("string-length: Not a string: " + v0.String())
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
					return ctx.Error("string-ref: Invalid code point in string at index: " + v1.String())
				}
				ctx.MultiVals = []Val{big.NewInt(int64(size))}
				return &Char{Value: ch}, 2
			}
			return ctx.Error("string-ref: Out of range: " + v1.String())
		}
		return ctx.Error("string-ref: Not an exact integer index: " + v1.String())
	}
	return ctx.Error("string-ref: Not a string: " + v0.String())
}

func primStringCompare(ctx *Scheme, args []Val) (Val, int) {
	v0 := args[0]
	v1 := args[1]
	s0, ok0 := v0.(*Str)
	if !ok0 {
		return ctx.Error("sint:string-compare: not a string: " + v0.String())
	}
	s1, ok1 := v1.(*Str)
	if !ok1 {
		return ctx.Error("sint:string-compare: not a string: " + v1.String())
	}
	return big.NewInt(int64(strings.Compare(s0.Value, s1.Value))), 1
}

func primStringAppend(ctx *Scheme, args []Val) (Val, int) {
	// TODO: Maybe use a strings.Builder instead?  Depends on the typical number
	// of strings that are appended.
	s := ""
	for _, v := range args {
		s2, ok := v.(*Str)
		if !ok {
			return ctx.Error("string-append: Not a string: " + v.String())
		}
		s = s + s2.Value
	}
	return &Str{Value: s}, 1
}

func primSubstring(ctx *Scheme, args []Val) (Val, int) {
	v0 := args[0]
	v1 := args[1]
	v2 := args[2]
	s0, ok0 := v0.(*Str)
	if !ok0 {
		return ctx.Error("substring: not a string: " + v0.String())
	}
	i1, i2, ok := bothInt(v1, v2)
	if !ok {
		return ctx.Error("substring: invalid indices: " + v1.String() + " " + v2.String())
	}
	if i1.IsInt64() && i1.Int64() >= 0 && i1.Int64() < int64(len(s0.Value)) &&
		i2.IsInt64() && i2.Int64() >= 0 && i2.Int64() < int64(len(s0.Value)) &&
		i1.Int64() <= i2.Int64() {
		return &Str{Value: s0.Value[i1.Int64():i2.Int64()]}, 1
	} else {
		return ctx.Error("substring: indices out of range: " + v1.String() + " " + v2.String())
	}
}

// sint:list->string assumes the list is proper, but it does check that each value
// is a char.
//
// TODO: Again, may be interesting to use a strings.Builder here for efficiency.
// This might be different than for string-append.

func primList2String(ctx *Scheme, args []Val) (Val, int) {
	v := args[0]
	s := ""
	for {
		if v == ctx.NullVal {
			return &Str{Value: s}, 1
		}
		c := v.(*Cons)
		ch, ok := c.Car.(*Char)
		if !ok {
			return ctx.Error("sint:list->string: not a character: " + c.Car.String())
		}
		s = s + string(ch.Value)
		v = c.Cdr
	}
}
