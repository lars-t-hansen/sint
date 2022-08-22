package runtime

import (
	"math/big"
	. "sint/core"
)

func addPrimitive(c *Scheme, name string, fixed int, rest bool, primop func(*Scheme, []Val) Val) {
	sym := c.Intern(name)
	sym.Value = &Procedure{Lam: &Lambda{Fixed: fixed, Rest: rest, Body: nil}, Env: nil, Primop: primop}
}

// These are primitives that do not invoke procedures.  All primitives that do invoke
// procedures must be supported differently, to ensure proper tail recursion.  Not sure
// how to do that yet, but the most obvious way is to have hand-written procedures that
// invoke special instructions that do the right thing.  Then the evaluator will insert
// the correct glue.  apply is obvious, as in Larceny.  eval I'm not sure about.  it is
// a function that compiles an expression and then runs it.  Probably hte compiled
// expression is delivered in the form of a thunk.  Then the thunk can be invoked in the
// manner of a tail call.  So eval is user code that uses a primitive to create the thunk,
// and then just invokes it.

func InitPrimitives(c *Scheme) {
	// TODO: These could go in a table, it doesn't have to be code

	// Fundamental predicates
	addPrimitive(c, "null?", 1, false, primNullp)
	addPrimitive(c, "pair?", 1, false, primPairp)
	addPrimitive(c, "symbol?", 1, false, primSymbolp)
	addPrimitive(c, "number?", 1, false, primNumberp)
	addPrimitive(c, "procedure?", 1, false, primProcedurep)

	// Pairs and lists
	addPrimitive(c, "cons", 2, false, primCons)
	addPrimitive(c, "car", 1, false, primCar)
	addPrimitive(c, "cdr", 1, false, primCdr)
	addPrimitive(c, "set-car!", 2, false, primSetcar)
	addPrimitive(c, "set-cdr!", 2, false, primSetcdr)

	// Numbers
	addPrimitive(c, "real?", 1, false, primRealp)
	addPrimitive(c, "integer?", 1, false, primIntegerp)
	addPrimitive(c, "+", 0, true, primAdd)
	addPrimitive(c, "-", 1, true, primSub)
	addPrimitive(c, "<", 2, true, primLess)
	addPrimitive(c, "=", 2, true, primEqual)

	// Sundry
	addPrimitive(c, "compile-toplevel-phrase", 1, false, primCompileToplevel)

	// eqv?
	// eq?
	// *
	// /
	// quotient
	// (other numerics as required)
	// <=
	// >
	// >=
	// exact->inexact
	// inexact->exact
	// string?
	// symbol->string
	// string->symbol
	// (Anything to do with characters, which we don't have yet but must have)
	// (Many string functions, ditto)

	// eval + apply are NOT here, see comments in the evaluator
}

func primNullp(ctx *Scheme, args []Val) Val {
	if args[0] == ctx.NullVal {
		return ctx.TrueVal
	}
	return ctx.FalseVal
}

func primPairp(ctx *Scheme, args []Val) Val {
	if _, ok := args[0].(*Cons); ok {
		return ctx.TrueVal
	}
	return ctx.FalseVal
}

func primSymbolp(ctx *Scheme, args []Val) Val {
	if _, ok := args[0].(*Symbol); ok {
		return ctx.TrueVal
	}
	return ctx.FalseVal
}

func primProcedurep(ctx *Scheme, args []Val) Val {
	if _, ok := args[0].(*Symbol); ok {
		return ctx.TrueVal
	}
	return ctx.FalseVal
}

func primNumberp(ctx *Scheme, args []Val) Val {
	switch args[0].(type) {
	case *big.Int:
		return ctx.TrueVal
	case *big.Float:
		return ctx.TrueVal
	default:
		return ctx.FalseVal
	}
}

func primRealp(ctx *Scheme, args []Val) Val {
	switch args[0].(type) {
	case *big.Int:
		return ctx.TrueVal
	case *big.Float:
		return ctx.TrueVal
	default:
		panic("real?: not a number: " + args[0].String())
	}
}

func primIntegerp(ctx *Scheme, args []Val) Val {
	switch args[0].(type) {
	case *big.Int:
		return ctx.TrueVal
	case *big.Float:
		return ctx.FalseVal
	default:
		panic("integer?: not a number: " + args[0].String())
	}
}

func checkCons(v Val, fn string) *Cons {
	if c, ok := v.(*Cons); ok {
		return c
	}
	panic(fn + ": Not a pair: " + v.String())
}

