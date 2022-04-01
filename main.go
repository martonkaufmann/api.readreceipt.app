package main

import (
	"context"
	"log"

	"github.com/readreceipt/api/handler"
	"github.com/readreceipt/api/server"
	"github.com/readreceipt/api/service/cache"
	"github.com/readreceipt/api/service/database"
	"github.com/readreceipt/api/service/monitoring"
)

func main() {
	if err := monitoring.Init(); err != nil {
		log.Fatalf("monitoring.Init: %s", err)
	}

	if err := cache.Init(); err != nil {
		log.Fatalf("cache.Init: %s", err)
	}

	if err := database.Init(); err != nil {
		log.Fatalf("database.Init: %s", err)
	}
	defer database.Client().Disconnect(context.TODO())

	e := server.New()

	handler.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":80"))
}
