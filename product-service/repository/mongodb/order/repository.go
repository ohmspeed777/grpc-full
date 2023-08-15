package order

import (
	"app/internal/core/domain"
	"app/internal/core/ports"
	"app/repository/mongodb/common"
	"context"
	"sync"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	collectionName = "orders"
)

type Repo struct {
	transformer *transformer
	collection  *mongo.Collection
	mux         sync.Mutex
}

type Dependencies struct {
	DB *mongo.Database
}

func NewRepo(d Dependencies) ports.OrderRepository {
	return &Repo{
		transformer: new(transformer),
		collection:  d.DB.Collection(collectionName),
	}
}

func (r *Repo) FindAll(ctx context.Context, q domain.Query) (empty []domain.Order, counter int64, err error) {
	var (
		model = []OrderLookedUp{}
	)

	filter := r.transformer.buildQueryFilter(q)
	cursor, err := r.collection.Find(ctx, filter, common.GetPaginationOption(q))

	if err != nil {
		return empty, counter, errors.WithStack(err)
	}

	err = cursor.All(ctx, &model)
	if err != nil {
		return empty, counter, errors.WithStack(err)
	}

	counter, err = r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return empty, counter, errors.WithStack(err)
	}

	return r.transformer.toManyDomain(model), counter, nil
}

func (r *Repo) Create(ctx context.Context, e domain.Order) (empty domain.Order, err error) {
	model := r.transformer.toModel(e)
	_, err = r.collection.InsertOne(ctx, &model)

	if err != nil {
		return empty, errors.WithStack(err)
	}

	return e, nil
}
