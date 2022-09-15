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

type ReadError struct {
	message string
}

func (r *ReadError) Error() string {
	return r.message
}

type reader struct {
	s   *Scheme
	rdr InputStream
}

func Read(s *Scheme, rdr InputStream) (Val, error) {
	r := &reader{s: s, rdr: rdr}
	return r.read()
}

func (r *reader) readError(msg string) *ReadError {
	return &ReadError{message: msg}
}

func (r *reader) read() (Val, error) {
	c, atEOF, err := r.skipWhitespace()
	if err != nil {
		return nil, err
	}
	if atEOF {
		return r.s.EofVal, nil
	}
	switch c {
	case '(':
		return r.readList()
	case ')':
		return nil, r.readError("Unmatched ')'")
	case '.':
		d, _, err := r.rdr.ReadRune()
		if err != nil {
			if e := r.handleErrorIgnoreEOF(err); e != nil {
				return nil, e
			}
			return r.s.Shared.DotSym, nil
		}
		r.rdr.UnreadRune()
		// TODO: Maybe .37 is valid syntax for 0.37
		if isSymbolSubsequent(d) {
			return r.readSymbol(c)
		}
		return r.s.Shared.DotSym, nil
	case '\'':
		v, err := r.read()
		if err != nil {
			return nil, err
		}
		return &Cons{Car: r.s.Shared.QuoteSym, Cdr: &Cons{Car: v, Cdr: r.s.NullVal}}, nil
	case '#':
		d, _, err := r.rdr.ReadRune()
		if err != nil {
			return r.handleError(err)
		}
		if d == 't' {
			return r.s.TrueVal, nil
		}
		if d == 'f' {
			return r.s.FalseVal, nil
		}
		if d == 'x' {
			return r.readHexNumber()
		}
		if d == '/' {
			// TODO: Implement this
			return nil, r.readError("No syntax for regular expressions yet")
		}
		if d == '\\' {
			return r.readCharacter()
		}
		return nil, r.readError("Unknown # syntax")
	case '"':
		return r.readString()
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return r.readDecimalNumber(c)
	default:
		if isSymbolInitial(c) {
			// TODO: quasiquote, unquote, unquote-splicing
			return r.readSymbol(c)
		}
		return nil, r.readError("Unknown character " + string(c))
	}
}

func (r *reader) readNotEOF() (Val, error) {
	w, err := r.read()
	if err != nil {
		return nil, err
	}
	if w == r.s.EofVal {
		return nil, r.readError("EOF not allowed here")
	}
	return w, nil
}

// Initial paren has been consumed
func (r *reader) readList() (Val, error) {
	var l *Cons
	var last *Cons
	for {
		canRead, err := r.canReadRightParen()
		if err != nil {
			return nil, err
		}
		if canRead {
			break
		}
		v, err := r.readNotEOF()
		if err != nil {
			return nil, err
		}
		if v == r.s.Shared.DotSym {
			if last == nil {
				return nil, r.readError("Illegal '.' in list")
			}
			w, err := r.readNotEOF()
			if err != nil {
				return nil, err
			}
			last.Cdr = w
			canRead, err := r.canReadRightParen()
			if err != nil {
				return nil, err
			}
			if !canRead {
				return nil, r.readError("Expected ')'")
			}
			break
		}
		p := &Cons{Car: v, Cdr: r.s.NullVal}
		if last != nil {
			last.Cdr = p
		} else {
			l = p
		}
		last = p
	}
	if l == nil {
		return r.s.NullVal, nil
	}
	return l, nil
}

func (r *reader) canReadRightParen() (bool, error) {
	c, atEOF, err := r.skipWhitespace()
	if err != nil {
		return false, err
	}
	if atEOF {
		return false, r.readError("EOF in datum")
	}
	if c == ')' {
		return true, nil
	}
	r.rdr.UnreadRune()
	return false, nil
}

// Leading #x has been consumed
func (r *reader) readHexNumber() (Val, error) {
	return nil, r.readError("Hex numbers not supported yet")
}

