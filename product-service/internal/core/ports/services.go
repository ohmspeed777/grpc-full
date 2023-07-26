package ports

import (
	"app/internal/core/domain"
	"context"
)

type ProductService interface {
	FindAll(ctx context.Context, req domain.Query) (res domain.ResponseInfo[[]domain.Product], err error)
	Create(ctx context.Context, req domain.Product) (res domain.Product, err error)
}
