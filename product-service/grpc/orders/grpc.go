package orders

import (
	"app/internal/core/ports"
	"app/protobufs/common"
	pb "app/protobufs/orders"
	context "context"
	"crypto/rsa"
	"errors"
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/ohmspeed777/go-pkg/jwtx"
	codes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	status "google.golang.org/grpc/status"
)

type Dependencies struct {
	OrderService ports.OrderService
	Key          *rsa.PublicKey
}

type GRPC struct {
	OrderService ports.OrderService
	transformer  *transformer
	pb.UnimplementedOrderServiceServer
	Key *rsa.PublicKey
}

func NewGRPC(d Dependencies) *GRPC {
	return &GRPC{
		OrderService: d.OrderService,
		transformer:  new(transformer),
		Key:          d.Key,
	}
}

const (
	authorizationHeader = "authorization"
)

func (g *GRPC) GetMyOrder(ctx context.Context, req *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	e := g.transformer.toQueryRequest(req)

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.PermissionDenied, errors.New("invalid context").Error())
	}

	values := md.Get(authorizationHeader)
	if len(values) <= 0 {
		return nil, status.Error(codes.PermissionDenied, errors.New("invalid context").Error())
	}

	authHeader := values[0]
	fields := strings.Fields(authHeader)
	if len(fields) <= 1 {
		return nil, status.Error(codes.PermissionDenied, errors.New("invalid context").Error())
	}

	accessToken := fields[1]
	user, err := mapClaims(accessToken, g.Key)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	e.UserID = user.ID

	entity, err := g.OrderService.FindAll(ctx, e)
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

func verifyToken(tokenString string, publicKey *rsa.PublicKey) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func mapClaims(tokenString string, publicKey *rsa.PublicKey) (*jwtx.User, error) {
	user := &jwtx.User{}
	token, err := verifyToken(tokenString, publicKey)
	if err != nil {
		return nil, errors.New("verifyToken is not pass")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("claims is not jwt.MapClaims type")
	}

	user.ID = claims["id"].(string)
	user.Role = int(claims["role"].(float64))

	return user, nil
}
