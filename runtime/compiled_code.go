// This file will be replaced by generated code, what's here is a placeholder

package runtime

import . "sint/core"

func gcall1(c *Scheme, s string, e1 Code) Code {
	return &Call{Exprs: []Code{&Global{Name: c.Intern(s)}, e1}}
}

func InitCompiled(c *Scheme) {
	c.Intern("cadr").Value =
		&Procedure{
			Lam: &Lambda{
				Fixed: 1,
				Rest:  false,
				Body:  gcall1(c, "car", gcall1(c, "cadr", &Lexical{Levels: 0, Offset: 0}))},
			Env:    nil,
			Primop: nil}
}
