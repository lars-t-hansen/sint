// Definitions of data structures representing Scheme values.

package core

import "fmt"

// Values.
//
// type Val union {
//   *Cons,
//   *Symbol,
//   *Procedure,
//   *Char,
//   *Str,
//   *Null,			// Singleton
//   *True			// Singleton
//   *False			// Singleton
//   *Unspecified,	// Singleton
//   *Undefined		// Singleton
//   *EofObject     // Singleton
//   *big.Int,      // Exact integer
//   *big.Float,    // Inexact real (rational?)
// }

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
	Env    *lexenv                  // closed-over lexical environment, nil for global procedures and primitives
	Primop func(*Scheme, []Val) Val // nil for non-primitives
}

func (c *Procedure) String() string {
	return "procedure"
}

type Char struct {
	// The value may not be a unicode code point, so 'rune' is only suggestive.
	// Also see Str below.
	Value rune
}

func (c *Char) String() string {
	return fmt.Sprint("[char ", c.Value, "]")
}

// Sint strings are Go strings, ie, they are immutable byte slices holding
// UTF-8 encoded Unicode code points.  This is nonstandard.  See the Design
// section of README.md.
type Str struct {
	Value string
}

func (s *Str) String() string {
	return fmt.Sprint("[string ", s.Value, "]")
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
