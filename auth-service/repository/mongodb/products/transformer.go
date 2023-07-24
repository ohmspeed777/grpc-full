package products

import (
	"app/internal/core/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type transformer struct {
}

func (t *transformer) toModel() {}

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
