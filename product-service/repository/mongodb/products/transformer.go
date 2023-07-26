package products

import (
	"app/internal/core/domain"
	"app/repository/mongodb/common"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type transformer struct {
}

func (t *transformer) toModel(e domain.Product) Product {
	return Product{
		Name:  e.Name,
		Price: e.Price,
		Model: common.Model{
			CreatedAt: e.CreatedAt,
			UpdatedAt: e.UpdatedAt,
		},
	}
}

func (t *transformer) toDomain(p Product) domain.Product {
	return domain.Product{
		ID:        p.ID.Hex(),
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		Name:      p.Name,
		Price:     p.Price,
	}
}

func (t *transformer) toManyDomain(p []Product) []domain.Product {
	resp := []domain.Product{}
	for _, v := range p {
		resp = append(resp, t.toDomain(v))
	}

	return resp
}

func (t *transformer) buildQueryFilter(q domain.Query) primitive.M {
	filter := primitive.M{}
	return filter
}
