package runtime

import (
	. "sint/core"
)

func initEquivalence(c *Scheme) {
	code1 :=
		&Setglobal{Name: c.Intern("equal?"), Rhs: &Lambda{
			Fixed: 2, Rest: false,
			Body: &Let{Exprs: []Code{
				&Call{Exprs: []Code{
					&Global{Name: c.Intern("eqv?")},
					&Lexical{Levels: 0, Offset: 0},
					&Lexical{Levels: 0, Offset: 1},
				}},
			}, Body: &If{
				Test:       &Lexical{Levels: 0, Offset: 0},
				Consequent: &Lexical{Levels: 0, Offset: 0},
				Alternate: &If{
					Test: &Call{Exprs: []Code{
						&Global{Name: c.Intern("pair?")},
						&Lexical{Levels: 1, Offset: 0},
					}},
					Consequent: &If{
						Test: &Call{Exprs: []Code{
							&Global{Name: c.Intern("pair?")},
							&Lexical{Levels: 1, Offset: 1},
						}},
						Consequent: &If{
							Test: &Call{Exprs: []Code{
								&Global{Name: c.Intern("equal?")},
								&Call{Exprs: []Code{
									&Global{Name: c.Intern("car")},
									&Lexical{Levels: 1, Offset: 0},
								}},
								&Call{Exprs: []Code{
									&Global{Name: c.Intern("car")},
									&Lexical{Levels: 1, Offset: 1},
								}},
							}},
							Consequent: &Call{Exprs: []Code{
								&Global{Name: c.Intern("equal?")},
								&Call{Exprs: []Code{
									&Global{Name: c.Intern("cdr")},
									&Lexical{Levels: 1, Offset: 0},
								}},
								&Call{Exprs: []Code{
									&Global{Name: c.Intern("cdr")},
									&Lexical{Levels: 1, Offset: 1},
								}},
							}},
							Alternate: &Quote{Value: c.FalseVal},
						},
						Alternate: &Quote{Value: c.FalseVal},
					},
					Alternate: &Quote{Value: c.FalseVal},
				},
			}}}}
	c.EvalToplevel(code1)
}