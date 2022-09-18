package post

import "context"

type PsqlRepository struct {
}

func NewPsqlRepository() PsqlRepository {
	return PsqlRepository{}
}

func (r PsqlRepository) Insert(ctx context.Context, post Post) error {
	return nil
}
