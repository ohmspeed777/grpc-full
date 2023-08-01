package user

import (
	"app/internal/core/domain"
	"app/internal/core/ports"
	"context"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/jinzhu/copier"
	"github.com/ohmspeed777/go-pkg/errorx"
	"github.com/ohmspeed777/go-pkg/logx"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type Dependencies struct {
	UserRepository ports.UserRepository
	Key            string
}

type service struct {
	Key            string
	UserRepository ports.UserRepository
}

func NewService(d Dependencies) ports.UserService {
	return &service{
		UserRepository: d.UserRepository,
		Key:            d.Key,
	}
}

func (s *service) SignUp(ctx context.Context, req domain.SignUpReq) error {
	entity := domain.User{}
	copier.Copy(&entity, &req)
	entity.CreatedAt = time.Now()
	entity.UpdatedAt = time.Now()

	old, err := s.UserRepository.FindOneByEmail(ctx, req.Email)
	if err != nil && err != mongo.ErrNoDocuments {
		return errorx.New(http.StatusBadRequest, "can not create user", err)
	}

	if old.Email == req.Email {
		return errorx.New(http.StatusBadRequest, "this email have already used", err)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 12)
	if err != nil {
		return errorx.New(http.StatusBadRequest, "can not encrypt password", err)
	}

	entity.Password = string(hash)
	_, err = s.UserRepository.Create(ctx, entity)
	if err != nil {
		return errorx.New(http.StatusBadRequest, "can not create user", err)
	}

	return nil
}

func (s *service) SignIn(ctx context.Context, req domain.SignInReq) (*domain.SignInRes, error) {
	user, err := s.UserRepository.FindOneByEmail(ctx, req.Email)
	if err != nil {
		return nil, errorx.New(http.StatusBadRequest, "invalid email or password", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
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
		return nil, errorx.New(http.StatusBadRequest, "can not generate refresh token", err)
	}

	return &domain.SignInRes{
		Token:        accessToken,
		RefreshToken: refreshToken,
	}, nil
}
