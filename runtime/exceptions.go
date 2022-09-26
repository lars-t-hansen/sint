
// Generated from runtime/exceptions.sch
package runtime
import (
	. "sint/core"
	"math/big"
)
// Make sure the imports are used, or the Go compiler barfs.
var _ Val = big.NewInt(0)
func initExceptions(c *Scheme) {
code1 := 
&Setglobal{Name:c.Intern("error"), Rhs:&Lambda{
Fixed:1, Rest:true,
Body:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("not")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("string?")},
&Lexical{Levels:0, Offset:0},
}},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("error")},
&Quote{Value:&Str{Value:"error: the first argument must be a string"}},
&Lexical{Levels:0, Offset:0},
}},
Alternate:&Call{Exprs:[]Code{
&Global{Name:c.Intern("sint:report-error")},
&Lexical{Levels:0, Offset:0},
&Lexical{Levels:0, Offset:1},
}},
},
Name:"error"}}
_, unwcode1 := c.EvalToplevel(code1)
if unwcode1 != nil { panic(unwcode1.String()) }
}
