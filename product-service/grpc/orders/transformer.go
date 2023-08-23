package orders

import (
	"app/internal/core/domain"
	"app/protobufs/common"
	"app/protobufs/foods"
	pb "app/protobufs/orders"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type transformer struct{}

func (t *transformer) toQueryRequest(req *pb.GetAllRequest) domain.Query {
	e := domain.Query{
		Limit: 20,
		Page:  1,
	}
	if req.Query != nil {
		e.Limit = uint(req.Query.Limit)
		e.Page = uint(req.Query.Page)
	}

	return e
}

func (t *transformer) toResponse(p domain.Order) *pb.Order {
	items := []*pb.OrderItem{}
	for _, v := range p.Items {
		food := v.Product
		items = append(items, &pb.OrderItem{
			Quantity:  int64(v.Quantity),
			ProductId: v.ProductID,
			Product: &foods.Food{
				Id:        food.ID,
				CreatedAt: timestamppb.New(food.CreatedAt),
				UpdatedAt: timestamppb.New(food.UpdatedAt),
				Name:      food.Name,
				Price:     food.Price,
			},
		})
	}
	resp := &pb.Order{
		Id:        p.ID,
		CreatedAt: timestamppb.New(p.CreatedAt),
		UpdatedAt: timestamppb.New(p.UpdatedAt),
		Total:     p.Total,
		UserId:    p.UserID,
		Items:     items,
	}

	return resp
}

func (t *transformer) toPaginateResponse(entities []domain.Order, info *common.PageInfo) pb.GetAllResponse {
	res := pb.GetAllResponse{}
	data := make([]*pb.Order, len(entities))
	for i, v := range entities {
		data[i] = t.toResponse(v)
	}
	res.Entities = data
	res.PageInfo = info
	return res
}
