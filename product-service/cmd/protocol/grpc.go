package protocol

import (
	"app/grpc/foods"
	"app/grpc/orders"
	_foods "app/protobufs/foods"
	_orders "app/protobufs/orders"
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

	orderGroup := orders.NewGRPC(orders.Dependencies{
		Key:          &app.key.PublicKey,
		OrderService: app.Service.OrderService,
	})

	// regis service
	_foods.RegisterFoodsServiceServer(server, foodGroup)
	_orders.RegisterOrderServiceServer(server, orderGroup)

	logx.GetLog().Infof("grpc server starting on port: %d", app.Config.APP.GRPCPort)
	server.Serve(lis)
}