// "initial" is the leading digit
func (r *reader) readDecimalNumber(initial rune) (Val, error) {
	s := string(initial)
	isFloating := false

	// Integer part
	digs, any, err := r.readDecimalDigits()
	if err != nil {
		return nil, err
	}
	if any {
		s = s + digs
	}

	// Optional fractional part
	{
		d, _, err := r.rdr.ReadRune()
		if err != nil {
			if e := r.handleErrorIgnoreEOF(err); e != nil {
				return nil, e
			}
			goto eofAfterDatum
		}
		if d == '.' {
			isFloating = true
			s = s + "."
			digs, any, err := r.readDecimalDigits()
			if err != nil {
				return nil, err
			}
			if !any {
				return nil, r.readError("Digits required after decimal point")
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
			if e := r.handleErrorIgnoreEOF(err); e != nil {
				return nil, e
			}
			goto eofAfterDatum
		}
		if d == 'e' || d == 'E' {
			isFloating = true
			s = s + string(d)
			x, _, err := r.rdr.ReadRune()
			if err != nil {
				if e := r.handleErrorIgnoreEOF(err); e != nil {
					return nil, e
				}
				return nil, r.readError("EOF in datum")
			}
			if x == '+' || x == '-' {
				s = s + string(x)
			} else {
				r.rdr.UnreadRune()
			}
			digs, any, err := r.readDecimalDigits()
			if err != nil {
				return nil, err
			}
			if !any {
				return nil, r.readError("Digits required in exponent")
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
		return &f, nil
	}
	var i big.Int
	i.SetString(s, 10)
	return &i, nil
}

func (r *reader) readDecimalDigits() (s string, any bool, rdrErr error) {
	for {
		c, _, err := r.rdr.ReadRune()
		if err != nil {
			if e := r.handleErrorIgnoreEOF(err); e != nil {
				rdrErr = e
			}
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

func (r *reader) readCharacter() (Val, error) {
	e, _, err := r.rdr.ReadRune()
	if err != nil {
		if e := r.handleErrorIgnoreEOF(err); e != nil {
			return nil, e
		}
		return nil, r.readError("EOF in character")
	}
	switch e {
	case 'n', 'r', 's', 't':
		f, _, err := r.rdr.ReadRune()
		if err != nil {
			if e := r.handleErrorIgnoreEOF(err); e != nil {
				return nil, e
			}
			break
		}
		r.rdr.UnreadRune()
		if !isAlphabetic(f) {
			break
		}
		name, err := r.readSymbol(e)
		if err != nil {
			return nil, err
		}
		if name == r.s.Shared.NewlineSym {
			return &Char{Value: '\n'}, nil
		}
		if name == r.s.Shared.ReturnSym {
			return &Char{Value: '\r'}, nil
		}
		if name == r.s.Shared.TabSym {
			return &Char{Value: '\t'}, nil
		}
		if name == r.s.Shared.SpaceSym {
			return &Char{Value: ' '}, nil
		}
		return nil, r.readError("Illegal character name: " + name.Name)
	default:
		break
	}
	return &Char{Value: e}, nil
}

func (r *reader) readString() (*Str, error) {
	s := ""
	for {
		c, _, err := r.rdr.ReadRune()
		if err != nil {
			if e := r.handleErrorIgnoreEOF(err); e != nil {
				return nil, e
			}
			return nil, r.readError("EOF in string")
		}
		if c == '"' {
			return &Str{Value: s}, nil
		}
		if c == '\\' {
			d, _, err := r.rdr.ReadRune()
			if err != nil {
				if e := r.handleErrorIgnoreEOF(err); e != nil {
					return nil, e
				}
				return nil, r.readError("EOF in string")
			}
			switch d {
			case 'n':
				c = '\n'
			case 'r':
				c = '\r'
			case 't':
				c = '\t'
			case '\\':
				c = '\\'
			default:
				// TODO: \x, \u, probably others
				return nil, r.readError("Unsupported escape sequence in string")
			}
		}
		// TODO: Check for invalid code point?
		s = s + string(c)
	}
}

func (r *reader) readSymbol(initial rune) (*Symbol, error) {
	s := string(initial)
	for {
		d, _, err := r.rdr.ReadRune()
		if err != nil {
			if e := r.handleErrorIgnoreEOF(err); e != nil {
				return nil, e
			}
			break
		}
		if isSymbolSubsequent(d) {
			s = s + string(d)
		} else {
			r.rdr.UnreadRune()
			break
		}
	}
	return r.s.Intern(s), nil
}

// Skip whitespace and comments.  Throws on I/O error.  If EOF is encountered,
// atEOF is true and the ch is garbage.  Otherwise, atEOF is false and ch has the
// first nonblank character.
func (r *reader) skipWhitespace() (ch rune, atEOF bool, rdrErr error) {
again:
	c, _, err := r.rdr.ReadRune()
	if err != nil {
		if e := r.handleErrorIgnoreEOF(err); e != nil {
			rdrErr = e
		} else {
			atEOF = true
		}
		return
	}
	if isSpace(c) {
		goto again
	}
	if c == ';' {
		for {
			d, _, err := r.rdr.ReadRune()
			if err != nil {
				if e := r.handleErrorIgnoreEOF(err); e != nil {
					rdrErr = e
				} else {
					atEOF = true
				}
				return
			}
			if d == '\n' {
				goto again
			}
		}
	}
	ch = c
	return
}

var charTable [128]byte

const (
	kInitial    = 1
	kSubsequent = 2
	kSpace      = 4
	kAlphabetic = 8
)

func init() {
	for c := 'a'; c <= 'z'; c++ {
		charTable[c] = kInitial | kSubsequent | kAlphabetic
	}
	for c := 'A'; c <= 'Z'; c++ {
		charTable[c] = kInitial | kSubsequent | kAlphabetic
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
	charTable['!'] = kInitial | kSubsequent
	charTable[':'] = kInitial | kSubsequent
	charTable[','] = kInitial | kSubsequent
	charTable['@'] = kInitial | kSubsequent
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

func isAlphabetic(c rune) bool {
	return c < 128 && (charTable[c]&kAlphabetic) != 0
}

func (r *reader) handleError(err error) (Val, error) {
	if e := r.handleErrorIgnoreEOF(err); e != nil {
		return nil, e
	}
	return r.s.EofVal, nil
}

func (r *reader) handleErrorIgnoreEOF(err error) error {
	if err == io.EOF {
		return nil
	}
	return r.readError("I/O error")
}
