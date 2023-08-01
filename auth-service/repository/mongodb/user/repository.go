package user

import (
	"app/internal/core/domain"
	"app/internal/core/ports"
	"app/repository/mongodb/common"
	"context"
	"sync"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	collectionName = "users"
)

type Repo struct {
	transformer *transformer
	collection  *mongo.Collection
	mux         sync.Mutex
}

type Dependencies struct {
	DB *mongo.Database
}

func NewRepo(d Dependencies) ports.UserRepository {
	return &Repo{
		transformer: new(transformer),
		collection:  d.DB.Collection(collectionName),
	}
}

func (r *Repo) FindAll(ctx context.Context, q domain.Query) (empty []domain.User, counter int64, err error) {
	var (
		model = []User{}
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

func (r *Repo) Create(ctx context.Context, e domain.User) (empty domain.User, err error) {
	model := r.transformer.toModel(e)
	_, err = r.collection.InsertOne(ctx, &model)

	if err != nil {
		return empty, errors.WithStack(err)
	}

	return r.transformer.toDomain(model), nil
}

func (r *Repo) FindOneByID(ctx context.Context, id string) (empty domain.User, err error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return empty, errors.WithStack(err)
	}

	model := User{}
	err = r.collection.FindOne(ctx, primitive.M{"_id": _id}).Decode(&model)
	if err != nil {
		return empty, errors.WithStack(err)
	}

	return r.transformer.toDomain(model), nil
}

func (r *Repo) FindOneByEmail(ctx context.Context, email string) (empty domain.User, err error) {
	model := User{}
	err = r.collection.FindOne(ctx, primitive.M{"email": email}).Decode(&model)
	if err != nil {
		return empty, err
	}

	return r.transformer.toDomain(model), nil
}
