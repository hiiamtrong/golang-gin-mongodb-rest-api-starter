package dto

import (
	"encoding/json"

	"github.com/hiiamtrong/golang-gin-mongodb-rest-api-starter/model"
)

type CreateTodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (req *CreateTodoRequest) Validate() error {
	return nil
}

func (req *CreateTodoRequest) ToModel() model.Todo {
	todo := model.NewTodo()
	todo.Title = req.Title
	todo.Description = req.Description
	return todo
}

type UpdateTodoRequest struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
}

func (req *UpdateTodoRequest) Validate() error {
	return nil
}

func (req *UpdateTodoRequest) ToMap() map[string]interface{} {
	data, err := json.Marshal(req)

	if err != nil {
		return nil
	}

	var newMap map[string]interface{}

	err = json.Unmarshal(data, &newMap)
	if err != nil {
		return nil
	}
	return newMap
}
