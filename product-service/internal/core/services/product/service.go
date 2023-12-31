package product

import (
	"app/internal/core/domain"
	"app/internal/core/ports"
	"context"
	"net/http"
	"time"

	"github.com/ohmspeed777/go-pkg/errorx"
)

type Dependencies struct {
	ProductRepository ports.ProductRepository
}

type service struct {
	ProductRepository ports.ProductRepository
}

func NewService(d Dependencies) ports.ProductService {
	return &service{
		ProductRepository: d.ProductRepository,
	}
}

func (s *service) FindAll(ctx context.Context, req domain.Query) (res domain.ResponseInfo[[]domain.Product], err error) {
	entity, counter, err := s.ProductRepository.FindAll(ctx, req)
	if err != nil {
		return res, errorx.New(http.StatusUnprocessableEntity, "can not find products")
	}

	return domain.NewPagination(entity, req.GetPage(), req.GetLimit(), counter), nil
}

func (s *service) Create(ctx context.Context, req domain.Product) (res domain.Product, err error) {
	req.CreatedAt = time.Now()
	req.UpdatedAt = time.Now()
	entity, err := s.ProductRepository.Create(ctx, req)
	if err != nil {
		return res, errorx.New(http.StatusUnprocessableEntity, "can not create products")
	}

	return entity, nil
}
