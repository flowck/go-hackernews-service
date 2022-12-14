package main

import (
	"context"
	"go-hackernews-service/internal/app"
	"go-hackernews-service/internal/infra/config"
	"go-hackernews-service/internal/infra/database"
	"log"
	"os"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cfg := config.GetConfig()
	db := database.GetDb(cfg.DbDriverName, cfg.DbUrl, cfg.DbMaxConnections, cfg.DbMaxIdleConnections)
	application := app.NewApp(db)

	if err := application.PostService.PublishPost(ctx, "Hello World", "My very first post"); err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
