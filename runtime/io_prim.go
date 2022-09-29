// R7RS 6.13, Input and output, also see io.sch

package runtime

import (
	"bufio"
	"io"
	"math/big"
	"os"
	. "sint/core"
	"strings"
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
	addPrimitive(ctx, "sint:port-flags", 1, false, primPortFlags)
	addPrimitive(ctx, "read", 0, true, primRead)
	addPrimitive(ctx, "read-char", 0, true, primReadChar)
	addPrimitive(ctx, "peek-char", 0, true, primPeekChar)
	addPrimitive(ctx, "read-line", 0, true, primReadLine)
	addPrimitive(ctx, "open-input-file", 1, false, primOpenInputFile)
	addPrimitive(ctx, "close-input-port", 1, false, primCloseInputPort)
	addPrimitive(ctx, "open-output-file", 1, false, primOpenOutputFile)
	addPrimitive(ctx, "close-output-port", 1, false, primCloseOutputPort)
}

var portDiagnostics map[int]string = make(map[int]string)

func init() {
	portDiagnostics[IsBinaryPort] = "binary"
	portDiagnostics[IsTextPort] = "textual"
	portDiagnostics[IsInputPort] = "input"
	portDiagnostics[IsOutputPort] = "output"
}

// If maybePort is not undefined it must be a port, otherwise get the port parameter at
// the tlsKey.  The port must have the right direction and type.
//
// If the `v` value is not nil then (v, nv) is an error return and `port` should be ignored.

func getPort(ctx *Scheme, maybePort Val, name string, tlsKey int32, direction PortFlags, ty PortFlags) (port *Port, v Val, nv int) {
	ok := true
	if maybePort != ctx.UndefinedVal {
		if port, ok = maybePort.(*Port); ok {
			goto checkPort
		}
		v, nv = ctx.Error(name+": not an "+portDiagnostics[int(direction)]+" port", maybePort)
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
		v, nv = ctx.Error(name+": not an "+portDiagnostics[int(direction)]+" port", port)
		return
	}
	if (flags & ty) == 0 {
		v, nv = ctx.Error(name+": not a "+portDiagnostics[int(ty)]+" port", port)
		return
	}
	return
}

func primWrite(ctx *Scheme, a0, a1 Val, _ []Val) (Val, int) {
	p, v, nv := getPort(ctx, a1, "write", CurrentOutputPort, IsOutputPort, IsTextPort)
	if v != nil {
		return v, nv
	}
	{
		writer := p.AcquireOutputStream()
		Write(a0, true, writer)
		p.ReleaseOutputStream(writer)
	}
	return ctx.UnspecifiedVal, 1
}

func primWriteln(ctx *Scheme, a0, a1 Val, _ []Val) (Val, int) {
	p, v, nv := getPort(ctx, a1, "writeln", CurrentOutputPort, IsOutputPort, IsTextPort)
	if v != nil {
		return v, nv
	}
	{
		writer := p.AcquireOutputStream()
		Write(a0, true, writer)
		writer.WriteRune('\n')
		p.ReleaseOutputStream(writer)
	}
	return ctx.UnspecifiedVal, 1
}

func primDisplay(ctx *Scheme, a0, a1 Val, _ []Val) (Val, int) {
	p, v, nv := getPort(ctx, a1, "display", CurrentOutputPort, IsOutputPort, IsTextPort)
	if v != nil {
		return v, nv
	}
	{
		writer := p.AcquireOutputStream()
		Write(a0, false, writer)
		p.ReleaseOutputStream(writer)
	}
	return ctx.UnspecifiedVal, 1
}

