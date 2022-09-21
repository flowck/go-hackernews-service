package app

import (
	"database/sql"
	"go-hackernews-service/internal/components/post"
)

type App struct {
	PostService post.Service
}

func NewApp(db *sql.DB) App {
	// Assemble all the components here and expose their services
	var err error
	var postService post.Service

	if postService, err = post.NewService(post.NewPsqlRepository(db)); err != nil {
		panic(err)
	}

	return App{postService}
}
