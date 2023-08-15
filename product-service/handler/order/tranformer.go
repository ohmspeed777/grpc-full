package order

import (
	"app/handler/common"
	"app/handler/product"
	"app/internal/core/domain"
)

type transformer struct{}

func (t *transformer) toResponse(p domain.Order) Order {
	items := []*OrderItem{}
	for _, v := range p.Items {
		items = append(items, &OrderItem{
			Quantity:  v.Quantity,
			ProductID: v.ProductID,
			Product: &product.Product{
				ID:        v.Product.ID,
				CreatedAt: v.Product.CreatedAt,
				UpdatedAt: v.Product.UpdatedAt,
				Name:      v.Product.Name,
				Price:     v.Product.Price,
			},
		})
	}
	resp := Order{
		ID:        p.ID,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}

	return resp
}

func (t *transformer) toPaginateResponse(entities []domain.Order, info common.PageInfo) common.ResponseInfo {
	res := common.ResponseInfo{}
	data := make([]Order, len(entities))
	for i, v := range entities {
		data[i] = t.toResponse(v)
	}
	res.Entities = data
	res.PageInfo = info
	return res
}

func (t *transformer) toQueryRequest(req common.Query) domain.Query {
	return domain.Query{
		Page:  req.Page,
		Limit: req.Limit,
	}
}

func (t *transformer) toRequest(req Order) domain.Order {
	items := []*domain.OrderItem{}
	for _, v := range req.Items {
		items = append(items, &domain.OrderItem{
			Quantity:  v.Quantity,
			ProductID: v.ProductID,
		})
	}
	return domain.Order{
		Items:  items,
		Total:  req.Total,
		UserID: req.UserID,
	}
}
