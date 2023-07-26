package common

import (
	"app/internal/core/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Model struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty"`
	CreatedAt time.Time          `bson:"created_at,omitempty"`
}

func GetPaginationOption(q domain.Query) *options.FindOptions {
	return options.Find().SetSkip(q.GetSkip()).SetLimit(q.GetLimit())
}
