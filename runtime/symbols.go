
// Generated from runtime/symbols.sch
package runtime
import (
	. "sint/core"
	"math/big"
)
// Make sure the imports are used, or the Go compiler barfs.
var _ Val = big.NewInt(0)
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
&Quote{Value:&Str{Value:"symbol=?: not a symbol: "}},
&Lexical{Levels:0, Offset:0},
}},
Alternate:&Quote{Value:c.UnspecifiedVal},
},
Name:"symbol=? > check",
Signature:&Cons{Car:c.Intern("x"), Cdr:c.NullVal}},
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
},
Name:"symbol=? > loop",
Signature:&Cons{Car:c.Intern("sa"), Cdr:&Cons{Car:c.Intern("xs"), Cdr:c.NullVal}}},
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
}},
Name:"symbol=?",
Signature:&Cons{Car:c.Intern("a"), Cdr:&Cons{Car:c.Intern("b"), Cdr:c.Intern("xs")}}}}}
_, unwcode1 := c.EvalToplevel(code1)
if unwcode1 != nil { panic(unwcode1.String()) }
code2 := 
&Setglobal{Name:c.Intern("apropos"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Let{Exprs:[]Code{
&Call{Exprs:[]Code{
&Global{Name:c.Intern("list-sort!")},
&Lambda{
Fixed:2, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("string<?")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("symbol->string")},
&Lexical{Levels:0, Offset:0},
}},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("symbol->string")},
&Lexical{Levels:0, Offset:1},
}},
}},
Name:"apropos >  > xs",
Signature:&Cons{Car:c.Intern("a"), Cdr:&Cons{Car:c.Intern("b"), Cdr:c.NullVal}}},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("filter-global-variables")},
&Lexical{Levels:0, Offset:0},
}},
}},
}, Body:&Call{Exprs:[]Code{
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
Name:"apropos > ",
Signature:&Cons{Car:c.Intern("x"), Cdr:c.NullVal}},
&Lexical{Levels:0, Offset:0},
}}},
Name:"apropos",
Signature:&Cons{Car:c.Intern("pattern"), Cdr:c.NullVal}}}
_, unwcode2 := c.EvalToplevel(code2)
if unwcode2 != nil { panic(unwcode2.String()) }
}
