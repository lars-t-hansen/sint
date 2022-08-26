
// Generated from runtime/symbols.sch
package runtime
import (
	. "sint/core"
)
func initSymbols(c *Scheme) {
code1 := 
&Setglobal{Name:c.Intern("symbol=?"), Rhs:&Letrec{Exprs:[]Code{
&Lambda{
Fixed:1, Rest:false,
Body:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("not")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("symbol?")},
&Lexical{Levels:0, Offset:0},
}},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("error")},
&Quote{Value:&Str{Value:"[string symbol=?: not a symbol: ]"}},
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
Alternate:&Begin{Exprs:[]Code{
&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:1},
}},
}},
&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("string=?")},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("symbol->string")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:1},
}},
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
Alternate:&Quote{Value:c.FalseVal},
},
}},
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
&Let{Exprs:[]Code{
&Call{Exprs:[]Code{
&Global{Name:c.Intern("symbol->string")},
&Lexical{Levels:0, Offset:0},
}},
}, Body:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("string=?")},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("symbol->string")},
&Lexical{Levels:1, Offset:1},
}},
}},
Consequent:&Call{Exprs:[]Code{
&Lexical{Levels:2, Offset:1},
&Lexical{Levels:0, Offset:0},
&Lexical{Levels:1, Offset:2},
}},
Alternate:&Quote{Value:c.FalseVal},
}},
}}}}}
c.EvalToplevel(code1)
}
