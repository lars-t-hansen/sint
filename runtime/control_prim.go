// Control features primitive procedures.
//
// R7RS 6.10, Control features, also see control.sch

package runtime

import (
	"math"
	"math/big"
	"sint/compiler"
	. "sint/core"
)

func initControlPrimitives(ctx *Scheme) {
	addPrimitive(ctx, "procedure?", 1, false, primProcedurep)
	addPrimitive(ctx, "string-map", 2, true, primStringMap)
	addPrimitive(ctx, "string-for-each", 2, true, primStringForEach)
	addPrimitive(ctx, "values", 0, true, primValues)
	addPrimitive(ctx, "unspecified", 0, false, primUnspecified)
	addPrimitive(ctx, "sint:receive-values", 1, false, primReceiveValues)
	addPrimitive(ctx, "sint:new-tls-key", 0, false, primNewTlsKey)
	addPrimitive(ctx, "sint:read-tls-value", 1, false, primReadTlsValue)
	addPrimitive(ctx, "sint:write-tls-value", 2, false, primWriteTlsValue)
	addPrimitive(ctx, "sint:call-with-unwind-handler", 3, false, primUnwindHandler)
	addPrimitive(ctx, "sint:unwind", 2, false, primUnwind)

	// See runtime/control.sch.  This is a procedure with the signature (fn l)
	// where the `fn` must be a procedure and `l` must be a proper list.
	// It applies `fn` to the elements of `l` in a properly tail-recursive manner.
	sym := ctx.Intern("sint:apply")
	sym.Value = &Procedure{
		Lam: &Lambda{
			Fixed: 2,
			Rest:  false,
			Body:  &Apply{Proc: &Lexical{Levels: 0, Offset: 0}, Args: &Lexical{Levels: 0, Offset: 1}}},
		Env:    nil,
		Primop: nil}

	// See runtime/control.sch.  This treats its argument as a top-level program form
	// and returns a thunk that evaluates that form.
	addPrimitive(ctx, "sint:compile-toplevel-phrase", 1, false, primCompileToplevel)
}

func primProcedurep(ctx *Scheme, args []Val) (Val, int) {
	if _, ok := args[0].(*Procedure); ok {
		return ctx.TrueVal, 1
	}
	return ctx.FalseVal, 1
}

func primStringMap(ctx *Scheme, args []Val) (Val, int) {
	if len(args) > 2 {
		return ctx.Error("string-map: Only supported for one string for now")
	}
	p, pOk := args[0].(*Procedure)
	if !pOk {
		return ctx.Error("string-map: Not a procedure: " + args[0].String())
	}
	s, sOk := args[1].(*Str)
	if !sOk {
		return ctx.Error("string-map: Not a string: " + args[1].String())
	}
	var callArgs [1]Val
	result := ""
	for _, ch := range s.Value {
		callArgs[0] = &Char{Value: ch}
		res, unw := ctx.Invoke(p, callArgs[:])
		if unw != nil {
			return unw, EvalUnwind
		}
		nch, ok := res[0].(*Char)
		if !ok {
			return ctx.Error("string-map: not a character: " + nch.String())
		}
		result = result + string(nch.Value)
	}
	return &Str{Value: result}, 1
}

func primStringForEach(ctx *Scheme, args []Val) (Val, int) {
	if len(args) > 2 {
		return ctx.Error("string-for-each: Only supported for one string for now")
	}
	p, pOk := args[0].(*Procedure)
	if !pOk {
		return ctx.Error("string-for-each: Not a procedure: " + args[0].String())
	}
	s, sOk := args[1].(*Str)
	if !sOk {
		return ctx.Error("string-for-each: Not a string: " + args[1].String())
	}
	var callArgs [1]Val
	for _, ch := range s.Value {
		callArgs[0] = &Char{Value: ch}
		_, unw := ctx.Invoke(p, callArgs[:])
		if unw != nil {
			return unw, EvalUnwind
		}
	}
	return ctx.UnspecifiedVal, 1
}
func primValues(ctx *Scheme, args []Val) (Val, int) {
	if len(args) == 0 {
		ctx.MultiVals = []Val{}
		return ctx.UnspecifiedVal, 0
	}
	ctx.MultiVals = args[1:]
	return args[0], len(args)
}

func primReceiveValues(ctx *Scheme, args []Val) (Val, int) {
	results, unw := ctx.Invoke(args[0], []Val{})
	if unw != nil {
		return unw, EvalUnwind
	}
	l := ctx.NullVal
	for i := len(results) - 1; i >= 0; i-- {
		l = &Cons{Car: results[i], Cdr: l}
	}
	return l, 1
}

func primUnspecified(ctx *Scheme, args []Val) (Val, int) {
	return ctx.UnspecifiedVal, 1
}

func primNewTlsKey(ctx *Scheme, args []Val) (Val, int) {
	return big.NewInt(int64(ctx.AllocateTlsKey())), 1
}

func primReadTlsValue(ctx *Scheme, args []Val) (Val, int) {
	v := args[0]
	if iv, ok := v.(*big.Int); ok {
		if iv.IsInt64() {
			n := iv.Int64()
			if n >= 0 && n <= math.MaxInt32 {
				return ctx.GetTlsValue(int32(n)), 1
			}
		}
		return ctx.UnspecifiedVal, 1
	}
	return ctx.Error("sint:read-tls-value: key must be exact integer: " + v.String())
}

func primWriteTlsValue(ctx *Scheme, args []Val) (Val, int) {
	v0 := args[0]
	v1 := args[1]
	if iv, ok := v0.(*big.Int); ok {
		if iv.IsInt64() {
			n := iv.Int64()
			if n >= 0 && n <= math.MaxInt32 {
				ctx.SetTlsValue(int32(n), v1)
			}
		}
		return ctx.UnspecifiedVal, 1
	}
	return ctx.Error("sint:write-tls-value: key must be exact integer: " + v0.String())
}

// The documentation for the unwinding primitives is in control.sch

func primUnwindHandler(ctx *Scheme, args []Val) (Val, int) {
	// (sint:call-with-unwind-handler key thunk handler)
	filterKey := args[0]
	thunk := args[1]
	thunkProc, thunkOk := thunk.(*Procedure)
	if !thunkOk || thunkProc.Lam.Fixed != 0 {
		return ctx.Error("sint:unwind-handler: not a thunk: " + thunk.String())
	}
	handler := args[2]
	handlerProc, handlerOk := handler.(*Procedure)
	if !handlerOk || handlerProc.Lam.Fixed != 2 {
		return ctx.Error("sint:unwind-handler: not a handler: " + thunk.String())
	}
	return ctx.InvokeWithUnwindHandler(filterKey, thunkProc, handlerProc)
}

func primUnwind(ctx *Scheme, args []Val) (Val, int) {
	// (sint:unwind key payload)
	return ctx.NewUnwindPackage(args[0], args[1]), EvalUnwind
}

func primCompileToplevel(ctx *Scheme, args []Val) (Val, int) {
	// Compiles args[0] into a lambda and then creates a toplevel procedure
	// from that lambda, and returns the procedure
	// TODO: The compiler is stateless and thread-safe and can be cached on the engine
	comp := compiler.NewCompiler(ctx.Shared)
	prog, err := comp.CompileToplevel(args[0])
	if err != nil {
		return ctx.Error(err.Error())
	}
	return &Procedure{Lam: &Lambda{Fixed: 0, Rest: false, Body: prog}, Env: nil, Primop: nil}, 1
}
