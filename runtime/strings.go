
// Generated from runtime/strings.sch
package runtime
import (
	. "sint/core"
	"math/big"
)
func dummyStrings() {
	// Make sure the imports are used, or the Go compiler barfs.
	var _ Val = big.NewInt(0)
}
func initStrings(c *Scheme) {
code1 := 
&Setglobal{Name:c.Intern("string=?"), Rhs:&Letrec{Exprs:[]Code{
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
&Global{Name:c.Intern("=")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("sint:string-compare")},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:1},
}},
}},
&Quote{Value:big.NewInt(0)},
}},
Consequent:&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:1},
}},
}},
Alternate:&Quote{Value:c.FalseVal},
},
}},
}, Body:&Lambda{
Fixed:2, Rest:true,
Body:&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cons")},
&Lexical{Levels:0, Offset:1},
&Lexical{Levels:0, Offset:2},
}},
}}}}}
c.EvalToplevel(code1)
code2 := 
&Setglobal{Name:c.Intern("string<?"), Rhs:&Letrec{Exprs:[]Code{
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
&Global{Name:c.Intern("<")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("sint:string-compare")},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:1},
}},
}},
&Quote{Value:big.NewInt(0)},
}},
Consequent:&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:1},
}},
}},
Alternate:&Quote{Value:c.FalseVal},
},
}},
}, Body:&Lambda{
Fixed:2, Rest:true,
Body:&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cons")},
&Lexical{Levels:0, Offset:1},
&Lexical{Levels:0, Offset:2},
}},
}}}}}
c.EvalToplevel(code2)
code3 := 
&Setglobal{Name:c.Intern("string<=?"), Rhs:&Letrec{Exprs:[]Code{
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
&Global{Name:c.Intern("<=")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("sint:string-compare")},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:1},
}},
}},
&Quote{Value:big.NewInt(0)},
}},
Consequent:&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:1},
}},
}},
Alternate:&Quote{Value:c.FalseVal},
},
}},
}, Body:&Lambda{
Fixed:2, Rest:true,
Body:&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cons")},
&Lexical{Levels:0, Offset:1},
&Lexical{Levels:0, Offset:2},
}},
}}}}}
c.EvalToplevel(code3)
code4 := 
&Setglobal{Name:c.Intern("string>?"), Rhs:&Letrec{Exprs:[]Code{
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
&Global{Name:c.Intern(">")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("sint:string-compare")},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:1},
}},
}},
&Quote{Value:big.NewInt(0)},
}},
Consequent:&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:1},
}},
}},
Alternate:&Quote{Value:c.FalseVal},
},
}},
}, Body:&Lambda{
Fixed:2, Rest:true,
Body:&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cons")},
&Lexical{Levels:0, Offset:1},
&Lexical{Levels:0, Offset:2},
}},
}}}}}
c.EvalToplevel(code4)
code5 := 
&Setglobal{Name:c.Intern("string>=?"), Rhs:&Letrec{Exprs:[]Code{
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
&Global{Name:c.Intern(">=")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("sint:string-compare")},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:1},
}},
}},
&Quote{Value:big.NewInt(0)},
}},
Consequent:&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:1},
}},
}},
Alternate:&Quote{Value:c.FalseVal},
},
}},
}, Body:&Lambda{
Fixed:2, Rest:true,
Body:&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cons")},
&Lexical{Levels:0, Offset:1},
&Lexical{Levels:0, Offset:2},
}},
}}}}}
c.EvalToplevel(code5)
code6 := 
&Setglobal{Name:c.Intern("string->list"), Rhs:&Letrec{Exprs:[]Code{
&Lambda{
Fixed:4, Rest:false,
Body:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("=")},
&Lexical{Levels:0, Offset:3},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("string-length")},
&Lexical{Levels:0, Offset:2},
}},
}},
Consequent:&Lexical{Levels:0, Offset:0},
Alternate:&Call{Exprs:[]Code{
&Global{Name:c.Intern("call-with-values")},
&Lambda{
Fixed:0, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("string-ref")},
&Lexical{Levels:1, Offset:2},
&Lexical{Levels:1, Offset:3},
}}},
&Lambda{
Fixed:2, Rest:false,
Body:&Let{Exprs:[]Code{
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cons")},
&Lexical{Levels:0, Offset:0},
&Quote{Value:c.NullVal},
}},
}, Body:&Begin{Exprs:[]Code{
&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Lexical{Levels:2, Offset:0},
}},
Consequent:&Setlex{Levels:2, Offset:0, Rhs:&Lexical{Levels:0, Offset:0}},
Alternate:&Call{Exprs:[]Code{
&Global{Name:c.Intern("set-cdr!")},
&Lexical{Levels:2, Offset:1},
&Lexical{Levels:0, Offset:0},
}},
},
&Call{Exprs:[]Code{
&Lexical{Levels:3, Offset:0},
&Lexical{Levels:2, Offset:0},
&Lexical{Levels:0, Offset:0},
&Lexical{Levels:2, Offset:2},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("+")},
&Lexical{Levels:2, Offset:3},
&Lexical{Levels:1, Offset:1},
}},
}},
}}}},
}},
}},
}, Body:&Lambda{
Fixed:1, Rest:true,
Body:&Begin{Exprs:[]Code{
&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("not")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Lexical{Levels:0, Offset:1},
}},
}},
Consequent:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("not")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:1},
}},
}},
}},
Consequent:&Setlex{Levels:0, Offset:0, Rhs:&Call{Exprs:[]Code{
&Global{Name:c.Intern("substring")},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:1},
}},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cadr")},
&Lexical{Levels:0, Offset:1},
}},
}}},
Alternate:&Setlex{Levels:0, Offset:0, Rhs:&Call{Exprs:[]Code{
&Global{Name:c.Intern("substring")},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:1},
}},
}}},
},
Alternate:&Quote{Value:c.UnspecifiedVal},
},
&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Quote{Value:c.NullVal},
&Quote{Value:c.NullVal},
&Lexical{Levels:0, Offset:0},
&Quote{Value:big.NewInt(0)},
}},
}}}}}
c.EvalToplevel(code6)
code7 := 
&Setglobal{Name:c.Intern("string-copy"), Rhs:&Lambda{
Fixed:1, Rest:true,
Body:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("not")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Lexical{Levels:0, Offset:1},
}},
}},
Consequent:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("not")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:1},
}},
}},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("substring")},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:1},
}},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cadr")},
&Lexical{Levels:0, Offset:1},
}},
}},
Alternate:&Call{Exprs:[]Code{
&Global{Name:c.Intern("substring")},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:1},
}},
}},
},
Alternate:&Call{Exprs:[]Code{
&Global{Name:c.Intern("substring")},
&Lexical{Levels:0, Offset:0},
&Quote{Value:big.NewInt(0)},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("string-length")},
&Lexical{Levels:0, Offset:0},
}},
}},
}}}
c.EvalToplevel(code7)
}
