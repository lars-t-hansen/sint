
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
&Quote{Value:c.Intern("sint-0")},
&Global{Name:c.Intern(".1")},
}}}}
c.EvalToplevel(code1)
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
&Global{Name:c.Intern("eval")},
&Lexical{Levels:0, Offset:0},
}},
&Call{Exprs:[]Code{
&Lexical{Levels:2, Offset:0},
&Lexical{Levels:1, Offset:0},
}},
}},
Alternate:&Quote{Value:c.UnspecifiedVal},
}}},
}, Body:&Lambda{
Fixed:1, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("call-with-input-file")},
&Lexical{Levels:0, Offset:0},
&Lambda{
Fixed:1, Rest:false,
Body:&Call{Exprs:[]Code{
&Lexical{Levels:2, Offset:0},
&Lexical{Levels:0, Offset:0},
}}},
}}}}}
c.EvalToplevel(code2)
}
