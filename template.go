package mvc

import (
	"log"
	"reflect"

	"github.com/CloudyKit/jet"

	"github.com/mops1k/go-mvc/cli"
	"github.com/mops1k/go-mvc/http"
	"github.com/mops1k/go-mvc/service"
)

func init() {
	service.Template.AddFunc("path", templatePathFunc)
}

func templatePathFunc(a jet.Arguments) reflect.Value {
	var args []string
	a.RequireNumOfArguments("path", 1, -1)

	name := a.Get(0)
	if a.NumOfArguments() > 1 {
		for i := 1; i < a.NumOfArguments(); i++ {
			args = append(args, a.Get(i).String())
		}
	}

	url, err := http.Routing.Mux().Get(name.String()).URL(args...)
	if err != nil {
		cli.Logger.Get(cli.AppLog).(*log.Logger).Panic(err)
	}

	return reflect.ValueOf(url.String())
}
