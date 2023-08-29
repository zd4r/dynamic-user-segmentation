package main

import (
	"context"
	"log"

	"github.com/zd4r/dynamic-user-segmentation/internal/app"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to create new app: %s", err.Error())
	}

	if err := a.Run(); err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}
}
