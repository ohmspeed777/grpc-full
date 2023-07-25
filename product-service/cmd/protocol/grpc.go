package protocol

import (
	"app/grpc/foods"
	"fmt"
	"net"

	"github.com/ohmspeed777/go-pkg/logx"
	"google.golang.org/grpc"
)

func NewGRPC() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", app.Config.APP.GRPCPort))
	if err != nil {
		logx.GetLog().Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()

	// create groups
	foodGroup := foods.NewGRPC(foods.Dependencies{
		ProductService: app.Service.ProductService,
	})

	// regis service
	foods.RegisterFoodsServiceServer(server, foodGroup)

	logx.GetLog().Infof("grpc server starting on port: %d", app.Config.APP.GRPCPort)
	server.Serve(lis)
}
