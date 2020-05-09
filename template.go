package mvc

import (
	"log"
	"reflect"

	"github.com/CloudyKit/jet"
	"golang.org/x/text/language"

	"github.com/mops1k/go-mvc/cli"
	"github.com/mops1k/go-mvc/http"
	"github.com/mops1k/go-mvc/service"
)

func init() {
	service.Template.AddFunc("path", templatePathFunc)
	service.Template.AddFunc("translate", templateTransFunc)
}

func templatePathFunc(a jet.Arguments) reflect.Value {
	a.RequireNumOfArguments("path", 1, -1)
	var args []string

	name := a.Get(0)
	if a.NumOfArguments() > 1 {
		for i := 1; i < a.NumOfArguments(); i++ {
			args = append(args, a.Get(i).String())
		}
	}

	url, err := http.Routing.Mux().Get(name.String()).URL(args...)
	if err != nil {
		cli.Logger.Get(cli.ErrorLog).(*log.Logger).Panic(err)
	}

	return reflect.ValueOf(url.String())
}

func templateTransFunc(a jet.Arguments) reflect.Value {
	a.RequireNumOfArguments("translate", 1, -1)
	key := a.Get(0).String()
	var result string
	switch a.NumOfArguments() {
	case 1:
		result = service.Translation.Trans(key, nil)
	case 2:
		result = service.Translation.TransFor(key, nil, language.Make(a.Get(1).String()))
	default:
		hasLocale := true
		args := make(map[string]interface{})
		if a.NumOfArguments()%2 != 0 {
			hasLocale = false
		}

		count := a.NumOfArguments()
		if !hasLocale {
			count--
		}

		for i := 1; i < count; i++ {
			args[a.Get(i).String()] = a.Get(i + 1)
			i++
		}

		if hasLocale {
			result = service.Translation.TransFor(key, args, language.Make(a.Get(count-1).String()))
		} else {
			result = service.Translation.Trans(key, args)
		}
	}

	return reflect.ValueOf(result)
}
