package handler

import (
	"app/configs"
	"app/handler/product"
	"app/internal/core/services"
)

type Handler struct {
	Product *product.Handler
}

type Dependencies struct {
	Config  configs.Config
	Service *services.Service
}

func NewHandler(d Dependencies) *Handler {
	return &Handler{
		Product: product.NewHandler(product.Dependencies{
			ProductService: d.Service.ProductService,
		}),
	}
}
