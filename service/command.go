package service

import (
	"github.com/mops1k/go-mvc/service/command"
	"github.com/mops1k/go-mvc/tools"
)

type Command interface {
	Name() string
	Description() string
	Action(ctx command.Context)
}

type CommandCollection struct {
	*tools.MapIterator
	data map[string]Command
}

var Commands *CommandCollection

func init() {
	Commands = &CommandCollection{}
}

func (cc *CommandCollection) Add(c Command) *CommandCollection {
	if cc.data == nil {
		cc.data = make(map[string]Command)
	}

	cc.data[c.Name()] = c

	return cc
}

func (cc *CommandCollection) Get(name string) Command {
	return cc.data[name]
}

func (cc *CommandCollection) GetAll() map[string]Command {
	return cc.data
}

func (cc *CommandCollection) Has(name string) bool {
	return cc.data[name] != nil
}
