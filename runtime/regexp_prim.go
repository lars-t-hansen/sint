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
	s, sErr := ctx.CheckString(a0, "string->regexp")
	if sErr != nil {
		return ctx.SignalWrappedError(sErr)
	}
	if re, err := regexp.Compile(s); err == nil {
		return &Regexp{Value: re}, 1
	}
	return ctx.Error("string->regexp: Invalid regexp syntax", a0)
}

func primRegexpToString(ctx *Scheme, a0, _ Val, _ []Val) (Val, int) {
	re, reErr := ctx.CheckRegexp(a0, "regexp->string")
	if reErr != nil {
		return ctx.SignalWrappedError(reErr)
	}
	return &Str{Value: re.String()}, 1
}

func primRegexpFindAll(ctx *Scheme, a0, a1 Val, _ []Val) (Val, int) {
	re, reErr := ctx.CheckRegexp(a0, "regexp-find-all")
	if reErr != nil {
		return ctx.SignalWrappedError(reErr)
	}
	s, sErr := ctx.CheckString(a1, "regexp-find-all")
	if sErr != nil {
		return ctx.SignalWrappedError(sErr)
	}
	results := re.FindAll([]byte(s), -1)
	l := ctx.NullVal
	for i := len(results) - 1; i >= 0; i-- {
		r := results[i]
		l = &Cons{Car: &Str{Value: string(r)}, Cdr: l}
	}
	return l, 1
}
