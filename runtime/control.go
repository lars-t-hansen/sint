
// Generated from runtime/control.sch
package runtime
import (
	. "sint/core"
)
func initControl(c *Scheme) {
code1 := 
&Setglobal{Name:c.Intern("eval"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Call{Exprs:[]Code{
&Call{Exprs:[]Code{
&Global{Name:c.Intern("sint:compile-toplevel-phrase")},
&Lexical{Levels:0, Offset:0},
}},
}}}}
c.EvalToplevel(code1)
code2 := 
&Setglobal{Name:c.Intern("map"), Rhs:&Letrec{Exprs:[]Code{
&Lambda{
Fixed:2, Rest:false,
Body:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Lexical{Levels:0, Offset:1},
}},
Consequent:&Quote{Value:c.NullVal},
Alternate:&Call{Exprs:[]Code{
&Global{Name:c.Intern("cons")},
&Call{Exprs:[]Code{
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:1},
}},
}},
&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:1},
}},
}},
}},
}},
&Lambda{
Fixed:3, Rest:false,
Body:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Lexical{Levels:0, Offset:1},
}},
Consequent:&Quote{Value:c.NullVal},
Alternate:&Call{Exprs:[]Code{
&Global{Name:c.Intern("cons")},
&Call{Exprs:[]Code{
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:1},
}},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:2},
}},
}},
&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:1},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:1},
}},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:2},
}},
}},
}},
}},
&Lambda{
Fixed:2, Rest:false,
Body:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:1},
}},
}},
Consequent:&Quote{Value:c.NullVal},
Alternate:&Call{Exprs:[]Code{
&Global{Name:c.Intern("cons")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("apply")},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:1},
}},
}},
&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:2},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:1},
}},
}},
}},
}},
}, Body:&Lambda{
Fixed:2, Rest:true,
Body:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Lexical{Levels:0, Offset:2},
}},
Consequent:&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Lexical{Levels:0, Offset:0},
&Lexical{Levels:0, Offset:1},
}},
Alternate:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:2},
}},
}},
Consequent:&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:1},
&Lexical{Levels:0, Offset:0},
&Lexical{Levels:0, Offset:1},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:2},
}},
}},
Alternate:&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:2},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cons")},
&Lexical{Levels:0, Offset:1},
&Lexical{Levels:0, Offset:2},
}},
}},
},
}}}}
c.EvalToplevel(code2)
}