func primCons(_ *Scheme, args []Val) Val {
	return &Cons{Car: args[0], Cdr: args[1]}
}

func primCar(_ *Scheme, args []Val) Val {
	return checkCons(args[0], "car").Car
}

func primCdr(_ *Scheme, args []Val) Val {
	return checkCons(args[0], "cdr").Cdr
}

func primSetcar(ctx *Scheme, args []Val) Val {
	checkCons(args[0], "set-car!").Car = args[1]
	return ctx.UnspecifiedVal
}

func primSetcdr(ctx *Scheme, args []Val) Val {
	checkCons(args[0], "set-cdr!").Cdr = args[1]
	return ctx.UnspecifiedVal
}

func primAdd(c *Scheme, args []Val) Val {
	if len(args) == 0 {
		return c.Zero
	}
	if len(args) == 1 {
		return checkNumber(args[0], "+")
	}
	r := add2(args[0], args[1])
	for _, v := range args[2:] {
		r = add2(r, v)
	}
	return r
}

func primSub(_ *Scheme, args []Val) Val {
	if len(args) == 1 {
		switch v := args[0].(type) {
		case *big.Int:
			var r big.Int
			r.Neg(v)
			return &r
		case *big.Float:
			var r big.Float
			r.Neg(v)
			return &r
		default:
			panic("'-': Not a number: " + args[0].String())
		}
	}
	r := sub2(args[0], args[1])
	for _, v := range args[2:] {
		r = sub2(r, v)
	}
	return r
}

func primLess(c *Scheme, args []Val) Val {
	for i := 1; i < len(args); i++ {
		if cmp2(args[i-1], args[i], "<") != -1 {
			return c.FalseVal
		}
	}
	return c.TrueVal
}

func primEqual(c *Scheme, args []Val) Val {
	for i := 1; i < len(args); i++ {
		if cmp2(args[i-1], args[i], "=") != 0 {
			return c.FalseVal
		}
	}
	return c.TrueVal
}

func add2(a Val, b Val) Val {
	if ia, ib, ok := bothInt(a, b); ok {
		var z big.Int
		z.Add(ia, ib)
		return &z
	}
	fa, fb := bothFloat(a, b, "+")
	var z big.Float
	z.Add(fa, fb)
	return &z
}

func sub2(a Val, b Val) Val {
	if ia, ib, ok := bothInt(a, b); ok {
		var z big.Int
		z.Sub(ia, ib)
		return &z
	}
	fa, fb := bothFloat(a, b, "+")
	var z big.Float
	z.Sub(fa, fb)
	return &z
}

func cmp2(a Val, b Val, name string) int {
	if ia, ib, ok := bothInt(a, b); ok {
		return ia.Cmp(ib)
	}
	fa, fb := bothFloat(a, b, name)
	return fa.Cmp(fb)
}

func bothInt(a Val, b Val) (*big.Int, *big.Int, bool) {
	if ia, ok := a.(*big.Int); ok {
		if ib, ok := b.(*big.Int); ok {
			return ia, ib, true
		}
	}
	return nil, nil, false
}

// Coerce both values to float and return them
func bothFloat(a Val, b Val, name string) (*big.Float, *big.Float) {
	if fa, ok := a.(*big.Float); ok {
		if fb, ok := b.(*big.Float); ok {
			return fa, fb
		}
		if ib, ok := b.(*big.Int); ok {
			var fb big.Float
			fb.SetInt(ib)
			return fa, &fb
		}
		panic("'" + name + "': Not a number: " + b.String())
	}
	if ia, ok := a.(*big.Int); ok {
		var fa big.Float
		fa.SetInt(ia)
		if fb, ok := b.(*big.Float); ok {
			return &fa, fb
		}
		if ib, ok := b.(*big.Int); ok {
			var fb big.Float
			fb.SetInt(ib)
			return &fa, &fb
		}
		panic("'" + name + "': Not a number: " + b.String())
	}
	panic("'" + name + "': Not a number: " + a.String())
}

func checkNumber(v Val, s string) Val {
	if _, ok := v.(*big.Int); ok {
		return v
	}
	if _, ok := v.(*big.Float); ok {
		return v
	}
	panic("'" + s + ": Not a number: " + v.String())
}

func primCompileToplevel(c *Scheme, args []Val) Val {
	// Compiles args[0] into a lambda and then creates a toplevel procedure
	// from that lambda, and returns the procedure
	panic("compileToplevel not implemented yet")
}
