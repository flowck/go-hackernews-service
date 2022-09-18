package main

import (
	"context"
	"go-hackernews-service/internal/app"
	"log"
	"os"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	application := app.NewApp()

	if err := application.PostService.PublishPost(ctx, "Hello World", "My very first post"); err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
