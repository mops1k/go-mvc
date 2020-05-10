package mvc

import (
	"context"
	"fmt"
	"log"
	netHttp "net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/common-nighthawk/go-figure"
	"github.com/gookit/color"
	"github.com/spf13/cast"

	"github.com/mops1k/go-mvc/cli"
	cmd "github.com/mops1k/go-mvc/cli/command"
	"github.com/mops1k/go-mvc/http"
	"github.com/mops1k/go-mvc/service"
	"github.com/mops1k/go-mvc/service/command"
)

var (
	srv    *http.Server
	appLog *log.Logger
)

func init() {
	logo := figure.NewFigure("go-mvc", "isometric1", true)
	color.Green.Println(logo.String())

	cli.Logger.Set(cli.AppLog, log.New(os.Stdout, "[app] ", log.LstdFlags))
	appLog = cli.Logger.Get(cli.AppLog).(*log.Logger)
	srv = http.GetServer(
		service.Config.GetString("server.host"),
		service.Config.GetInt("server.port"),
		cli.Logger.Get(cli.ErrorLog).(*log.Logger),
	)
	srv.SetTimeouts(
		cast.ToUint16(service.Config.GetInt("server.timeout.read")),
		cast.ToUint16(service.Config.GetInt("server.timeout.write")),
		cast.ToUint16(service.Config.GetInt("server.timeout.idle")),
	)
}

func Run() {
	appLog.Printf("Application has started at %s\n", srv)

	go func() {
		go func() {
			if err := srv.ListenAndServe(); err != nil {
				appLog.Println(err)
			}
		}()

		appLog.Println(`Type "help" for list available commands`)
		cc := service.Commands
		cc.Add(&cmd.RoutingCommand{})
		cc.Add(&cmd.HelpCommand{})

		var c string
		for {
			color.Green.Print("> ")
			_, _ = fmt.Scanln(&c)
			parser := command.NewParser()
			parser.Parse(c)
			if cc.Has(parser.Ctx().Command()) {
				c := cc.Get(parser.Ctx().Command())
				c.Action(parser.Ctx())
			} else {
				appLog.Printf(`Unknown command "%s".`, parser.Ctx().Command())
			}
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-done
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		appLog.Println("Stopping application services")
		service.Manager.Close()
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		appLog.Fatalf("Application Stopping Failed:%+v", err)
	}
	appLog.Println("Application has stopped")
}

func HttpMiddleware(middleware http.Middleware) {
	srv.Middleware(middleware)
}

func HttpHandler(h netHttp.Handler) {
	srv.SetHandler(h)
}
