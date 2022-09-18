package post

import (
	"context"
	"errors"
	"log"
)

type Service struct {
	postRepository Repository
}

func NewService(postRepository Repository) (Service, error) {

	if postRepository == nil {
		return Service{}, errors.New("postRepository cannot be nil")
	}

	return Service{postRepository}, nil
}

func (s *Service) PublishPost(ctx context.Context, title string, content string) error {
	log.Println("Publishing post", title, content)
	return nil
}

func (s *Service) EditPost(ctx context.Context, id, title, content string) (Post, error) {
	return Post{}, nil
}
