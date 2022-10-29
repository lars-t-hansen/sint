
// Generated from runtime/equivalence.sch
package runtime
import (
	. "sint/core"
	"math/big"
)
// Make sure the imports are used, or the Go compiler barfs.
var _ Val = big.NewInt(0)
func initEquivalence(c *Scheme) {
code1 := 
&Setglobal{Name:c.Intern("equal?"), Rhs:&Lambda{
Fixed:2, Rest:false,
Body:&Let{Exprs:[]Code{
&Call{Exprs:[]Code{
&Global{Name:c.Intern("eqv?")},
&Lexical{Levels:0, Offset:0},
&Lexical{Levels:0, Offset:1},
}},
}, Body:&If{
Test:&Lexical{Levels:0, Offset:0},
Consequent:&Lexical{Levels:0, Offset:0},
Alternate:&Let{Exprs:[]Code{
&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("pair?")},
&Lexical{Levels:1, Offset:0},
}},
Consequent:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("pair?")},
&Lexical{Levels:1, Offset:1},
}},
Consequent:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("equal?")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:1, Offset:0},
}},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:1, Offset:1},
}},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("equal?")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:1, Offset:0},
}},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:1, Offset:1},
}},
}},
Alternate:&Quote{Value:c.FalseVal},
},
Alternate:&Quote{Value:c.FalseVal},
},
Alternate:&Quote{Value:c.FalseVal},
},
}, Body:&If{
Test:&Lexical{Levels:0, Offset:0},
Consequent:&Lexical{Levels:0, Offset:0},
Alternate:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("string?")},
&Lexical{Levels:2, Offset:0},
}},
Consequent:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("string?")},
&Lexical{Levels:2, Offset:1},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("string=?")},
&Lexical{Levels:2, Offset:0},
&Lexical{Levels:2, Offset:1},
}},
Alternate:&Quote{Value:c.FalseVal},
},
Alternate:&Quote{Value:c.FalseVal},
},
}},
}},
Name:"equal?",
Signature:&Cons{Car:c.Intern("a"), Cdr:&Cons{Car:c.Intern("b"), Cdr:c.NullVal}}}}
_, unwcode1 := c.EvalToplevel(code1)
if unwcode1 != nil { panic(unwcode1.String()) }
}
