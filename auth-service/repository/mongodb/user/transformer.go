package user

import (
	"app/internal/core/domain"
	"app/repository/mongodb/common"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type transformer struct {
}

func (t *transformer) toModel(e domain.User) User {
	return User{
		Email:     e.Email,
		Password:  e.Password,
		FirstName: e.FirstName,
		LastName:  e.LastName,
		Model: common.Model{
			CreatedAt: e.CreatedAt,
			UpdatedAt: e.UpdatedAt,
		},
	}
}

func (t *transformer) toDomain(p User) domain.User {
	return domain.User{
		ID:        p.ID.Hex(),
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		Email:     p.Email,
		Password:  p.Password,
		FirstName: p.FirstName,
		LastName:  p.LastName,
	}
}

func (t *transformer) toManyDomain(p []User) []domain.User {
	resp := []domain.User{}
	for _, v := range p {
		resp = append(resp, t.toDomain(v))
	}

	return resp
}

func (t *transformer) buildQueryFilter(q domain.Query) primitive.M {
	filter := primitive.M{}
	return filter
}
