package runtime

import (
	. "sint/core"
)

func initBooleans(c *Scheme) {
	code1 :=
		&Setglobal{Name: c.Intern("boolean?"), Rhs: &Lambda{
			Fixed: 1, Rest: false,
			Body: &Let{Exprs: []Code{
				&Call{Exprs: []Code{
					&Global{Name: c.Intern("eq?")},
					&Lexical{Levels: 0, Offset: 0},
					&Quote{Value: c.TrueVal},
				}},
			}, Body: &If{
				Test:       &Lexical{Levels: 0, Offset: 0},
				Consequent: &Lexical{Levels: 0, Offset: 0},
				Alternate: &Call{Exprs: []Code{
					&Global{Name: c.Intern("eq?")},
					&Lexical{Levels: 1, Offset: 0},
					&Quote{Value: c.FalseVal},
				}},
			}}}}
	c.EvalToplevel(code1)
	code2 :=
		&Setglobal{Name: c.Intern("not"), Rhs: &Lambda{
			Fixed: 1, Rest: false,
			Body: &If{
				Test:       &Lexical{Levels: 0, Offset: 0},
				Consequent: &Quote{Value: c.FalseVal},
				Alternate:  &Quote{Value: c.TrueVal},
			}}}
	c.EvalToplevel(code2)
}
