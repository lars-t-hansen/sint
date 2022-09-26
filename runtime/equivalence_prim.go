// Equivalence primitive procedures.
//
// R7R7 6.1, Equivalence predicates, also see equivalence.sch

package runtime

import (
	"math/big"
	. "sint/core"
)

func initEquivalencePrimitives(ctx *Scheme) {
	addPrimitive(ctx, "eq?", 2, false, primEqp)
	addPrimitive(ctx, "eqv?", 2, false, primEqvp)
}

func primEqp(ctx *Scheme, a0, a1 Val, _ []Val) (Val, int) {
	if a0 == a1 {
		return ctx.TrueVal, 1
	}
	return ctx.FalseVal, 1
}

func primEqvp(ctx *Scheme, a0, a1 Val, _ []Val) (Val, int) {
	if a0 == a1 {
		return ctx.TrueVal, 1
	}
	if n1, ok := a0.(*big.Int); ok {
		if n2, ok := a1.(*big.Int); ok {
			if n1.Cmp(n2) == 0 {
				return ctx.TrueVal, 1
			}
		}
		return ctx.FalseVal, 1
	}
	if n1, ok := a0.(*big.Float); ok {
		if n2, ok := a1.(*big.Float); ok {
			// TODO: Some fine points here around NaN?
			if n1.Cmp(n2) == 0 {
				return ctx.TrueVal, 1
			}
		}
		return ctx.FalseVal, 1
	}
	if c1, ok := a0.(*Char); ok {
		if c2, ok := a1.(*Char); ok {
			if c1.Value == c2.Value {
				return ctx.TrueVal, 1
			}
		}
		return ctx.FalseVal, 1
	}
	return ctx.FalseVal, 1
}
