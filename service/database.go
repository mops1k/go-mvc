package service

import (
	"errors"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/spf13/cast"

	"github.com/mops1k/go-mvc/cli"
)

type manager struct {
	instances map[string]*gorm.DB
}

var Manager *manager

func init() {
	databases := Config.GetStringMap("database")
	Manager = &manager{instances: make(map[string]*gorm.DB)}
	for name := range databases {
		Manager.set(name)
	}
}

func (m *manager) Get(name interface{}) (*gorm.DB, error) {
	if name == nil {
		if len(m.instances) == 1 {
			for _, instance := range m.instances {
				return instance, nil
			}
		}

		return nil, errors.New("manager name cannot be nil if database instances more than 1")
	}

	if m.instances[cast.ToString(name)] == nil {
		return nil, fmt.Errorf(`manager for "%s" are not found`, name)
	}

	return m.instances[cast.ToString(name)], nil
}

func (m *manager) set(name string) {
	var err error
	if !Config.GetBool("database." + name + ".enabled") {
		return
	}

	m.instances[name], err = gorm.Open(Config.GetString("database."+name+".type"), Config.GetString("database."+name+".url"))
	if err != nil {
		cli.Logger.Get(cli.DbLog).(*log.Logger).Panic(err)
	}
	m.instances[name].SetLogger(cli.Logger.Get(cli.DbLog).(*log.Logger))
	m.instances[name].Debug()
}

func (m *manager) AddModels(name string, models ...interface{}) {
	manager, err := m.Get(name)
	if err != nil {
		cli.Logger.Get(cli.DbLog).(*log.Logger).Print(err)
		return
	}

	manager.AutoMigrate(models...)
}

func (m *manager) Close() {
	for name, instance := range m.instances {
		err := instance.Close()
		if err != nil {
			cli.Logger.Get(cli.DbLog).(*log.Logger).Printf(`[database: %s] %s`, name, err)
		}
	}
}
