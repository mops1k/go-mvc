package service

import (
	"log"

	"github.com/mops1k/go-mvc/cli"
	"github.com/mops1k/go-mvc/service/command"
	"github.com/mops1k/go-mvc/tools/map"
)

type Command interface {
	Name() string
	Description() string
	Action(ctx command.Context)
}

type CommandCollection struct {
	_map.Collection
}

var Commands *CommandCollection

func init() {
	Commands = &CommandCollection{}
}

func (cc *CommandCollection) Add(c Command) *CommandCollection {
	err := cc.Collection.Add(c.Name(), c)
	if err != nil {
		cli.Logger.Get(cli.AppLog).(*log.Logger).Panic(err)
	}

	return cc
}

func (cc *CommandCollection) Get(name string) Command {
	value, err := cc.Collection.Get(name)
	if err != nil {
		cli.Logger.Get(cli.AppLog).(*log.Logger).Panic("command does not exists")
	}

	return value.(Command)
}
