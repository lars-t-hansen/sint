package runtime

import (
	"fmt"
	"math/big"
	"os"
	. "sint/core"
)

type Writer interface {
	WriteString(s string) (int, error)
	WriteRune(r rune) (int, error)
}

type StdoutWriter struct {
}

func (w *StdoutWriter) WriteString(s string) (int, error) {
	return os.Stdout.WriteString(s)
}

func (w *StdoutWriter) WriteRune(r rune) (int, error) {
	return os.Stdout.WriteString(string(r))
}

func Write(v Val, quoted bool, w Writer) {
	switch x := v.(type) {
	case *big.Int:
		w.WriteString(x.String())
	case *big.Float:
		w.WriteString(x.String())
	case *Char:
		if quoted {
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
				w.WriteString(fmt.Sprintf("#\\%c", x.Value))
			}
		} else {
			w.WriteRune(x.Value)
		}
	case *Str:
		// Hm, is this always the right syntax?
		if quoted {
			w.WriteString(fmt.Sprintf("%q", x.Value))
		} else {
			w.WriteString(x.Value)
		}
	case *Chan:
		w.WriteString(fmt.Sprintf("#<channel %d>", cap(x.Ch)))
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
		writeList(x, quoted, w)
	default:
		w.WriteString("#<weird>")
	}
}

func writeList(c *Cons, quoted bool, w Writer) {
	w.WriteRune('(')
	for {
		Write(c.Car, quoted, w)
		if _, isNull := c.Cdr.(*Null); isNull {
			break
		}
		next, isCons := c.Cdr.(*Cons)
		if !isCons {
			w.WriteString(" . ")
			Write(c.Cdr, quoted, w)
			break
		}
		w.WriteRune(' ')
		c = next
	}
	w.WriteRune(')')
}
