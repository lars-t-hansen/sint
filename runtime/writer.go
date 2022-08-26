package runtime

import (
	"bufio"
	"fmt"
	"math/big"
	. "sint/core"
)

func Write(v Val, w *bufio.Writer) {
	switch x := v.(type) {
	case *big.Int:
		w.WriteString(x.String())
	case *big.Float:
		w.WriteString(x.String())
	case *Char:
		switch x.Value {
		case ' ':
			w.WriteString("#\\space")
		case '\t':
			w.WriteString("#\\tab")
		case '\n':
			w.WriteString("#\\newline")
		case '\r':
			w.WriteString("#\\return")
		default:
			// Hm, maybe check if it's printable?
			fmt.Fprintf(w, "#\\%c", x.Value)
		}
	case *Str:
		// Hm, is this always the right syntax?
		fmt.Fprintf(w, "%q", x.Value)
	case *True:
		w.WriteString("#t")
	case *False:
		w.WriteString("#f")
	case *Unspecified:
		w.WriteString("#!unspecified")
	case *Undefined:
		w.WriteString("#!undefined")
	case *EofObject:
		w.WriteString("#!eof")
	case *Null:
		w.WriteString("()")
	case *Procedure:
		w.WriteString("#<procedure>")
	case *Symbol:
		w.WriteString(x.Name)
	case *Cons:
		writeList(x, w)
	default:
		w.WriteString("#<weird>")
	}
}

func writeList(c *Cons, w *bufio.Writer) {
	w.WriteRune('(')
	for {
		Write(c.Car, w)
		if _, isNull := c.Cdr.(*Null); isNull {
			break
		}
		next, isCons := c.Cdr.(*Cons)
		if !isCons {
			w.WriteString(" . ")
			Write(c.Cdr, w)
			break
		}
		w.WriteRune(' ')
		c = next
	}
	w.WriteRune(')')
}
