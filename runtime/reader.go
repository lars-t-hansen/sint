// The reader produces an sexpr from a rune input stream.
// This is not a standards-compliant reader, but it's moving in that direction.

package runtime

import (
	"io"
	"math/big"
	"regexp"
	. "sint/core"
	"unicode/utf8"
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

// Returns a number or nil on error (not numeric)
// radix is -10 for "default" decimal radix, otherwise the requested radix
func StringToNumber(s string, radix int) Val {
	if radix != -10 && radix != 10 {
		panic("Only radix 10 supported at this time")
	}
	isNumber, isFloating := hasNumberSyntax(s)
	if !isNumber {
		return nil
	}
	if isFloating {
		var f big.Float
		f.Parse(s, 10)
		return &f
	}
	var i big.Int
	i.SetString(s, 10)
	return &i
}

// Determine if a string can be parsed as a number.
// TODO: This must handle infinities, nan, radix prefix, exact/inexact prefix,
// and eventually complexes and so on.
func hasNumberSyntax(s string) (isNumber bool, isFloating bool) {
	idx := 0
	lim := len(s)
	hasIntegerPart := false
	hasFractionalPart := false
	hasExponentPart := false
	// Assume optimistically it's a number
	isNumber = true
	// Skip a leading sign
	{
		ch, size := utf8.DecodeRuneInString(s[idx:])
		if ch == '+' || ch == '-' {
			idx += size
		}
	}
	// Skip integer part, if present
	for idx < lim {
		ch, size := utf8.DecodeRuneInString(s[idx:])
		if ch < '0' || ch > '9' {
			break
		}
		idx += size
		hasIntegerPart = true
	}
	// Skip fractional part
	if idx < lim {
		ch, size := utf8.DecodeRuneInString(s[idx:])
		if ch == '.' {
			idx += size
			ndig := 0
			for idx < lim {
				ch, size := utf8.DecodeRuneInString(s[idx:])
				if ch < '0' || ch > '9' {
					break
				}
				idx += size
				ndig++
			}
			if ndig > 0 {
				hasFractionalPart = true
			} else {
				isNumber = false
			}
		}
	}
	// Skip exponent part
	if isNumber && idx < lim {
		ch, size := utf8.DecodeRuneInString(s[idx:])
		if ch == 'e' || ch == 'E' {
			idx += size
			if idx < lim {
				ch, size := utf8.DecodeRuneInString(s[idx:])
				if ch == '+' || ch == '-' {
					idx += size
				}
				ndig := 0
				for idx < lim {
					ch, size := utf8.DecodeRuneInString(s[idx:])
					if ch < '0' || ch > '9' {
						break
					}
					idx += size
					ndig++
				}
				if ndig > 0 {
					hasExponentPart = true
				} else {
					isNumber = false
				}
			} else {
				isNumber = false
			}
		}
	}
	// It's only a number if we got to the end
	if idx < lim {
		isNumber = false
	}
	if isNumber && !hasIntegerPart && !hasFractionalPart && !hasExponentPart {
		isNumber = false
	}
	if isNumber && (hasFractionalPart || hasExponentPart) {
		isFloating = true
	}
	return
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
		// Symbols starting with . is not a thing but it's useful.
		// Numbers can start with .
		if isSymbolSubsequent(d) {
			return r.readSymbolOrNumber(c)
		}
		return r.s.Shared.DotSym, nil
	case '\'':
		v, err := r.read()
		if err != nil {
			return nil, err
		}
		return &Cons{Car: r.s.Shared.QuoteSym, Cdr: &Cons{Car: v, Cdr: r.s.NullVal}}, nil
	case '`':
		return nil, r.readError("Unhandled reserved syntax '`'")
	case ',':
		d, _, err := r.rdr.ReadRune()
		if err != nil {
			if e := r.handleErrorIgnoreEOF(err); e != nil {
				return nil, e
			}
			return nil, r.readError("Unhandled reserved syntax ','")
		}
		r.rdr.UnreadRune()
		if d == '@' {
			return nil, r.readError("Unhandled reserved syntax ',@'")
		}
		return nil, r.readError("Unhandled reserved syntax ','")
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
			return r.readRegExp()
		}
		if d == '\\' {
			return r.readCharacter()
		}
		return nil, r.readError("Unknown # syntax")
	case '"':
		return r.readString()
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return r.readNumber(c)
	case '-', '+':
		return r.readSymbolOrNumber(c)
	default:
		if isSymbolInitial(c) {
			return r.readSymbolOrNumber(c)
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

func (r *reader) readRegExp() (Val, error) {
	s, err := r.readStringOrRegExp('/', "regular expression")
	if err != nil {
		return nil, err
	}
	re, err := regexp.Compile(s)
	if err != nil {
		return nil, r.readError("Bad regular expression: " + err.Error())
	}
	return &Regexp{Value: re}, nil
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
		name, err := r.readSymbolOrNumber(e)
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
		return nil, r.readError("Illegal character name: " + name.(*Symbol).Name)
	default:
		break
	}
	return &Char{Value: e}, nil
}

func (r *reader) readString() (Val, error) {
	s, err := r.readStringOrRegExp('"', "string")
	if err != nil {
		return nil, err
	}
	return &Str{Value: s}, nil
}

func (r *reader) readStringOrRegExp(terminator rune, kind string) (string, error) {
	s := ""
	for {
		c, _, err := r.rdr.ReadRune()
		if err != nil {
			if e := r.handleErrorIgnoreEOF(err); e != nil {
				return "", e
			}
			return "", r.readError("EOF in " + kind)
		}
		if c == terminator {
			return s, nil
		}
		if c == '\\' {
			d, _, err := r.rdr.ReadRune()
			if err != nil {
				if e := r.handleErrorIgnoreEOF(err); e != nil {
					return "", e
				}
				return "", r.readError("EOF in " + kind)
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
				if d == terminator {
					c = terminator
					break
				}
				// TODO: \x, \u, probably others
				return "", r.readError("Unsupported escape sequence in " + kind)
			}
		}
		// TODO: Check for invalid code point?
		s = s + string(c)
	}
}

func (r *reader) readNumber(initial rune) (Val, error) {
	v, err := r.readSymbolOrNumber(initial)
	if err != nil {
		return v, err
	}
	if _, isSym := v.(*Symbol); isSym {
		return nil, r.readError("Number expected here")
	}
	return v, nil
}

func (r *reader) readSymbolOrNumber(initial rune) (Val, error) {
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
	num := StringToNumber(s, -10)
	if num != nil {
		return num, nil
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
	charTable['%'] = kInitial | kSubsequent
	charTable['&'] = kInitial | kSubsequent
	charTable['*'] = kInitial | kSubsequent
	charTable['/'] = kInitial | kSubsequent
	charTable['<'] = kInitial | kSubsequent
	charTable['>'] = kInitial | kSubsequent
	charTable['='] = kInitial | kSubsequent
	charTable['?'] = kInitial | kSubsequent
	charTable['!'] = kInitial | kSubsequent
	charTable[':'] = kInitial | kSubsequent
	charTable['@'] = kInitial | kSubsequent
	charTable['^'] = kInitial | kSubsequent
	charTable['~'] = kInitial | kSubsequent
	for c := '0'; c <= '9'; c++ {
		charTable[c] = kSubsequent
	}
	charTable['.'] = kSubsequent
	charTable['+'] = kSubsequent
	charTable['-'] = kSubsequent
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
