package http

import (
	"log"

	"github.com/arthurkushman/pgo"
	"github.com/jinzhu/gorm"

	"github.com/mops1k/go-mvc/cli"
	"github.com/mops1k/go-mvc/service"
)

type Controller interface {
	Action(c *Context) string
	Name() string
}

type BaseController struct {
	Controller
}

type controllerCollection struct {
	controllers []Controller
}

var Controllers *controllerCollection

func init() {
	Controllers = &controllerCollection{}
}

func (bc *BaseController) Render(filename string, vars map[string]interface{}) string {
	return service.Template.RenderTemplate(filename, vars)
}

func (bc *BaseController) RenderString(content string, vars map[string]interface{}) string {
	return service.Template.RenderString(content, vars)
}

func (bc *BaseController) GetManager(name interface{}) *gorm.DB {
	manager, err := service.Manager.Get(name)
	if err != nil {
		cli.Logger.Get(cli.AppLog).(*log.Logger).Panic(err)
	}

	return manager
}

func (cc *controllerCollection) Add(c Controller) *controllerCollection {
	if !pgo.InArray(c, cc.controllers) {
		cc.controllers = append(cc.controllers, c)
	}

	return cc
}

func (cc *controllerCollection) All() []Controller {
	return cc.controllers
}
