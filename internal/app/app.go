package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	httpV1 "github.com/zd4r/dynamic-user-segmentation/internal/api/http/v1"
	"github.com/zd4r/dynamic-user-segmentation/internal/config"
	"github.com/zd4r/dynamic-user-segmentation/pkg/closer"
	"github.com/zd4r/dynamic-user-segmentation/pkg/httpserver"
)

const serviceName = "api"

type App struct {
	serviceProvider *serviceProvider

	httpServer *httpserver.Server
}

func NewApp(ctx context.Context) (*App, error) {
	var a App

	if err := a.initDeps(ctx); err != nil {
		return nil, err
	}

	return &a, nil
}

func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	a.runHTTPServer()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGTERM, os.Interrupt)

	select {
	case s := <-interrupt:
		log.Printf("interrupt signal: %s", s.String())
	case err := <-a.httpServer.Notify():
		log.Printf("a.httpServer.Notify error: %s", err.Error())
	}

	if err := a.httpServer.Shutdown(); err != nil {
		return err
	}

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		config.Init,
		a.initServiceProvider,
		a.initHTTPServer,
	}

	for _, f := range inits {
		if err := f(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = NewServiceProvider()

	return nil
}

func (a *App) initHTTPServer(ctx context.Context) error {
	handler := httpV1.NewRouter(
		a.serviceProvider.GetUserService(ctx),
		a.serviceProvider.GetSegmentService(ctx),
		a.serviceProvider.GetExperimentService(ctx),
		a.serviceProvider.GetReportService(ctx),
		a.serviceProvider.GetLogger(serviceName),
	)

	a.httpServer = httpserver.New(
		handler,
		httpserver.Port(a.serviceProvider.GetHTTPConfig().Port()),
	)

	return nil
}

func (a *App) runHTTPServer() {
	log.Printf("HTTP server is running on %s\n", a.serviceProvider.GetHTTPConfig().Port())
	a.httpServer.Start()
}
