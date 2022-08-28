package runtime

import (
	"sint/compiler"
	. "sint/core"
)

func InitPrimitives(c *Scheme) {
	initEquivalencePrimitives(c)
	initNumbersPrimitives(c)
	initPairPrimitives(c)
	initSymbolPrimitives(c)
	initCharPrimitives(c)
	initStringPrimitives(c)
	initControlPrimitives(c)
	initExceptionsPrimitives(c)
	initIoPrimitives(c)

	// R7RS 6.13, Input and output, also see io.sch
	addPrimitive(c, "eof-object?", 1, false, primEofObjectp)

	// See runtime/control.sch.  This treats its argument as a top-level program form
	// and returns a thunk that evaluates that form.
	addPrimitive(c, "sint:compile-toplevel-phrase", 1, false, primCompileToplevel)

	// See runtime/control.sch.  This is a one-instruction procedure with the signature (fn l count)
	// where the `fn` must be a procedure and `l` must appear to be a list up to at least `count` elements.
	// It applies `fn` to the `count` first elements of `l` in a properly tail-recursive manner.
	// The values are not arguments to the instruction but are taken from the environment, lexical offsets
	// 0, 1, and 2 at relative level 0.
	sym := c.Intern("sint:raw-apply")
	sym.Value = &Procedure{Lam: &Lambda{Fixed: 3, Rest: false, Body: &Apply{}}, Env: nil, Primop: nil}
}

func addPrimitive(c *Scheme, name string, fixed int, rest bool, primop func(*Scheme, []Val) Val) {
	sym := c.Intern(name)
	sym.Value = &Procedure{Lam: &Lambda{Fixed: fixed, Rest: rest, Body: nil}, Env: nil, Primop: primop}
}

func primEofObjectp(ctx *Scheme, args []Val) Val {
	if _, ok := args[0].(*EofObject); ok {
		return ctx.TrueVal
	}
	return ctx.FalseVal
}

func primCompileToplevel(c *Scheme, args []Val) Val {
	// Compiles args[0] into a lambda and then creates a toplevel procedure
	// from that lambda, and returns the procedure
	// TODO: The compiler is stateless and thread-safe and can be cached on the engine
	comp := compiler.NewCompiler(c)
	prog, err := comp.CompileToplevel(args[0])
	if err != nil {
		panic(err.Error())
	}
	return &Procedure{Lam: &Lambda{Fixed: 0, Rest: false, Body: prog}, Env: nil, Primop: nil}
}
