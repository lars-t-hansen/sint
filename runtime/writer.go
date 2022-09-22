package runtime

import (
	"fmt"
	"math/big"
	. "sint/core"
)

type Writer interface {
	WriteString(s string) (int, error)
	WriteRune(r rune) (int, error)
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
	case *Port:
		s := ""
		// It's not safe to call Flags() here because that may lock the port, and the
		// port may already be locked because we're writing to its output stream.
		flags := x.RacyFlags()
		if (flags & IsTextPort) != 0 {
			s = "textual "
		} else {
			if (flags & IsBinaryPort) == 0 {
				panic("Bad port state")
			}
			s = "binary "
		}
		if (flags & IsInputPort) != 0 {
			if (x.Flags() & IsOutputPort) != 0 {
				s = "input/output port"
			} else {
				s = "input port"
			}
		} else {
			s = "output port"
		}
		if (flags & IsClosedPort) != 0 {
			s = "closed " + s
		}
		if x.Name != "" {
			s = s + " " + x.Name
		}
		w.WriteString("#<" + s + ">")
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
		if x.Lam.Name != "" {
			w.WriteString(fmt.Sprintf("#<procedure %s>", x.Lam.Name))
		} else {
			w.WriteString("#<procedure>")
		}
	case *UnwindPkg:
		w.WriteString("#<unwind-package>")
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
