package infrastructure

import (
	"app/configs"
	"context"
	"time"

	"github.com/ohmspeed777/go-pkg/logx"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongo(conf configs.MongoDB) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	log := logx.GetLog()

	opts := options.Client()

	opts.ApplyURI(conf.ToURI())

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Errorf("cannot connect to mongodb err : %s", err)
	}

	return client
}
