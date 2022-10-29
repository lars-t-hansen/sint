// Given any compiled representation, emit Go code that will recreate it.

package compiler

import (
	"bufio"
	"fmt"
	"math/big"
	. "sint/core"
)

func EmitGo(expr Code, name string, w *bufio.Writer) {
	w.WriteString(name + " := \n")
	emit(expr, w)
	w.WriteString("\n")
}

func emit(expr Code, w *bufio.Writer) {
	switch e := expr.(type) {
	case *Quote:
		w.WriteString("&Quote{Value:")
		emitDatum(e.Value, w)
		w.WriteString("}")
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
	case *Apply:
		// There are no technical reasons why we couldn't just emit it, but
		// it should never appear in code we want to emit.
		panic("The Apply instruction is not supported by the emitter")
	case *Lambda:
		fmt.Fprintf(w, "&Lambda{\nFixed:%d, Rest:%t,\nBody:", e.Fixed, e.Rest)
		emit(e.Body, w)
		if e.Name != "" {
			fmt.Fprintf(w, ",\nName:%q", e.Name)
		}
		if e.Docstring != "" {
			fmt.Fprintf(w, ",\nDocstring:%q", e.Docstring)
		}
		w.WriteString("}")
	case *Let:
		w.WriteString("&Let{Exprs:")
		emitExprs(e.Exprs, w)
		w.WriteString(", Body:")
		emit(e.Body, w)
		w.WriteString("}")
	case *LetStar:
		w.WriteString("&LetStar{Exprs:")
		emitExprs(e.Exprs, w)
		w.WriteString(", Body:")
		emit(e.Body, w)
		w.WriteString("}")
	case *Letrec:
		w.WriteString("&Letrec{Exprs:")
		emitExprs(e.Exprs, w)
		w.WriteString(", Body:")
		emit(e.Body, w)
		w.WriteString("}")
	case *Lexical:
		fmt.Fprintf(w, "&Lexical{Levels:%d, Offset:%d}", e.Levels, e.Offset)
	case *Setlex:
		fmt.Fprintf(w, "&Setlex{Levels:%d, Offset:%d, Rhs:", e.Levels, e.Offset)
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
	w.WriteString("[]Code{\n")
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
	case *Str:
		fmt.Fprintf(w, "&Str{Value:%q}", d.Value)
	case *big.Int:
		if d.IsInt64() {
			fmt.Fprintf(w, "big.NewInt(%d)", d.Int64())
		} else {
			bytes, err := d.GobEncode()
			if err != nil {
				panic(err)
			}
			fmt.Fprintf(w, "c.DecodeInt([]byte{%s})", bytes)
		}
	case *big.Float:
		n, acc := d.Float64()
		if acc == big.Exact {
			fmt.Fprintf(w, "big.NewFloat(%g)", n)
		} else {
			bytes, err := d.GobEncode()
			if err != nil {
				panic(err)
			}
			fmt.Fprintf(w, "c.DecodeFloat([]byte{%s})", bytes)
		}
	case *Char:
		fmt.Fprintf(w, "&Char{Value:%d}", d.Value)
	case *Regexp:
		fmt.Fprintf(w, "&RegExp{Value:c.DecodeRegExp(%q)}", d.Value.String())
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
