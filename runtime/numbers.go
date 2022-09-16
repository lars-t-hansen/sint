
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
_, unwcode1 := c.EvalToplevel(code1)
if unwcode1 != nil { panic(unwcode1.String()) }
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
_, unwcode2 := c.EvalToplevel(code2)
if unwcode2 != nil { panic(unwcode2.String()) }
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
_, unwcode3 := c.EvalToplevel(code3)
if unwcode3 != nil { panic(unwcode3.String()) }
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
_, unwcode4 := c.EvalToplevel(code4)
if unwcode4 != nil { panic(unwcode4.String()) }
code5 := 
&Setglobal{Name:c.Intern("integer?"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("sint:exact-integer?")},
&Lexical{Levels:0, Offset:0},
}}}}
_, unwcode5 := c.EvalToplevel(code5)
if unwcode5 != nil { panic(unwcode5.String()) }
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
_, unwcode6 := c.EvalToplevel(code6)
if unwcode6 != nil { panic(unwcode6.String()) }
code7 := 
&Setglobal{Name:c.Intern("imag-part"), Rhs:&Lambda{
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
&Quote{Value:&Str{Value:"imag-part: not a number: "}},
&Lexical{Levels:0, Offset:0},
}},
Alternate:&Quote{Value:c.UnspecifiedVal},
},
&Quote{Value:big.NewInt(0)},
}}}}
_, unwcode7 := c.EvalToplevel(code7)
if unwcode7 != nil { panic(unwcode7.String()) }
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
_, unwcode8 := c.EvalToplevel(code8)
if unwcode8 != nil { panic(unwcode8.String()) }
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
_, unwcode9 := c.EvalToplevel(code9)
if unwcode9 != nil { panic(unwcode9.String()) }
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
_, unwcode10 := c.EvalToplevel(code10)
if unwcode10 != nil { panic(unwcode10.String()) }
code11 := 
&Setglobal{Name:c.Intern("exact->inexact"), Rhs:&Global{Name:c.Intern("inexact")}}
_, unwcode11 := c.EvalToplevel(code11)
if unwcode11 != nil { panic(unwcode11.String()) }
code12 := 
&Setglobal{Name:c.Intern("inexact->exact"), Rhs:&Global{Name:c.Intern("exact")}}
_, unwcode12 := c.EvalToplevel(code12)
if unwcode12 != nil { panic(unwcode12.String()) }
code13 := 
&Setglobal{Name:c.Intern("square"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("*")},
&Lexical{Levels:0, Offset:0},
&Lexical{Levels:0, Offset:0},
}}}}
_, unwcode13 := c.EvalToplevel(code13)
if unwcode13 != nil { panic(unwcode13.String()) }
code14 := 
&Setglobal{Name:c.Intern("nan?"), Rhs:&Lambda{
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
&Quote{Value:&Str{Value:"nan?: not a number: "}},
&Lexical{Levels:0, Offset:0},
}},
Alternate:&Quote{Value:c.UnspecifiedVal},
},
&Quote{Value:c.FalseVal},
}}}}
_, unwcode14 := c.EvalToplevel(code14)
if unwcode14 != nil { panic(unwcode14.String()) }
code15 := 
&Setglobal{Name:c.Intern("zero?"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("=")},
&Lexical{Levels:0, Offset:0},
&Quote{Value:big.NewInt(0)},
}}}}
_, unwcode15 := c.EvalToplevel(code15)
if unwcode15 != nil { panic(unwcode15.String()) }
code16 := 
&Setglobal{Name:c.Intern("positive?"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern(">")},
&Lexical{Levels:0, Offset:0},
&Quote{Value:big.NewInt(0)},
}}}}
_, unwcode16 := c.EvalToplevel(code16)
if unwcode16 != nil { panic(unwcode16.String()) }
code17 := 
&Setglobal{Name:c.Intern("negative?"), Rhs:&Lambda{
Fixed:1, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("<")},
&Lexical{Levels:0, Offset:0},
&Quote{Value:big.NewInt(0)},
}}}}
_, unwcode17 := c.EvalToplevel(code17)
if unwcode17 != nil { panic(unwcode17.String()) }
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
_, unwcode18 := c.EvalToplevel(code18)
if unwcode18 != nil { panic(unwcode18.String()) }
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
_, unwcode19 := c.EvalToplevel(code19)
if unwcode19 != nil { panic(unwcode19.String()) }
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
}, Body:&Let{Exprs:[]Code{
&If{
Test:&Lexical{Levels:1, Offset:0},
Consequent:&Lexical{Levels:1, Offset:0},
Alternate:&Call{Exprs:[]Code{
&Global{Name:c.Intern("inexact?")},
&Lexical{Levels:0, Offset:0},
}},
},
}, Body:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern(">")},
&Lexical{Levels:1, Offset:0},
&Lexical{Levels:2, Offset:1},
}},
Consequent:&Call{Exprs:[]Code{
&Lexical{Levels:3, Offset:0},
&Lexical{Levels:0, Offset:0},
&Lexical{Levels:1, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:2, Offset:2},
}},
}},
Alternate:&Call{Exprs:[]Code{
&Lexical{Levels:3, Offset:0},
&Lexical{Levels:0, Offset:0},
&Lexical{Levels:2, Offset:1},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:2, Offset:2},
}},
}},
}}},
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
_, unwcode20 := c.EvalToplevel(code20)
if unwcode20 != nil { panic(unwcode20.String()) }
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
}, Body:&Let{Exprs:[]Code{
&If{
Test:&Lexical{Levels:1, Offset:0},
Consequent:&Lexical{Levels:1, Offset:0},
Alternate:&Call{Exprs:[]Code{
&Global{Name:c.Intern("inexact?")},
&Lexical{Levels:0, Offset:0},
}},
},
}, Body:&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("<")},
&Lexical{Levels:1, Offset:0},
&Lexical{Levels:2, Offset:1},
}},
Consequent:&Call{Exprs:[]Code{
&Lexical{Levels:3, Offset:0},
&Lexical{Levels:0, Offset:0},
&Lexical{Levels:1, Offset:0},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:2, Offset:2},
}},
}},
Alternate:&Call{Exprs:[]Code{
&Lexical{Levels:3, Offset:0},
&Lexical{Levels:0, Offset:0},
&Lexical{Levels:2, Offset:1},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("cdr")},
&Lexical{Levels:2, Offset:2},
}},
}},
}}},
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
_, unwcode21 := c.EvalToplevel(code21)
if unwcode21 != nil { panic(unwcode21.String()) }
}
