// Definitions of data structures representing Scheme values.

package core

import "fmt"

// Values.
//
// Values are represented as pointers, ensuring that data races at least will
// not expose torn values.
//
// type Val union {
//   *Cons,			// Pair
//   *Symbol,		// Symbol (interned or not)
//   *Procedure,	// Procedure: lambda + environment
//   *Char,			// Unicode character
//   *Str,			// Immutable UTF-8 encoded Unicode code points
//   *Chan,			// Channel that can transmit any Val
//   *UnwindPkg,    // Internal value that transmits unwind information
//   *Null,			// The () singleton
//   *True,			// The #t singleton
//   *False,		// The #f singleton
//   *Unspecified,  // The #!unspecified singleton
//   *Undefined,    // The undefined value singleton
//   *EofObject,    // The #!eof singleton
//   *big.Int,		// Exact integer
//   *big.Float,    // Inexact rational
// }
//
// Eventually the UnwindPkg could become an instance of a Record type.

type Val interface {
	fmt.Stringer
}

type Cons struct {
	Car Val
	Cdr Val
}

func (c *Cons) String() string {
	return "[cons " + c.Car.String() + " " + c.Cdr.String() + "]"
}

type Symbol struct {
	Name  string
	Value Val // if the symbol is a global variable, otherwise c.undefined
}

func (c *Symbol) String() string {
	return "[symbol " + c.Name + "]"
}

type Procedure struct {
	Lam    *Lambda
	Env    *lexenv                         // closed-over lexical environment, nil for global procedures and primitives
	Primop func(*Scheme, []Val) (Val, int) // nil for non-primitives
}

func (c *Procedure) String() string {
	return "procedure"
}

// A char value is a Unicode code point always: character literals are
// restricted to code points, individual-char getters from strings
// will never return non-code points, integer->char checks that its
// input is a valid code point, and (in the future) read-char checks
// that it is reading a code point.
type Char struct {
	Value rune
}

func (c *Char) String() string {
	return fmt.Sprint("[char ", c.Value, "]")
}

// Sint strings are Go strings, ie, they are immutable byte slices holding
// UTF-8 encoded Unicode code points.  This is nonstandard.  See MANUAL.md.
type Str struct {
	Value string
}

func (s *Str) String() string {
	return fmt.Sprint("[string ", s.Value, "]")
}

type Chan struct {
	Ch chan Val
}

func (s *Chan) String() string {
	return "channel"
}

type UnwindPkg struct {
	Key     Val
	Payload Val
}

func (u *UnwindPkg) String() string {
	return "unwind-package: " + u.Payload.String()
}

type Null struct{}

func (c *Null) String() string {
	return "null"
}

type True struct{}

func (c *True) String() string {
	return "true"
}

type False struct{}

func (c *False) String() string {
	return "false"
}

type Unspecified struct{}

func (c *Unspecified) String() string {
	return "unspecified"
}

type Undefined struct{}

func (c *Undefined) String() string {
	return "undefined"
}

type EofObject struct{}

func (c *EofObject) String() string {
	return "eof-object"
}
