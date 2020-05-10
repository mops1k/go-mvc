package command

type Context struct {
	options  *Option
	argument *Argument
	command  string
}

func (c Context) Command() string {
	return c.command
}

func (c Context) Options() *Option {
	return c.options
}

func (c Context) Argument() *Argument {
	return c.argument
}
