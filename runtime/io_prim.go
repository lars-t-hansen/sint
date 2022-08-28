package runtime

import (
	. "sint/core"
)

func initIoPrimitives(ctx *Scheme) {
	addPrimitive(ctx, "display", 1, true, primDisplay)
	addPrimitive(ctx, "newline", 0, true, primNewline)
	addPrimitive(ctx, "write", 1, true, primWrite)
	addPrimitive(ctx, "writeln", 1, true, primWriteln)
}

func primWrite(ctx *Scheme, args []Val) (Val, int) {
	// TODO: Need to handle the port, but for now always use stdout.
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
