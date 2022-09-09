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
	addPrimitive(ctx, "port?", 1, false, primPortp)
	addPrimitive(ctx, "close-port", 1, false, primClosePort)
	addPrimitive(ctx, "sint:port-flags", 1, false, primPortFlags)
}

func primWrite(ctx *Scheme, args []Val) (Val, int) {
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

func primPortp(ctx *Scheme, args []Val) (Val, int) {
	/*
		if _, ok := args[0].(*EofObject); ok {
			return ctx.TrueVal, 1
		}
	*/
	// FIXME
	return ctx.FalseVal, 1
}

func primPortFlags(ctx *Scheme, args []Val) (Val, int) {
	/*
		if _, ok := args[0].(*EofObject); ok {
			return ctx.TrueVal, 1
		}
	*/
	// FIXME
	return ctx.Zero, 1
}

func primClosePort(ctx *Scheme, args []Val) (Val, int) {
	/*
		if _, ok := args[0].(*EofObject); ok {
			return ctx.TrueVal, 1
		}
	*/
	// FIXME
	return ctx.Zero, 1
}
