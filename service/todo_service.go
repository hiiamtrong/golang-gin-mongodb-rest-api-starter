package service

import (
	"context"
	objectidpkg "github.com/hiiamtrong/golang-gin-mongodb-rest-api-starter/pkg/object_id"
	"github.com/hiiamtrong/golang-gin-mongodb-rest-api-starter/repository"

	"github.com/hiiamtrong/golang-gin-mongodb-rest-api-starter/model"
)

type TodoService interface {
	Create(ctx context.Context, todo model.Todo) (*model.Todo, error)
	List(ctx context.Context) ([]model.Todo, error)
	Update(ctx context.Context, todoID string, data map[string]interface{}) (*model.Todo, error)
	Delete(ctx context.Context, todoID string) error
}

type todoServiceImpl struct {
	todoRepo repository.TodoRepository
}

func NewTodoService(todoRepo repository.TodoRepository) TodoService {
	return &todoServiceImpl{
		todoRepo: todoRepo,
	}
}

func (svc *todoServiceImpl) Create(ctx context.Context, todo model.Todo) (*model.Todo, error) {
	newTodo, err := svc.todoRepo.Create(ctx, &todo)
	if err != nil {
		return nil, err
	}
	return newTodo, nil
}

func (svc *todoServiceImpl) List(ctx context.Context) ([]model.Todo, error) {
	todos, err := svc.todoRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (svc *todoServiceImpl) Update(ctx context.Context, todoID string, data map[string]interface{}) (*model.Todo, error) {
	todoObjectID := objectidpkg.ObjectIDFromHex(todoID)
	newTodo, err := svc.todoRepo.Update(ctx, todoObjectID, data)
	if err != nil {
		return nil, err
	}
	return newTodo, nil
}

func (svc *todoServiceImpl) Delete(ctx context.Context, todoID string) error {
	todoObjectID := objectidpkg.ObjectIDFromHex(todoID)
	err := svc.todoRepo.Delete(ctx, todoObjectID)
	if err != nil {
		return err
	}
	return nil
}
