package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/arthurkushman/pgo"
	"github.com/jinzhu/gorm"

	"github.com/mops1k/go-mvc/cli"
	"github.com/mops1k/go-mvc/service"
	"github.com/mops1k/go-mvc/tools/array"
)

type Controller interface {
	Action(c *Context) (string, error)
	Name() string
	Error(w http.ResponseWriter, err error, code int)
}

type BaseController struct {
	Controller
}

type controllerCollection struct {
	*array.Iterator
	data []Controller
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

func (bc *BaseController) RenderJson(c interface{}, ctx *Context) (string, error) {
	ctx.Header("Content-Type", "application/json")
	result, err := json.Marshal(c)

	return bytes.NewBuffer(result).String(), err
}

func (bc *BaseController) Error(w http.ResponseWriter, err error, code int) {
	http.Error(w, err.Error(), code)
}

func (bc *BaseController) NotFound(err string, ctx *Context) (string, error) {
	ctx.statusCode = http.StatusNotFound

	return "", errors.New(err)
}

func (bc *BaseController) GetManager(name interface{}) *gorm.DB {
	manager, err := service.Manager.Get(name)
	if err != nil {
		cli.Logger.Get(cli.ErrorLog).(*log.Logger).Panic(err)
	}

	return manager
}

func (cc *controllerCollection) Add(c Controller) *controllerCollection {
	if !pgo.InArray(c, cc.data) {
		cc.data = append(cc.data, c)
	}

	return cc
}

func (cc *controllerCollection) All() []Controller {
	return cc.data
}
