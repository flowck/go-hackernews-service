package main

import (
	"context"
	"fmt"
	"go-hackernews-service/internal/app"
	"go-hackernews-service/internal/infra/config"
	"go-hackernews-service/internal/infra/database"
	"log"
	"os"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	application := app.NewApp()

	cfg := config.GetConfig()
	db := database.GetDb(cfg.DbDriverName, cfg.DbUrl, cfg.DbMaxConnections, cfg.DbMaxIdleConnections)

	fmt.Println(db)

	if err := application.PostService.PublishPost(ctx, "Hello World", "My very first post"); err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
