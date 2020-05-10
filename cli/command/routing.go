package command

import (
	"os"

	"github.com/gorilla/mux"

	"github.com/jedib0t/go-pretty/table"

	"github.com/mops1k/go-mvc/http"
	"github.com/mops1k/go-mvc/service/command"
)

type RoutingCommand struct{}

func (e *RoutingCommand) Name() string {
	return "routing"
}

func (e *RoutingCommand) Description() string {
	return "List registered routes"
}

func (e *RoutingCommand) Action(ctx command.Context) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Name", "URI Template", "Methods"})
	_ = http.Routing.Mux().Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, err := route.GetPathTemplate()
		methods, _ := route.GetMethods()
		if err != nil {
			return err
		}

		t.AppendRows([]table.Row{
			{route.GetName(), path, methods},
		})

		return nil
	})
	t.Render()
}
