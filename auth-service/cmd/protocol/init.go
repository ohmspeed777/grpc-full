package protocol

import (
	"app/configs"
	"app/infrastructure"
	"app/internal/core/services"
	"app/repository/mongodb"
	"crypto/rsa"

	"github.com/golang-jwt/jwt"
	"github.com/ohmspeed777/go-pkg/logx"
)

var app *application

type application struct {
	Config  *configs.Config
	Service *services.Service
	key     *rsa.PrivateKey
}

func Init() {
	app = &application{}
	app.Config = &configs.Config{}
	app.Service = &services.Service{}
	app.Config = configs.NewConfig()

	mongo := infrastructure.NewMongo(app.Config.MongoDB)

	repo := mongodb.NewRepository(mongodb.Dependencies{
		Config:      app.Config,
		MongoClient: mongo,
	})

	app.Service = services.NewService(services.Dependencies{
		Conf:       app.Config,
		Repository: repo,
	})

	pkbytes := []byte(app.Config.JWT.PRIV)
	privateKeyImported, err := jwt.ParseRSAPrivateKeyFromPEM(pkbytes)
	if err != nil {
		logx.GetLog().Fatal(err)
	}

	app.key = privateKeyImported
}
