// The reader produces an sexpr from a rune input stream.

package runtime

import (
	"io"
	"math/big"
	. "sint/core"
)

// Matches bufio.Reader and strings.Reader
type InputStream interface {
	ReadRune() (rune, int, error)
	UnreadRune() error
}

type reader struct {
	c      *Scheme
	rdr    InputStream
	symDot *Symbol
}

func Read(c *Scheme, rdr InputStream) Val {
	r := &reader{c: c, rdr: rdr, symDot: c.Intern(".")}
	return r.read()
}

func (r *reader) read() Val {
	c, atEOF := r.skipWhitespace()
	if atEOF {
		return r.c.EofVal
	}
	switch c {
	case '(':
		return r.readList()
	case '.':
		d, _, err := r.rdr.ReadRune()
		if err != nil {
			r.handleErrorIgnoreEOF(err)
			return r.symDot
		}
		r.rdr.UnreadRune()
		// TODO: Maybe .37 is valid syntax for 0.37
		if isSymbolSubsequent(d) {
			return r.readSymbol(c)
		}
		return r.symDot
	case '#':
		d, _, err := r.rdr.ReadRune()
		if err != nil {
			return r.handleError(err)
		}
		if d == 't' {
			return r.c.TrueVal
		}
		if d == 'f' {
			return r.c.FalseVal
		}
		if d == 'x' {
			return r.readHexNumber()
		}
		if d == '!' {
			panic("No syntax for #!unspecified and suchlike yet")
		}
		if d == '\\' {
			panic("Characters not yet supported")
		}
		panic("Unknown # syntax")
	case '"':
		panic("Strings not yet supported")
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return r.readDecimalNumber(c)
	default:
		if isSymbolInitial(c) {
			return r.readSymbol(c)
		}
		panic("Unknown character")
	}
}

func (r *reader) readNotEOF() Val {
	w := r.read()
	if w == r.c.EofVal {
		panic("EOF not allowed here")
	}
	return w
}

// Initial paren has been consumed
func (r *reader) readList() Val {
	var l *Cons
	var last *Cons
	for {
		if r.canReadRightParen() {
			break
		}
		v := r.readNotEOF()
		if v == r.symDot {
			if last == nil {
				panic("Illegal '.' in list")
			}
			w := r.readNotEOF()
			last.Cdr = w
			if !r.canReadRightParen() {
				panic("Expected ')'")
			}
			break
		}
		p := &Cons{Car: v, Cdr: r.c.NullVal}
		if last != nil {
			last.Cdr = p
		} else {
			l = p
		}
		last = p
	}
	if l == nil {
		return r.c.NullVal
	}
	return l
}

func (r *reader) canReadRightParen() bool {
	c, atEOF := r.skipWhitespace()
	if atEOF {
		panic("EOF in datum")
	}
	if c == ')' {
		return true
	}
	r.rdr.UnreadRune()
	return false
}

// Leading #x has been consumed
func (r *reader) readHexNumber() Val {
	panic("Hex numbers not supported yet")
}

// "initial" is the leading digit
func (r *reader) readDecimalNumber(initial rune) Val {
	s := string(initial)
	isFloating := false

	// Integer part
	if digs, any := r.readDecimalDigits(); any {
		s = s + digs
	}

	// Optional fractional part
	{
		d, _, err := r.rdr.ReadRune()
		if err != nil {
			r.handleErrorIgnoreEOF(err)
			goto eofAfterDatum
		}
		if d == '.' {
			isFloating = true
			s = s + "."
			digs, any := r.readDecimalDigits()
			if !any {
				panic("Digits required after decimal point")
			}
			s = s + digs
		} else {
			r.rdr.UnreadRune()
		}
	}

	// Optional exponential part
	{
		d, _, err := r.rdr.ReadRune()
		if err != nil {
			r.handleErrorIgnoreEOF(err)
			goto eofAfterDatum
		}
		if d == 'e' || d == 'E' {
			isFloating = true
			s = s + string(d)
			x, _, err := r.rdr.ReadRune()
			if err != nil {
				r.handleErrorIgnoreEOF(err)
				panic("EOF in datum")
			}
			if x == '+' || x == '-' {
				s = s + string(x)
			} else {
				r.rdr.UnreadRune()
			}
			digs, any := r.readDecimalDigits()
			if !any {
				panic("Digits required in exponent")
			}
			s = s + digs
		} else {
			r.rdr.UnreadRune()
		}
	}

eofAfterDatum:
	if isFloating {
		var f big.Float
		f.Parse(s, 10)
		return &f
	}
	var i big.Int
	i.SetString(s, 10)
	return &i
}

func (r *reader) readDecimalDigits() (s string, any bool) {
	for {
		c, _, err := r.rdr.ReadRune()
		if err != nil {
			r.handleErrorIgnoreEOF(err)
			return
		}
		if c < '0' || c > '9' {
			r.rdr.UnreadRune()
			return
		}
		any = true
		s = s + string(c)
	}
}

func (r *reader) readSymbol(initial rune) Val {
	s := string(initial)
	for {
		d, _, err := r.rdr.ReadRune()
		if err != nil {
			r.handleErrorIgnoreEOF(err)
			break
		}
		if isSymbolSubsequent(d) {
			s = s + string(d)
		} else {
			r.rdr.UnreadRune()
			break
		}
	}
	return r.c.Intern(s)
}

// Skip whitespace.  Throws on I/O error.  If EOF is encountered, atEOF is
// true and the ch is garbage.  Otherwise, atEOF is false and ch has the
// first nonblank character.
func (r *reader) skipWhitespace() (ch rune, atEOF bool) {
	for {
		c, _, err := r.rdr.ReadRune()
		if err != nil {
			r.handleErrorIgnoreEOF(err)
			atEOF = true
			return
		}
		if !isSpace(c) {
			ch = c
			return
		}
	}
}

var charTable [128]byte

const (
	kInitial    = 1
	kSubsequent = 2
	kSpace      = 4
)

func init() {
	for c := 'a'; c <= 'z'; c++ {
		charTable[c] = kInitial | kSubsequent
	}
	for c := 'A'; c <= 'Z'; c++ {
		charTable[c] = kInitial | kSubsequent
	}
	charTable['_'] = kInitial | kSubsequent
	charTable['$'] = kInitial | kSubsequent
	charTable['+'] = kInitial | kSubsequent
	charTable['-'] = kInitial | kSubsequent
	charTable['*'] = kInitial | kSubsequent
	charTable['/'] = kInitial | kSubsequent
	charTable['<'] = kInitial | kSubsequent
	charTable['>'] = kInitial | kSubsequent
	charTable['='] = kInitial | kSubsequent
	charTable['?'] = kInitial | kSubsequent
	for c := '0'; c <= '9'; c++ {
		charTable[c] = kSubsequent
	}
	charTable[' '] = kSpace
	charTable['\n'] = kSpace
	charTable['\r'] = kSpace
	charTable['\t'] = kSpace
}

func isSymbolInitial(c rune) bool {
	return c < 128 && (charTable[c]&kInitial) != 0
}

func isSymbolSubsequent(c rune) bool {
	return c < 128 && (charTable[c]&kSubsequent) != 0
}

func isSpace(c rune) bool {
	return c < 128 && (charTable[c]&kSpace) != 0
}

func (r *reader) handleError(err error) Val {
	r.handleErrorIgnoreEOF(err)
	return r.c.EofVal
}

func (r *reader) handleErrorIgnoreEOF(err error) {
	if err == io.EOF {
		return
	}
	panic("I/O error")
}
