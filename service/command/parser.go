package command

import (
	"regexp"
	"strings"
)

type Parser struct {
	ctx Context
}

func (p *Parser) Ctx() Context {
	return p.ctx
}

func NewParser() *Parser {
	return &Parser{ctx: Context{options: &Option{}, argument: &Argument{}}}
}

func (p *Parser) Parse(str string) {
	if str == "" {
		return
	}

	pattern, err := regexp.Compile(`([-]{1,2}[\w\d-]+)[\s|=]?([\s\w\d]*)`)
	if err != nil {
		panic(err)
	}

	matches := pattern.FindAllStringSubmatch(str, -1)
	for _, match := range matches {
		name := match[1]
		var value interface{}
		value = match[2]
		if value == "" {
			value = false
		}

		p.ctx.options.set(name, value)

		str = strings.Replace(str, match[1], "", -1)
		str = strings.Replace(str, match[2], "", -1)
	}

	pattern, err = regexp.Compile(`^([\w\d-]+)[\s]?.*`)
	if err != nil {
		panic(err)
	}
	matches = pattern.FindAllStringSubmatch(str, 1)

	p.ctx.command = matches[0][1]
	str = strings.Replace(str, matches[0][1], "", -1)

	pattern, err = regexp.Compile(`([\w\d-]+)?.*`)
	if err != nil {
		panic(err)
	}
	matches = pattern.FindAllStringSubmatch(str, 1)

	p.ctx.argument.value = matches[0][1]
}
