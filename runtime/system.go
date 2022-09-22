
// Generated from runtime/system.sch
package runtime
import (
	. "sint/core"
	"math/big"
)
func dummySystem() {
	// Make sure the imports are used, or the Go compiler barfs.
	var _ Val = big.NewInt(0)
}
func initSystem(c *Scheme) {
code1 := 
&Setglobal{Name:c.Intern("features"), Rhs:&Lambda{
Fixed:0, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("list")},
&Quote{Value:c.Intern("sint")},
&Quote{Value:c.Intern("sint-0.1")},
}},
Name:"features"}}
_, unwcode1 := c.EvalToplevel(code1)
if unwcode1 != nil { panic(unwcode1.String()) }
code2 := 
&Setglobal{Name:c.Intern("load"), Rhs:&Letrec{Exprs:[]Code{
&Lambda{
Fixed:1, Rest:false,
Body:&Let{Exprs:[]Code{
&Call{Exprs:[]Code{
&Global{Name:c.Intern("read")},
&Lexical{Levels:0, Offset:0},
}},
}, Body:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("not")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("eof-object?")},
&Lexical{Levels:0, Offset:0},
}},
}},
Consequent:&Begin{Exprs:[]Code{
&Call{Exprs:[]Code{
&Global{Name:c.Intern("call-with-values")},
&Lambda{
Fixed:0, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("eval")},
&Lexical{Levels:1, Offset:0},
}},
Name:"load > loop > [lambda]"},
&Lambda{
Fixed:0, Rest:true,
Body:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("not")},
&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("=")},
&Quote{Value:big.NewInt(1)},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("length")},
&Lexical{Levels:0, Offset:0},
}},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("eq?")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("unspecified")},
}},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:0},
}},
}},
Alternate:&Quote{Value:c.FalseVal},
},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("for-each")},
&Lambda{
Fixed:1, Rest:false,
Body:&Begin{Exprs:[]Code{
&Call{Exprs:[]Code{
&Global{Name:c.Intern("display")},
&Lexical{Levels:0, Offset:0},
}},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("newline")},
}},
}},
Name:"load > loop > [lambda] > [lambda]"},
&Lexical{Levels:0, Offset:0},
}},
Alternate:&Quote{Value:c.UnspecifiedVal},
},
Name:"load > loop > [lambda]"},
}},
&Call{Exprs:[]Code{
&Lexical{Levels:2, Offset:0},
&Lexical{Levels:1, Offset:0},
}},
}},
Alternate:&Quote{Value:c.UnspecifiedVal},
}},
Name:"load > loop"},
}, Body:&Lambda{
Fixed:1, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("call-with-input-file")},
&Lexical{Levels:0, Offset:0},
&Lexical{Levels:1, Offset:0},
}},
Name:"load"}}}
_, unwcode2 := c.EvalToplevel(code2)
if unwcode2 != nil { panic(unwcode2.String()) }
}
