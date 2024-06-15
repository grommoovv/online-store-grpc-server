package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"log/slog"
	"online-store-server/internal/domain"
)

type (
	Product interface {
		GetAll(ctx context.Context) ([]domain.Product, error)
		GetByID(ctx context.Context, id int) (domain.Product, error)
	}

	Repository struct {
		Product
	}
)

func New(log *slog.Logger, db *sqlx.DB) *Repository {
	return &Repository{
		Product: NewProductRepository(log, db),
	}
}
