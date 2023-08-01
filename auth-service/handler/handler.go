package handler

import (
	"app/configs"
	"app/handler/user"
	"app/internal/core/services"
)

type Handler struct {
	User *user.Handler
}

type Dependencies struct {
	Config  configs.Config
	Service *services.Service
}

func NewHandler(d Dependencies) *Handler {
	return &Handler{
		User: user.NewHandler(user.Dependencies{
			UserService: d.Service.UserService,
		}), 
	}
}
