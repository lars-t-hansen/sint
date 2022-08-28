
// Generated from runtime/exceptions.sch
package runtime
import (
	. "sint/core"
	"math/big"
)
func dummyExceptions() {
	// Make sure the imports are used, or the Go compiler barfs.
	var _ Val = big.NewInt(0)
}
func initExceptions(c *Scheme) {
code1 := 
&Setglobal{Name:c.Intern("error"), Rhs:&Letrec{Exprs:[]Code{
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
&Global{Name:c.Intern("string-append")},
&Lexical{Levels:0, Offset:1},
&Quote{Value:&Str{Value:" "}},
&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:1},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:0},
}},
}},
}},
}},
}},
&Lambda{
Fixed:1, Rest:false,
Body:&Let{Exprs:[]Code{
&Call{Exprs:[]Code{
&Global{Name:c.Intern("string?")},
&Lexical{Levels:0, Offset:0},
}},
}, Body:&If{
Test:&Lexical{Levels:0, Offset:0},
Consequent:&Lexical{Levels:0, Offset:0},
Alternate:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("number?")},
&Lexical{Levels:1, Offset:0},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("number->string")},
&Lexical{Levels:1, Offset:0},
}},
Alternate:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("symbol?")},
&Lexical{Levels:1, Offset:0},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("symbol->string")},
&Lexical{Levels:1, Offset:0},
}},
Alternate:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("char?")},
&Lexical{Levels:1, Offset:0},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("string")},
&Lexical{Levels:1, Offset:0},
}},
Alternate:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("eq?")},
&Lexical{Levels:1, Offset:0},
&Quote{Value:c.TrueVal},
}},
Consequent:&Quote{Value:&Str{Value:"#t"}},
Alternate:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("eq?")},
&Lexical{Levels:1, Offset:0},
&Quote{Value:c.FalseVal},
}},
Consequent:&Quote{Value:&Str{Value:"#f"}},
Alternate:&Quote{Value:&Str{Value:"#<weird>"}},
},
},
},
},
},
}}},
}, Body:&Lambda{
Fixed:1, Rest:true,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("sint:throw-string")},
&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Lexical{Levels:0, Offset:1},
&Lexical{Levels:0, Offset:0},
}},
}}}}}
c.EvalToplevel(code1)
}
