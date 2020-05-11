package service

import (
	"bytes"
	"log"

	"github.com/tyler-sommer/stick"
	"github.com/tyler-sommer/stick/twig"

	"github.com/mops1k/go-mvc/cli"
)

type template struct {
	env     *stick.Env
	globals map[string]interface{}
}

var Template *template

func init() {
	Template = &template{env: stick.New(stick.NewFilesystemLoader("./templates"))}
	Template.AddExtension(twig.NewAutoEscapeExtension())
}

func (t *template) AddFunc(key string, handler stick.Func) {
	t.env.Functions[key] = handler
}

func (t *template) AddFilter(key string, handler stick.Filter) {
	t.env.Filters[key] = handler
}

func (t *template) UseFuncAsFilter(key string, handler stick.Func) {
	fn := func(ctx stick.Context, val stick.Value, args ...stick.Value) stick.Value {
		arguments := append([]stick.Value{val}, args...)

		return handler(ctx, arguments...)
	}

	t.env.Filters[key] = fn
}

func (t *template) AddGlobal(key string, value interface{}) {
	t.globals[key] = value
}

func (t *template) AddExtension(e stick.Extension) {
	err := t.env.Register(e)
	if err != nil {
		cli.Logger.Get(cli.ErrorLog).(*log.Logger).Panic(err)
	}
}

func (t *template) render(content string, vars map[string]interface{}) string {
	args := make(map[string]stick.Value)
	for key, value := range vars {
		args[key] = value
	}

	var w bytes.Buffer
	err := t.env.Execute(content, &w, args)
	if err != nil {
		cli.Logger.Get(cli.ErrorLog).(*log.Logger).Panic(err)
	}

	return w.String()
}

func (t *template) RenderTemplate(filename string, vars map[string]interface{}) string {
	return t.render(filename, vars)
}

func (t *template) RenderString(content string, vars map[string]interface{}) string {
	prevLoader := t.env.Loader

	t.env.Loader = &stick.StringLoader{}
	result := t.render(content, vars)
	t.env.Loader = prevLoader

	return result
}
