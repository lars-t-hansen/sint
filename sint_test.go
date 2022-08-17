package sint

import (
	"math/big"
	"testing"
)

func call1(p Code, v0 Code) Code {
	return &Call{[]Code{p, v0}}
}

func call2(p Code, v0 Code, v1 Code) Code {
	return &Call{[]Code{p, v0, v1}}
}

func glob(c *Scheme, s string) Code {
	return &Global{c.intern(s)}
}

func exact(n int64) Code {
	return &Quote{big.NewInt(n)}
}

func TestFib(t *testing.T) {
	c := NewScheme()
	n := &Lexical{0, 0}
	lam := &Lambda{
		1,
		false,
		&If{call2(glob(c, "<"), n, exact(2)),
			n,
			call2(
				glob(c, "+"),
				call1(glob(c, "fib"), call2(glob(c, "-"), n, exact(1))),
				call1(glob(c, "fib"), call2(glob(c, "-"), n, exact(2))))}}
	prog := &Begin{[]Code{
		&Setglobal{c.intern("fib"), lam},
		call1(glob(c, "fib"), exact(10))}}
	v := c.eval(prog, nil)
	if v.(*big.Int).Cmp(big.NewInt(55)) != 0 {
		t.Fatal("Wrong answer from fib")
	}
	t.Log(v)
}
