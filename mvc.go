package mvc

import (
	"log"
	netHttp "net/http"
	"os"

	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cast"

	"github.com/mops1k/go-mvc/cli"
	"github.com/mops1k/go-mvc/http"
	"github.com/mops1k/go-mvc/service"
)

var (
	srv    *http.Server
	appLog *log.Logger
)

func init() {
	logo := figure.NewFigure("go-mvc", "isometric1", true)
	logo.Print()

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
	defer service.Manager.Close()
	appLog.Printf("Application has started at %s\n", srv)

	if err := srv.ListenAndServe(); err != nil {
		appLog.Println(err)
	}
}

func HttpMiddleware(middleware http.Middleware) {
	srv.Middleware(middleware)
}

func HttpHandler(h netHttp.Handler) {
	srv.SetHandler(h)
}
