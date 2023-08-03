package protocol

import (
	"app/grpc/user"
	"fmt"
	"net"
	_user "app/protobufs/user"
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
	userGroup := user.NewGRPC(user.Dependencies{
		UserService: app.Service.UserService,
	})

	// regis service
	_user.RegisterUserServiceServer(server, userGroup)

	logx.GetLog().Infof("grpc server starting on port: %d", app.Config.APP.GRPCPort)
	server.Serve(lis)
}
