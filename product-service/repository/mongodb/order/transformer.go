package order

import (
	"app/internal/core/domain"
	"app/repository/mongodb/common"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type transformer struct {
}

func (t *transformer) toModel(e domain.Order) Order {
	items := []*OrderItem{}
	for _, v := range e.Items {
		pid, _ := primitive.ObjectIDFromHex(v.ProductID)
		items = append(items, &OrderItem{
			Quantity:  v.Quantity,
			ProductID: pid,
		})
	}
	uid, _ := primitive.ObjectIDFromHex(e.UserID)
	return Order{
		Items:  items,
		UserID: uid,
		Total:  e.Total,
		Model: common.Model{
			CreatedAt: e.CreatedAt,
			UpdatedAt: e.UpdatedAt,
		},
	}
}

func (t *transformer) toDomain(p OrderLookedUp) domain.Order {
	items := []*domain.OrderItem{}
	for _, item := range p.Items {
		for _, product := range p.ProductsJoined {
			if product.ID.Hex() == item.ProductID.Hex() {
				items = append(items, &domain.OrderItem{
					Quantity:  item.Quantity,
					ProductID: item.ProductID.Hex(),
					Product: &domain.Product{
						ID:        product.ID.Hex(),
						CreatedAt: product.CreatedAt,
						UpdatedAt: product.UpdatedAt,
						Name:      product.Name,
						Price:     product.Price,
					},
				})
			}
		}
	}

	return domain.Order{
		ID:        p.ID.Hex(),
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		Items:     items,
		Total:     p.Total,
		UserID:    p.UserID.Hex(),
	}
}

func (t *transformer) toManyDomain(p []OrderLookedUp) []domain.Order {
	resp := []domain.Order{}
	for _, v := range p {
		resp = append(resp, t.toDomain(v))
	}

	return resp
}

func (t *transformer) buildQueryFilter(q domain.Query) primitive.M {
	filter := primitive.M{}
	if q.UserID != "" {
		filter["user_id"] = q.UserID
	}
	return filter
}
