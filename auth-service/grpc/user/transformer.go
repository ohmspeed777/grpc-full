package user

import (
	"app/protobufs/common"
	"app/internal/core/domain"
	"app/protobufs/foods"
	pb "app/protobufs/user"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type transformer struct{}

func (t *transformer) toSignInReqEntity(req *pb.SignInReq) domain.SignInReq {
	return domain.SignInReq{
		Email:    req.Email,
		Password: req.Password,
	}
}

func (t *transformer) toSignInRespProto(e *domain.SignInRes) *pb.SignInResp {
	return &pb.SignInResp{
		Token:        e.Token,
		RefreshToken: e.RefreshToken,
	}
}

func (t *transformer) toGetMyOrderRespProto(p *domain.Order) *pb.Order {
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
	return &pb.Order{
		Id:        p.ID,
		CreatedAt: timestamppb.New(p.CreatedAt),
		UpdatedAt: timestamppb.New(p.UpdatedAt),
		Total:     p.Total,
		User: &pb.User{
			UpdatedAt: timestamppb.New(p.User.UpdatedAt),
			CreatedAt: timestamppb.New(p.User.CreatedAt),
			Id:        p.User.ID,
			Email:     p.User.Email,
			FirstName: p.User.FirstName,
			LastName:  p.User.LastName,
		},
		Items: items,
	}

}

func (t *transformer) toPaginateResponse(entities []*domain.Order, info *common.PageInfo) *pb.GetMyOrderResp {
	res := pb.GetMyOrderResp{}
	data := make([]*pb.Order, len(entities))
	for i, v := range entities {
		data[i] = t.toGetMyOrderRespProto(v)
	}
	res.Entities = data
	res.PageInfo = info
	return &res
}
