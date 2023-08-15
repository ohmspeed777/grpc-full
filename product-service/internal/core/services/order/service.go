package order

import (
	"app/internal/core/domain"
	"app/internal/core/ports"
	"context"
	"net/http"
	"time"

	"github.com/ohmspeed777/go-pkg/errorx"
)

type Dependencies struct {
	OrderRepository ports.OrderRepository
}

type service struct {
	OrderRepository ports.OrderRepository
}

func NewService(d Dependencies) ports.OrderService {
	return &service{
		OrderRepository: d.OrderRepository,
	}
}

func (s *service) FindAll(ctx context.Context, req domain.Query) (res domain.ResponseInfo[[]domain.Order], err error) {
	entity, counter, err := s.OrderRepository.FindAll(ctx, req)
	if err != nil {
		return res, errorx.New(http.StatusUnprocessableEntity, "can not find products")
	}

	return domain.NewPagination(entity, req.GetPage(), req.GetLimit(), counter), nil
}

func (s *service) Create(ctx context.Context, req domain.Order) (res domain.Order, err error) {
	req.CreatedAt = time.Now()
	req.UpdatedAt = time.Now()
	entity, err := s.OrderRepository.Create(ctx, req)
	if err != nil {
		return res, errorx.New(http.StatusUnprocessableEntity, "can not create products")
	}

	return entity, nil
}
