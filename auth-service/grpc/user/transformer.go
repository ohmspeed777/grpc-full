package user

import (
	"app/internal/core/domain"
	pb "app/protobufs/user"
)

type transformer struct{}

func (t *transformer) toSignInReqEntity(req *pb.SignInReq) domain.SignInReq {
	return domain.SignInReq{}
}

func (t *transformer) toSignInRespProto(e *domain.SignInRes) *pb.SignInResp {
	return &pb.SignInResp{}
}
