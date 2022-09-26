
// Generated from runtime/io.sch
package runtime
import (
	//lint:ignore ST1001 dot import
	. "sint/core"
	"math/big"
)
// Make sure the imports are used, or the Go compiler barfs.
var _ Val = big.NewInt(0)
func initIo(c *Scheme) {
code1 := 
&Setglobal{Name:c.Intern("*current-input-port-key*"), Rhs:&Quote{Value:big.NewInt(1)}}
_, unwcode1 := c.EvalToplevel(code1)
if unwcode1 != nil { panic(unwcode1.String()) }
code2 := 
&Setglobal{Name:c.Intern("*current-output-port-key*"), Rhs:&Quote{Value:big.NewInt(2)}}
_, unwcode2 := c.EvalToplevel(code2)
if unwcode2 != nil { panic(unwcode2.String()) }
code3 := 
&Setglobal{Name:c.Intern("*current-error-port-key*"), Rhs:&Quote{Value:big.NewInt(3)}}
_, unwcode3 := c.EvalToplevel(code3)
if unwcode3 != nil { panic(unwcode3.String()) }
code4 := 
&Setglobal{Name:c.Intern("*input-port-flag*"), Rhs:&Quote{Value:big.NewInt(1)}}
_, unwcode4 := c.EvalToplevel(code4)
if unwcode4 != nil { panic(unwcode4.String()) }
code5 := 
&Setglobal{Name:c.Intern("*output-port-flag*"), Rhs:&Quote{Value:big.NewInt(2)}}
_, unwcode5 := c.EvalToplevel(code5)
if unwcode5 != nil { panic(unwcode5.String()) }
code6 := 
&Setglobal{Name:c.Intern("*textual-port-flag*"), Rhs:&Quote{Value:big.NewInt(4)}}
_, unwcode6 := c.EvalToplevel(code6)
if unwcode6 != nil { panic(unwcode6.String()) }
code7 := 
&Setglobal{Name:c.Intern("*binary-port-flag*"), Rhs:&Quote{Value:big.NewInt(8)}}
_, unwcode7 := c.EvalToplevel(code7)
if unwcode7 != nil { panic(unwcode7.String()) }
code8 := 
&Setglobal{Name:c.Intern("*closed-port-flag*"), Rhs:&Quote{Value:big.NewInt(16)}}
_, unwcode8 := c.EvalToplevel(code8)
if unwcode8 != nil { panic(unwcode8.String()) }
code9 := 
&Setglobal{Name:c.Intern("input-port?"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("port?")},
&Lexical{Levels:0, Offset:0},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("not")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("zero?")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("bitwise-and")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("sint:port-flags")},
&Lexical{Levels:0, Offset:0},
}},
&Global{Name:c.Intern("*input-port-flag*")},
}},
}},
}},
Alternate:&Quote{Value:c.FalseVal},
},
Name:"input-port?"}}
_, unwcode9 := c.EvalToplevel(code9)
if unwcode9 != nil { panic(unwcode9.String()) }
code10 := 
&Setglobal{Name:c.Intern("output-port?"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("port?")},
&Lexical{Levels:0, Offset:0},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("not")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("zero?")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("bitwise-and")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("sint:port-flags")},
&Lexical{Levels:0, Offset:0},
}},
&Global{Name:c.Intern("*output-port-flag*")},
}},
}},
}},
Alternate:&Quote{Value:c.FalseVal},
},
Name:"output-port?"}}
_, unwcode10 := c.EvalToplevel(code10)
if unwcode10 != nil { panic(unwcode10.String()) }
code11 := 
&Setglobal{Name:c.Intern("textual-port?"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("port?")},
&Lexical{Levels:0, Offset:0},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("not")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("zero?")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("bitwise-and")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("sint:port-flags")},
&Lexical{Levels:0, Offset:0},
}},
&Global{Name:c.Intern("*textual-port-flag*")},
}},
}},
}},
Alternate:&Quote{Value:c.FalseVal},
},
Name:"textual-port?"}}
_, unwcode11 := c.EvalToplevel(code11)
if unwcode11 != nil { panic(unwcode11.String()) }
code12 := 
&Setglobal{Name:c.Intern("binary-port?"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("port?")},
&Lexical{Levels:0, Offset:0},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("not")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("zero?")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("bitwise-and")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("sint:port-flags")},
&Lexical{Levels:0, Offset:0},
}},
&Global{Name:c.Intern("*binary-port-flag*")},
}},
}},
}},
Alternate:&Quote{Value:c.FalseVal},
},
Name:"binary-port?"}}
_, unwcode12 := c.EvalToplevel(code12)
if unwcode12 != nil { panic(unwcode12.String()) }
code13 := 
&Setglobal{Name:c.Intern("input-port-open?"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Begin{Exprs:[]Code{
&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("not")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("input-port?")},
&Lexical{Levels:0, Offset:0},
}},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("error")},
&Quote{Value:&Str{Value:"input-port-open?: Not an input port: "}},
&Lexical{Levels:0, Offset:0},
}},
Alternate:&Quote{Value:c.UnspecifiedVal},
},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("zero?")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("bitwise-and")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("sint:port-flags")},
&Lexical{Levels:0, Offset:0},
}},
&Global{Name:c.Intern("*closed-port-flag*")},
}},
}},
}},
Name:"input-port-open?"}}
_, unwcode13 := c.EvalToplevel(code13)
if unwcode13 != nil { panic(unwcode13.String()) }
code14 := 
&Setglobal{Name:c.Intern("output-port-open?"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Begin{Exprs:[]Code{
&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("not")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("output-port?")},
&Lexical{Levels:0, Offset:0},
}},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("error")},
&Quote{Value:&Str{Value:"output-port-open?: Not an output port: "}},
&Lexical{Levels:0, Offset:0},
}},
Alternate:&Quote{Value:c.UnspecifiedVal},
},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("zero?")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("bitwise-and")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("sint:port-flags")},
&Lexical{Levels:0, Offset:0},
}},
&Global{Name:c.Intern("*closed-port-flag*")},
}},
}},
}},
Name:"output-port-open?"}}
_, unwcode14 := c.EvalToplevel(code14)
if unwcode14 != nil { panic(unwcode14.String()) }
code15 := 
&Setglobal{Name:c.Intern("current-input-port"), Rhs:&Call{Exprs:[]Code{
&Global{Name:c.Intern("sint:make-parameter-function")},
&Global{Name:c.Intern("*current-input-port-key*")},
&Lambda{
Fixed:1, Rest:false,
Body:&Begin{Exprs:[]Code{
&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("not")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("input-port?")},
&Lexical{Levels:0, Offset:0},
}},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("error")},
&Quote{Value:&Str{Value:"Cannot set current input port: "}},
&Lexical{Levels:0, Offset:0},
}},
Alternate:&Quote{Value:c.UnspecifiedVal},
},
&Lexical{Levels:0, Offset:0},
}},
Name:"current-input-port"},
}}}
_, unwcode15 := c.EvalToplevel(code15)
if unwcode15 != nil { panic(unwcode15.String()) }
code16 := 
&Setglobal{Name:c.Intern("current-output-port"), Rhs:&Call{Exprs:[]Code{
&Global{Name:c.Intern("sint:make-parameter-function")},
&Global{Name:c.Intern("*current-output-port-key*")},
&Lambda{
Fixed:1, Rest:false,
Body:&Begin{Exprs:[]Code{
&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("not")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("output-port?")},
&Lexical{Levels:0, Offset:0},
}},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("error")},
&Quote{Value:&Str{Value:"Cannot set current output port: "}},
&Lexical{Levels:0, Offset:0},
}},
Alternate:&Quote{Value:c.UnspecifiedVal},
},
&Lexical{Levels:0, Offset:0},
}},
Name:"current-output-port"},
}}}
_, unwcode16 := c.EvalToplevel(code16)
if unwcode16 != nil { panic(unwcode16.String()) }
code17 := 
&Setglobal{Name:c.Intern("current-error-port"), Rhs:&Call{Exprs:[]Code{
&Global{Name:c.Intern("sint:make-parameter-function")},
&Global{Name:c.Intern("*current-error-port-key*")},
&Lambda{
Fixed:1, Rest:false,
Body:&Begin{Exprs:[]Code{
&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("not")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("output-port?")},
&Lexical{Levels:0, Offset:0},
}},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("error")},
&Quote{Value:&Str{Value:"Cannot set current error port: "}},
&Lexical{Levels:0, Offset:0},
}},
Alternate:&Quote{Value:c.UnspecifiedVal},
},
&Lexical{Levels:0, Offset:0},
}},
Name:"current-error-port"},
}}}
_, unwcode17 := c.EvalToplevel(code17)
if unwcode17 != nil { panic(unwcode17.String()) }
code18 := 
&Setglobal{Name:c.Intern("call-with-input-file"), Rhs:&Lambda{
Fixed:2, Rest:false,
Body:&Let{Exprs:[]Code{
&Call{Exprs:[]Code{
&Global{Name:c.Intern("open-input-file")},
&Lexical{Levels:0, Offset:0},
}},
}, Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("dynamic-wind")},
&Lambda{
Fixed:0, Rest:false,
Body:&Quote{Value:c.TrueVal},
Name:"call-with-input-file > "},
&Lambda{
Fixed:0, Rest:false,
Body:&Call{Exprs:[]Code{
&Lexical{Levels:2, Offset:1},
&Lexical{Levels:1, Offset:0},
}},
Name:"call-with-input-file > "},
&Lambda{
Fixed:0, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("close-input-port")},
&Lexical{Levels:1, Offset:0},
}},
Name:"call-with-input-file > "},
}}},
Name:"call-with-input-file"}}
_, unwcode18 := c.EvalToplevel(code18)
if unwcode18 != nil { panic(unwcode18.String()) }
code19 := 
&Setglobal{Name:c.Intern("call-with-output-file"), Rhs:&Lambda{
Fixed:2, Rest:false,
Body:&Let{Exprs:[]Code{
&Call{Exprs:[]Code{
&Global{Name:c.Intern("open-output-file")},
&Lexical{Levels:0, Offset:0},
}},
}, Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("dynamic-wind")},
&Lambda{
Fixed:0, Rest:false,
Body:&Quote{Value:c.TrueVal},
Name:"call-with-output-file > "},
&Lambda{
Fixed:0, Rest:false,
Body:&Call{Exprs:[]Code{
&Lexical{Levels:2, Offset:1},
&Lexical{Levels:1, Offset:0},
}},
Name:"call-with-output-file > "},
&Lambda{
Fixed:0, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("close-output-port")},
&Lexical{Levels:1, Offset:0},
}},
Name:"call-with-output-file > "},
}}},
Name:"call-with-output-file"}}
_, unwcode19 := c.EvalToplevel(code19)
if unwcode19 != nil { panic(unwcode19.String()) }
code20 := 
&Setglobal{Name:c.Intern("call-with-port"), Rhs:&Lambda{
Fixed:2, Rest:false,
Body:&Begin{Exprs:[]Code{
&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("not")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("port?")},
&Lexical{Levels:0, Offset:0},
}},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("error")},
&Quote{Value:&Str{Value:"call-with-port: Not a port: "}},
&Lexical{Levels:0, Offset:0},
}},
Alternate:&Quote{Value:c.UnspecifiedVal},
},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("dynamic-wind")},
&Lambda{
Fixed:0, Rest:false,
Body:&Quote{Value:c.TrueVal},
Name:"call-with-port > "},
&Lambda{
Fixed:0, Rest:false,
Body:&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:1},
&Lexical{Levels:1, Offset:0},
}},
Name:"call-with-port > "},
&Lambda{
Fixed:0, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("close-port")},
&Lexical{Levels:1, Offset:0},
}},
Name:"call-with-port > "},
}},
}},
Name:"call-with-port"}}
_, unwcode20 := c.EvalToplevel(code20)
if unwcode20 != nil { panic(unwcode20.String()) }
code21 := 
&Setglobal{Name:c.Intern("close-port"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Begin{Exprs:[]Code{
&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("not")},
&Let{Exprs:[]Code{
&Call{Exprs:[]Code{
&Global{Name:c.Intern("input-port?")},
&Lexical{Levels:0, Offset:0},
}},
}, Body:&If{
Test:&Lexical{Levels:0, Offset:0},
Consequent:&Lexical{Levels:0, Offset:0},
Alternate:&Call{Exprs:[]Code{
&Global{Name:c.Intern("output-port?")},
&Lexical{Levels:1, Offset:0},
}},
}},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("error")},
&Quote{Value:&Str{Value:"close-port: Not a port"}},
&Lexical{Levels:0, Offset:0},
}},
Alternate:&Quote{Value:c.UnspecifiedVal},
},
&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("input-port?")},
&Lexical{Levels:0, Offset:0},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("close-input-port")},
&Lexical{Levels:0, Offset:0},
}},
Alternate:&Quote{Value:c.UnspecifiedVal},
},
&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("output-port?")},
&Lexical{Levels:0, Offset:0},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("close-output-port")},
&Lexical{Levels:0, Offset:0},
}},
Alternate:&Quote{Value:c.UnspecifiedVal},
},
}},
Name:"close-port"}}
_, unwcode21 := c.EvalToplevel(code21)
if unwcode21 != nil { panic(unwcode21.String()) }
}
