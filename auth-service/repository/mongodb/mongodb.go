package mongodb

import (
	"app/configs"
	"app/internal/core/ports"
	"app/repository/mongodb/user"

	"go.mongodb.org/mongo-driver/mongo"
)

type Dependencies struct {
	Config      *configs.Config
	MongoClient *mongo.Client
}

type Repository struct {
	UserRepository ports.UserRepository
}

func NewRepository(d Dependencies) *Repository {
	return &Repository{
		UserRepository: user.NewRepo(user.Dependencies{
			DB: d.MongoClient.Database(d.Config.MongoDB.Database),
		}),
	}
}
