package mvc

import (
	"fmt"
	"log"

	"github.com/tyler-sommer/stick"
	"golang.org/x/text/language"

	"github.com/mops1k/go-mvc/cli"
	"github.com/mops1k/go-mvc/http"
	"github.com/mops1k/go-mvc/service"
)

func init() {
	service.Template.AddFunc("path", templatePathFunc)
	service.Template.AddFunc("trans", templateTransFunc)
	service.Template.UseFuncAsFilter("path", templatePathFunc)
	service.Template.UseFuncAsFilter("trans", templateTransFunc)
}

func templatePathFunc(ctx stick.Context, args ...stick.Value) stick.Value {
	var values []string
	path := stick.CoerceString(args[0])
	for i := len(args); i < 3; i++ {
		args = append(args, stick.Value(nil))
	}

	if args[1] != nil && stick.IsMap(args[1]) {
		_, err := stick.Iterate(args[1], func(k, v stick.Value, l stick.Loop) (brk bool, err error) {
			values = append(values, stick.CoerceString(k), stick.CoerceString(v))

			return true, nil
		})

		if err != nil {
			cli.Logger.Get(cli.ErrorLog).(*log.Logger).Panic(err)
		}
	}

	route := http.Routing.Mux().GetRoute(path)

	url, err := route.URL(values...)
	if err != nil {
		cli.Logger.Get(cli.ErrorLog).(*log.Logger).Panic(err)
	}

	if args[2] != nil && stick.IsMap(args[2]) {
		_, err := stick.Iterate(args[2], func(k, v stick.Value, l stick.Loop) (brk bool, err error) {
			url.RawQuery += fmt.Sprintf("%s=%s", stick.CoerceString(k), stick.CoerceString(v))
			if !l.Last {
				url.RawQuery += "&"
			}

			return true, nil
		})

		if err != nil {
			cli.Logger.Get(cli.ErrorLog).(*log.Logger).Panic(err)
		}
	}

	return stick.Value(url.String())
}

func templateTransFunc(ctx stick.Context, args ...stick.Value) stick.Value {
	for i := len(args); i < 3; i++ {
		args = append(args, stick.Value(nil))
	}

	var values map[string]interface{}
	key := stick.CoerceString(args[0])

	var result string
	if args[1] != nil && stick.IsMap(args[1]) {
		values = make(map[string]interface{})
		_, err := stick.Iterate(args[1], func(k, v stick.Value, l stick.Loop) (brk bool, err error) {
			values[stick.CoerceString(k)] = stick.CoerceString(v)

			return true, nil
		})

		if err != nil {
			cli.Logger.Get(cli.ErrorLog).(*log.Logger).Panic(err)
		}
	}

	if args[2] == nil {
		result = service.Translation.Trans(key, values)
	} else {
		locale := stick.CoerceString(args[2])
		result = service.Translation.TransFor(key, values, language.Make(locale))
	}

	return stick.Value(result)
}
