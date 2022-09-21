
// Generated from runtime/sort.sch
package runtime
import (
	. "sint/core"
	"math/big"
)
func dummySort() {
	// Make sure the imports are used, or the Go compiler barfs.
	var _ Val = big.NewInt(0)
}
func initSort(c *Scheme) {
code1 := 
&Setglobal{Name:c.Intern("list-sort!"), Rhs:&Letrec{Exprs:[]Code{
&Lambda{
Fixed:4, Rest:false,
Body:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Lexical{Levels:0, Offset:1},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("set-cdr!")},
&Lexical{Levels:0, Offset:3},
&Lexical{Levels:0, Offset:2},
}},
Alternate:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Lexical{Levels:0, Offset:2},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("set-cdr!")},
&Lexical{Levels:0, Offset:3},
&Lexical{Levels:0, Offset:1},
}},
Alternate:&If{
Test:&Call{Exprs:[]Code{
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
Consequent:&Let{Exprs:[]Code{
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:1},
}},
}, Body:&Begin{Exprs:[]Code{
&Call{Exprs:[]Code{
&Global{Name:c.Intern("set-cdr!")},
&Lexical{Levels:1, Offset:1},
&Quote{Value:c.NullVal},
}},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("set-cdr!")},
&Lexical{Levels:1, Offset:3},
&Lexical{Levels:1, Offset:1},
}},
&Call{Exprs:[]Code{
&Lexical{Levels:2, Offset:0},
&Lexical{Levels:1, Offset:0},
&Lexical{Levels:0, Offset:0},
&Lexical{Levels:1, Offset:2},
&Lexical{Levels:1, Offset:1},
}},
}}},
Alternate:&Let{Exprs:[]Code{
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:2},
}},
}, Body:&Begin{Exprs:[]Code{
&Call{Exprs:[]Code{
&Global{Name:c.Intern("set-cdr!")},
&Lexical{Levels:1, Offset:2},
&Quote{Value:c.NullVal},
}},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("set-cdr!")},
&Lexical{Levels:1, Offset:3},
&Lexical{Levels:1, Offset:2},
}},
&Call{Exprs:[]Code{
&Lexical{Levels:2, Offset:0},
&Lexical{Levels:1, Offset:0},
&Lexical{Levels:1, Offset:1},
&Lexical{Levels:0, Offset:0},
&Lexical{Levels:1, Offset:2},
}},
}}},
},
},
}},
&Lambda{
Fixed:4, Rest:false,
Body:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Lexical{Levels:0, Offset:2},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("values")},
&Quote{Value:c.NullVal},
&Quote{Value:c.NullVal},
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
&Global{Name:c.Intern("values")},
&Lexical{Levels:0, Offset:2},
&Quote{Value:c.NullVal},
}},
Alternate:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("=")},
&Lexical{Levels:0, Offset:0},
&Quote{Value:big.NewInt(1)},
}},
Consequent:&Let{Exprs:[]Code{
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:0, Offset:2},
}},
}, Body:&Begin{Exprs:[]Code{
&Call{Exprs:[]Code{
&Global{Name:c.Intern("set-cdr!")},
&Lexical{Levels:1, Offset:2},
&Quote{Value:c.NullVal},
}},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("values")},
&Lexical{Levels:1, Offset:2},
&Lexical{Levels:0, Offset:0},
}},
}}},
Alternate:&Call{Exprs:[]Code{
&Global{Name:c.Intern("call-with-values")},
&Lambda{
Fixed:0, Rest:false,
Body:&Call{Exprs:[]Code{
&Lexical{Levels:2, Offset:1},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("-")},
&Lexical{Levels:1, Offset:0},
&Quote{Value:big.NewInt(1)},
}},
&Lexical{Levels:1, Offset:1},
&Lexical{Levels:1, Offset:2},
&Lexical{Levels:1, Offset:3},
}}},
&Lambda{
Fixed:2, Rest:false,
Body:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Lexical{Levels:0, Offset:1},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("values")},
&Lexical{Levels:0, Offset:0},
&Quote{Value:c.NullVal},
}},
Alternate:&Call{Exprs:[]Code{
&Global{Name:c.Intern("call-with-values")},
&Lambda{
Fixed:0, Rest:false,
Body:&Call{Exprs:[]Code{
&Lexical{Levels:3, Offset:1},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("-")},
&Lexical{Levels:2, Offset:0},
&Quote{Value:big.NewInt(1)},
}},
&Lexical{Levels:2, Offset:1},
&Lexical{Levels:1, Offset:1},
&Lexical{Levels:2, Offset:3},
}}},
&Lambda{
Fixed:2, Rest:false,
Body:&Begin{Exprs:[]Code{
&Call{Exprs:[]Code{
&Global{Name:c.Intern("set-cdr!")},
&Lexical{Levels:2, Offset:3},
&Quote{Value:c.NullVal},
}},
&Call{Exprs:[]Code{
&Lexical{Levels:3, Offset:0},
&Lexical{Levels:2, Offset:1},
&Lexical{Levels:1, Offset:0},
&Lexical{Levels:0, Offset:0},
&Lexical{Levels:2, Offset:3},
}},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("values")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:2, Offset:3},
}},
&Lexical{Levels:0, Offset:1},
}},
}}},
}},
}},
}},
},
},
}},
}, Body:&Lambda{
Fixed:2, Rest:false,
Body:&Begin{Exprs:[]Code{
&Quote{Value:&Str{Value:"Given a binary predicate `<?` and a list `xs`, sort the `xs` in-place and return the new list."}},
&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:1},
&Quote{Value:big.NewInt(32)},
&Lexical{Levels:0, Offset:0},
&Lexical{Levels:0, Offset:1},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cons")},
&Quote{Value:c.FalseVal},
&Quote{Value:c.FalseVal},
}},
}},
}}}}}
_, unwcode1 := c.EvalToplevel(code1)
if unwcode1 != nil { panic(unwcode1.String()) }
code2 := 
&Setglobal{Name:c.Intern("list-sorted?"), Rhs:&Lambda{
Fixed:2, Rest:false,
Body:&Begin{Exprs:[]Code{
&Quote{Value:&Str{Value:"Return #t iff the list `xs` are sorted according to the binary predicate `<?`."}},
&Let{Exprs:[]Code{
&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Lexical{Levels:0, Offset:1},
}},
}, Body:&If{
Test:&Lexical{Levels:0, Offset:0},
Consequent:&Lexical{Levels:0, Offset:0},
Alternate:&Let{Exprs:[]Code{
&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:1, Offset:1},
}},
}},
}, Body:&If{
Test:&Lexical{Levels:0, Offset:0},
Consequent:&Lexical{Levels:0, Offset:0},
Alternate:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("not")},
&Call{Exprs:[]Code{
&Lexical{Levels:2, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cadr")},
&Lexical{Levels:2, Offset:1},
}},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:2, Offset:1},
}},
}},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("list-sorted?")},
&Lexical{Levels:2, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:2, Offset:1},
}},
}},
Alternate:&Quote{Value:c.FalseVal},
},
}},
}},
}},
Name:"list-sorted?"}}
_, unwcode2 := c.EvalToplevel(code2)
if unwcode2 != nil { panic(unwcode2.String()) }
}
