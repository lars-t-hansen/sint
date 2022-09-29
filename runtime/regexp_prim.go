package runtime

import (
	"regexp"
	. "sint/core"
)

func initRegexpPrimitives(ctx *Scheme) {
	addPrimitive(ctx, "regexp?", 1, false, primRegexpp)
	addPrimitive(ctx, "string->regexp", 1, false, primStringToRegexp)
	addPrimitive(ctx, "regexp->string", 1, false, primRegexpToString)
	addPrimitive(ctx, "regexp-find-all", 2, false, primRegexpFindAll)
}

func primRegexpp(ctx *Scheme, a0, _ Val, _ []Val) (Val, int) {
	if _, ok := a0.(*Regexp); ok {
		return ctx.TrueVal, 1
	}
	return ctx.FalseVal, 1
}

func primStringToRegexp(ctx *Scheme, a0, _ Val, _ []Val) (Val, int) {
	if s, ok := a0.(*Str); ok {
		if re, err := regexp.Compile(s.Value); err == nil {
			return &Regexp{Value: re}, 1
		}
		return ctx.Error("string->regexp: Invalid regexp syntax", a0)
	}
	return ctx.Error("string->regexp: Not a string", a0)
}

func primRegexpToString(ctx *Scheme, a0, _ Val, _ []Val) (Val, int) {
	if re, ok := a0.(*Regexp); ok {
		return &Str{Value: re.Value.String()}, 1
	}
	return ctx.Error("regexp->string: Not a regular expression", a0)
}

func primRegexpFindAll(ctx *Scheme, a0, a1 Val, _ []Val) (Val, int) {
	if re, ok := a0.(*Regexp); ok {
		if s, ok := a1.(*Str); ok {
			results := re.Value.FindAll([]byte(s.Value), -1)
			l := ctx.NullVal
			for i := len(results) - 1; i >= 0; i-- {
				r := results[i]
				l = &Cons{Car: &Str{Value: string(r)}, Cdr: l}
			}
			return l, 1
		}
		return ctx.Error("regexp-find-all: Not a string", a1)
	}
	return ctx.Error("regexp->string: Not a regular expression", a0)
}
