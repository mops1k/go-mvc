package http

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"

	"github.com/gorilla/mux"

	"github.com/mops1k/go-mvc/cli"
)

type routing struct {
	mux *mux.Router
}

var Routing *routing

func init() {
	Routing = &routing{mux: mux.NewRouter()}
	staticDirName := "static"
	Routing.mux.PathPrefix("/" + staticDirName).
		Handler(http.StripPrefix(
			"/"+staticDirName+"/",
			http.FileServer(http.Dir("./"+staticDirName+"/")))).Name(staticDirName)

}

func (r *routing) HandleControllers() {
	for _, controller := range Controllers.All() {
		r.addController(controller)
	}
}

func (r *routing) Mux() *mux.Router {
	return r.mux
}

func (r *routing) addController(c Controller) {
	path, pathName, methods := r.getPathInfoForController(c)
	r.mux.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
		var context = &Context{response: writer, request: request, statusCode: http.StatusOK, headers: make(map[string]string)}

		content := c.Action(context)

		if _, exists := context.headers["Content-Type"]; !exists {
			context.headers["Content-Type"] = "text/html"
		}

		if context.headers != nil {
			for name, value := range context.headers {
				writer.Header().Add(name, value)
			}
		}
		writer.WriteHeader(context.statusCode)

		_, err := fmt.Fprint(writer, content)
		if err != nil {
			cli.Logger.Get(cli.HttpLog).(*log.Logger).Fatal(err)
		}
	}).Methods(methods...).Name(pathName)
}

func (r *routing) setDefaultMethods(methods []string) []string {
	if len(methods) == 0 {
		methods = append(methods, http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete)
	}

	return methods
}

func (r *routing) Path(name string, args map[string]string) (string, error) {
	var opts []string
	for name, value := range args {
		opts = append(opts, name, value)
	}

	url, err := r.mux.Get(name).URL(opts...)
	if err != nil {
		return "", err
	}

	return url.String(), nil
}

func (r *routing) getPathInfoForController(c Controller) (path string, name string, methods []string) {
	v := reflect.ValueOf(c)
	i := reflect.Indirect(v)
	s := i.Type()
	field, ok := s.FieldByName("PathInfo")
	if !ok {
		return
	}
	path, ok = field.Tag.Lookup("path")
	if !ok {
		return
	}
	name, ok = field.Tag.Lookup("name")
	if !ok {
		return
	}
	methodsString, ok := field.Tag.Lookup("methods")
	if !ok {
		return
	}
	methods = strings.Split(methodsString, ",")

	return
}
