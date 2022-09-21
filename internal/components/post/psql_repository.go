package post

import (
	"context"
	"database/sql"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"go-hackernews-service/internal/infra/models"
)

type PsqlRepository struct {
	db *sql.DB
}

func NewPsqlRepository(db *sql.DB) PsqlRepository {
	if db == nil {
		panic("db pointer is missing")
	}

	return PsqlRepository{db}
}

func (r PsqlRepository) Insert(ctx context.Context, post Post) error {
	row := models.Post{ID: post.Id, Title: post.Title}

	if err := row.Insert(ctx, r.db, boil.Infer()); err != nil {
		return err
	}

	return nil
}
