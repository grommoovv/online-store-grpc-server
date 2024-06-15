package service

import (
	"context"
	"log/slog"
	"online-store-server/internal/domain"
	"online-store-server/internal/repository"
)

type ProductService struct {
	log  *slog.Logger
	repo repository.Product
}

func NewProductService(log *slog.Logger, repo repository.Product) *ProductService {
	return &ProductService{
		log:  log,
		repo: repo,
	}
}

func (p *ProductService) GetAll(ctx context.Context) ([]domain.Product, error) {
	return p.repo.GetAll(ctx)
}

func (p *ProductService) GetByID(ctx context.Context, id int) (domain.Product, error) {
	return p.repo.GetByID(ctx, id)
}
