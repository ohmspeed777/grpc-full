package ports

import (
	"app/internal/core/domain"
	"context"
)

type ProductService interface {
	FindAll(ctx context.Context, req domain.Query) (res domain.ResponseInfo[[]domain.Product], err error)
}

type UserService interface {
	SignIn(ctx context.Context, req domain.SignInReq) (*domain.SignInRes, error)
	SignUp(ctx context.Context, req domain.SignUpReq) error
	GetMyOrder(ctx context.Context, userID string) ([]*domain.Order, error)
	ClientStream(ctx context.Context) error
	BidirectionalStream(ctx context.Context) error
}
