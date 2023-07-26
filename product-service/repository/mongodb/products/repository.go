package products

import (
	"app/internal/core/domain"
	"app/internal/core/ports"
	"app/repository/mongodb/common"
	"context"
	"fmt"
	"sync"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	collectionName = "products"
)

type Repo struct {
	transformer *transformer
	collection  *mongo.Collection
	mux         sync.Mutex
}

type Dependencies struct {
	DB *mongo.Database
}

func NewRepo(d Dependencies) ports.ProductRepository {
	return &Repo{
		transformer: new(transformer),
		collection:  d.DB.Collection(collectionName),
	}
}

func (r *Repo) FindAll(ctx context.Context, q domain.Query) (empty []domain.Product, counter int64, err error) {
	var (
		model = []Product{}
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

func (r *Repo) Create(ctx context.Context, e domain.Product) (empty domain.Product, err error) {
	model := r.transformer.toModel(e)
	_, err = r.collection.InsertOne(ctx, &model)

	if err != nil {
		fmt.Println()
		fmt.Println()
		fmt.Println(err)
		fmt.Println()
		return empty, errors.WithStack(err)
	}

	return r.transformer.toDomain(model), nil
}
