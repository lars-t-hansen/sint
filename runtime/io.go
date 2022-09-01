
// Generated from runtime/io.sch
package runtime
import (
	. "sint/core"
	"math/big"
)
func dummyIo() {
	// Make sure the imports are used, or the Go compiler barfs.
	var _ Val = big.NewInt(0)
}
func initIo(c *Scheme) {
code1 := 
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
Body:&Quote{Value:c.TrueVal}},
&Lambda{
Fixed:0, Rest:false,
Body:&Call{Exprs:[]Code{
&Lexical{Levels:2, Offset:1},
&Lexical{Levels:1, Offset:0},
}}},
&Lambda{
Fixed:0, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("close-input-port")},
&Lexical{Levels:1, Offset:0},
}}},
}}}}}
c.EvalToplevel(code1)
code2 := 
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
Body:&Quote{Value:c.TrueVal}},
&Lambda{
Fixed:0, Rest:false,
Body:&Call{Exprs:[]Code{
&Lexical{Levels:2, Offset:1},
&Lexical{Levels:1, Offset:0},
}}},
&Lambda{
Fixed:0, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("close-output-port")},
&Lexical{Levels:1, Offset:0},
}}},
}}}}}
c.EvalToplevel(code2)
}
