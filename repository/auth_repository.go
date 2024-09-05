package repository

import (
	"context"

	"github.com/hiiamtrong/golang-gin-mongodb-rest-api-starter/model"
	"github.com/hiiamtrong/golang-gin-mongodb-rest-api-starter/pkg/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	UserCollection = "user"
)

type AuthRepository interface {
	FindByUsernameAndPassword(ctx context.Context, username, password string) (*model.User, error)
	FindByUsername(ctx context.Context, username string) (*model.User, error)
	Create(ctx context.Context, user *model.User) (*model.User, error)
}

type authRepositoryImpl struct {
	Mongodb *database.Mongodb
}

func NewAuthRepository(
	mongodb *database.Mongodb,
) AuthRepository {
	return &authRepositoryImpl{
		Mongodb: mongodb,
	}
}

func (repo *authRepositoryImpl) FindByUsernameAndPassword(ctx context.Context, username, password string) (*model.User, error) {
	user := &model.User{}
	err := repo.Mongodb.Database.Collection(UserCollection).FindOne(ctx, bson.D{{"username", username}, {"password", password}}).Decode(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *authRepositoryImpl) FindByUsername(ctx context.Context, username string) (*model.User, error) {
	user := &model.User{}
	err := repo.Mongodb.Database.Collection(UserCollection).FindOne(ctx, bson.D{{"username", username}}).Decode(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *authRepositoryImpl) Create(ctx context.Context, user *model.User) (*model.User, error) {
	r, err := repo.Mongodb.Database.Collection(UserCollection).InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = r.InsertedID.(primitive.ObjectID)
	return user, nil
}
