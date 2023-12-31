package foods

import (
	"app/internal/core/domain"
	"app/internal/core/ports"
	"app/protobufs/common"
	pb "app/protobufs/foods"
	context "context"
	"fmt"
	"io"
	"time"

	"github.com/ohmspeed777/go-pkg/logx"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type Dependencies struct {
	ProductService ports.ProductService
}

type GRPC struct {
	ProductService ports.ProductService
	transformer    *transformer
	pb.UnimplementedFoodsServiceServer
}

func NewGRPC(d Dependencies) *GRPC {
	return &GRPC{
		ProductService: d.ProductService,
		transformer:    new(transformer),
	}
}

func (g *GRPC) GetAll(ctx context.Context, req *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	entity, err := g.ProductService.FindAll(ctx, g.transformer.toQueryRequest(req))
	if err != nil {
		logx.GetLog().Error(err)
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

func (g *GRPC) GetAllStream(req *common.Empty, srv pb.FoodsService_GetAllStreamServer) error {
	entity, err := g.ProductService.FindAll(srv.Context(), domain.Query{})
	if err != nil {
		logx.GetLog().Error(err)
		return status.Error(codes.Internal, err.Error())
	}

	for _, v := range entity.Entities {
		time.Sleep(1 * time.Second)
		err := srv.Send(g.transformer.toResponse(v))
		if err != nil {
			logx.GetLog().Error(err)
			return status.Error(codes.Internal, err.Error())
		}
	}

	return nil
}

func (g *GRPC) BiDirectionalStream(srv pb.FoodsService_BiDirectionalStreamServer) error {
	for {
		food, err := srv.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			logx.GetLog().Println(err.Error())
			return err
		}

		fmt.Println("food: ", food)

		srv.Send(food)
	}
}

func (g *GRPC) SendStream(srv pb.FoodsService_SendStreamServer) error {
	foods := []*pb.Food{}
	resp := &pb.GetAllResponse{}

	for {
		food, err := srv.Recv()

		if err == io.EOF {
			resp.Entities = foods
			err := srv.SendAndClose(resp)
			if err != nil {
				logx.GetLog().Println(err.Error())
				return err
			}
			return nil
		}

		if err != nil {
			logx.GetLog().Println(err.Error())
			return err
		}

		fmt.Println("food: ", food)
		foods = append(foods, food)
	}
}
