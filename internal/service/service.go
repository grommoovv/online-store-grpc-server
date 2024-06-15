package service

import (
	"context"
	"log/slog"
	"online-store-server/internal/domain"
	"online-store-server/internal/repository"
)

type (
	Product interface {
		GetAll(ctx context.Context) ([]domain.Product, error)
		GetByID(ctx context.Context, id int) (domain.Product, error)
	}

	Service struct {
		Product Product
	}
)

func New(log *slog.Logger, repo *repository.Repository) *Service {
	return &Service{
		Product: NewProductService(log, repo),
	}
}
