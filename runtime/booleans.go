
// Generated from runtime/booleans.sch
package runtime
import (
	. "sint/core"
	"math/big"
)
func dummyBooleans() {
	// Make sure the imports are used, or the Go compiler barfs.
	var _ Val = big.NewInt(0)
}
func initBooleans(c *Scheme) {
code1 := 
&Setglobal{Name:c.Intern("boolean?"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Let{Exprs:[]Code{
&Call{Exprs:[]Code{
&Global{Name:c.Intern("eq?")},
&Lexical{Levels:0, Offset:0},
&Quote{Value:c.TrueVal},
}},
}, Body:&If{
Test:&Lexical{Levels:0, Offset:0},
Consequent:&Lexical{Levels:0, Offset:0},
Alternate:&Call{Exprs:[]Code{
&Global{Name:c.Intern("eq?")},
&Lexical{Levels:1, Offset:0},
&Quote{Value:c.FalseVal},
}},
}}}}
_, unwcode1 := c.EvalToplevel(code1)
if unwcode1 != nil { panic(unwcode1.String()) }
code2 := 
&Setglobal{Name:c.Intern("not"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&If{
Test:&Lexical{Levels:0, Offset:0},
Consequent:&Quote{Value:c.FalseVal},
Alternate:&Quote{Value:c.TrueVal},
}}}
_, unwcode2 := c.EvalToplevel(code2)
if unwcode2 != nil { panic(unwcode2.String()) }
code3 := 
&Setglobal{Name:c.Intern("boolean=?"), Rhs:&Letrec{Exprs:[]Code{
&Lambda{
Fixed:1, Rest:false,
Body:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("not")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("boolean?")},
&Lexical{Levels:0, Offset:0},
}},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("error")},
&Quote{Value:&Str{Value:"boolean=?: not a boolean: "}},
&Lexical{Levels:0, Offset:0},
}},
Alternate:&Quote{Value:c.UnspecifiedVal},
}},
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
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:1},
}},
}},
Consequent:&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:1},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:1},
}},
}},
Alternate:&Begin{Exprs:[]Code{
&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:1},
}},
}},
&Quote{Value:c.FalseVal},
}},
},
}},
}, Body:&Lambda{
Fixed:2, Rest:true,
Body:&Begin{Exprs:[]Code{
&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Lexical{Levels:0, Offset:0},
}},
&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Lexical{Levels:0, Offset:1},
}},
&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("eq?")},
&Lexical{Levels:0, Offset:0},
&Lexical{Levels:0, Offset:1},
}},
Consequent:&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:1},
&Lexical{Levels:0, Offset:0},
&Lexical{Levels:0, Offset:2},
}},
Alternate:&Quote{Value:c.FalseVal},
},
}}}}}
_, unwcode3 := c.EvalToplevel(code3)
if unwcode3 != nil { panic(unwcode3.String()) }
}
