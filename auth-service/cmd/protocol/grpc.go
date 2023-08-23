package protocol

import (
	"app/grpc/user"
	_user "app/protobufs/user"
	"fmt"
	"net"

	"github.com/ohmspeed777/go-pkg/logx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGRPC() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", app.Config.APP.GRPCPort))
	if err != nil {
		logx.GetLog().Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()

	// create groups
	userGroup := user.NewGRPC(user.Dependencies{
		UserService: app.Service.UserService,
		Key: &app.key.PublicKey,
	})

	// regis service
	_user.RegisterUserServiceServer(server, userGroup)

	logx.GetLog().Infof("grpc server starting on port: %d", app.Config.APP.GRPCPort)
	server.Serve(lis)
}

func newOrderGrpcClient() *grpc.ClientConn {
	conn, err := grpc.Dial(app.Config.GRPC.Order, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logx.GetLog().Fatalf("did not connect grpc server: %v", err)
	}

	return conn
}
