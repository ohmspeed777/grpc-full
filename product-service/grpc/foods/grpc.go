package foods

import (
	"app/grpc/common"
	"app/internal/core/domain"
	"app/internal/core/ports"
	context "context"
	"fmt"
	"time"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type Dependencies struct {
	ProductService ports.ProductService
}

type GRPC struct {
	ProductService ports.ProductService
	transformer    *transformer
	UnimplementedFoodsServiceServer
}

func NewGRPC(d Dependencies) *GRPC {
	return &GRPC{
		ProductService: d.ProductService,
		transformer:    new(transformer),
	}
}

func (g *GRPC) GetAll(ctx context.Context, req *GetAllRequest) (*GetAllResponse, error) {
	entity, err := g.ProductService.FindAll(ctx, g.transformer.toQueryRequest(req))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	resp := g.transformer.toPaginateResponse(entity.Entities, &common.PageInfo{
		Page:          entity.PageInfo.Page,
		Size:          entity.PageInfo.Size,
		AllOfEntities: entity.PageInfo.AllOfEntities,
		NumOfPages:    entity.PageInfo.NumOfPages,
	})

	return &resp, nil
}

func (g *GRPC) GetAllStream(req *common.Empty, srv FoodsService_GetAllStreamServer) error {
	entity, err := g.ProductService.FindAll(srv.Context(), domain.Query{})
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	for _, v := range entity.Entities {
		time.Sleep(1 * time.Second)
		fmt.Println(v.Name)
		err := srv.Send(g.transformer.toResponse(v))
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}
	}

	return nil
}
