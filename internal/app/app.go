package app

import (
	"context"
	"log"
	"net/http"
)

type App struct {
	serviceProvider *serviceProvider

	httpServer *http.Server
}

func NewApp(ctx context.Context) (*App, error) {
	var a App

	//err := a.initDeps(ctx)
	//if err != nil {
	//	return nil, err
	//}

	return &a, nil
}

func (a *App) Run() error {
	log.Println("App is running")
	return nil
}
