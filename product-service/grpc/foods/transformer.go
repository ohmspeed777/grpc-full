package foods

import (
	"app/protobufs/common"
	"app/internal/core/domain"
	pb "app/protobufs/foods"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type transformer struct{}

func (t *transformer) toQueryRequest(req *pb.GetAllRequest) domain.Query {
	return domain.Query{
		Limit: uint(req.Query.Limit),
		Page:  uint(req.Query.Page),
	}
}

func (t *transformer) toResponse(p domain.Product) *pb.Food {
	return &pb.Food{
		Id:        p.ID,
		CreatedAt: timestamppb.New(p.CreatedAt),
		UpdatedAt: timestamppb.New(p.UpdatedAt),
		Name:      p.Name,
		Price:     p.Price,
	}
}

func (t *transformer) toPaginateResponse(entities []domain.Product, info *common.PageInfo) pb.GetAllResponse {
	res := pb.GetAllResponse{}
	data := make([]*pb.Food, len(entities))
	for i, v := range entities {
		data[i] = t.toResponse(v)
	}
	res.Entities = data
	res.PageInfo = info
	return res
}
