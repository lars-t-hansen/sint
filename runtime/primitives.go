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
	initConcurrencyPrimitives(c)

	// See runtime/control.sch.  This treats its argument as a top-level program form
	// and returns a thunk that evaluates that form.
	addPrimitive(c, "sint:compile-toplevel-phrase", 1, false, primCompileToplevel)
}

func addPrimitive(c *Scheme, name string, fixed int, rest bool, primop func(*Scheme, []Val) (Val, int)) {
	sym := c.Intern(name)
	sym.Value = &Procedure{Lam: &Lambda{Fixed: fixed, Rest: rest, Body: nil}, Env: nil, Primop: primop}
}

func primCompileToplevel(c *Scheme, args []Val) (Val, int) {
	// Compiles args[0] into a lambda and then creates a toplevel procedure
	// from that lambda, and returns the procedure
	// TODO: The compiler is stateless and thread-safe and can be cached on the engine
	comp := compiler.NewCompiler(c.Shared)
	prog, err := comp.CompileToplevel(args[0])
	if err != nil {
		return c.Error(err.Error())
	}
	return &Procedure{Lam: &Lambda{Fixed: 0, Rest: false, Body: prog}, Env: nil, Primop: nil}, 1
}
