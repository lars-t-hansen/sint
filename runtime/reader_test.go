package runtime

import (
	"math/big"
	. "sint/core"
	"strconv"
	"strings"
	"testing"
)

// None of these tests are whitebox; they could be moved out of the package,
// to avoid polluting it.

func TestReader(t *testing.T) {
	c := NewScheme(nil)
	{
		in := strings.NewReader(" \t\n\r")
		expectEOF(c, in, t)
	}
	{
		in := strings.NewReader("abc")
		expectSymbol(c, in, "abc", t)
		expectEOF(c, in, t)
	}
	{
		in := strings.NewReader(" \t\n\rabc \t\n\r")
		expectSymbol(c, in, "abc", t)
		expectEOF(c, in, t)
	}
	{
		in := strings.NewReader("178")
		expectExact(c, in, 178, t)
		expectEOF(c, in, t)
	}
	{
		in := strings.NewReader("17.75")
		expectInexact(c, in, 17.75, t)
		expectEOF(c, in, t)
	}
	// TODO: Lots more
}

func expectEOF(c *Scheme, in InputStream, t *testing.T) {
	v, err := Read(c, in)
	if err != nil {
		t.Fatal(err.Error())
	}
	if v != c.EofVal {
		t.Fatal("Expected EOF object")
	}
}

func expectSymbol(c *Scheme, in InputStream, s string, t *testing.T) {
	v, err := Read(c, in)
	if err != nil {
		t.Fatal(err.Error())
	}
	if sym, ok := v.(*Symbol); ok {
		if sym.Name == s {
			return
		}
	}
	t.Fatal("Expected symbol: " + s)
}

func expectExact(c *Scheme, in InputStream, i int, t *testing.T) {
	v, err := Read(c, in)
	if err != nil {
		t.Fatal(err.Error())
	}
	if n, ok := v.(*big.Int); ok {
		if n.Cmp(big.NewInt(int64(i))) == 0 {
			return
		}
	}
	t.Fatal("Expected exact number: " + strconv.Itoa(i))
}

func expectInexact(c *Scheme, in InputStream, d float64, t *testing.T) {
	v, err := Read(c, in)
	if err != nil {
		t.Fatal(err.Error())
	}
	if n, ok := v.(*big.Float); ok {
		if n.Cmp(big.NewFloat(d)) == 0 {
			return
		}
	}
	t.Fatal("Expected inexact number: " + strconv.FormatFloat(d, 'g', -1, 64))
}
