package post

import (
	"context"
	"errors"
	"github.com/google/uuid"
)

type Post struct {
	Id      string
	Title   string
	Content string
}

type Repository interface {
	Insert(ctx context.Context, post Post) error
}

func NewPost(title, content string) (Post, error) {
	if title == "" {
		return Post{}, errors.New("title cannot be empty")
	}

	if content == "" {
		return Post{}, errors.New("content cannot be empty")
	}

	return Post{
		Id:      uuid.NewString(),
		Title:   title,
		Content: content,
	}, nil
}
