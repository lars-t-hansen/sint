// R7RS 6.13, Input and output, also see io.sch

package runtime

import (
	"io"
	"math/big"
	. "sint/core"
)

func initIoPrimitives(ctx *Scheme) {
	addPrimitive(ctx, "display", 1, true, primDisplay)
	addPrimitive(ctx, "newline", 0, true, primNewline)
	addPrimitive(ctx, "write", 1, true, primWrite)
	addPrimitive(ctx, "writeln", 1, true, primWriteln)
	addPrimitive(ctx, "write-char", 1, true, primWriteChar)
	addPrimitive(ctx, "eof-object", 0, false, primEofObject)
	addPrimitive(ctx, "eof-object?", 1, false, primEofObjectp)
	addPrimitive(ctx, "port?", 1, false, primPortp)
	addPrimitive(ctx, "close-port", 1, false, primClosePort)
	addPrimitive(ctx, "sint:port-flags", 1, false, primPortFlags)
	addPrimitive(ctx, "read", 0, true, primRead)
	addPrimitive(ctx, "read-char", 0, true, primReadChar)
}

func currentInputPort(ctx *Scheme) *Port {
	p := ctx.GetTlsValue(CurrentInputPort)
	if port, ok := p.(*Port); ok {
		return port
	}
	panic("currentInputPort: no port set")
}

var portDiagnostics map[int]string = make(map[int]string)

func init() {
	portDiagnostics[IsBinaryPort] = "binary"
	portDiagnostics[IsTextPort] = "textual"
	portDiagnostics[IsInputPort] = "input"
	portDiagnostics[IsOutputPort] = "output"
}

// If args[maybePort] exists it must be a port, otherwise get the port parameter at
// the tlsKey.  The port must have the right direction and type.
//
// If the `v` value is not nil then (v, nv) is an error return and `port` should be ignored.

func getPort(ctx *Scheme, args []Val, maybePort int, name string, tlsKey int32, direction PortFlags, ty PortFlags) (port *Port, v Val, nv int) {
	ok := true
	if len(args) > maybePort {
		if port, ok = args[maybePort].(*Port); ok {
			goto checkPort
		}
		v, nv = ctx.Error(name + ": not an " + portDiagnostics[int(direction)] + " port: " + args[maybePort].String())
		return
	}
	{
		p := ctx.GetTlsValue(tlsKey)
		port, ok = p.(*Port)
		if !ok {
			panic(portDiagnostics[int(direction)] + " port: no current port set")
		}
	}
checkPort:
	flags := port.Flags()
	if (flags & direction) == 0 {
		v, nv = ctx.Error(name + ": not an " + portDiagnostics[int(direction)] + " port: " + port.String())
		return
	}
	if (flags & ty) == 0 {
		v, nv = ctx.Error(name + ": not a " + portDiagnostics[int(ty)] + " port: " + port.String())
		return
	}
	return
}

func primWrite(ctx *Scheme, args []Val) (Val, int) {
	p, v, nv := getPort(ctx, args, 1, "write", CurrentOutputPort, IsOutputPort, IsTextPort)
	if v != nil {
		return v, nv
	}
	{
		writer := p.AcquireOutputStream()
		Write(args[0], false, writer)
		p.ReleaseOutputStream(writer)
	}
	return ctx.UnspecifiedVal, 1
}

func primWriteln(ctx *Scheme, args []Val) (Val, int) {
	p, v, nv := getPort(ctx, args, 1, "writeln", CurrentOutputPort, IsOutputPort, IsTextPort)
	if v != nil {
		return v, nv
	}
	{
		writer := p.AcquireOutputStream()
		Write(args[0], false, writer)
		writer.WriteRune('\n')
		p.ReleaseOutputStream(writer)
	}
	return ctx.UnspecifiedVal, 1
}

func primDisplay(ctx *Scheme, args []Val) (Val, int) {
	p, v, nv := getPort(ctx, args, 1, "display", CurrentOutputPort, IsOutputPort, IsTextPort)
	if v != nil {
		return v, nv
	}
	{
		writer := p.AcquireOutputStream()
		Write(args[0], true, writer)
		p.ReleaseOutputStream(writer)
	}
	return ctx.UnspecifiedVal, 1
}

func primNewline(ctx *Scheme, args []Val) (Val, int) {
	p, v, nv := getPort(ctx, args, 0, "newline", CurrentOutputPort, IsOutputPort, IsTextPort)
	if v != nil {
		return v, nv
	}
	{
		writer := p.AcquireOutputStream()
		writer.WriteRune('\n')
		p.ReleaseOutputStream(writer)
	}
	return ctx.UnspecifiedVal, 1
}

func primWriteChar(ctx *Scheme, args []Val) (Val, int) {
	c, ok := args[0].(*Char)
	if !ok {
		return ctx.Error("write-char: not a character: " + args[0].String())
	}
	p, v, nv := getPort(ctx, args, 1, "write-char", CurrentOutputPort, IsOutputPort, IsTextPort)
	if v != nil {
		return v, nv
	}
	{
		writer := p.AcquireOutputStream()
		writer.WriteRune(c.Value)
		p.ReleaseOutputStream(writer)
	}
	return ctx.UnspecifiedVal, 1
}

func primRead(ctx *Scheme, args []Val) (Val, int) {
	p, v, nv := getPort(ctx, args, 0, "read", CurrentInputPort, IsInputPort, IsTextPort)
	if v != nil {
		return v, nv
	}
	reader := p.AcquireInputStream()
	readv, readErr := Read(ctx, reader)
	p.ReleaseInputStream(reader)
	if readErr != nil {
		return ctx.Error(readErr.Error())
	}
	return readv, 1
}

func primReadChar(ctx *Scheme, args []Val) (Val, int) {
	p, v, nv := getPort(ctx, args, 0, "read-char", CurrentInputPort, IsInputPort, IsTextPort)
	if v != nil {
		return v, nv
	}
	reader := p.AcquireInputStream()
	readv, _, readErr := reader.ReadRune()
	p.ReleaseInputStream(reader)
	if readErr == io.EOF {
		return ctx.EofVal, 1
	}
	if readErr != nil {
		return ctx.Error(readErr.Error())
	}
	// TODO: Do we need to range check the value?
	return &Char{Value: readv}, 1
}
func primEofObject(ctx *Scheme, args []Val) (Val, int) {
	return ctx.EofVal, 1
}

func primEofObjectp(ctx *Scheme, args []Val) (Val, int) {
	if _, ok := args[0].(*EofObject); ok {
		return ctx.TrueVal, 1
	}
	return ctx.FalseVal, 1
}

func primPortp(ctx *Scheme, args []Val) (Val, int) {
	if _, ok := args[0].(*Port); ok {
		return ctx.TrueVal, 1
	}
	return ctx.FalseVal, 1
}

func primPortFlags(ctx *Scheme, args []Val) (Val, int) {
	if p, ok := args[0].(*Port); ok {
		return big.NewInt(int64(p.Flags())), 1
	}
	return ctx.Zero, 1
}

// Is this properly defined?  If it's an input-and-output port, are
// we closing both sides?  Surely not.
func primClosePort(ctx *Scheme, args []Val) (Val, int) {
	/*
		if _, ok := args[0].(*EofObject); ok {
			return ctx.TrueVal, 1
		}
	*/
	// FIXME
	return ctx.Zero, 1
}
