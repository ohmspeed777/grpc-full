package user

import (
	"app/internal/core/ports"
	"app/protobufs/common"
	pb "app/protobufs/user"
	context "context"
	"crypto/rsa"
	"errors"
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/ohmspeed777/go-pkg/jwtx"
	"github.com/ohmspeed777/go-pkg/logx"
	codes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	status "google.golang.org/grpc/status"
)

const (
	authorizationHeader = "authorization"
)

type Dependencies struct {
	UserService ports.UserService
	Key         *rsa.PublicKey
}

type GRPC struct {
	UserService ports.UserService
	transformer *transformer
	pb.UnimplementedUserServiceServer
	Key *rsa.PublicKey
}

func NewGRPC(d Dependencies) *GRPC {
	return &GRPC{
		UserService: d.UserService,
		transformer: new(transformer),
		Key:         d.Key,
	}
}

func (g *GRPC) SignIn(ctx context.Context, req *pb.SignInReq) (*pb.SignInResp, error) {
	entity, err := g.UserService.SignIn(ctx, g.transformer.toSignInReqEntity(req))
	if err != nil {
		logx.GetLog().Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return g.transformer.toSignInRespProto(entity), nil
}

func (g *GRPC) GetMyOrder(ctx context.Context, req *pb.GetMyOrderReq) (*pb.GetMyOrderResp, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		err := errors.New("invalid context")
		logx.GetLog().Error(err)
		return nil, status.Error(codes.PermissionDenied, err.Error())
	}

	values := md.Get(authorizationHeader)
	if len(values) <= 0 {
		err := errors.New("invalid context")
		logx.GetLog().Error(err)
		return nil, status.Error(codes.PermissionDenied, err.Error())
	}

	authHeader := values[0]
	fields := strings.Fields(authHeader)
	if len(fields) <= 1 {
		return nil, status.Error(codes.PermissionDenied, errors.New("invalid context").Error())
	}

	accessToken := fields[1]
	user, err := mapClaims(accessToken, g.Key)
	if err != nil {
		logx.GetLog().Error(err)
		return nil, status.Error(codes.PermissionDenied, err.Error())
	}

	entity, err := g.UserService.GetMyOrder(ctx, user.ID)
	if err != nil {
		logx.GetLog().Error(err)
		return nil, status.Error(codes.Unavailable, err.Error())
	}

	return g.transformer.toPaginateResponse(entity, &common.PageInfo{}), nil
}

func verifyToken(tokenString string, publicKey *rsa.PublicKey) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})
	if err != nil {
		logx.GetLog().Error(err)
		return nil, err
	}
	return token, nil
}

func mapClaims(tokenString string, publicKey *rsa.PublicKey) (*jwtx.User, error) {
	user := &jwtx.User{}
	token, err := verifyToken(tokenString, publicKey)
	if err != nil {
		logx.GetLog().Error(err)
		return nil, errors.New("verifyToken is not pass")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		logx.GetLog().Error(err)
		return nil, errors.New("claims is not jwt.MapClaims type")
	}

	user.ID = claims["id"].(string)
	user.Role = int(claims["role"].(float64))

	return user, nil
}
