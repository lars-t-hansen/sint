
// Generated from runtime/numbers.sch
package runtime
import (
	. "sint/core"
	"math/big"
)
func dummyNumbers() {
	// Make sure the imports are used, or the Go compiler barfs.
	var _ Val = big.NewInt(0)
}
func initNumbers(c *Scheme) {
code1 := 
&Setglobal{Name:c.Intern("number?"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Let{Exprs:[]Code{
&Call{Exprs:[]Code{
&Global{Name:c.Intern("sint:exact-integer?")},
&Lexical{Levels:0, Offset:0},
}},
}, Body:&If{
Test:&Lexical{Levels:0, Offset:0},
Consequent:&Lexical{Levels:0, Offset:0},
Alternate:&Call{Exprs:[]Code{
&Global{Name:c.Intern("sint:inexact-float?")},
&Lexical{Levels:1, Offset:0},
}},
}}}}
c.EvalToplevel(code1)
code2 := 
&Setglobal{Name:c.Intern("complex?"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Let{Exprs:[]Code{
&Call{Exprs:[]Code{
&Global{Name:c.Intern("sint:exact-integer?")},
&Lexical{Levels:0, Offset:0},
}},
}, Body:&If{
Test:&Lexical{Levels:0, Offset:0},
Consequent:&Lexical{Levels:0, Offset:0},
Alternate:&Call{Exprs:[]Code{
&Global{Name:c.Intern("sint:inexact-float?")},
&Lexical{Levels:1, Offset:0},
}},
}}}}
c.EvalToplevel(code2)
code3 := 
&Setglobal{Name:c.Intern("real?"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Let{Exprs:[]Code{
&Call{Exprs:[]Code{
&Global{Name:c.Intern("sint:exact-integer?")},
&Lexical{Levels:0, Offset:0},
}},
}, Body:&If{
Test:&Lexical{Levels:0, Offset:0},
Consequent:&Lexical{Levels:0, Offset:0},
Alternate:&Call{Exprs:[]Code{
&Global{Name:c.Intern("sint:inexact-float?")},
&Lexical{Levels:1, Offset:0},
}},
}}}}
c.EvalToplevel(code3)
code4 := 
&Setglobal{Name:c.Intern("rational?"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Let{Exprs:[]Code{
&Call{Exprs:[]Code{
&Global{Name:c.Intern("sint:exact-integer?")},
&Lexical{Levels:0, Offset:0},
}},
}, Body:&If{
Test:&Lexical{Levels:0, Offset:0},
Consequent:&Lexical{Levels:0, Offset:0},
Alternate:&Call{Exprs:[]Code{
&Global{Name:c.Intern("sint:inexact-float?")},
&Lexical{Levels:1, Offset:0},
}},
}}}}
c.EvalToplevel(code4)
code5 := 
&Setglobal{Name:c.Intern("integer?"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("sint:exact-integer?")},
&Lexical{Levels:0, Offset:0},
}}}}
c.EvalToplevel(code5)
code6 := 
&Setglobal{Name:c.Intern("real-part"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Begin{Exprs:[]Code{
&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("not")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("number?")},
&Lexical{Levels:0, Offset:0},
}},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("error")},
&Quote{Value:&Str{Value:"real-part: not a number: "}},
&Lexical{Levels:0, Offset:0},
}},
Alternate:&Quote{Value:c.UnspecifiedVal},
},
&Lexical{Levels:0, Offset:0},
}}}}
c.EvalToplevel(code6)
code7 := 
&Setglobal{Name:c.Intern("imag-part"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Begin{Exprs:[]Code{
&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("not")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("number?")},
&Global{Name:c.Intern("z")},
}},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("error")},
&Quote{Value:&Str{Value:"imag-part: not a number: "}},
&Global{Name:c.Intern("z")},
}},
Alternate:&Quote{Value:c.UnspecifiedVal},
},
&Quote{Value:big.NewInt(0)},
}}}}
c.EvalToplevel(code7)
code8 := 
&Setglobal{Name:c.Intern("exact?"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Begin{Exprs:[]Code{
&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("not")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("number?")},
&Lexical{Levels:0, Offset:0},
}},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("error")},
&Quote{Value:&Str{Value:"exact?: not a number: "}},
&Lexical{Levels:0, Offset:0},
}},
Alternate:&Quote{Value:c.UnspecifiedVal},
},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("sint:exact-integer?")},
&Lexical{Levels:0, Offset:0},
}},
}}}}
c.EvalToplevel(code8)
code9 := 
&Setglobal{Name:c.Intern("inexact?"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Begin{Exprs:[]Code{
&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("not")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("number?")},
&Lexical{Levels:0, Offset:0},
}},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("error")},
&Quote{Value:&Str{Value:"inexact?: not a number: "}},
&Lexical{Levels:0, Offset:0},
}},
Alternate:&Quote{Value:c.UnspecifiedVal},
},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("sint:inexact-float?")},
&Lexical{Levels:0, Offset:0},
}},
}}}}
c.EvalToplevel(code9)
code10 := 
&Setglobal{Name:c.Intern("exact-integer?"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Begin{Exprs:[]Code{
&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("not")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("number?")},
&Lexical{Levels:0, Offset:0},
}},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("error")},
&Quote{Value:&Str{Value:"exact-integer?: not a number: "}},
&Lexical{Levels:0, Offset:0},
}},
Alternate:&Quote{Value:c.UnspecifiedVal},
},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("sint:exact-integer?")},
&Lexical{Levels:0, Offset:0},
}},
}}}}
c.EvalToplevel(code10)
code11 := 
&Setglobal{Name:c.Intern("exact->inexact"), Rhs:&Global{Name:c.Intern("inexact")}}
c.EvalToplevel(code11)
code12 := 
&Setglobal{Name:c.Intern("inexact->exact"), Rhs:&Global{Name:c.Intern("exact")}}
c.EvalToplevel(code12)
code13 := 
&Setglobal{Name:c.Intern("square"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("*")},
&Lexical{Levels:0, Offset:0},
&Lexical{Levels:0, Offset:0},
}}}}
c.EvalToplevel(code13)
code14 := 
&Setglobal{Name:c.Intern("nan?"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Begin{Exprs:[]Code{
&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("not")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("number?")},
&Global{Name:c.Intern("z")},
}},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("error")},
&Quote{Value:&Str{Value:"nan?: not a number: "}},
&Global{Name:c.Intern("z")},
}},
Alternate:&Quote{Value:c.UnspecifiedVal},
},
&Quote{Value:c.FalseVal},
}}}}
c.EvalToplevel(code14)
code15 := 
&Setglobal{Name:c.Intern("zero?"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("=")},
&Lexical{Levels:0, Offset:0},
&Quote{Value:big.NewInt(0)},
}}}}
c.EvalToplevel(code15)
code16 := 
&Setglobal{Name:c.Intern("positive"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern(">")},
&Lexical{Levels:0, Offset:0},
&Quote{Value:big.NewInt(0)},
}}}}
c.EvalToplevel(code16)
code17 := 
&Setglobal{Name:c.Intern("negative?"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("<")},
&Lexical{Levels:0, Offset:0},
&Quote{Value:big.NewInt(0)},
}}}}
c.EvalToplevel(code17)
code18 := 
&Setglobal{Name:c.Intern("odd?"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Begin{Exprs:[]Code{
&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("not")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("exact-integer?")},
&Lexical{Levels:0, Offset:0},
}},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("error")},
&Quote{Value:&Str{Value:"odd?: not an exact integer: "}},
&Lexical{Levels:0, Offset:0},
}},
Alternate:&Quote{Value:c.UnspecifiedVal},
},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("not")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("zero?")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("remainder")},
&Lexical{Levels:0, Offset:0},
&Quote{Value:big.NewInt(2)},
}},
}},
}},
}}}}
c.EvalToplevel(code18)
code19 := 
&Setglobal{Name:c.Intern("even?"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Begin{Exprs:[]Code{
&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("not")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("exact-integer?")},
&Lexical{Levels:0, Offset:0},
}},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("error")},
&Quote{Value:&Str{Value:"even?: not an exact integer: "}},
&Lexical{Levels:0, Offset:0},
}},
Alternate:&Quote{Value:c.UnspecifiedVal},
},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("zero?")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("remainder")},
&Lexical{Levels:0, Offset:0},
&Quote{Value:big.NewInt(2)},
}},
}},
}}}}
c.EvalToplevel(code19)
code20 := 
&Setglobal{Name:c.Intern("max"), Rhs:&Letrec{Exprs:[]Code{
&Lambda{
Fixed:3, Rest:false,
Body:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Lexical{Levels:0, Offset:2},
}},
Consequent:&If{
Test:&Lexical{Levels:0, Offset:0},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("inexact")},
&Lexical{Levels:0, Offset:1},
}},
Alternate:&Lexical{Levels:0, Offset:1},
},
Alternate:&Let{Exprs:[]Code{
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:2},
}},
}, Body:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern(">")},
&Lexical{Levels:0, Offset:0},
&Lexical{Levels:1, Offset:1},
}},
Consequent:&Call{Exprs:[]Code{
&Lexical{Levels:2, Offset:0},
&If{
Test:&Lexical{Levels:1, Offset:0},
Consequent:&Lexical{Levels:1, Offset:0},
Alternate:&Call{Exprs:[]Code{
&Global{Name:c.Intern("inexact?")},
&Lexical{Levels:0, Offset:0},
}},
},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:1, Offset:2},
}},
}},
Alternate:&Call{Exprs:[]Code{
&Lexical{Levels:2, Offset:0},
&Lexical{Levels:1, Offset:0},
&Lexical{Levels:1, Offset:1},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:1, Offset:2},
}},
}},
}},
}},
}, Body:&Lambda{
Fixed:1, Rest:true,
Body:&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("inexact?")},
&Lexical{Levels:0, Offset:0},
}},
&Lexical{Levels:0, Offset:0},
&Lexical{Levels:0, Offset:1},
}}}}}
c.EvalToplevel(code20)
code21 := 
&Setglobal{Name:c.Intern("min"), Rhs:&Letrec{Exprs:[]Code{
&Lambda{
Fixed:3, Rest:false,
Body:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Lexical{Levels:0, Offset:2},
}},
Consequent:&If{
Test:&Lexical{Levels:0, Offset:0},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("inexact")},
&Lexical{Levels:0, Offset:1},
}},
Alternate:&Lexical{Levels:0, Offset:1},
},
Alternate:&Let{Exprs:[]Code{
&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:2},
}},
}, Body:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("<")},
&Lexical{Levels:0, Offset:0},
&Lexical{Levels:1, Offset:1},
}},
Consequent:&Call{Exprs:[]Code{
&Lexical{Levels:2, Offset:0},
&If{
Test:&Lexical{Levels:1, Offset:0},
Consequent:&Lexical{Levels:1, Offset:0},
Alternate:&Call{Exprs:[]Code{
&Global{Name:c.Intern("inexact?")},
&Lexical{Levels:0, Offset:0},
}},
},
&Lexical{Levels:0, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:1, Offset:2},
}},
}},
Alternate:&Call{Exprs:[]Code{
&Lexical{Levels:2, Offset:0},
&Lexical{Levels:1, Offset:0},
&Lexical{Levels:1, Offset:1},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:1, Offset:2},
}},
}},
}},
}},
}, Body:&Lambda{
Fixed:1, Rest:true,
Body:&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("inexact?")},
&Lexical{Levels:0, Offset:0},
}},
&Lexical{Levels:0, Offset:0},
&Lexical{Levels:0, Offset:1},
}}}}}
c.EvalToplevel(code21)
}