func primNewline(ctx *Scheme, a0, _ Val, _ []Val) (Val, int) {
	p, v, nv := getPort(ctx, a0, "newline", CurrentOutputPort, IsOutputPort, IsTextPort)
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

func primWriteChar(ctx *Scheme, a0, a1 Val, _ []Val) (Val, int) {
	c, cErr := checkChar(ctx, a0, "write-char")
	if cErr != nil {
		return ctx.SignalWrappedError(cErr)
	}
	p, v, nv := getPort(ctx, a1, "write-char", CurrentOutputPort, IsOutputPort, IsTextPort)
	if v != nil {
		return v, nv
	}
	{
		writer := p.AcquireOutputStream()
		writer.WriteRune(c)
		p.ReleaseOutputStream(writer)
	}
	return ctx.UnspecifiedVal, 1
}

func primRead(ctx *Scheme, a0, _ Val, _ []Val) (Val, int) {
	p, v, nv := getPort(ctx, a0, "read", CurrentInputPort, IsInputPort, IsTextPort)
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

func primReadChar(ctx *Scheme, a0, _ Val, _ []Val) (Val, int) {
	p, v, nv := getPort(ctx, a0, "read-char", CurrentInputPort, IsInputPort, IsTextPort)
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

// TODO: This is just read-char + unread, it would be nice to merge the two functions.
func primPeekChar(ctx *Scheme, a0, a1 Val, _ []Val) (Val, int) {
	p, v, nv := getPort(ctx, a0, "peek-char", CurrentInputPort, IsInputPort, IsTextPort)
	if v != nil {
		return v, nv
	}
	reader := p.AcquireInputStream()
	readv, _, readErr := reader.ReadRune()
	reader.UnreadRune()
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

func primReadLine(ctx *Scheme, a0, a1 Val, _ []Val) (Val, int) {
	p, v, nv := getPort(ctx, a0, "read-line", CurrentInputPort, IsInputPort, IsTextPort)
	if v != nil {
		return v, nv
	}
	reader := p.AcquireInputStream()
	var buf strings.Builder
	var readv rune
	var readErr error
	for {
		readv, _, readErr = reader.ReadRune()
		if readErr != nil || readv == '\n' {
			break
		}
		buf.WriteRune(readv)
	}
	p.ReleaseInputStream(reader)
	if readErr == io.EOF {
		return ctx.EofVal, 1
	}
	if readErr != nil {
		return ctx.Error(readErr.Error())
	}
	return &Str{Value: buf.String()}, 1
}

func primEofObject(ctx *Scheme, _, _ Val, _ []Val) (Val, int) {
	return ctx.EofVal, 1
}

func primEofObjectp(ctx *Scheme, a0, _ Val, _ []Val) (Val, int) {
	if _, ok := a0.(*EofObject); ok {
		return ctx.TrueVal, 1
	}
	return ctx.FalseVal, 1
}

func primPortp(ctx *Scheme, a0, _ Val, _ []Val) (Val, int) {
	if _, ok := a0.(*Port); ok {
		return ctx.TrueVal, 1
	}
	return ctx.FalseVal, 1
}

func primPortFlags(ctx *Scheme, a0, _ Val, _ []Val) (Val, int) {
	if p, ok := a0.(*Port); ok {
		return big.NewInt(int64(p.Flags())), 1
	}
	return ctx.Zero, 1
}

type SchemeFile struct {
	handle    *os.File
	instream  *bufio.Reader
	outstream *bufio.Writer
}

func (f *SchemeFile) ReadRune() (rune, int, error) {
	return f.instream.ReadRune()
}

func (f *SchemeFile) UnreadRune() error {
	return f.instream.UnreadRune()
}

func (f *SchemeFile) WriteString(s string) (int, error) {
	return f.outstream.WriteString(s)
}

func (f *SchemeFile) WriteRune(r rune) (int, error) {
	return f.outstream.WriteRune(r)
}

func (f *SchemeFile) Flush() {
	f.outstream.Flush()
}

func (f *SchemeFile) Close() {
	if f.outstream != nil {
		f.outstream.Flush()
	}
	f.handle.Close()
}

func primOpenInputFile(ctx *Scheme, a0, _ Val, _ []Val) (Val, int) {
	fn, fnErr := checkString(ctx, a0, "open-input-file")
	if fnErr != nil {
		return ctx.SignalWrappedError(fnErr)
	}
	input, inErr := os.Open(fn)
	if inErr != nil {
		return ctx.Error("open-input-file: can't open file: "+inErr.Error(), a0)
	}
	f := &SchemeFile{handle: input, instream: bufio.NewReader(input)}
	return NewInputPort(f, true, fn), 1
}

func primOpenOutputFile(ctx *Scheme, a0, _ Val, _ []Val) (Val, int) {
	fn, fnErr := checkString(ctx, a0, "open-output-file")
	if fnErr != nil {
		return ctx.SignalWrappedError(fnErr)
	}
	output, outErr := os.Create(fn)
	if outErr != nil {
		return ctx.Error("open-output-file: can't open file: "+outErr.Error(), a0)
	}
	f := &SchemeFile{handle: output, outstream: bufio.NewWriter(output)}
	return NewOutputPort(f, true, fn), 1
}

func primCloseInputPort(ctx *Scheme, a0, _ Val, _ []Val) (Val, int) {
	port, portOk := a0.(*Port)
	if !portOk {
		return ctx.Error("close-input-port: not a port", a0)
	}
	f := port.Flags()
	if (f & IsInputPort) == 0 {
		return ctx.Error("close-input-port: not an input port", port)
	}
	s := port.AcquireInputStream() // The port is now locked
	if (port.RacyFlags() & IsClosedPort) == 0 {
		s.Close()
		port.RacySetClosed()
	}
	port.ReleaseInputStream(s)
	return ctx.UnspecifiedVal, 1
}

func primCloseOutputPort(ctx *Scheme, a0, _ Val, _ []Val) (Val, int) {
	port, portOk := a0.(*Port)
	if !portOk {
		return ctx.Error("close-output-port: not a port", a0)
	}
	f := port.Flags()
	if (f & IsOutputPort) == 0 {
		return ctx.Error("close-output-port: not an output port", port)
	}
	s := port.AcquireOutputStream() // The port is now locked
	if (port.RacyFlags() & IsClosedPort) == 0 {
		s.Close()
		port.RacySetClosed()
	}
	port.ReleaseOutputStream(s)
	return ctx.UnspecifiedVal, 1
}
