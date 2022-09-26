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
	addPrimitive(ctx, "procedure-name", 1, false, primProcedureName)
	addPrimitive(ctx, "procedure-arity", 1, false, primProcedureArity)
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

func primProcedurep(ctx *Scheme, a0, _ Val, _ []Val) (Val, int) {
	if _, ok := a0.(*Procedure); ok {
		return ctx.TrueVal, 1
	}
	return ctx.FalseVal, 1
}

func primProcedureName(ctx *Scheme, a0, _ Val, _ []Val) (Val, int) {
	if proc, ok := a0.(*Procedure); ok {
		return &Str{Value: proc.Lam.Name}, 1
	}
	return ctx.Error("procedure-name: Not a procedure", a0)
}

func primProcedureArity(ctx *Scheme, a0, _ Val, _ []Val) (Val, int) {
	if proc, ok := a0.(*Procedure); ok {
		if proc.Lam.Rest {
			return big.NewFloat(float64(proc.Lam.Fixed)), 1
		}
		return big.NewInt(int64(proc.Lam.Fixed)), 1
	}
	return ctx.Error("procedure-arity: Not a procedure", a0)
}

func primStringMap(ctx *Scheme, a0, a1 Val, rest []Val) (Val, int) {
	if len(rest) > 0 {
		return ctx.Error("string-map: Only supported for one string for now")
	}
	p, pOk := a0.(*Procedure)
	if !pOk {
		return ctx.Error("string-map: Not a procedure", a0)
	}
	s, sOk := a1.(*Str)
	if !sOk {
		return ctx.Error("string-map: Not a string", a1)
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
			return ctx.Error("string-map: not a character", nch)
		}
		result = result + string(nch.Value)
	}
	return &Str{Value: result}, 1
}

func primStringForEach(ctx *Scheme, a0, a1 Val, rest []Val) (Val, int) {
	if len(rest) > 2 {
		return ctx.Error("string-for-each: Only supported for one string for now")
	}
	p, pOk := a0.(*Procedure)
	if !pOk {
		return ctx.Error("string-for-each: Not a procedure: ", a0)
	}
	s, sOk := a1.(*Str)
	if !sOk {
		return ctx.Error("string-for-each: Not a string: ", a1)
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
func primValues(ctx *Scheme, a0, a1 Val, rest []Val) (Val, int) {
	if a0 == ctx.UndefinedVal {
		return ctx.UnspecifiedVal, 0
	}
	if a1 == ctx.UndefinedVal {
		return a0, 1
	}
	// This is nuts, we should reuse whatever array is there
	ctx.MultiVals = make([]Val, len(rest)+1)
	ctx.MultiVals[0] = a1
	copy(ctx.MultiVals[1:], rest)
	return a0, len(rest) + 2
}

func primReceiveValues(ctx *Scheme, a0, _ Val, rest []Val) (Val, int) {
	results, unw := ctx.Invoke(a0, []Val{})
	if unw != nil {
		return unw, EvalUnwind
	}
	l := ctx.NullVal
	for i := len(results) - 1; i >= 0; i-- {
		l = &Cons{Car: results[i], Cdr: l}
	}
	return l, 1
}

func primUnspecified(ctx *Scheme, _, _ Val, rest []Val) (Val, int) {
	return ctx.UnspecifiedVal, 1
}

func primNewTlsKey(ctx *Scheme, _, _ Val, rest []Val) (Val, int) {
	return big.NewInt(int64(ctx.AllocateTlsKey())), 1
}

func primReadTlsValue(ctx *Scheme, a0, _ Val, rest []Val) (Val, int) {
	if iv, ok := a0.(*big.Int); ok {
		if iv.IsInt64() {
			n := iv.Int64()
			if n >= 0 && n <= math.MaxInt32 {
				return ctx.GetTlsValue(int32(n)), 1
			}
		}
		return ctx.UnspecifiedVal, 1
	}
	return ctx.Error("sint:read-tls-value: key must be exact integer", a0)
}

func primWriteTlsValue(ctx *Scheme, a0, a1 Val, rest []Val) (Val, int) {
	if iv, ok := a0.(*big.Int); ok {
		if iv.IsInt64() {
			n := iv.Int64()
			if n >= 0 && n <= math.MaxInt32 {
				ctx.SetTlsValue(int32(n), a1)
			}
		}
		return ctx.UnspecifiedVal, 1
	}
	return ctx.Error("sint:write-tls-value: key must be exact integer", a0)
}

// The documentation for the unwinding primitives is in control.sch

func primUnwindHandler(ctx *Scheme, a0, a1 Val, rest []Val) (Val, int) {
	a2 := rest[0]
	// (sint:call-with-unwind-handler key thunk handler)
	filterKey := a0
	thunk := a1
	thunkProc, thunkOk := thunk.(*Procedure)
	if !thunkOk || thunkProc.Lam.Fixed != 0 {
		return ctx.Error("sint:unwind-handler: not a thunk", thunk)
	}
	handler := a2
	handlerProc, handlerOk := handler.(*Procedure)
	if !handlerOk || handlerProc.Lam.Fixed != 2 {
		return ctx.Error("sint:unwind-handler: not a handler", thunk)
	}
	return ctx.InvokeWithUnwindHandler(filterKey, thunkProc, handlerProc)
}

func primUnwind(ctx *Scheme, a0, a1 Val, rest []Val) (Val, int) {
	// (sint:unwind key payload)
	return ctx.NewUnwindPackage(a0, a1), EvalUnwind
}

func primCompileToplevel(ctx *Scheme, a0, _ Val, rest []Val) (Val, int) {
	// Compiles args[0] into a lambda and then creates a toplevel procedure
	// from that lambda, and returns the procedure
	// TODO: The compiler is stateless and thread-safe and can be cached on the engine
	comp := compiler.NewCompiler(ctx.Shared)
	prog, err := comp.CompileToplevel(a0)
	if err != nil {
		return ctx.Error(err.Error())
	}
	return &Procedure{Lam: &Lambda{Fixed: 0, Rest: false, Body: prog}, Env: nil, Primop: nil}, 1
}
