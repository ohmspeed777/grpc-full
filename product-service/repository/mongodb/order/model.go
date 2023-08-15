package order

import (
	"app/repository/mongodb/common"
	"app/repository/mongodb/products"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	common.Model `bson:",inline"`
	Items        []*OrderItem       `bson:"items"`
	Total        float64            `bson:"total"`
	UserID       primitive.ObjectID `bson:"user_id"`
}

type OrderItem struct {
	Quantity  uint               `bson:"quantity"`
	ProductID primitive.ObjectID `bson:"product_id"`
}

type OrderLookedUp struct {
	Order          `bson:",inline"`
	ProductsJoined []*products.Product ` bson:"products_joined"`
}
