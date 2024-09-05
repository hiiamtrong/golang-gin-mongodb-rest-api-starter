package repository

import (
	"context"

	"github.com/hiiamtrong/golang-gin-mongodb-rest-api-starter/model"
	"github.com/hiiamtrong/golang-gin-mongodb-rest-api-starter/pkg/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const TodoCollection = "todo"

type TodoRepository interface {
	Create(ctx context.Context, todo *model.Todo) (*model.Todo, error)
	List(ctx context.Context) ([]model.Todo, error)
	Update(ctx context.Context, todoID primitive.ObjectID, data map[string]interface{}) (*model.Todo, error)
	Delete(ctx context.Context, todoID primitive.ObjectID) error
}

type todoRepositoryImpl struct {
	Mongodb *database.Mongodb
}

func NewTodoRepository(
	mongodb *database.Mongodb,
) TodoRepository {
	return &todoRepositoryImpl{
		Mongodb: mongodb,
	}
}

func (repo *todoRepositoryImpl) Create(ctx context.Context, todo *model.Todo) (*model.Todo, error) {
	r, err := repo.Mongodb.Database.Collection(TodoCollection).InsertOne(ctx, todo)
	if err != nil {
		return nil, err
	}
	todo.ID = r.InsertedID.(primitive.ObjectID)
	return todo, nil
}

func (repo *todoRepositoryImpl) List(ctx context.Context) ([]model.Todo, error) {
	cursor, err := repo.Mongodb.Database.Collection(TodoCollection).Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	var todos []model.Todo
	if err = cursor.All(ctx, &todos); err != nil {
		return nil, err
	}
	return todos, nil
}

func (repo *todoRepositoryImpl) Update(ctx context.Context, todoID primitive.ObjectID, data map[string]interface{}) (*model.Todo, error) {
	result := repo.Mongodb.Database.Collection(TodoCollection).FindOneAndUpdate(ctx, bson.D{{"_id", todoID}}, bson.M{"$set": data})
	if result.Err() != nil {
		return nil, result.Err()
	}

	updateTodo := &model.Todo{}
	err := result.Decode(updateTodo)
	if err != nil {
		return nil, err
	}
	return updateTodo, nil
}

func (repo *todoRepositoryImpl) Delete(ctx context.Context, todoID primitive.ObjectID) error {
	_, err := repo.Mongodb.Database.Collection(TodoCollection).DeleteOne(ctx, bson.D{{"_id", todoID}})
	if err != nil {
		return err
	}
	return nil
}
