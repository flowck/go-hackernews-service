package app

import "go-hackernews-service/internal/components/post"

type App struct {
	PostService post.Service
}

func NewApp() App {
	// Assemble all the components here and expose their services
	var err error
	var postService post.Service

	if postService, err = post.NewService(post.NewPsqlRepository()); err != nil {
		panic(err)
	}

	return App{postService}
}
