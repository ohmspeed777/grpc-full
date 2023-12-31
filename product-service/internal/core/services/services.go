package services

import (
	"app/configs"
	"app/internal/core/ports"
	"app/internal/core/services/order"
	"app/internal/core/services/product"
	"app/repository/mongodb"
)

type Dependencies struct {
	Conf       *configs.Config
	Repository *mongodb.Repository
}

type Service struct {
	ProductService ports.ProductService
	OrderService   ports.OrderService
}

func NewService(d Dependencies) *Service {
	return &Service{
		ProductService: product.NewService(product.Dependencies{
			ProductRepository: d.Repository.ProductRepository,
		}),
		OrderService: order.NewService(order.Dependencies{
			OrderRepository: d.Repository.OrderRepository,
		}),
	}
}
