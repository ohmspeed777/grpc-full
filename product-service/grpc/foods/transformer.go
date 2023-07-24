package foods

import (
	"app/grpc/common"
	"app/internal/core/domain"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type transformer struct{}

func (t *transformer) toQueryRequest(req *GetAllRequest) domain.Query {
	return domain.Query{
		Limit: uint(req.Query.Limit),
		Page:  uint(req.Query.Page),
	}
}

func (t *transformer) toResponse(p domain.Product) *Food {
	return &Food{
		Id:        p.ID,
		CreatedAt: timestamppb.New(p.CreatedAt),
		UpdatedAt: timestamppb.New(p.UpdatedAt),
		Name:      p.Name,
		Price:     p.Price,
	}
}

func (t *transformer) toPaginateResponse(entities []domain.Product, info *common.PageInfo) GetAllResponse {
	res := GetAllResponse{}
	data := make([]*Food, len(entities))
	for i, v := range entities {
		data[i] = t.toResponse(v)
	}
	res.Entities = data
	res.PageInfo = info
	return res
}
