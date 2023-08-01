package services

import (
	"app/configs"
	"app/internal/core/ports"
	"app/internal/core/services/user"
	"app/repository/mongodb"
)

type Dependencies struct {
	Conf       *configs.Config
	Repository *mongodb.Repository
}

type Service struct {
	UserService ports.UserService
}

func NewService(d Dependencies) *Service {
	return &Service{
		UserService: user.NewService(user.Dependencies{
			UserRepository: d.Repository.UserRepository,
			Key:            d.Conf.JWT.PRIV,
		}),
	}
}
