package ports

import (
	"app/internal/core/domain"
	"context"
)

type ProductRepository interface {
	FindAll(ctx context.Context, q domain.Query) (empty []domain.Product, counter int64, err error)
	Create(ctx context.Context, e domain.Product) (empty domain.Product, err error)
}
