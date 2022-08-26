package runtime

import (
	"math/big"
	. "sint/core"
)

// R7R7 6.1, Equivalence predicates, also see equivalence.sch
func initEquivalencePrimitives(c *Scheme) {
	addPrimitive(c, "eq?", 2, false, primEqp)
	addPrimitive(c, "eqv?", 2, false, primEqvp)
}

func primEqp(ctx *Scheme, args []Val) Val {
	if args[0] == args[1] {
		return ctx.TrueVal
	}
	return ctx.FalseVal
}

func primEqvp(ctx *Scheme, args []Val) Val {
	if args[0] == args[1] {
		return ctx.TrueVal
	}
	if n1, ok := args[0].(*big.Int); ok {
		if n2, ok := args[1].(*big.Int); ok {
			if n1.Cmp(n2) == 0 {
				return ctx.TrueVal
			}
		}
		return ctx.FalseVal
	}
	if n1, ok := args[0].(*big.Float); ok {
		if n2, ok := args[1].(*big.Float); ok {
			// TODO: Some fine points here around NaN?
			if n1.Cmp(n2) == 0 {
				return ctx.TrueVal
			}
		}
		return ctx.FalseVal
	}
	if c1, ok := args[0].(*Char); ok {
		if c2, ok := args[1].(*Char); ok {
			if c1.Value == c2.Value {
				return ctx.TrueVal
			}
		}
		return ctx.FalseVal
	}
	return ctx.FalseVal
}
