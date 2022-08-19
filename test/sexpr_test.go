// Test cases that compile sexprs into programs and then run those programs.

package sint

import (
	"math/big"
	"sint/compiler"
	. "sint/core"
	"sint/runtime"
	"testing"
)

var nullVal Val

func TestFibSexpr(t *testing.T) {
	c := NewScheme()
	nullVal = c.NullVal
	comp := compiler.NewCompiler(c)
	runtime.InitPrimitives(c)
	runtime.InitCompiled(c)
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
	defnProg := comp.CompileToplevel(defn)
	c.EvalToplevel(defnProg)
	invoke := list(symFib, big.NewInt(10))
	invokeProg := comp.CompileToplevel(invoke)
	v := c.EvalToplevel(invokeProg)
	if v.(*big.Int).Cmp(big.NewInt(55)) != 0 {
		t.Fatal("Wrong answer from fib")
	}
}

func list(vs ...Val) Val {
	v := nullVal
	for i := len(vs) - 1; i >= 0; i-- {
		v = &Cons{Car: vs[i], Cdr: v}
	}
	return v
}
