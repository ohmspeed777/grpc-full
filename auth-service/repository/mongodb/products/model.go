package products

import "app/repository/mongodb/common"

type Product struct {
	common.Model `bson:",inline"`
	Name         string  `bson:"name"`
	Price        float64 `bson:"price"`
}
