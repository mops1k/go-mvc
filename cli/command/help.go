package command

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/table"

	"github.com/mops1k/go-mvc/service"
	"github.com/mops1k/go-mvc/service/command"
)

type HelpCommand struct{}

func (h HelpCommand) Name() string {
	return "help"
}

func (h HelpCommand) Description() string {
	return "Show all commands"
}

func (h HelpCommand) Action(ctx command.Context) {
	fmt.Println("Available commands:")
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Name", "Description"})
	for service.Commands.Next() {
		if cmd, ok := service.Commands.Current().(service.Command); ok {
			if cmd.Name() == h.Name() {
				continue
			}

			t.AppendRows([]table.Row{
				{cmd.Name(), cmd.Description()},
			})
		}
	}

	t.Render()
}
