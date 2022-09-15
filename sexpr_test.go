// Test cases that compile sexprs into programs and then run those programs.

package main

import (
	"math/big"
	"sint/compiler"
	. "sint/core"
	"sint/runtime"
	"testing"
)

var nullVal Val

func TestFibSexpr(t *testing.T) {
	c := NewScheme(nil)
	runtime.StandardInitialization(c)
	nullVal = c.NullVal
	comp := compiler.NewCompiler(c.Shared)
	symDefine := c.Intern("define")
	symLess := c.Intern("<")
	symFib := c.Intern("fib")
	symPlus := c.Intern("+")
	symMinus := c.Intern("-")
	symIf := c.Intern("if")
	symN := c.Intern("n")
	defn :=
		list(symDefine, list(symFib, symN),
			list(symIf, list(symLess, symN, big.NewInt(2)),
				/* then */
				symN,
				/* else */
				list(symPlus,
					list(symFib, list(symMinus, symN, big.NewInt(1))),
					list(symFib, list(symMinus, symN, big.NewInt(2))))))
	defnProg, defnErr := comp.CompileToplevel(defn)
	if defnErr != nil {
		panic(defnErr.Error())
	}
	_, unw1 := c.EvalToplevel(defnProg)
	if unw1 != nil {
		panic("Error: " + unw1.String())
	}
	invoke := list(symFib, big.NewInt(10))
	invokeProg, invokeErr := comp.CompileToplevel(invoke)
	if invokeErr != nil {
		panic(invokeErr.Error())
	}
	v, unw2 := c.EvalToplevel(invokeProg)
	if unw2 != nil {
		// FIXME
		panic("Error: " + unw2.String())
	}
	if v[0].(*big.Int).Cmp(big.NewInt(55)) != 0 {
		t.Fatal("Wrong answer from fib")
	}
}

/*
func TestLetSexpr(t *testing.T) {
	c := NewScheme()
	nullVal = c.NullVal
	comp := compiler.NewCompiler(c)
	runtime.InitPrimitives(c)
	runtime.InitCompiled(c)
	symLet := c.Intern("let")
	syma := c.Intern("a")
	symb := c.Intern("b")
	symPlus := c.Intern("+")
	prog := list(symLet, list(list(syma, big.NewInt(10))),
		list(symLet, list(list(symb, big.NewInt(20))),
			list(symPlus, syma, symb)))
	code := comp.CompileToplevel(prog)
	t.Fatal(code.String())
}
*/

func list(vs ...Val) Val {
	v := nullVal
	for i := len(vs) - 1; i >= 0; i-- {
		v = &Cons{Car: vs[i], Cdr: v}
	}
	return v
}
