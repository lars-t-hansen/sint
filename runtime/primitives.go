package runtime

import (
	"math/big"
	"sint/compiler"
	. "sint/core"
)

func addPrimitive(c *Scheme, name string, fixed int, rest bool, primop func(*Scheme, []Val) Val) {
	sym := c.Intern(name)
	sym.Value = &Procedure{Lam: &Lambda{Fixed: fixed, Rest: rest, Body: nil}, Env: nil, Primop: primop}
}

func InitPrimitives(c *Scheme) {
	// R7R7 6.1, Equivalence predicates, also see equivalence.sch
	addPrimitive(c, "eq?", 2, false, primEqp)
	addPrimitive(c, "eqv?", 2, false, primEqvp)

	// R7RS 6.2, Numbers, also see numbers.sch
	// TODO: /
	// TODO: quotient
	// TODO: remainder
	// TODO: exact
	// TODO: inexact
	// TODO: (other numerics as required)
	addPrimitive(c, "sint:inexact-float?", 1, false, primInexactFloatp)
	addPrimitive(c, "sint:exact-integer?", 1, false, primExactIntegerp)
	addPrimitive(c, "+", 0, true, primAdd)
	addPrimitive(c, "-", 1, true, primSub)
	addPrimitive(c, "*", 0, true, primMul)
	addPrimitive(c, "<", 2, true, primLess)
	addPrimitive(c, "<=", 2, true, primLessOrEqual)
	addPrimitive(c, "=", 2, true, primEqual)
	addPrimitive(c, ">", 2, true, primGreater)
	addPrimitive(c, ">=", 2, true, primGreaterOrEqual)

	// R7RS 6.3, Booleans, see booleans.sch

	// R7RS 6.4, Pairs and lists, also see pairs.sch
	addPrimitive(c, "null?", 1, false, primNullp)
	addPrimitive(c, "pair?", 1, false, primPairp)
	addPrimitive(c, "cons", 2, false, primCons)
	addPrimitive(c, "car", 1, false, primCar)
	addPrimitive(c, "cdr", 1, false, primCdr)
	addPrimitive(c, "set-car!", 2, false, primSetcar)
	addPrimitive(c, "set-cdr!", 2, false, primSetcdr)

	// R7RS 6.5, Symbols
	// TODO: symbol->string
	// TODO: string->symbol
	// TODO: symbol=? which is not the same as eq? and might be defined in scheme
	//       in terms of string=?
	addPrimitive(c, "symbol?", 1, false, primSymbolp)
	addPrimitive(c, "gensym", 0, false, primGensym)

	// R7RS 6.6, Characters
	// TODO: char=?
	// TODO: char>?
	// TODO: char>=?
	// TODO: char<?
	// TODO: char<=?
	// TODO: (and probably many others)
	addPrimitive(c, "char?", 1, false, primCharp)
	addPrimitive(c, "char->integer", 1, false, primChar2Int)
	addPrimitive(c, "integer->char", 1, false, primInt2Char)

	// R7RS 6.7, Strings

	// R7RS 6.10, Control features, also see control.sch
	addPrimitive(c, "procedure?", 1, false, primProcedurep)

	// R7RS 6.13, Input and output, also see io.sch
	addPrimitive(c, "eof-object?", 1, false, primEofObjectp)

	// See runtime/control.sch.  This treats its argument as a top-level program form
	// and returns a thunk that evaluates that form.
	addPrimitive(c, "sint:compile-toplevel-phrase", 1, false, primCompileToplevel)

	// See runtime/control.sch.  This is a one-instruction procedure with the signature (fn l count)
	// where the `fn` must be a procedure and `l` must appear to be a list up to at least `count` elements.
	// It applies `fn` to the `count` first elements of `l` in a properly tail-recursive manner.
	// The values are not arguments to the instruction but are taken from the environment, lexical offsets
	// 0, 1, and 2 at relative level 0.
	sym := c.Intern("sint:raw-apply")
	sym.Value = &Procedure{Lam: &Lambda{Fixed: 3, Rest: false, Body: &Apply{}}, Env: nil, Primop: nil}
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

func primCharp(ctx *Scheme, args []Val) Val {
	if _, ok := args[0].(*Char); ok {
		return ctx.TrueVal
	}
	return ctx.FalseVal
}

func primEofObjectp(ctx *Scheme, args []Val) Val {
	if _, ok := args[0].(*EofObject); ok {
		return ctx.TrueVal
	}
	return ctx.FalseVal
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

func primInexactFloatp(ctx *Scheme, args []Val) Val {
	if _, ok := args[0].(*big.Float); ok {
		return ctx.TrueVal
	}
	return ctx.FalseVal
}

func primExactIntegerp(ctx *Scheme, args []Val) Val {
	if _, ok := args[0].(*big.Int); ok {
		return ctx.TrueVal
	}
	return ctx.FalseVal
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

func primMul(c *Scheme, args []Val) Val {
	if len(args) == 0 {
		return big.NewInt(1)
	}
	if len(args) == 1 {
		return checkNumber(args[0], "*")
	}
	r := mul2(args[0], args[1])
	for _, v := range args[2:] {
		r = mul2(r, v)
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

func primLessOrEqual(c *Scheme, args []Val) Val {
	for i := 1; i < len(args); i++ {
		if cmp2(args[i-1], args[i], "<=") == 1 {
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

func primGreater(c *Scheme, args []Val) Val {
	for i := 1; i < len(args); i++ {
		if cmp2(args[i-1], args[i], ">") != 1 {
			return c.FalseVal
		}
	}
	return c.TrueVal
}

func primGreaterOrEqual(c *Scheme, args []Val) Val {
	for i := 1; i < len(args); i++ {
		if cmp2(args[i-1], args[i], ">=") != -1 {
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

func mul2(a Val, b Val) Val {
	if ia, ib, ok := bothInt(a, b); ok {
		var z big.Int
		z.Mul(ia, ib)
		return &z
	}
	fa, fb := bothFloat(a, b, "*")
	var z big.Float
	z.Mul(fa, fb)
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
	if !isNumber(v) {
		panic("'" + s + ": Not a number: " + v.String())
	}
	return v
}

func isNumber(v Val) bool {
	if _, ok := v.(*big.Int); ok {
		return true
	}
	if _, ok := v.(*big.Float); ok {
		return true
	}
	return false
}

func primCompileToplevel(c *Scheme, args []Val) Val {
	// Compiles args[0] into a lambda and then creates a toplevel procedure
	// from that lambda, and returns the procedure
	// TODO: The compiler is stateless and thread-safe and can be cached on the engine
	comp := compiler.NewCompiler(c)
	prog := comp.CompileToplevel(args[0])
	return &Procedure{Lam: &Lambda{Fixed: 0, Rest: false, Body: prog}, Env: nil, Primop: nil}
}

func primGensym(c *Scheme, _ []Val) Val {
	return c.Gensym("S")
}

func primChar2Int(c *Scheme, args []Val) Val {
	if ch, ok := args[0].(*Char); ok {
		return big.NewInt(int64(ch.Value))
	}
	panic("char->integer: Not a character: " + args[0].String())
}

func primInt2Char(c *Scheme, args []Val) Val {
	if n, ok := args[0].(*big.Int); ok {
		if !n.IsInt64() {
			panic("char->integer: Integer outside character range: " + args[0].String())
		}
		k := n.Int64()
		// TODO: Is this right?
		if k < 0 || k > 0xDFFF {
			panic("char->integer: Integer outside character range: " + args[0].String())
		}
		return &Char{Value: rune(n.Int64())}
	}
	panic("char->integer: Not an exact integer: " + args[0].String())
}
