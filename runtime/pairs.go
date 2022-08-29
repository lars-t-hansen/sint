
// Generated from runtime/pairs.sch
package runtime
import (
	. "sint/core"
	"math/big"
)
func dummyPairs() {
	// Make sure the imports are used, or the Go compiler barfs.
	var _ Val = big.NewInt(0)
}
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
}}}}
c.EvalToplevel(code1)
code2 := 
&Setglobal{Name:c.Intern("cadr"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:0},
}},
}}}}
c.EvalToplevel(code2)
code3 := 
&Setglobal{Name:c.Intern("cdar"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:0},
}},
}}}}
c.EvalToplevel(code3)
code4 := 
&Setglobal{Name:c.Intern("cddr"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:0},
}},
}}}}
c.EvalToplevel(code4)
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
}}}}
c.EvalToplevel(code5)
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
}}}}
c.EvalToplevel(code6)
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
}}}}
c.EvalToplevel(code7)
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
}}}}
c.EvalToplevel(code8)
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
}}}}
c.EvalToplevel(code9)
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
}}}}
c.EvalToplevel(code10)
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
}}}}
c.EvalToplevel(code11)
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
}}}}
c.EvalToplevel(code12)
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
}}}}
c.EvalToplevel(code13)
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
}}}}
c.EvalToplevel(code14)
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
}}}}
c.EvalToplevel(code15)
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
}}}}
c.EvalToplevel(code16)
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
}}}}
c.EvalToplevel(code17)
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
}}}}
c.EvalToplevel(code18)
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
}}}}
c.EvalToplevel(code19)
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
}}}}
c.EvalToplevel(code20)
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
}}}}
c.EvalToplevel(code21)
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
}}}}
c.EvalToplevel(code22)
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
}}}}
c.EvalToplevel(code23)
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
}}}}
c.EvalToplevel(code24)
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
}}}}
c.EvalToplevel(code25)
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
}}}}
c.EvalToplevel(code26)
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
}}}}
c.EvalToplevel(code27)
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
}}}}
c.EvalToplevel(code28)
code29 := 
&Setglobal{Name:c.Intern("list"), Rhs:&Lambda{
Fixed:0, Rest:true,
Body:&Lexical{Levels:0, Offset:0}}}
c.EvalToplevel(code29)
code30 := 
&Setglobal{Name:c.Intern("list?"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Lexical{Levels:0, Offset:0},
}},
Consequent:&Quote{Value:c.TrueVal},
Alternate:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("pair?")},
&Lexical{Levels:0, Offset:0},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("list?")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:0},
}},
}},
Alternate:&Quote{Value:c.FalseVal},
},
}}}
c.EvalToplevel(code30)
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
}},
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
}}}}}
c.EvalToplevel(code31)
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
}},
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
}},
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
}}}}
c.EvalToplevel(code32)
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
}},
}, Body:&Lambda{
Fixed:2, Rest:false,
Body:&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Lexical{Levels:0, Offset:0},
&Lexical{Levels:0, Offset:1},
}}}}}
c.EvalToplevel(code33)
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
}}}}
c.EvalToplevel(code34)
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
}}}}
c.EvalToplevel(code35)
code36 := 
&Setglobal{Name:c.Intern("list-copy"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("append")},
&Lexical{Levels:0, Offset:0},
&Quote{Value:c.NullVal},
}}}}
c.EvalToplevel(code36)
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
}},
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
}}}}}
c.EvalToplevel(code37)
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
}},
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
}}}}}
c.EvalToplevel(code38)
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
}},
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
}}}}}
c.EvalToplevel(code39)
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
}},
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
}}}}}
c.EvalToplevel(code40)
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
}},
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
}}}}}
c.EvalToplevel(code41)
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
}},
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
}}}}}
c.EvalToplevel(code42)
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
}},
}, Body:&Lambda{
Fixed:1, Rest:false,
Body:&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Lexical{Levels:0, Offset:0},
&Quote{Value:big.NewInt(0)},
}}}}}
c.EvalToplevel(code43)
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
}},
}, Body:&Lambda{
Fixed:1, Rest:false,
Body:&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Lexical{Levels:0, Offset:0},
&Quote{Value:c.NullVal},
}}}}}
c.EvalToplevel(code44)
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
}},
}, Body:&Lambda{
Fixed:2, Rest:false,
Body:&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Lexical{Levels:0, Offset:0},
&Lexical{Levels:0, Offset:1},
}}}}}
c.EvalToplevel(code45)
}
