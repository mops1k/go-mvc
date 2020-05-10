package mvc

import (
	"context"
	"log"
	netHttp "net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	appLog.Printf("Application has started at %s\n", srv)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			appLog.Println(err)
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
