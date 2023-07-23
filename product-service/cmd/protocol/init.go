package protocol

import (
	"app/configs"
	"app/infrastructure"
	"app/internal/core/services"
	"app/repository/mongodb"
)

var app *application

type application struct {
	Config  *configs.Config
	Service *services.Service
}

func init() {
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

}
