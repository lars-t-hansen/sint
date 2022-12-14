
// Generated from runtime/control.sch
package runtime
import (
	. "sint/core"
	"math/big"
)
// Make sure the imports are used, or the Go compiler barfs.
var _ Val = big.NewInt(0)
func initControl(c *Scheme) {
code1 := 
&Setglobal{Name:c.Intern("eval"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Call{Exprs:[]Code{
&Call{Exprs:[]Code{
&Global{Name:c.Intern("sint:compile-toplevel-phrase")},
&Lexical{Levels:0, Offset:0},
}},
}},
Name:"eval",
Signature:&Cons{Car:c.Intern("x"), Cdr:c.NullVal}}}
_, unwcode1 := c.EvalToplevel(code1)
if unwcode1 != nil { panic(unwcode1.String()) }
code2 := 
&Setglobal{Name:c.Intern("apply"), Rhs:&Letrec{Exprs:[]Code{
&Lambda{
Fixed:2, Rest:false,
Body:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Lexical{Levels:0, Offset:1},
}},
Consequent:&Begin{Exprs:[]Code{
&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("not")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("list?")},
&Lexical{Levels:0, Offset:0},
}},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("error")},
&Quote{Value:&Str{Value:"apply: expected list: "}},
&Lexical{Levels:0, Offset:0},
}},
Alternate:&Quote{Value:c.UnspecifiedVal},
},
&Lexical{Levels:0, Offset:0},
}},
Alternate:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:1},
}},
}},
Consequent:&Begin{Exprs:[]Code{
&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("not")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("list?")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:1},
}},
}},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("error")},
&Quote{Value:&Str{Value:"apply: expected list: "}},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:1},
}},
}},
Alternate:&Quote{Value:c.UnspecifiedVal},
},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cons")},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:1},
}},
}},
}},
Alternate:&Let{Exprs:[]Code{
&Call{Exprs:[]Code{
&Global{Name:c.Intern("reverse")},
&Lexical{Levels:0, Offset:1},
}},
}, Body:&Begin{Exprs:[]Code{
&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("not")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("list?")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:0},
}},
}},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("error")},
&Quote{Value:&Str{Value:"apply: expected list: "}},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:0},
}},
}},
Alternate:&Quote{Value:c.UnspecifiedVal},
},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cons")},
&Global{Name:c.Intern("x")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("reverse-append")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:0},
}},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:0},
}},
}},
}},
}}},
},
},
Name:"apply > build-apply-args",
Signature:&Cons{Car:c.Intern("fst"), Cdr:&Cons{Car:c.Intern("rest"), Cdr:c.NullVal}}},
}, Body:&Lambda{
Fixed:2, Rest:true,
Body:&Begin{Exprs:[]Code{
&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("not")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("procedure?")},
&Lexical{Levels:0, Offset:0},
}},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("error")},
&Quote{Value:&Str{Value:"apply: expected procedure"}},
}},
Alternate:&Quote{Value:c.UnspecifiedVal},
},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("sint:apply")},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Lexical{Levels:0, Offset:1},
&Lexical{Levels:0, Offset:2},
}},
}},
}},
Name:"apply",
Signature:&Cons{Car:c.Intern("fn"), Cdr:&Cons{Car:c.Intern("x"), Cdr:c.Intern("rest")}}}}}
_, unwcode2 := c.EvalToplevel(code2)
if unwcode2 != nil { panic(unwcode2.String()) }
code3 := 
&Setglobal{Name:c.Intern("call-with-values"), Rhs:&Lambda{
Fixed:2, Rest:false,
Body:&Begin{Exprs:[]Code{
&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("not")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("procedure?")},
&Lexical{Levels:0, Offset:0},
}},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("error")},
&Quote{Value:&Str{Value:"call-with-values: expected procedure: "}},
&Lexical{Levels:0, Offset:0},
}},
Alternate:&Quote{Value:c.UnspecifiedVal},
},
&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("not")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("procedure?")},
&Lexical{Levels:0, Offset:1},
}},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("error")},
&Quote{Value:&Str{Value:"call-with-values: expected procedure: "}},
&Lexical{Levels:0, Offset:1},
}},
Alternate:&Quote{Value:c.UnspecifiedVal},
},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("sint:apply")},
&Lexical{Levels:0, Offset:1},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("sint:receive-values")},
&Lexical{Levels:0, Offset:0},
}},
}},
}},
Name:"call-with-values",
Signature:&Cons{Car:c.Intern("thunk"), Cdr:&Cons{Car:c.Intern("receiver"), Cdr:c.NullVal}}}}
_, unwcode3 := c.EvalToplevel(code3)
if unwcode3 != nil { panic(unwcode3.String()) }
code4 := 
&Setglobal{Name:c.Intern("make-parameter"), Rhs:&Lambda{
Fixed:1, Rest:true,
Body:&Let{Exprs:[]Code{
&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Lexical{Levels:0, Offset:1},
}},
Consequent:&Lambda{
Fixed:1, Rest:false,
Body:&Lexical{Levels:0, Offset:0},
Name:"make-parameter >  > conv",
Signature:&Cons{Car:c.Intern("x"), Cdr:c.NullVal}},
Alternate:&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:1},
}},
},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("sint:new-tls-key")},
}},
}, Body:&Begin{Exprs:[]Code{
&Call{Exprs:[]Code{
&Global{Name:c.Intern("sint:write-tls-value")},
&Lexical{Levels:0, Offset:1},
&Call{Exprs:[]Code{
&Lexical{Levels:0, Offset:0},
&Lexical{Levels:1, Offset:0},
}},
}},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("sint:make-parameter-function")},
&Lexical{Levels:0, Offset:1},
&Lexical{Levels:0, Offset:0},
}},
}}},
Name:"make-parameter",
Signature:&Cons{Car:c.Intern("init"), Cdr:c.Intern("rest")}}}
_, unwcode4 := c.EvalToplevel(code4)
if unwcode4 != nil { panic(unwcode4.String()) }
code5 := 
&Setglobal{Name:c.Intern("sint:make-parameter-function"), Rhs:&Lambda{
Fixed:2, Rest:false,
Body:&Lambda{
Fixed:0, Rest:true,
Body:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Lexical{Levels:0, Offset:0},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("sint:read-tls-value")},
&Lexical{Levels:1, Offset:0},
}},
Alternate:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:0},
}},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("sint:write-tls-value")},
&Lexical{Levels:1, Offset:0},
&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:1},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:0},
}},
}},
}},
Alternate:&Call{Exprs:[]Code{
&Global{Name:c.Intern("error")},
&Quote{Value:&Str{Value:"Invalid call to parameter function"}},
}},
},
},
Name:"sint:make-parameter-function > ",
Signature:c.Intern("args")},
Name:"sint:make-parameter-function",
Signature:&Cons{Car:c.Intern("key"), Cdr:&Cons{Car:c.Intern("conv"), Cdr:c.NullVal}}}}
_, unwcode5 := c.EvalToplevel(code5)
if unwcode5 != nil { panic(unwcode5.String()) }
code6 := 
&Setglobal{Name:c.Intern("call-with-current-continuation"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Let{Exprs:[]Code{
&Call{Exprs:[]Code{
&Global{Name:c.Intern("goroutine-id")},
}},
&Quote{Value:c.FalseVal},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cons")},
&Quote{Value:c.Intern("call/cc")},
&Quote{Value:c.NullVal},
}},
}, Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("sint:call-with-unwind-handler")},
&Lexical{Levels:0, Offset:2},
&Lambda{
Fixed:0, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("call-with-values")},
&Lambda{
Fixed:0, Rest:false,
Body:&Call{Exprs:[]Code{
&Lexical{Levels:3, Offset:0},
&Lambda{
Fixed:0, Rest:true,
Body:&Begin{Exprs:[]Code{
&If{
Test:&Lexical{Levels:3, Offset:1},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("error")},
&Quote{Value:&Str{Value:"call-with-current-continuation: already returned"}},
}},
Alternate:&Quote{Value:c.UnspecifiedVal},
},
&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("not")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("=")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("goroutine-id")},
}},
&Lexical{Levels:3, Offset:0},
}},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("error")},
&Quote{Value:&Str{Value:"call-with-current-continuation: different goroutine"}},
}},
Alternate:&Quote{Value:c.UnspecifiedVal},
},
&Setlex{Levels:3, Offset:1, Rhs:&Quote{Value:c.TrueVal}},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("sint:unwind")},
&Lexical{Levels:3, Offset:2},
&Lexical{Levels:0, Offset:0},
}},
}},
Name:"call-with-current-continuation >  > [lambda] > [lambda]",
Signature:c.Intern("vs")},
}},
Name:"call-with-current-continuation >  > [lambda]",
Signature:c.NullVal},
&Lambda{
Fixed:0, Rest:true,
Body:&Begin{Exprs:[]Code{
&Setlex{Levels:2, Offset:1, Rhs:&Quote{Value:c.TrueVal}},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("apply")},
&Global{Name:c.Intern("values")},
&Lexical{Levels:0, Offset:0},
}},
}},
Name:"call-with-current-continuation >  > [lambda]",
Signature:c.Intern("vs")},
}},
Name:"call-with-current-continuation > ",
Signature:c.NullVal},
&Lambda{
Fixed:2, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("apply")},
&Global{Name:c.Intern("values")},
&Lexical{Levels:0, Offset:1},
}},
Name:"call-with-current-continuation > ",
Signature:&Cons{Car:c.Intern("k"), Cdr:&Cons{Car:c.Intern("vs"), Cdr:c.NullVal}}},
}}},
Name:"call-with-current-continuation",
Signature:&Cons{Car:c.Intern("proc"), Cdr:c.NullVal}}}
_, unwcode6 := c.EvalToplevel(code6)
if unwcode6 != nil { panic(unwcode6.String()) }
code7 := 
&Setglobal{Name:c.Intern("call/cc"), Rhs:&Global{Name:c.Intern("call-with-current-continuation")}}
_, unwcode7 := c.EvalToplevel(code7)
if unwcode7 != nil { panic(unwcode7.String()) }
code8 := 
&Setglobal{Name:c.Intern("dynamic-wind"), Rhs:&Lambda{
Fixed:3, Rest:false,
Body:&Begin{Exprs:[]Code{
&Call{Exprs:[]Code{
&Lexical{Levels:0, Offset:0},
}},
&Let{Exprs:[]Code{
&Quote{Value:c.FalseVal},
}, Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("sint:call-with-unwind-handler")},
&Quote{Value:c.FalseVal},
&Lambda{
Fixed:0, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("call-with-values")},
&Lexical{Levels:2, Offset:1},
&Lambda{
Fixed:0, Rest:true,
Body:&Begin{Exprs:[]Code{
&Setlex{Levels:2, Offset:0, Rhs:&Quote{Value:c.TrueVal}},
&Call{Exprs:[]Code{
&Lexical{Levels:3, Offset:2},
}},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("apply")},
&Global{Name:c.Intern("values")},
&Lexical{Levels:0, Offset:0},
}},
}},
Name:"dynamic-wind >  > [lambda]",
Signature:c.Intern("vs")},
}},
Name:"dynamic-wind > ",
Signature:c.NullVal},
&Lambda{
Fixed:2, Rest:false,
Body:&Begin{Exprs:[]Code{
&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("not")},
&Lexical{Levels:1, Offset:0},
}},
Consequent:&Call{Exprs:[]Code{
&Lexical{Levels:2, Offset:2},
}},
Alternate:&Quote{Value:c.UnspecifiedVal},
},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("sint:unwind")},
&Lexical{Levels:0, Offset:0},
&Lexical{Levels:0, Offset:1},
}},
}},
Name:"dynamic-wind > ",
Signature:&Cons{Car:c.Intern("key"), Cdr:&Cons{Car:c.Intern("vs"), Cdr:c.NullVal}}},
}}},
}},
Name:"dynamic-wind",
Signature:&Cons{Car:c.Intern("before"), Cdr:&Cons{Car:c.Intern("during"), Cdr:&Cons{Car:c.Intern("after"), Cdr:c.NullVal}}}}}
_, unwcode8 := c.EvalToplevel(code8)
if unwcode8 != nil { panic(unwcode8.String()) }
code9 := 
&Setglobal{Name:c.Intern("map"), Rhs:&Letrec{Exprs:[]Code{
&Lambda{
Fixed:2, Rest:false,
Body:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Lexical{Levels:0, Offset:1},
}},
Consequent:&Quote{Value:c.NullVal},
Alternate:&Call{Exprs:[]Code{
&Global{Name:c.Intern("cons")},
&Call{Exprs:[]Code{
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:1},
}},
}},
&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:1},
}},
}},
}},
},
Name:"map > map1",
Signature:&Cons{Car:c.Intern("fn"), Cdr:&Cons{Car:c.Intern("l0"), Cdr:c.NullVal}}},
&Lambda{
Fixed:3, Rest:false,
Body:&If{
Test:&Let{Exprs:[]Code{
&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Lexical{Levels:0, Offset:1},
}},
}, Body:&If{
Test:&Lexical{Levels:0, Offset:0},
Consequent:&Lexical{Levels:0, Offset:0},
Alternate:&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Lexical{Levels:1, Offset:2},
}},
}},
Consequent:&Quote{Value:c.NullVal},
Alternate:&Call{Exprs:[]Code{
&Global{Name:c.Intern("cons")},
&Call{Exprs:[]Code{
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:1},
}},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:2},
}},
}},
&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:1},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:1},
}},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:2},
}},
}},
}},
},
Name:"map > map2",
Signature:&Cons{Car:c.Intern("fn"), Cdr:&Cons{Car:c.Intern("l0"), Cdr:&Cons{Car:c.Intern("l1"), Cdr:c.NullVal}}}},
&Lambda{
Fixed:2, Rest:false,
Body:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("some?")},
&Global{Name:c.Intern("null?")},
&Lexical{Levels:0, Offset:1},
}},
Consequent:&Quote{Value:c.NullVal},
Alternate:&Call{Exprs:[]Code{
&Global{Name:c.Intern("cons")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("apply")},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:1},
}},
}},
&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:2},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:1},
}},
}},
}},
},
Name:"map > mapn",
Signature:&Cons{Car:c.Intern("fn"), Cdr:&Cons{Car:c.Intern("ls"), Cdr:c.NullVal}}},
}, Body:&Lambda{
Fixed:2, Rest:true,
Body:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Lexical{Levels:0, Offset:2},
}},
Consequent:&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Lexical{Levels:0, Offset:0},
&Lexical{Levels:0, Offset:1},
}},
Alternate:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:2},
}},
}},
Consequent:&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:1},
&Lexical{Levels:0, Offset:0},
&Lexical{Levels:0, Offset:1},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:2},
}},
}},
Alternate:&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:2},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cons")},
&Lexical{Levels:0, Offset:1},
&Lexical{Levels:0, Offset:2},
}},
}},
},
},
Name:"map",
Signature:&Cons{Car:c.Intern("fn"), Cdr:&Cons{Car:c.Intern("l0"), Cdr:c.Intern("rest")}}}}}
_, unwcode9 := c.EvalToplevel(code9)
if unwcode9 != nil { panic(unwcode9.String()) }
code10 := 
&Setglobal{Name:c.Intern("for-each"), Rhs:&Letrec{Exprs:[]Code{
&Lambda{
Fixed:2, Rest:false,
Body:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Lexical{Levels:0, Offset:1},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("unspecified")},
}},
Alternate:&Begin{Exprs:[]Code{
&Call{Exprs:[]Code{
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:1},
}},
}},
&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:1},
}},
}},
}},
},
Name:"for-each > each1",
Signature:&Cons{Car:c.Intern("fn"), Cdr:&Cons{Car:c.Intern("l0"), Cdr:c.NullVal}}},
&Lambda{
Fixed:3, Rest:false,
Body:&If{
Test:&Let{Exprs:[]Code{
&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Lexical{Levels:0, Offset:1},
}},
}, Body:&If{
Test:&Lexical{Levels:0, Offset:0},
Consequent:&Lexical{Levels:0, Offset:0},
Alternate:&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Lexical{Levels:1, Offset:2},
}},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("unspecified")},
}},
Alternate:&Begin{Exprs:[]Code{
&Call{Exprs:[]Code{
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:1},
}},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:2},
}},
}},
&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:1},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:1},
}},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:2},
}},
}},
}},
},
Name:"for-each > each2",
Signature:&Cons{Car:c.Intern("fn"), Cdr:&Cons{Car:c.Intern("l0"), Cdr:&Cons{Car:c.Intern("l1"), Cdr:c.NullVal}}}},
&Lambda{
Fixed:2, Rest:false,
Body:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("some?")},
&Global{Name:c.Intern("null?")},
&Lexical{Levels:0, Offset:1},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("unspecified")},
}},
Alternate:&Begin{Exprs:[]Code{
&Call{Exprs:[]Code{
&Global{Name:c.Intern("apply")},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:1},
}},
}},
&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:2},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:1},
}},
}},
}},
},
Name:"for-each > eachn",
Signature:&Cons{Car:c.Intern("fn"), Cdr:&Cons{Car:c.Intern("ls"), Cdr:c.NullVal}}},
}, Body:&Lambda{
Fixed:2, Rest:true,
Body:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Lexical{Levels:0, Offset:2},
}},
Consequent:&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Lexical{Levels:0, Offset:0},
&Lexical{Levels:0, Offset:1},
}},
Alternate:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:2},
}},
}},
Consequent:&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:1},
&Lexical{Levels:0, Offset:0},
&Lexical{Levels:0, Offset:1},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:2},
}},
}},
Alternate:&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:2},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cons")},
&Lexical{Levels:0, Offset:1},
&Lexical{Levels:0, Offset:2},
}},
}},
},
},
Name:"for-each",
Signature:&Cons{Car:c.Intern("fn"), Cdr:&Cons{Car:c.Intern("l0"), Cdr:c.Intern("rest")}}}}}
_, unwcode10 := c.EvalToplevel(code10)
if unwcode10 != nil { panic(unwcode10.String()) }
code11 := 
&Setglobal{Name:c.Intern("every?"), Rhs:&Letrec{Exprs:[]Code{
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
&Global{Name:c.Intern("not")},
&Call{Exprs:[]Code{
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:1},
}},
}},
}},
Consequent:&Quote{Value:c.FalseVal},
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
Name:"every? > loop",
Signature:&Cons{Car:c.Intern("p"), Cdr:&Cons{Car:c.Intern("l"), Cdr:c.NullVal}}},
}, Body:&Lambda{
Fixed:2, Rest:false,
Body:&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Lexical{Levels:0, Offset:0},
&Lexical{Levels:0, Offset:1},
}},
Name:"every?",
Signature:&Cons{Car:c.Intern("p"), Cdr:&Cons{Car:c.Intern("l"), Cdr:c.NullVal}}}}}
_, unwcode11 := c.EvalToplevel(code11)
if unwcode11 != nil { panic(unwcode11.String()) }
code12 := 
&Setglobal{Name:c.Intern("some?"), Rhs:&Letrec{Exprs:[]Code{
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
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:1},
}},
}},
Consequent:&Quote{Value:c.TrueVal},
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
Name:"some? > loop",
Signature:&Cons{Car:c.Intern("p"), Cdr:&Cons{Car:c.Intern("l"), Cdr:c.NullVal}}},
}, Body:&Lambda{
Fixed:2, Rest:false,
Body:&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Lexical{Levels:0, Offset:0},
&Lexical{Levels:0, Offset:1},
}},
Name:"some?",
Signature:&Cons{Car:c.Intern("p"), Cdr:&Cons{Car:c.Intern("l"), Cdr:c.NullVal}}}}}
_, unwcode12 := c.EvalToplevel(code12)
if unwcode12 != nil { panic(unwcode12.String()) }
code13 := 
&Setglobal{Name:c.Intern("filter"), Rhs:&Letrec{Exprs:[]Code{
&Lambda{
Fixed:2, Rest:false,
Body:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Lexical{Levels:0, Offset:1},
}},
Consequent:&Quote{Value:c.NullVal},
Alternate:&If{
Test:&Call{Exprs:[]Code{
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:1},
}},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("cons")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:1},
}},
&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:1},
}},
}},
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
Name:"filter > loop",
Signature:&Cons{Car:c.Intern("p"), Cdr:&Cons{Car:c.Intern("l"), Cdr:c.NullVal}}},
}, Body:&Lambda{
Fixed:2, Rest:false,
Body:&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Lexical{Levels:0, Offset:0},
&Lexical{Levels:0, Offset:1},
}},
Name:"filter",
Signature:&Cons{Car:c.Intern("p"), Cdr:&Cons{Car:c.Intern("l"), Cdr:c.NullVal}}}}}
_, unwcode13 := c.EvalToplevel(code13)
if unwcode13 != nil { panic(unwcode13.String()) }
}
