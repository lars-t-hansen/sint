// Control features primitive procedures.
//
// R7RS 6.10, Control features, also see control.sch

package runtime

import (
	. "sint/core"
)

func initControlPrimitives(c *Scheme) {
	addPrimitive(c, "procedure?", 1, false, primProcedurep)
	addPrimitive(c, "string-map", 2, true, primStringMap)
	addPrimitive(c, "values", 0, true, primValues)
	// call-with-values is tricky, it needs to be properly tail-recursive.
	// basically, it is like apply.  so we need an apply-like primitive for
	// it, or we can implement it in terms of sint:apply?
}

func primProcedurep(ctx *Scheme, args []Val) (Val, int) {
	if _, ok := args[0].(*Symbol); ok {
		return ctx.TrueVal, 1
	}
	return ctx.FalseVal, 1
}

func primStringMap(ctx *Scheme, args []Val) (Val, int) {
	if len(args) > 2 {
		panic("string-map: Only supported for one string for now")
	}
	p, pOk := args[0].(*Procedure)
	if !pOk {
		panic("string-map: Not a procedure: " + args[0].String())
	}
	s, sOk := args[1].(*Str)
	if !sOk {
		panic("string-map: Not a string: " + args[1].String())
	}
	var callArgs [1]Val
	result := ""
	for _, ch := range s.Value {
		callArgs[0] = &Char{Value: ch}
		res := ctx.Invoke(p, callArgs[:])
		nch, ok := res[0].(*Char)
		if !ok {
			panic("string-map: not a character: " + nch.String())
		}
		result = result + string(nch.Value)
	}
	return &Str{Value: result}, 1
}

func primValues(ctx *Scheme, args []Val) (Val, int) {
	if len(args) == 0 {
		ctx.MultiVals = []Val{}
		return ctx.UnspecifiedVal, 0
	}
	ctx.MultiVals = args[1:]
	return args[0], len(args)
}
