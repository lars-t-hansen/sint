// Basic input/output stream implementations for streams that can be embedded in Port values.
// This feels rough still.

package runtime

import (
	"bufio"
	"os"
	"sint/core"
)

// StdWriter is a ClosableFlushableOutputStream that writes to stdout/stderr, see core/values.go

type stdWriter struct {
	stream *os.File
}

func NewStdoutWriter() core.ClosableFlushableOutputStream {
	return &stdWriter{stream: os.Stdout}
}

func NewStderrWriter() core.ClosableFlushableOutputStream {
	return &stdWriter{stream: os.Stderr}
}

func (w *stdWriter) WriteString(s string) (int, error) {
	return w.stream.WriteString(s)
}

func (w *stdWriter) WriteRune(r rune) (int, error) {
	return w.stream.WriteString(string(r))
}

func (w *stdWriter) Flush() {
	// Do nothing
}

func (w *stdWriter) Close() {
	// Do nothing?
}

// StdReader is a ClosableInputStream that reads from stdin, see core/values.go
//
// TODO: This probably has to be a singleton per underlying stream, or there
// will be multiple buffers, which will be a mess.

type stdReader struct {
	stream *bufio.Reader
}

func NewStdinReader() core.ClosableInputStream {
	return &stdReader{stream: bufio.NewReader(os.Stdin)}
}

func (r *stdReader) ReadRune() (rune, int, error) {
	return r.stream.ReadRune()
}

func (r *stdReader) UnreadRune() error {
	return r.stream.UnreadRune()
}

func (r *stdReader) Close() {
	// Nothing
}
