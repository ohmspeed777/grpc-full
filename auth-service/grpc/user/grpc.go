package user

import (
	"app/internal/core/ports"
	pb "app/protobufs/user"
	context "context"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type Dependencies struct {
	UserService ports.UserService
}

type GRPC struct {
	UserService ports.UserService
	transformer *transformer
	pb.UnimplementedUserServiceServer
}

func NewGRPC(d Dependencies) *GRPC {
	return &GRPC{
		UserService: d.UserService,
		transformer: new(transformer),
	}
}

func (g *GRPC) SignIn(ctx context.Context, req *pb.SignInReq) (*pb.SignInResp, error) {
	entity, err := g.UserService.SignIn(ctx, g.transformer.toSignInReqEntity(req))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return g.transformer.toSignInRespProto(entity), nil
}
