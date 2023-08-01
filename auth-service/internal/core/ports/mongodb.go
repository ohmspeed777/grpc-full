package ports

import (
	"app/internal/core/domain"
	"context"
)

type ProductRepository interface {
	FindAll(ctx context.Context, q domain.Query) (empty []domain.Product, counter int64, err error)
}

type UserRepository interface {
	FindOneByID(ctx context.Context, id string) (empty domain.User, err error) 
	Create(ctx context.Context, e domain.User) (empty domain.User, err error)
	FindOneByEmail(ctx context.Context, email string) (empty domain.User, err error)
}
