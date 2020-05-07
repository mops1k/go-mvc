package service

import (
	"bytes"
	"log"

	"github.com/CloudyKit/jet"
	"github.com/arthurkushman/pgo"

	"github.com/mops1k/go-mvc/cli"
)

type template struct {
	view *jet.Set
}

var Template *template

func init() {
	Template = &template{view: jet.NewHTMLSet(`./templates`)}
}

func (t *template) AddFunc(key string, fn jet.Func) {
	t.view.AddGlobalFunc(key, fn)
}

func (t *template) AddGlobal(key string, value interface{}) {
	t.view.AddGlobal(key, value)
}

func (t *template) render(template *jet.Template, vars map[string]interface{}) string {
	varMap := make(jet.VarMap)
	for name, value := range vars {
		varMap.Set(name, value)
	}

	var w bytes.Buffer
	err := template.Execute(&w, varMap, nil)
	if err != nil {
		cli.Logger.Get(cli.ErrorLog).(*log.Logger).Panic(err)
	}

	return w.String()
}

func (t *template) RenderTemplate(filename string, vars map[string]interface{}) string {
	template, err := t.view.GetTemplate(filename)
	if err != nil {
		cli.Logger.Get(cli.ErrorLog).(*log.Logger).Panic(err)
	}

	return t.render(template, vars)
}

func (t *template) RenderString(content string, vars map[string]interface{}) string {
	template, err := t.view.Parse(pgo.Md5(content), content)
	if err != nil {
		cli.Logger.Get(cli.ErrorLog).(*log.Logger).Panic(err)
	}

	return t.render(template, vars)
}
