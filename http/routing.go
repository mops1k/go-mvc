package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/mops1k/go-mvc/cli"
	"github.com/mops1k/go-mvc/service"
)

type Routing struct {
	mux *mux.Router
}

func GetRouting() *Routing {
	r := &Routing{mux: mux.NewRouter()}
	staticDirName := "static"
	r.mux.PathPrefix("/" + staticDirName).
		Handler(http.StripPrefix(
			"/"+staticDirName+"/",
			http.FileServer(http.Dir("./"+staticDirName+"/")))).Name(staticDirName)

	return r
}

func (r *Routing) addController(c Controller, methods ...string) {
	methods = r.setDefaultMethods(methods)
	pathName, path := c.Name(), service.Config.GetString(c.Name()+".path")
	r.mux.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
		var context = &Context{response: writer, request: request, statusCode: http.StatusOK}
		content := c.Action(context)

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

func (r *Routing) setDefaultMethods(methods []string) []string {
	if len(methods) == 0 {
		methods = append(methods, http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete)
	}

	return methods
}

func (r *Routing) HandleControllers() {
	for _, controller := range Controllers.All() {
		methods := service.Config.GetStringSlice(controller.Name() + ".methods")
		r.addController(controller, methods...)
	}
}
