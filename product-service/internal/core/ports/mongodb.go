package ports

import (
	"app/internal/core/domain"
	"context"
)

type ProductRepository interface {
	FindAll(ctx context.Context, q domain.Query) (empty []domain.Product, counter int64, err error)
	Create(ctx context.Context, e domain.Product) (empty domain.Product, err error)
}

type OrderRepository interface {
	Create(ctx context.Context, e domain.Order) (empty domain.Order, err error)
	FindAll(ctx context.Context, q domain.Query) (empty []domain.Order, counter int64, err error)
}
