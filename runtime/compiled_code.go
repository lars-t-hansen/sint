// This file will be replaced by generated code, what's here is a placeholder

package runtime

import . "sint/core"

func gcall1(c *Scheme, s string, e1 Code) Code {
	return &Call{[]Code{&Global{c.Intern(s)}, e1}}
}

func InitCompiled(c *Scheme) {
	c.Intern("cadr").Value =
		&Procedure{
			&Lambda{1, false, gcall1(c, "car", gcall1(c, "cadr", &Lexical{0, 0}))},
			nil, nil}
}
