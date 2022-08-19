// The reader produces an sexpr from an input stream

package runtime

import (
	"bufio"
	. "sint/core"
)

// The input source here is probably a Reader and the basic accessor methods
// are probably ReadRune/UnreadRune on the assumption that the input is UTF8.
// That assumption may be wrong; discuss.
//
// The basic character type is rune.  There is a specific EOF error that we use
// to detect EOF, io.EOF.  Though in principle we may see EOF while there are
// also characters in the buffer, ReadRune appears to return a rune and not
// EOF if those characters supply a whole rune.

func Read(c *Scheme, rdr *bufio.Reader) Val {
	return c.NullVal
}
