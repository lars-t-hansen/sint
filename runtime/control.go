package runtime

import (
	. "sint/core"
)

func initControl(c *Scheme) {
	code1 :=
		&Setglobal{Name: c.Intern("eval"), Rhs: &Lambda{
			Fixed: 1, Rest: false,
			Body: &Call{Exprs: []Code{
				&Call{Exprs: []Code{
					&Global{Name: c.Intern("sint:compile-toplevel-phrase")},
					&Lexical{Levels: 0, Offset: 0},
				}},
			}}}}
	c.EvalToplevel(code1)
}
