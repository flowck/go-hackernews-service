package post

import (
	"context"
	"errors"
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
	newPost, err := NewPost(title, content)
	if err != nil {
		return err
	}

	return s.postRepository.Insert(ctx, newPost)
}

func (s *Service) EditPost(ctx context.Context, id, title, content string) (Post, error) {
	return Post{}, nil
}
