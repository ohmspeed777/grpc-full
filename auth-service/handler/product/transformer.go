package product

import (
	"app/handler/common"
	"app/internal/core/domain"
)

type transformer struct{}

func (t *transformer) toResponse(p domain.Product) Product {
	return Product{
		ID:        p.ID,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		Name:      p.Name,
		Price:     p.Price,
	}
}

func (t *transformer) toPaginateResponse(entities []domain.Product, info common.PageInfo) common.ResponseInfo {
	res := common.ResponseInfo{}
	data := make([]Product, len(entities))
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
