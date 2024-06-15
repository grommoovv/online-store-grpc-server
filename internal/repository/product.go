package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"online-store-server/internal/database/postgres"
	"online-store-server/internal/domain"

	"github.com/jmoiron/sqlx"
)

var (
	ErrProductsNotFound = errors.New("products not found")
	ErrProductNotFound  = errors.New("product not found")
)

type ProductRepository struct {
	log *slog.Logger
	db  *sqlx.DB
}

func NewProductRepository(log *slog.Logger, db *sqlx.DB) *ProductRepository {
	return &ProductRepository{
		log: log,
		db:  db,
	}
}

func (p *ProductRepository) GetAll(ctx context.Context) ([]domain.Product, error) {
	const op = "ProductRepository.GetAll"

	var res []domain.Product

	q := fmt.Sprintf("SELECT * FROM %s ORDER BY id DESC", postgres.ProductsTable)
	if err := p.db.SelectContext(ctx, &res, q); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, ErrProductsNotFound)
		}

		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return res, nil
}

func (p *ProductRepository) GetByID(ctx context.Context, id int) (domain.Product, error) {
	const op = "ProductRepository.GetByID"

	var res domain.Product

	q := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", postgres.ProductsTable)

	if err := p.db.QueryRowContext(ctx, q, id).Scan(&res.ID, &res.Title, &res.Description, &res.Price, &res.ImageURL, &res.Category); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Product{}, fmt.Errorf("%s: %w", op, ErrProductNotFound)
		}

		return domain.Product{}, fmt.Errorf("%s: %w", op, err)
	}

	return res, nil
}
