package runtime

import (
	. "sint/core"
)

func StandardInitialization(engine *Scheme) (reader ClosableInputStream, writer ClosableFlushableOutputStream,
	errWriter ClosableFlushableOutputStream) {
	initPrimitives(engine)
	initCompiledCode(engine)
	reader = NewStdinReader()
	engine.SetTlsValue(CurrentInputPort, NewInputPort(reader, true /* isText */, "<standard input>"))
	writer = NewStdoutWriter()
	engine.SetTlsValue(CurrentOutputPort, NewOutputPort(writer, true /* isText */, "<standard output>"))
	errWriter = NewStderrWriter()
	engine.SetTlsValue(CurrentErrorPort, NewOutputPort(errWriter, true /* isText */, "<standard error>"))
	return
}

// Code compiled from Scheme to Go is initialized here.  An alternative would be to just emit an init()
// function in each file; that might not work for the builtins but it might work for other things - not
// sure yet.

func initCompiledCode(c *Scheme) {
	// Fundamental stuff.  These should not reference each other during
	// initialization and can be in alpha order.
	initBooleans(c)
	initControl(c)
	initEquivalence(c)
	initExceptions(c)
	initNumbers(c)
	initPairs(c)
	initStrings(c)
	initSymbols(c)
	initSystem(c)

	// Higher-level stuff.  These can reference definitions from the previous set
	// during initialization.
	initIo(c)

	// Library stuff
	initGenerator(c)
	initSort(c)
}

func initPrimitives(ctx *Scheme) {
	initEquivalencePrimitives(ctx)
	initNumbersPrimitives(ctx)
	initPairPrimitives(ctx)
	initSymbolPrimitives(ctx)
	initCharPrimitives(ctx)
	initStringPrimitives(ctx)
	initControlPrimitives(ctx)
	initExceptionsPrimitives(ctx)
	initIoPrimitives(ctx)
	initConcurrencyPrimitives(ctx)
}

func addPrimitive(ctx *Scheme, name string, fixed int, rest bool, primop func(*Scheme, []Val) (Val, int)) {
	sym := ctx.Intern(name)
	sym.Value = &Procedure{Lam: &Lambda{Fixed: fixed, Rest: rest, Body: nil}, Env: nil, Primop: primop}
}
