package services

import (
	"app/configs"
	"app/internal/core/ports"
	"app/internal/core/services/user"
	orders "app/protobufs/orders"
	"app/repository/mongodb"

	"google.golang.org/grpc"
)

type Dependencies struct {
	Conf       *configs.Config
	Repository *mongodb.Repository
	CC         grpc.ClientConnInterface
}

type Service struct {
	UserService ports.UserService
}

func NewService(d Dependencies) *Service {
	return &Service{
		UserService: user.NewService(user.Dependencies{
			UserRepository: d.Repository.UserRepository,
			Key:            d.Conf.JWT.PRIV,
			OrderGRPC:      orders.NewOrderServiceClient(d.CC),
		}),
	}
}
