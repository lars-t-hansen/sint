// Definitions of data types representing compiled code.

package core

import (
	"fmt"
	"strconv"
)

// Compiled code
//
// type Code union {
//   *If,
//   *Begin,
//   *Quote,
//   *Call,
//   *Apply,
//   *Lambda,
//   *Let,
//   *Letrec,
//   *LetStar,
//   *LetValues,
//   *Lexical,
//   *Setlex,
//   *Global,
//   *Setglobal
// }

type Code interface {
	// TODO: Documentation: each expression should carry its source location
	fmt.Stringer
}

type Quote struct {
	Value Val
}

func (c *Quote) String() string {
	return "(quote " + c.Value.String() + ")"
}

type If struct {
	Test       Code
	Consequent Code
	Alternate  Code
}

func (c *If) String() string {
	return "(if " + c.Test.String() + " " + c.Consequent.String() + " " + c.Alternate.String() + ")"
}

type Begin struct {
	Exprs []Code
}

func (c *Begin) String() string {
	return "(begin " + stringifyExprs(c.Exprs) + ")"
}

func stringifyExprs(es []Code) string {
	s := es[0].String()
	for _, e := range es[1:] {
		s = s + " " + e.String()
	}
	return s
}

type Call struct {
	Exprs []Code
}

func (c *Call) String() string {
	return "(" + stringifyExprs(c.Exprs) + ")"
}

type Apply struct {
	Proc Code
	Args Code
}

func (c *Apply) String() string {
	return "(sint:apply " + c.Proc.String() + " " + c.Args.String() + ")"
}

type Lambda struct {
	Fixed int
	Rest  bool
	Body  Code
	Name  string
	// TODO: Documentation: this should carry the doc string and the source code
	// TODO: Documentation: This should carry the names of locals in the rib
}

func (c *Lambda) String() string {
	return "(lambda " + strconv.Itoa(c.Fixed) + " " + strconv.FormatBool(c.Rest) + " " + c.Body.String() + ")"
}

type Let struct {
	Exprs []Code
	Body  Code
	// TODO: Documentation: This should carry the names of locals in the rib
}

func (c *Let) String() string {
	return "(let (" + stringifyExprs(c.Exprs) + ") " + c.Body.String() + ")"
}

type LetStar struct {
	Exprs []Code
	Body  Code
	// TODO: Documentation: This should carry the names of locals in the rib
}

func (c *LetStar) String() string {
	return "(let* (" + stringifyExprs(c.Exprs) + ") " + c.Body.String() + ")"
}

type Letrec struct {
	Exprs []Code
	Body  Code
	// TODO: Documentation: This should carry the names of locals in the rib
}

func (c *Letrec) String() string {
	return "(letrec (" + stringifyExprs(c.Exprs) + ") " + c.Body.String() + ")"
}

type LetValues struct {
	Exprs    []Code // rhs expressions
	NumVars  int    // total number of variables bound
	Bindings []int  // a bitvector <number-of-fixed-args:wordwidth-1><rest-arg:1>, same length as Exprs
	Body     Code
	// TODO: Documentation: This should carry the names of locals in the rib
}

func (c *LetValues) String() string {
	return "(let-values (" + stringifyExprs(c.Exprs) + ") " + c.Body.String() + ")"
}

type Lexical struct {
	Levels int
	Offset int
	// TODO: Documentation: This should carry the name of the variable
}

func (c *Lexical) String() string {
	return "(lexical " + strconv.Itoa(c.Levels) + " " + strconv.Itoa(c.Offset) + ")"
}

// SelectCase = union { SelectSend, SelectReceive, SelectDefault }
type SelectCase any

type SelectSend struct {
	ChanLevels int
	ChanOffset int
	ValLevels  int
	ValOffset  int
	Body       Code
}

type SelectReceive struct {
	ChanLevels int
	ChanOffset int
	Body       Code
}

type SelectDefault struct {
	Body Code
}

type Select struct {
	cases []SelectCase
}

func (c *Select) String() string {
	// TODO: fix this
	return "select"
}

type Setlex struct {
	Levels int
	Offset int
	Rhs    Code
	// TODO: Documentation: This should carry the name of the variable
}

func (c *Setlex) String() string {
	return "(setlex " + strconv.Itoa(c.Levels) + " " + strconv.Itoa(c.Offset) + " " + c.Rhs.String() + ")"
}

type Global struct {
	Name *Symbol
}

func (c *Global) String() string {
	return "(global " + c.Name.Name + ")"
}

type Setglobal struct {
	Name *Symbol
	Rhs  Code
}

func (c *Setglobal) String() string {
	return "(setglobal " + c.Name.Name + " " + c.Rhs.String() + ")"
}
