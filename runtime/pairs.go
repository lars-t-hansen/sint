
// Generated from runtime/pairs.sch
package runtime
import (
	. "sint/core"
	"math/big"
)
// Make sure the imports are used, or the Go compiler barfs.
var _ Val = big.NewInt(0)
func initPairs(c *Scheme) {
code1 := 
&Setglobal{Name:c.Intern("caar"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:0},
}},
}},
Name:"caar",
Signature:&Cons{Car:c.Intern("x"), Cdr:c.NullVal}}}
_, unwcode1 := c.EvalToplevel(code1)
if unwcode1 != nil { panic(unwcode1.String()) }
code2 := 
&Setglobal{Name:c.Intern("cadr"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:0},
}},
}},
Name:"cadr",
Signature:&Cons{Car:c.Intern("x"), Cdr:c.NullVal}}}
_, unwcode2 := c.EvalToplevel(code2)
if unwcode2 != nil { panic(unwcode2.String()) }
code3 := 
&Setglobal{Name:c.Intern("cdar"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:0},
}},
}},
Name:"cdar",
Signature:&Cons{Car:c.Intern("x"), Cdr:c.NullVal}}}
_, unwcode3 := c.EvalToplevel(code3)
if unwcode3 != nil { panic(unwcode3.String()) }
code4 := 
&Setglobal{Name:c.Intern("cddr"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:0},
}},
}},
Name:"cddr",
Signature:&Cons{Car:c.Intern("x"), Cdr:c.NullVal}}}
_, unwcode4 := c.EvalToplevel(code4)
if unwcode4 != nil { panic(unwcode4.String()) }
code5 := 
&Setglobal{Name:c.Intern("caaar"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:0},
}},
}},
}},
Name:"caaar",
Signature:&Cons{Car:c.Intern("x"), Cdr:c.NullVal}}}
_, unwcode5 := c.EvalToplevel(code5)
if unwcode5 != nil { panic(unwcode5.String()) }
code6 := 
&Setglobal{Name:c.Intern("caadr"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:0},
}},
}},
}},
Name:"caadr",
Signature:&Cons{Car:c.Intern("x"), Cdr:c.NullVal}}}
_, unwcode6 := c.EvalToplevel(code6)
if unwcode6 != nil { panic(unwcode6.String()) }
code7 := 
&Setglobal{Name:c.Intern("cadar"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:0},
}},
}},
}},
Name:"cadar",
Signature:&Cons{Car:c.Intern("x"), Cdr:c.NullVal}}}
_, unwcode7 := c.EvalToplevel(code7)
if unwcode7 != nil { panic(unwcode7.String()) }
code8 := 
&Setglobal{Name:c.Intern("caddr"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:0},
}},
}},
}},
Name:"caddr",
Signature:&Cons{Car:c.Intern("x"), Cdr:c.NullVal}}}
_, unwcode8 := c.EvalToplevel(code8)
if unwcode8 != nil { panic(unwcode8.String()) }
code9 := 
&Setglobal{Name:c.Intern("cdaar"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:0},
}},
}},
}},
Name:"cdaar",
Signature:&Cons{Car:c.Intern("x"), Cdr:c.NullVal}}}
_, unwcode9 := c.EvalToplevel(code9)
if unwcode9 != nil { panic(unwcode9.String()) }
code10 := 
&Setglobal{Name:c.Intern("cdadr"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:0},
}},
}},
}},
Name:"cdadr",
Signature:&Cons{Car:c.Intern("x"), Cdr:c.NullVal}}}
_, unwcode10 := c.EvalToplevel(code10)
if unwcode10 != nil { panic(unwcode10.String()) }
code11 := 
&Setglobal{Name:c.Intern("cddar"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:0},
}},
}},
}},
Name:"cddar",
Signature:&Cons{Car:c.Intern("x"), Cdr:c.NullVal}}}
_, unwcode11 := c.EvalToplevel(code11)
if unwcode11 != nil { panic(unwcode11.String()) }
code12 := 
&Setglobal{Name:c.Intern("cdddr"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:0},
}},
}},
}},
Name:"cdddr",
Signature:&Cons{Car:c.Intern("x"), Cdr:c.NullVal}}}
_, unwcode12 := c.EvalToplevel(code12)
if unwcode12 != nil { panic(unwcode12.String()) }
code13 := 
&Setglobal{Name:c.Intern("caaaar"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:0},
}},
}},
}},
}},
Name:"caaaar",
Signature:&Cons{Car:c.Intern("x"), Cdr:c.NullVal}}}
_, unwcode13 := c.EvalToplevel(code13)
if unwcode13 != nil { panic(unwcode13.String()) }
code14 := 
&Setglobal{Name:c.Intern("caaadr"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:0},
}},
}},
}},
}},
Name:"caaadr",
Signature:&Cons{Car:c.Intern("x"), Cdr:c.NullVal}}}
_, unwcode14 := c.EvalToplevel(code14)
if unwcode14 != nil { panic(unwcode14.String()) }
code15 := 
&Setglobal{Name:c.Intern("caadar"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:0},
}},
}},
}},
}},
Name:"caadar",
Signature:&Cons{Car:c.Intern("x"), Cdr:c.NullVal}}}
_, unwcode15 := c.EvalToplevel(code15)
if unwcode15 != nil { panic(unwcode15.String()) }
code16 := 
&Setglobal{Name:c.Intern("caaddr"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:0},
}},
}},
}},
}},
Name:"caaddr",
Signature:&Cons{Car:c.Intern("x"), Cdr:c.NullVal}}}
_, unwcode16 := c.EvalToplevel(code16)
if unwcode16 != nil { panic(unwcode16.String()) }
code17 := 
&Setglobal{Name:c.Intern("cadaar"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:0},
}},
}},
}},
}},
Name:"cadaar",
Signature:&Cons{Car:c.Intern("x"), Cdr:c.NullVal}}}
_, unwcode17 := c.EvalToplevel(code17)
if unwcode17 != nil { panic(unwcode17.String()) }
code18 := 
&Setglobal{Name:c.Intern("cadadr"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:0},
}},
}},
}},
}},
Name:"cadadr",
Signature:&Cons{Car:c.Intern("x"), Cdr:c.NullVal}}}
_, unwcode18 := c.EvalToplevel(code18)
if unwcode18 != nil { panic(unwcode18.String()) }
code19 := 
&Setglobal{Name:c.Intern("caddar"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:0},
}},
}},
}},
}},
Name:"caddar",
Signature:&Cons{Car:c.Intern("x"), Cdr:c.NullVal}}}
_, unwcode19 := c.EvalToplevel(code19)
if unwcode19 != nil { panic(unwcode19.String()) }
code20 := 
&Setglobal{Name:c.Intern("cadddr"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:0},
}},
}},
}},
}},
Name:"cadddr",
Signature:&Cons{Car:c.Intern("x"), Cdr:c.NullVal}}}
_, unwcode20 := c.EvalToplevel(code20)
if unwcode20 != nil { panic(unwcode20.String()) }
code21 := 
&Setglobal{Name:c.Intern("cdaaar"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:0},
}},
}},
}},
}},
Name:"cdaaar",
Signature:&Cons{Car:c.Intern("x"), Cdr:c.NullVal}}}
_, unwcode21 := c.EvalToplevel(code21)
if unwcode21 != nil { panic(unwcode21.String()) }
code22 := 
&Setglobal{Name:c.Intern("cdaadr"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:0},
}},
}},
}},
}},
Name:"cdaadr",
Signature:&Cons{Car:c.Intern("x"), Cdr:c.NullVal}}}
_, unwcode22 := c.EvalToplevel(code22)
if unwcode22 != nil { panic(unwcode22.String()) }
code23 := 
&Setglobal{Name:c.Intern("cdadar"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:0},
}},
}},
}},
}},
Name:"cdadar",
Signature:&Cons{Car:c.Intern("x"), Cdr:c.NullVal}}}
_, unwcode23 := c.EvalToplevel(code23)
if unwcode23 != nil { panic(unwcode23.String()) }
code24 := 
&Setglobal{Name:c.Intern("cdaddr"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:0},
}},
}},
}},
}},
Name:"cdaddr",
Signature:&Cons{Car:c.Intern("x"), Cdr:c.NullVal}}}
_, unwcode24 := c.EvalToplevel(code24)
if unwcode24 != nil { panic(unwcode24.String()) }
code25 := 
&Setglobal{Name:c.Intern("cddaar"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:0},
}},
}},
}},
}},
Name:"cddaar",
Signature:&Cons{Car:c.Intern("x"), Cdr:c.NullVal}}}
_, unwcode25 := c.EvalToplevel(code25)
if unwcode25 != nil { panic(unwcode25.String()) }
code26 := 
&Setglobal{Name:c.Intern("cddadr"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:0},
}},
}},
}},
}},
Name:"cddadr",
Signature:&Cons{Car:c.Intern("x"), Cdr:c.NullVal}}}
_, unwcode26 := c.EvalToplevel(code26)
if unwcode26 != nil { panic(unwcode26.String()) }
code27 := 
&Setglobal{Name:c.Intern("cdddar"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:0},
}},
}},
}},
}},
Name:"cdddar",
Signature:&Cons{Car:c.Intern("x"), Cdr:c.NullVal}}}
_, unwcode27 := c.EvalToplevel(code27)
if unwcode27 != nil { panic(unwcode27.String()) }
code28 := 
&Setglobal{Name:c.Intern("cddddr"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:0},
}},
}},
}},
}},
Name:"cddddr",
Signature:&Cons{Car:c.Intern("x"), Cdr:c.NullVal}}}
_, unwcode28 := c.EvalToplevel(code28)
if unwcode28 != nil { panic(unwcode28.String()) }
code29 := 
&Setglobal{Name:c.Intern("list"), Rhs:&Lambda{
Fixed:0, Rest:true,
Body:&Lexical{Levels:0, Offset:0},
Name:"list",
Signature:c.Intern("xs")}}
_, unwcode29 := c.EvalToplevel(code29)
if unwcode29 != nil { panic(unwcode29.String()) }
code30 := 
&Setglobal{Name:c.Intern("list?"), Rhs:&Letrec{Exprs:[]Code{
&Lambda{
Fixed:2, Rest:false,
Body:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Lexical{Levels:0, Offset:1},
}},
Consequent:&Quote{Value:c.TrueVal},
Alternate:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("eq?")},
&Lexical{Levels:0, Offset:0},
&Lexical{Levels:0, Offset:1},
}},
Consequent:&Quote{Value:c.FalseVal},
Alternate:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("not")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("pair?")},
&Lexical{Levels:0, Offset:1},
}},
}},
Consequent:&Quote{Value:c.FalseVal},
Alternate:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:1},
}},
}},
Consequent:&Quote{Value:c.TrueVal},
Alternate:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("not")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("pair?")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:1},
}},
}},
}},
Consequent:&Quote{Value:c.FalseVal},
Alternate:&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:0},
}},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cddr")},
&Lexical{Levels:0, Offset:1},
}},
}},
},
},
},
},
},
Name:"list? > loop",
Signature:&Cons{Car:c.Intern("slow"), Cdr:&Cons{Car:c.Intern("fast"), Cdr:c.NullVal}}},
}, Body:&Lambda{
Fixed:1, Rest:false,
Body:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Lexical{Levels:0, Offset:0},
}},
Consequent:&Quote{Value:c.TrueVal},
Alternate:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("not")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("pair?")},
&Lexical{Levels:0, Offset:0},
}},
}},
Consequent:&Quote{Value:c.FalseVal},
Alternate:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:0},
}},
}},
Consequent:&Quote{Value:c.TrueVal},
Alternate:&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:0},
}},
}},
},
},
},
Name:"list?",
Signature:&Cons{Car:c.Intern("x"), Cdr:c.NullVal}}}}
_, unwcode30 := c.EvalToplevel(code30)
if unwcode30 != nil { panic(unwcode30.String()) }
code31 := 
&Setglobal{Name:c.Intern("make-list"), Rhs:&Letrec{Exprs:[]Code{
&Lambda{
Fixed:3, Rest:false,
Body:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("<=")},
&Lexical{Levels:0, Offset:0},
&Quote{Value:big.NewInt(0)},
}},
Consequent:&Lexical{Levels:0, Offset:2},
Alternate:&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("-")},
&Lexical{Levels:0, Offset:0},
&Quote{Value:big.NewInt(1)},
}},
&Lexical{Levels:0, Offset:1},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cons")},
&Lexical{Levels:0, Offset:1},
&Lexical{Levels:0, Offset:2},
}},
}},
},
Name:"make-list > loop",
Signature:&Cons{Car:c.Intern("k"), Cdr:&Cons{Car:c.Intern("fill"), Cdr:&Cons{Car:c.Intern("l"), Cdr:c.NullVal}}}},
}, Body:&Lambda{
Fixed:1, Rest:true,
Body:&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Lexical{Levels:0, Offset:0},
&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Lexical{Levels:0, Offset:1},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("unspecified")},
}},
Alternate:&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:1},
}},
},
&Quote{Value:c.NullVal},
}},
Name:"make-list",
Signature:&Cons{Car:c.Intern("k"), Cdr:c.Intern("rest")}}}}
_, unwcode31 := c.EvalToplevel(code31)
if unwcode31 != nil { panic(unwcode31.String()) }
code32 := 
&Setglobal{Name:c.Intern("append"), Rhs:&Letrec{Exprs:[]Code{
&Lambda{
Fixed:2, Rest:false,
Body:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Lexical{Levels:0, Offset:1},
}},
Consequent:&Lexical{Levels:0, Offset:0},
Alternate:&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:1},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("reverse")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:1},
}},
}},
&Lexical{Levels:0, Offset:0},
}},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:1},
}},
}},
},
Name:"append > loop",
Signature:&Cons{Car:c.Intern("result"), Cdr:&Cons{Car:c.Intern("lists"), Cdr:c.NullVal}}},
&Lambda{
Fixed:2, Rest:false,
Body:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Lexical{Levels:0, Offset:0},
}},
Consequent:&Lexical{Levels:0, Offset:1},
Alternate:&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:1},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:0},
}},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cons")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:0},
}},
&Lexical{Levels:0, Offset:1},
}},
}},
},
Name:"append > loop2",
Signature:&Cons{Car:c.Intern("rev"), Cdr:&Cons{Car:c.Intern("result"), Cdr:c.NullVal}}},
}, Body:&Lambda{
Fixed:0, Rest:true,
Body:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Lexical{Levels:0, Offset:0},
}},
Consequent:&Quote{Value:c.NullVal},
Alternate:&Let{Exprs:[]Code{
&Call{Exprs:[]Code{
&Global{Name:c.Intern("reverse")},
&Lexical{Levels:0, Offset:0},
}},
}, Body:&Call{Exprs:[]Code{
&Lexical{Levels:2, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:0},
}},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:0},
}},
}}},
},
Name:"append",
Signature:c.Intern("lists")}}}
_, unwcode32 := c.EvalToplevel(code32)
if unwcode32 != nil { panic(unwcode32.String()) }
code33 := 
&Setglobal{Name:c.Intern("list-tail"), Rhs:&Letrec{Exprs:[]Code{
&Lambda{
Fixed:2, Rest:false,
Body:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("<=")},
&Lexical{Levels:0, Offset:1},
&Quote{Value:big.NewInt(0)},
}},
Consequent:&Lexical{Levels:0, Offset:0},
Alternate:&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:0},
}},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("-")},
&Lexical{Levels:0, Offset:1},
&Quote{Value:big.NewInt(1)},
}},
}},
},
Name:"list-tail > loop",
Signature:&Cons{Car:c.Intern("l"), Cdr:&Cons{Car:c.Intern("k"), Cdr:c.NullVal}}},
}, Body:&Lambda{
Fixed:2, Rest:false,
Body:&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Lexical{Levels:0, Offset:0},
&Lexical{Levels:0, Offset:1},
}},
Name:"list-tail",
Signature:&Cons{Car:c.Intern("l"), Cdr:&Cons{Car:c.Intern("k"), Cdr:c.NullVal}}}}}
_, unwcode33 := c.EvalToplevel(code33)
if unwcode33 != nil { panic(unwcode33.String()) }
code34 := 
&Setglobal{Name:c.Intern("list-ref"), Rhs:&Lambda{
Fixed:2, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("list-tail")},
&Lexical{Levels:0, Offset:0},
&Lexical{Levels:0, Offset:1},
}},
}},
Name:"list-ref",
Signature:&Cons{Car:c.Intern("l"), Cdr:&Cons{Car:c.Intern("k"), Cdr:c.NullVal}}}}
_, unwcode34 := c.EvalToplevel(code34)
if unwcode34 != nil { panic(unwcode34.String()) }
code35 := 
&Setglobal{Name:c.Intern("list-set!"), Rhs:&Lambda{
Fixed:3, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("set-car!")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("list-tail")},
&Lexical{Levels:0, Offset:0},
&Lexical{Levels:0, Offset:1},
}},
&Lexical{Levels:0, Offset:2},
}},
Name:"list-set!",
Signature:&Cons{Car:c.Intern("l"), Cdr:&Cons{Car:c.Intern("k"), Cdr:&Cons{Car:c.Intern("v"), Cdr:c.NullVal}}}}}
_, unwcode35 := c.EvalToplevel(code35)
if unwcode35 != nil { panic(unwcode35.String()) }
code36 := 
&Setglobal{Name:c.Intern("list-copy"), Rhs:&Letrec{Exprs:[]Code{
&Lambda{
Fixed:2, Rest:false,
Body:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("not")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("pair?")},
&Lexical{Levels:0, Offset:0},
}},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("set-cdr!")},
&Lexical{Levels:0, Offset:1},
&Lexical{Levels:0, Offset:0},
}},
Alternate:&Let{Exprs:[]Code{
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cons")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:0},
}},
&Quote{Value:c.NullVal},
}},
}, Body:&Begin{Exprs:[]Code{
&Call{Exprs:[]Code{
&Global{Name:c.Intern("set-cdr!")},
&Lexical{Levels:1, Offset:1},
&Lexical{Levels:0, Offset:0},
}},
&Call{Exprs:[]Code{
&Lexical{Levels:2, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:1, Offset:0},
}},
&Lexical{Levels:0, Offset:0},
}},
}}},
},
Name:"list-copy > loop",
Signature:&Cons{Car:c.Intern("l"), Cdr:&Cons{Car:c.Intern("last"), Cdr:c.NullVal}}},
}, Body:&Lambda{
Fixed:1, Rest:false,
Body:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("not")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("pair?")},
&Lexical{Levels:0, Offset:0},
}},
}},
Consequent:&Lexical{Levels:0, Offset:0},
Alternate:&Let{Exprs:[]Code{
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cons")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:0},
}},
&Quote{Value:c.NullVal},
}},
}, Body:&Begin{Exprs:[]Code{
&Call{Exprs:[]Code{
&Lexical{Levels:2, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:1, Offset:0},
}},
&Lexical{Levels:0, Offset:0},
}},
&Lexical{Levels:0, Offset:0},
}}},
},
Name:"list-copy",
Signature:&Cons{Car:c.Intern("l"), Cdr:c.NullVal}}}}
_, unwcode36 := c.EvalToplevel(code36)
if unwcode36 != nil { panic(unwcode36.String()) }
code37 := 
&Setglobal{Name:c.Intern("memq"), Rhs:&Letrec{Exprs:[]Code{
&Lambda{
Fixed:2, Rest:false,
Body:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Lexical{Levels:0, Offset:1},
}},
Consequent:&Quote{Value:c.FalseVal},
Alternate:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("eq?")},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:1},
}},
}},
Consequent:&Lexical{Levels:0, Offset:1},
Alternate:&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:1},
}},
}},
},
},
Name:"memq > loop",
Signature:&Cons{Car:c.Intern("obj"), Cdr:&Cons{Car:c.Intern("list"), Cdr:c.NullVal}}},
}, Body:&Lambda{
Fixed:2, Rest:false,
Body:&Begin{Exprs:[]Code{
&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("not")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("list?")},
&Lexical{Levels:0, Offset:1},
}},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("error")},
&Quote{Value:&Str{Value:"memq: not a list: "}},
&Lexical{Levels:0, Offset:1},
}},
Alternate:&Quote{Value:c.UnspecifiedVal},
},
&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Lexical{Levels:0, Offset:0},
&Lexical{Levels:0, Offset:1},
}},
}},
Name:"memq",
Signature:&Cons{Car:c.Intern("obj"), Cdr:&Cons{Car:c.Intern("list"), Cdr:c.NullVal}}}}}
_, unwcode37 := c.EvalToplevel(code37)
if unwcode37 != nil { panic(unwcode37.String()) }
code38 := 
&Setglobal{Name:c.Intern("memv"), Rhs:&Letrec{Exprs:[]Code{
&Lambda{
Fixed:2, Rest:false,
Body:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Lexical{Levels:0, Offset:1},
}},
Consequent:&Quote{Value:c.FalseVal},
Alternate:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("eqv?")},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:1},
}},
}},
Consequent:&Lexical{Levels:0, Offset:1},
Alternate:&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:1},
}},
}},
},
},
Name:"memv > loop",
Signature:&Cons{Car:c.Intern("obj"), Cdr:&Cons{Car:c.Intern("list"), Cdr:c.NullVal}}},
}, Body:&Lambda{
Fixed:2, Rest:false,
Body:&Begin{Exprs:[]Code{
&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("not")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("list?")},
&Lexical{Levels:0, Offset:1},
}},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("error")},
&Quote{Value:&Str{Value:"memv: not a list: "}},
&Lexical{Levels:0, Offset:1},
}},
Alternate:&Quote{Value:c.UnspecifiedVal},
},
&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Lexical{Levels:0, Offset:0},
&Lexical{Levels:0, Offset:1},
}},
}},
Name:"memv",
Signature:&Cons{Car:c.Intern("obj"), Cdr:&Cons{Car:c.Intern("list"), Cdr:c.NullVal}}}}}
_, unwcode38 := c.EvalToplevel(code38)
if unwcode38 != nil { panic(unwcode38.String()) }
code39 := 
&Setglobal{Name:c.Intern("member"), Rhs:&Letrec{Exprs:[]Code{
&Lambda{
Fixed:3, Rest:false,
Body:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Lexical{Levels:0, Offset:1},
}},
Consequent:&Quote{Value:c.FalseVal},
Alternate:&If{
Test:&Call{Exprs:[]Code{
&Lexical{Levels:0, Offset:2},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:1},
}},
}},
Consequent:&Lexical{Levels:0, Offset:1},
Alternate:&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:1},
}},
&Lexical{Levels:0, Offset:2},
}},
},
},
Name:"member > loop",
Signature:&Cons{Car:c.Intern("obj"), Cdr:&Cons{Car:c.Intern("list"), Cdr:&Cons{Car:c.Intern("same?"), Cdr:c.NullVal}}}},
}, Body:&Lambda{
Fixed:2, Rest:true,
Body:&Begin{Exprs:[]Code{
&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("not")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("list?")},
&Lexical{Levels:0, Offset:1},
}},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("error")},
&Quote{Value:&Str{Value:"member: not a list: "}},
&Lexical{Levels:0, Offset:1},
}},
Alternate:&Quote{Value:c.UnspecifiedVal},
},
&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Lexical{Levels:0, Offset:0},
&Lexical{Levels:0, Offset:1},
&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Lexical{Levels:0, Offset:2},
}},
Consequent:&Global{Name:c.Intern("equal?")},
Alternate:&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:2},
}},
},
}},
}},
Name:"member",
Signature:&Cons{Car:c.Intern("obj"), Cdr:&Cons{Car:c.Intern("list"), Cdr:c.Intern("rest")}}}}}
_, unwcode39 := c.EvalToplevel(code39)
if unwcode39 != nil { panic(unwcode39.String()) }
code40 := 
&Setglobal{Name:c.Intern("assq"), Rhs:&Letrec{Exprs:[]Code{
&Lambda{
Fixed:2, Rest:false,
Body:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Lexical{Levels:0, Offset:1},
}},
Consequent:&Quote{Value:c.FalseVal},
Alternate:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("eq?")},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("caar")},
&Lexical{Levels:0, Offset:1},
}},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:1},
}},
Alternate:&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:1},
}},
}},
},
},
Name:"assq > loop",
Signature:&Cons{Car:c.Intern("obj"), Cdr:&Cons{Car:c.Intern("alist"), Cdr:c.NullVal}}},
}, Body:&Lambda{
Fixed:2, Rest:false,
Body:&Begin{Exprs:[]Code{
&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("not")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("list?")},
&Lexical{Levels:0, Offset:1},
}},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("error")},
&Quote{Value:&Str{Value:"assq: not a list: "}},
&Lexical{Levels:0, Offset:1},
}},
Alternate:&Quote{Value:c.UnspecifiedVal},
},
&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Lexical{Levels:0, Offset:0},
&Lexical{Levels:0, Offset:1},
}},
}},
Name:"assq",
Signature:&Cons{Car:c.Intern("obj"), Cdr:&Cons{Car:c.Intern("alist"), Cdr:c.NullVal}}}}}
_, unwcode40 := c.EvalToplevel(code40)
if unwcode40 != nil { panic(unwcode40.String()) }
code41 := 
&Setglobal{Name:c.Intern("assv"), Rhs:&Letrec{Exprs:[]Code{
&Lambda{
Fixed:2, Rest:false,
Body:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Lexical{Levels:0, Offset:1},
}},
Consequent:&Quote{Value:c.FalseVal},
Alternate:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("eqv?")},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("caar")},
&Lexical{Levels:0, Offset:1},
}},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:1},
}},
Alternate:&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:1},
}},
}},
},
},
Name:"assv > loop",
Signature:&Cons{Car:c.Intern("obj"), Cdr:&Cons{Car:c.Intern("alist"), Cdr:c.NullVal}}},
}, Body:&Lambda{
Fixed:2, Rest:false,
Body:&Begin{Exprs:[]Code{
&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("not")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("list?")},
&Lexical{Levels:0, Offset:1},
}},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("error")},
&Quote{Value:&Str{Value:"assv: not a list: "}},
&Lexical{Levels:0, Offset:1},
}},
Alternate:&Quote{Value:c.UnspecifiedVal},
},
&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Lexical{Levels:0, Offset:0},
&Lexical{Levels:0, Offset:1},
}},
}},
Name:"assv",
Signature:&Cons{Car:c.Intern("obj"), Cdr:&Cons{Car:c.Intern("alist"), Cdr:c.NullVal}}}}}
_, unwcode41 := c.EvalToplevel(code41)
if unwcode41 != nil { panic(unwcode41.String()) }
code42 := 
&Setglobal{Name:c.Intern("assoc"), Rhs:&Letrec{Exprs:[]Code{
&Lambda{
Fixed:3, Rest:false,
Body:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Lexical{Levels:0, Offset:1},
}},
Consequent:&Quote{Value:c.FalseVal},
Alternate:&If{
Test:&Call{Exprs:[]Code{
&Lexical{Levels:0, Offset:2},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("caar")},
&Lexical{Levels:0, Offset:1},
}},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:1},
}},
Alternate:&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:1},
}},
&Lexical{Levels:0, Offset:2},
}},
},
},
Name:"assoc > loop",
Signature:&Cons{Car:c.Intern("obj"), Cdr:&Cons{Car:c.Intern("alist"), Cdr:&Cons{Car:c.Intern("same?"), Cdr:c.NullVal}}}},
}, Body:&Lambda{
Fixed:2, Rest:true,
Body:&Begin{Exprs:[]Code{
&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("not")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("list?")},
&Lexical{Levels:0, Offset:1},
}},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("error")},
&Quote{Value:&Str{Value:"assoc: not a list: "}},
&Lexical{Levels:0, Offset:1},
}},
Alternate:&Quote{Value:c.UnspecifiedVal},
},
&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Lexical{Levels:0, Offset:0},
&Lexical{Levels:0, Offset:1},
&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Lexical{Levels:0, Offset:2},
}},
Consequent:&Global{Name:c.Intern("equal?")},
Alternate:&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:2},
}},
},
}},
}},
Name:"assoc",
Signature:&Cons{Car:c.Intern("obj"), Cdr:&Cons{Car:c.Intern("alist"), Cdr:c.Intern("rest")}}}}}
_, unwcode42 := c.EvalToplevel(code42)
if unwcode42 != nil { panic(unwcode42.String()) }
code43 := 
&Setglobal{Name:c.Intern("length"), Rhs:&Letrec{Exprs:[]Code{
&Lambda{
Fixed:2, Rest:false,
Body:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Lexical{Levels:0, Offset:0},
}},
Consequent:&Lexical{Levels:0, Offset:1},
Alternate:&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:0},
}},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("+")},
&Lexical{Levels:0, Offset:1},
&Quote{Value:big.NewInt(1)},
}},
}},
},
Name:"length > loop",
Signature:&Cons{Car:c.Intern("l"), Cdr:&Cons{Car:c.Intern("k"), Cdr:c.NullVal}}},
}, Body:&Lambda{
Fixed:1, Rest:false,
Body:&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Lexical{Levels:0, Offset:0},
&Quote{Value:big.NewInt(0)},
}},
Name:"length",
Signature:&Cons{Car:c.Intern("l"), Cdr:c.NullVal}}}}
_, unwcode43 := c.EvalToplevel(code43)
if unwcode43 != nil { panic(unwcode43.String()) }
code44 := 
&Setglobal{Name:c.Intern("reverse"), Rhs:&Letrec{Exprs:[]Code{
&Lambda{
Fixed:2, Rest:false,
Body:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Lexical{Levels:0, Offset:0},
}},
Consequent:&Lexical{Levels:0, Offset:1},
Alternate:&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:0},
}},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cons")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:0},
}},
&Lexical{Levels:0, Offset:1},
}},
}},
},
Name:"reverse > loop",
Signature:&Cons{Car:c.Intern("l"), Cdr:&Cons{Car:c.Intern("r"), Cdr:c.NullVal}}},
}, Body:&Lambda{
Fixed:1, Rest:false,
Body:&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Lexical{Levels:0, Offset:0},
&Quote{Value:c.NullVal},
}},
Name:"reverse",
Signature:&Cons{Car:c.Intern("l"), Cdr:c.NullVal}}}}
_, unwcode44 := c.EvalToplevel(code44)
if unwcode44 != nil { panic(unwcode44.String()) }
code45 := 
&Setglobal{Name:c.Intern("reverse-append"), Rhs:&Letrec{Exprs:[]Code{
&Lambda{
Fixed:2, Rest:false,
Body:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Lexical{Levels:0, Offset:0},
}},
Consequent:&Lexical{Levels:0, Offset:1},
Alternate:&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:0},
}},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cons")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:0},
}},
&Lexical{Levels:0, Offset:1},
}},
}},
},
Name:"reverse-append > loop",
Signature:&Cons{Car:c.Intern("xs"), Cdr:&Cons{Car:c.Intern("l"), Cdr:c.NullVal}}},
}, Body:&Lambda{
Fixed:2, Rest:false,
Body:&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Lexical{Levels:0, Offset:0},
&Lexical{Levels:0, Offset:1},
}},
Name:"reverse-append",
Signature:&Cons{Car:c.Intern("xs"), Cdr:&Cons{Car:c.Intern("l"), Cdr:c.NullVal}}}}}
_, unwcode45 := c.EvalToplevel(code45)
if unwcode45 != nil { panic(unwcode45.String()) }
}
