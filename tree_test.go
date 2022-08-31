// Test cases that run preconstructed program trees through the evaluator.

package main

import (
	"math/big"
	. "sint/core"
	"sint/runtime"
	"testing"
)

func call1(p Code, v0 Code) Code {
	return &Call{Exprs: []Code{p, v0}}
}

func call2(p Code, v0 Code, v1 Code) Code {
	return &Call{Exprs: []Code{p, v0, v1}}
}

func glob(c *Scheme, s string) Code {
	return &Global{Name: c.Intern(s)}
}

func exact(n int64) Code {
	return &Quote{Value: big.NewInt(n)}
}

func TestFibTree(t *testing.T) {
	c := NewScheme(nil)
	runtime.InitPrimitives(c)
	runtime.InitCompiled(c)
	n := &Lexical{Levels: 0, Offset: 0}
	lam := &Lambda{
		Fixed: 1,
		Rest:  false,
		Body: &If{
			Test:       call2(glob(c, "<"), n, exact(2)),
			Consequent: n,
			Alternate: call2(
				glob(c, "+"),
				call1(glob(c, "fib"), call2(glob(c, "-"), n, exact(1))),
				call1(glob(c, "fib"), call2(glob(c, "-"), n, exact(2))))}}
	prog := &Begin{Exprs: []Code{
		&Setglobal{Name: c.Intern("fib"), Rhs: lam},
		call1(glob(c, "fib"), exact(10))}}
	v := c.EvalToplevel(prog)
	if v[0].(*big.Int).Cmp(big.NewInt(55)) != 0 {
		t.Fatal("Wrong answer from fib")
	}
}
