package mvc

import (
	"bufio"
	"context"
	"flag"
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
	appCfg "github.com/mops1k/go-mvc/config"
	"github.com/mops1k/go-mvc/http"
	"github.com/mops1k/go-mvc/service"
	"github.com/mops1k/go-mvc/service/command"
)

var (
	srv    *http.Server
	appLog *log.Logger
	config *appCfg.ServerConfig
)

// init mvc
func init() {
	logo := figure.NewFigure("go-mvc", "isometric1", true)
	color.Green.Println(logo.String())

	appLog = cli.Logger.Get(cli.AppLog).(*log.Logger)

	config = &appCfg.ServerConfig{
		Port:     service.Config.GetInt("server.port"),
		Host:     service.Config.GetString("server.host"),
		CertFile: service.Config.GetString("server.tls.cert_file"),
		KeyFile:  service.Config.GetString("server.tls.key_file"),
		Timeouts: make(map[string]uint16),
	}

	srv = http.GetServer(
		config,
		cli.Logger.Get(cli.ErrorLog).(*log.Logger),
	)

	srv.SetTimeouts(
		cast.ToUint16(service.Config.GetInt("server.timeout.read")),
		cast.ToUint16(service.Config.GetInt("server.timeout.write")),
		cast.ToUint16(service.Config.GetInt("server.timeout.idle")),
	)
}

// Run mvc application
func Run() {
	appLog.Printf("Application has started at %s\n", srv)
	flag.Parse()

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

		parser := command.NewParser()
		go func() {
			for _, flagCmd := range flag.Args() {
				parser.Parse(flagCmd)
				if cc.Exists(parser.Ctx().Command()) {
					cliCmd := cc.Get(parser.Ctx().Command())
					cliCmd.Action(parser.Ctx())
				} else {
					appLog.Printf(`Unknown command "%s".`, parser.Ctx().Command())
				}
			}
		}()

		var c string
		for {
			color.Green.Print("> ")
			reader := bufio.NewReader(os.Stdin)
			c, _ = reader.ReadString('\n')

			parser.Parse(c[:])
			if cc.Exists(parser.Ctx().Command()) {
				cliCmd := cc.Get(parser.Ctx().Command())
				cliCmd.Action(parser.Ctx())
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

// Add middleware to http server
func HttpMiddleware(middleware http.Middleware) {
	srv.Middleware(middleware)
}

// Update custom http hanler
func HttpHandler(h netHttp.Handler) {
	srv.SetHandler(h)
}
