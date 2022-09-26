
// Generated from runtime/generator.sch
package runtime
import (
	. "sint/core"
	"math/big"
)
// Make sure the imports are used, or the Go compiler barfs.
var _ Val = big.NewInt(0)
func initGenerator(c *Scheme) {
code1 := 
&Setglobal{Name:c.Intern("make-generator"), Rhs:&Lambda{
Fixed:1, Rest:true,
Body:&Begin{Exprs:[]Code{
&Quote{Value:&Str{Value:"make-generator takes a procedure `p` of one argument, and optionally a value `end`, and returns a\nthunk, the generator.  The generator is invoked to obtain the values yielded by `p`.  `p` will be\ncalled once with a a function of one argument that is used by `p` to yield its values.  Once `p`\nreturns, the value `end` is yielded by the generator once, followed by a stream of #!unspecified.\n\n`p` is run on a concurrent thread, and care should be taken by both `p` and the consumer when\nupdating shared state.  The communication channel is unbuffered, so `p` and the consumer run\nsomewhat in lockstep, but they are not synchronized: `p` is working on the next item while the\nconsumer is processing the previous one.\n"}},
&Let{Exprs:[]Code{
&Call{Exprs:[]Code{
&Global{Name:c.Intern("make-channel")},
}},
&If{
Test:&Call{Exprs:[]Code{
&Global{Name:c.Intern("not")},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("null?")},
&Lexical{Levels:0, Offset:1},
}},
}},
Consequent:&Call{Exprs:[]Code{
&Global{Name:c.Intern("car")},
&Lexical{Levels:0, Offset:1},
}},
Alternate:&Call{Exprs:[]Code{
&Global{Name:c.Intern("unspecified")},
}},
},
}, Body:&Begin{Exprs:[]Code{
&Call{Exprs:[]Code{
&Global{Name:c.Intern("sint:go")},
&Let{Exprs:[]Code{
&Lambda{
Fixed:0, Rest:false,
Body:&Begin{Exprs:[]Code{
&Call{Exprs:[]Code{
&Lexical{Levels:2, Offset:0},
&Lambda{
Fixed:1, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("channel-send")},
&Lexical{Levels:2, Offset:0},
&Lexical{Levels:0, Offset:0},
}},
Name:"make-generator >  > .G1001.GO > [lambda]"},
}},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("channel-send")},
&Lexical{Levels:1, Offset:0},
&Lexical{Levels:1, Offset:1},
}},
&Call{Exprs:[]Code{
&Global{Name:c.Intern("close-channel")},
&Lexical{Levels:1, Offset:0},
}},
}},
Name:"make-generator >  > .G1001.GO"},
}, Body:&Lambda{
Fixed:0, Rest:false,
Body:&Call{Exprs:[]Code{
&Lexical{Levels:1, Offset:0},
}},
Name:"make-generator > "}},
}},
&Lambda{
Fixed:0, Rest:false,
Body:&Call{Exprs:[]Code{
&Global{Name:c.Intern("channel-receive")},
&Lexical{Levels:1, Offset:0},
}},
Name:"make-generator > "},
}}},
}},
Name:"make-generator"}}
_, unwcode1 := c.EvalToplevel(code1)
if unwcode1 != nil { panic(unwcode1.String()) }
}
