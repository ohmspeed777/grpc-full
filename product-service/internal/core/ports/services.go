package ports

import (
	"app/internal/core/domain"
	"context"
)

type ProductService interface {
	FindAll(ctx context.Context, req domain.Query) (res domain.ResponseInfo[[]domain.Product], err error)
	Create(ctx context.Context, req domain.Product) (res domain.Product, err error)
}

type OrderService interface {
	FindAll(ctx context.Context, req domain.Query) (res domain.ResponseInfo[[]domain.Order], err error)
	Create(ctx context.Context, req domain.Order) (res domain.Order, err error)
}