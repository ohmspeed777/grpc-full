package mongodb

import (
	"app/configs"
	"app/internal/core/ports"
	"app/repository/mongodb/order"
	"app/repository/mongodb/products"

	"go.mongodb.org/mongo-driver/mongo"
)

type Dependencies struct {
	Config      *configs.Config
	MongoClient *mongo.Client
}

type Repository struct {
	ProductRepository ports.ProductRepository
	OrderRepository   ports.OrderRepository
}

func NewRepository(d Dependencies) *Repository {
	return &Repository{
		ProductRepository: products.NewRepo(products.Dependencies{
			DB: d.MongoClient.Database(d.Config.MongoDB.Database),
		}),
		OrderRepository: order.NewRepo(order.Dependencies{
			DB: d.MongoClient.Database(d.Config.MongoDB.Database),
		}),
	}
}
