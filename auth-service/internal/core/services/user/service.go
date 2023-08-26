package user

import (
	"app/internal/core/domain"
	"app/internal/core/ports"
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	food_pb "app/protobufs/foods"
	pb "app/protobufs/orders"

	"github.com/golang-jwt/jwt"
	"github.com/jinzhu/copier"
	"github.com/ohmspeed777/go-pkg/errorx"
	"github.com/ohmspeed777/go-pkg/logx"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Dependencies struct {
	UserRepository ports.UserRepository
	Key            string
	OrderGRPC      pb.OrderServiceClient
	FoodGRPC       food_pb.FoodsServiceClient
}

type service struct {
	Key            string
	UserRepository ports.UserRepository
	OrderGRPC      pb.OrderServiceClient
	FoodGRPC       food_pb.FoodsServiceClient
}

func NewService(d Dependencies) ports.UserService {
	return &service{
		UserRepository: d.UserRepository,
		Key:            d.Key,
		OrderGRPC:      d.OrderGRPC,
		FoodGRPC:       d.FoodGRPC,
	}
}

func (s *service) SignUp(ctx context.Context, req domain.SignUpReq) error {
	entity := domain.User{}
	copier.Copy(&entity, &req)
	entity.CreatedAt = time.Now()
	entity.UpdatedAt = time.Now()

	old, err := s.UserRepository.FindOneByEmail(ctx, req.Email)
	if err != nil && err != mongo.ErrNoDocuments {
		logx.GetLog().Error(err)
		return errorx.New(http.StatusBadRequest, "can not create user", err)
	}

	if old.Email == req.Email {
		logx.GetLog().Error(err)
		return errorx.New(http.StatusBadRequest, "this email have already used", err)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 12)
	if err != nil {
		logx.GetLog().Error(err)
		return errorx.New(http.StatusBadRequest, "can not encrypt password", err)
	}

	entity.Password = string(hash)
	_, err = s.UserRepository.Create(ctx, entity)
	if err != nil {
		logx.GetLog().Error(err)
		return errorx.New(http.StatusBadRequest, "can not create user", err)
	}

	return nil
}

func (s *service) SignIn(ctx context.Context, req domain.SignInReq) (*domain.SignInRes, error) {
	user, err := s.UserRepository.FindOneByEmail(ctx, req.Email)
	if err != nil {
		logx.GetLog().Error(err)
		return nil, errorx.New(http.StatusBadRequest, "invalid email or password", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		logx.GetLog().Error(err)
		return nil, errorx.New(http.StatusBadRequest, "invalid email or password", err)
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(s.Key))
	if err != nil {
		logx.GetLog().Errorf("jwt getting key was error: %v", err)
		return nil, errorx.New(http.StatusBadRequest, "jwt getting key was error", err)
	}

	uuidToken := uuid.NewV4().String()
	uuidRTToken := uuid.NewV4().String()

	token := jwt.New(jwt.SigningMethodRS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = "access_token"
	claims["iss"] = "app"
	claims["jti"] = uuidToken
	claims["iat"] = time.Now().Local().Unix()
	claims["id"] = user.ID
	claims["role"] = 1
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	accessToken, err := token.SignedString(key)
	if err != nil {
		logx.GetLog().Error(err)
		return nil, errorx.New(http.StatusBadRequest, "can not generate access token", err)
	}

	// create refresh token
	rtToken := jwt.New(jwt.SigningMethodRS256)
	rtClaims := rtToken.Claims.(jwt.MapClaims)

	rtClaims["id"] = user.ID
	rtClaims["sub"] = "refresh_token"
	rtClaims["iss"] = "app"
	rtClaims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()
	rtClaims["jti"] = uuidRTToken

	refreshToken, err := rtToken.SignedString(key)
	if err != nil {
		logx.GetLog().Error(err)
		return nil, errorx.New(http.StatusBadRequest, "can not generate refresh token", err)
	}

	return &domain.SignInRes{
		Token:        accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *service) GetMyOrder(ctx context.Context, userID string) ([]*domain.Order, error) {
	user, err := s.UserRepository.FindOneByID(ctx, userID)
	if err != nil {
		logx.GetLog().Error(err)
		return nil, err
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("invalid context")
	}

	fmt.Println(md)

	ctx = metadata.NewOutgoingContext(ctx, md)

	resp, err := s.OrderGRPC.GetMyOrder(ctx, &pb.GetAllRequest{})
	if err != nil {
		logx.GetLog().Error(err)
		return nil, err
	}

	entity := []*domain.Order{}

	for _, e := range resp.Entities {
		entity = append(entity, &domain.Order{
			UpdatedAt: e.UpdatedAt.AsTime(),
			CreatedAt: e.CreatedAt.AsTime(),
			User:      &user,
			Total:     e.Total,
			ID:        e.Id,
			Items: func(items []*pb.OrderItem) []*domain.OrderItem {
				entities := []*domain.OrderItem{}
				for _, item := range items {
					entities = append(entities, &domain.OrderItem{
						Quantity:  uint(item.Quantity),
						ProductID: item.ProductId,
						Product: &domain.Product{
							ID:        item.Product.Id,
							Price:     item.Product.Price,
							Name:      item.Product.Name,
							UpdatedAt: item.Product.UpdatedAt.AsTime(),
							CreatedAt: item.Product.CreatedAt.AsTime(),
						},
					})
				}
				return entities
			}(e.Items),
		})
	}

	return entity, nil
}

func (s *service) ClientStream(ctx context.Context) error {
	stream, err := s.FoodGRPC.SendStream(ctx)
	if err != nil {
		logx.GetLog().Error(err)
		return err
	}

	foods := []*food_pb.Food{
		{
			Id:        "1",
			Name:      "น้ำตก",
			Price:     200,
			CreatedAt: timestamppb.Now(),
			UpdatedAt: timestamppb.Now(),
		},
		{
			Id:        "2",
			Name:      "กล้วย",
			Price:     5000000000,
			CreatedAt: timestamppb.Now(),
			UpdatedAt: timestamppb.Now(),
		},
	}

	for _, v := range foods {
		time.Sleep(1 * time.Second)
		err := stream.Send(v)
		if err != nil {
			logx.GetLog().Error(err)
			return err
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		logx.GetLog().Error(err)
		return err
	}

	fmt.Println(resp)

	return nil
}

func (s *service) BidirectionalStream(ctx context.Context) error {
	stream, err := s.FoodGRPC.BiDirectionalStream(ctx)
	if err != nil {
		logx.GetLog().Error(err)
		return err
	}

	foods := []*food_pb.Food{
		{
			Id:        "1",
			Name:      "น้ำตก",
			Price:     200,
			CreatedAt: timestamppb.Now(),
			UpdatedAt: timestamppb.Now(),
		},
		{
			Id:        "2",
			Name:      "กล้วย",
			Price:     5000000000,
			CreatedAt: timestamppb.Now(),
			UpdatedAt: timestamppb.Now(),
		},
	}

	go func() {
		for {
			resp, err := stream.Recv()
			if err != nil {
				logx.GetLog().Error(err)
				return
			}

			fmt.Println("resp ", resp)
		}
	}()

	for _, v := range foods {
		time.Sleep(1 * time.Second)
		err := stream.Send(v)
		if err != nil {
			logx.GetLog().Error(err)
			return err
		}
	}

	err = stream.CloseSend()
	if err != nil {
		logx.GetLog().Error(err)
		return err
	}
	return nil
}
