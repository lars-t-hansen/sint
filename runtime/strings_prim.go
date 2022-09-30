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

func primStringp(ctx *Scheme, a0, _ Val, _ []Val) (Val, int) {
	if _, ok := a0.(*Str); ok {
		return ctx.TrueVal, 1
	}
	return ctx.FalseVal, 1
}

func primStringLength(ctx *Scheme, a0, _ Val, _ []Val) (Val, int) {
	s, sErr := ctx.CheckString(a0, "string-length")
	if sErr != nil {
		return ctx.SignalWrappedError(sErr)
	}
	return big.NewInt(int64(len(s))), 1
}

func primStringRef(ctx *Scheme, a0, a1 Val, _ []Val) (Val, int) {
	s, sErr := ctx.CheckString(a0, "string-ref")
	if sErr != nil {
		return ctx.SignalWrappedError(sErr)
	}
	ix, ixErr := ctx.CheckExactIntInRange(a1, "string-ref", 0, int64(len(s)))
	if ixErr != nil {
		return ctx.SignalWrappedError(ixErr)
	}
	ch, size := utf8.DecodeRuneInString(s[int(ix):])
	if ch == utf8.RuneError {
		// This can happen when indexing into the middle of a char, for example.
		return ctx.Error("string-ref: Invalid code point in string at index", a1)
	}
	ctx.MultiVals = []Val{big.NewInt(int64(size))}
	return &Char{Value: ch}, 2
}

func primStringCompare(ctx *Scheme, a0, a1 Val, _ []Val) (Val, int) {
	s0, sErr := ctx.CheckString(a0, "sint:string-compare")
	if sErr != nil {
		return ctx.SignalWrappedError(sErr)
	}
	s1, sErr := ctx.CheckString(a1, "sint:string-compare")
	if sErr != nil {
		return ctx.SignalWrappedError(sErr)
	}
	return big.NewInt(int64(strings.Compare(s0, s1))), 1
}

func primStringAppend(ctx *Scheme, a0, a1 Val, rest []Val) (Val, int) {
	// TODO: Maybe use a strings.Builder instead?  Depends on the typical number
	// of strings that are appended.
	if a0 == ctx.UndefinedVal {
		return &Str{Value: ""}, 1
	}
	str0, sErr := ctx.CheckString(a0, "string-append")
	if sErr != nil {
		return ctx.SignalWrappedError(sErr)
	}
	if a1 == ctx.UndefinedVal {
		return a0, 1
	}
	s := str0
	str1, sErr := ctx.CheckString(a1, "string-append")
	if sErr != nil {
		return ctx.SignalWrappedError(sErr)
	}
	s = s + str1
	for _, v := range rest {
		str2, sErr := ctx.CheckString(v, "string-append")
		if sErr != nil {
			return ctx.SignalWrappedError(sErr)
		}
		s = s + str2
	}
	return &Str{Value: s}, 1
}

func primSubstring(ctx *Scheme, a0, a1 Val, rest []Val) (Val, int) {
	a2 := rest[0]
	s0, sErr := ctx.CheckString(a0, "substring")
	if sErr != nil {
		return ctx.SignalWrappedError(sErr)
	}
	i1, i1Err := ctx.CheckExactIntInRange(a1, "substring", 0, int64(len(s0)))
	if i1Err != nil {
		return ctx.SignalWrappedError(i1Err)
	}
	i2, i2Err := ctx.CheckExactIntInRange(a2, "substring", 0, int64(len(s0)))
	if i2Err != nil {
		return ctx.SignalWrappedError(i2Err)
	}
	return &Str{Value: s0[i1:i2]}, 1
}

// sint:list->string assumes the list is proper, but it does check that each value
// is a char.
//
// TODO: Again, may be interesting to use a strings.Builder here for efficiency.
// This might be different than for string-append.

func primList2String(ctx *Scheme, a0, _ Val, _ []Val) (Val, int) {
	s := ""
	for {
		if a0 == ctx.NullVal {
			return &Str{Value: s}, 1
		}
		c := a0.(*Cons)
		ch, ok := c.Car.(*Char)
		if !ok {
			return ctx.Error("sint:list->string: not a character", c.Car)
		}
		s = s + string(ch.Value)
		a0 = c.Cdr
	}
}
