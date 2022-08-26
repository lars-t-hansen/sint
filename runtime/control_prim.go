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
}

func primProcedurep(ctx *Scheme, args []Val) Val {
	if _, ok := args[0].(*Symbol); ok {
		return ctx.TrueVal
	}
	return ctx.FalseVal
}

func primStringMap(ctx *Scheme, args []Val) Val {
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
		nch, ok := res.(*Char)
		if !ok {
			panic("string-map: not a character: " + nch.String())
		}
		result = result + string(nch.Value)
	}
	return &Str{Value: result}
}
