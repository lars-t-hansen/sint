// Given any compiled representation, emit Go code that will recreate it.

package compiler

import (
	"bufio"
	"fmt"
	"math/big"
	. "sint/core"
)

func EmitGo(expr Code, name string, w *bufio.Writer) {
	w.WriteString("var " + name + " Code = \n")
	emit(expr, w)
	w.WriteString("\n")
}

func emit(expr Code, w *bufio.Writer) {
	switch e := expr.(type) {
	case *Quote:
		emitDatum(e.Value, w)
	case *If:
		w.WriteString("&If{\nTest:")
		emit(e.Test, w)
		w.WriteString(",\nConsequent:")
		emit(e.Consequent, w)
		w.WriteString(",\nAlternate:")
		emit(e.Alternate, w)
		w.WriteString(",\n}")
	case *Begin:
		w.WriteString("&Begin{Exprs:")
		emitExprs(e.Exprs, w)
		w.WriteString("}")
	case *Call:
		w.WriteString("&Call{Exprs:")
		emitExprs(e.Exprs, w)
		w.WriteString("}")
	case *Lambda:
		w.WriteString(fmt.Sprintf("&Lambda{Fixed:%d, Rest:%t, Body:", e.Fixed, e.Rest))
		emit(e.Body, w)
		w.WriteString("}")
	case *Let:
		w.WriteString("&Let{Bindings:")
		emitExprs(e.Exprs, w)
		w.WriteString(", Body:")
		emit(e.Body, w)
		w.WriteString("}")
	case *Letrec:
		w.WriteString("&Letrec{Bindings:")
		emitExprs(e.Exprs, w)
		w.WriteString(", Body:")
		emit(e.Body, w)
		w.WriteString("}")
	case *Lexical:
		w.WriteString(fmt.Sprintf("&Lexical{Levels:%d, Offset:%d}", e.Levels, e.Offset))
	case *Setlex:
		w.WriteString(fmt.Sprintf("&Setlex{Levels:%d, Offset:%d, Rhs:", e.Levels, e.Offset))
		emit(e.Rhs, w)
		w.WriteString("}")
	case *Global:
		w.WriteString("&Global{Name:")
		emitSymbol(e.Name, w)
		w.WriteString("}")
	case *Setglobal:
		w.WriteString("&Setglobal{Name:")
		emitSymbol(e.Name, w)
		w.WriteString(", Rhs:")
		emit(e.Rhs, w)
		w.WriteString("}")
	default:
		panic("Bad expression: " + expr.String())
	}

}

func emitExprs(es []Code, w *bufio.Writer) {
	w.WriteString("&Code[]{")
	for _, e := range es {
		emit(e, w)
		w.WriteString(",\n")
	}
	w.WriteString("}")
}

func emitDatum(v Val, w *bufio.Writer) {
	switch d := v.(type) {
	case *Undefined:
		w.WriteString("c.UndefinedVal")
	case *Unspecified:
		w.WriteString("c.UnspecifiedVal")
	case *Null:
		w.WriteString("c.NullVal")
	case *EofObject:
		w.WriteString("c.EofVal")
	case *True:
		w.WriteString("c.TrueVal")
	case *False:
		w.WriteString("c.FalseVal")
	case *Symbol:
		emitSymbol(d, w)
	case *big.Int:
		// FIXME: Different depending on whether the value is int64 or not
		w.WriteString("...")
	case *big.Float:
		// FIXME: Different depending on whether the value is float64 or not
		w.WriteString("...")
	case *Cons:
		w.WriteString("&Cons{Car:")
		emitDatum(d.Car, w)
		w.WriteString(", Cdr:")
		emitDatum(d.Cdr, w)
		w.WriteString("}")
	default:
		panic("Unknown datum type")
	}
}

func emitSymbol(s *Symbol, w *bufio.Writer) {
	w.WriteString(fmt.Sprintf("c.Intern(\"%s\")", s.Name))
}
