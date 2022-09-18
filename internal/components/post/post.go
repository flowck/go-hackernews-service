package post

import "context"

type Post struct {
	Id string
}

type Repository interface {
	Insert(ctx context.Context, post Post) error
}
