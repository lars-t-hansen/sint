// R7RS 6.13, Input and output, also see io.sch

package runtime

import (
	. "sint/core"
)

func initIoPrimitives(ctx *Scheme) {
	addPrimitive(ctx, "display", 1, true, primDisplay)
	addPrimitive(ctx, "newline", 0, true, primNewline)
	addPrimitive(ctx, "write", 1, true, primWrite)
	addPrimitive(ctx, "writeln", 1, true, primWriteln)
	addPrimitive(ctx, "eof-object", 0, false, primEof)
	addPrimitive(ctx, "eof-object?", 1, false, primEofObjectp)
}

func primWrite(ctx *Scheme, args []Val) (Val, int) {
	// TODO: Need to handle the port, but for now always use stdout.
	//
	// If the port argument is absent, go to the parameters on the ctx.
	// parameters are special functions that use the parameter key with a
	// primitive to read or write a value
	//
	// a call to a parameter procedure ends up as the primitive (sint:read-parameter <key>)
	// and the parameterize syntax ends up as the primitive (sint:write-parameter <key> <value>)
	// where the value has already been subjected to conversion.
	//
	// parameters are per-thread and are stored in the ctx.
	//
	// To get the current output port here, we can read the parameter directly and use its
	// value.
	//
	// The conversion procedure for the port probably ensures that the value in question is
	// an appropriate port, ie, "conversion" is really also type checking?  Note the port
	// can subsequently be closed, we must deal with that.
	//
	// parameterize also entails unwind-protect.
	//
	// If we're going to use 'panic' for call/cc and general error reporting, then we have
	// a slight problem in that true panics - when the system is in an unstable state -
	// will not allow unwind handlers to run.
	//
	// The alternative is to perform all unwinds in a controlled way and let panics be panics.
	// This could be a
	// return (nil, -1) for example, with the unwind reason stored on the context, or it
	// could be (val, -1) where the value carries the unwind reason.  Mostly only the evaluator
	// needs to worry about this, plus built-in higher-order functions (those that call Invoke)
	// and the repl perhaps.
	//
	// In the end, error handling will just be invoking the error continuation and all unwinds
	// will really be invoking a continuation.  The "val" carried could be a list of values
	// to return from call/cc / apply the continuation to.  In this setup, the Invoke call
	// within sint:unwind-protect would intercept the throw, invoke the unwind handler, and
	// then restart the throw.
	//
	// So, before we do I/O we should do parameters and dynamic-wind and call/cc, probably?

	writer := &StdoutWriter{}
	Write(args[0], false, writer)
	return ctx.UnspecifiedVal, 1
}

func primWriteln(ctx *Scheme, args []Val) (Val, int) {
	// TODO: Need to handle the port, but for now always use stdout.
	writer := &StdoutWriter{}
	Write(args[0], false, writer)
	writer.WriteRune('\n')
	return ctx.UnspecifiedVal, 1
}

func primDisplay(ctx *Scheme, args []Val) (Val, int) {
	// TODO: Need to handle the port, but for now always use stdout.
	writer := &StdoutWriter{}
	Write(args[0], true, writer)
	return ctx.UnspecifiedVal, 1
}

func primNewline(ctx *Scheme, args []Val) (Val, int) {
	// TODO: Need to handle the port, but for now always use stdout.
	writer := &StdoutWriter{}
	writer.WriteRune('\n')
	return ctx.UnspecifiedVal, 1
}

func primEof(ctx *Scheme, args []Val) (Val, int) {
	return ctx.UnspecifiedVal, 1
}

func primEofObjectp(ctx *Scheme, args []Val) (Val, int) {
	if _, ok := args[0].(*EofObject); ok {
		return ctx.TrueVal, 1
	}
	return ctx.FalseVal, 1
}
