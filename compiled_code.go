// This file will be replaced by generated code, what's here is a placeholder

package sint

func (c *Scheme) gcall1(s string, e1 Code) Code {
	return &Call{[]Code{c.intern(s), e1}}
}

func (c *Scheme) initCompiled() {
	c.intern("cadr").value =
		&Procedure{
			&Lambda{1, false, c.gcall1("car", c.gcall1("cadr", &Lexical{0, 0}))},
			nil, nil}
}
