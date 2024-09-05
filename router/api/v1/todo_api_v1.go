package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	errpkg "github.com/hiiamtrong/golang-gin-mongodb-rest-api-starter/pkg/error"

	"github.com/hiiamtrong/golang-gin-mongodb-rest-api-starter/router/api"
	"github.com/hiiamtrong/golang-gin-mongodb-rest-api-starter/router/dto"
	"github.com/hiiamtrong/golang-gin-mongodb-rest-api-starter/service"
)

type TodoAPIV1 struct {
	todoSvc service.TodoService
}

func NewTodoAPIV1(todoSvc service.TodoService) *TodoAPIV1 {
	return &TodoAPIV1{todoSvc: todoSvc}
}

// List Todo
// @Tags Todo
// @Summary List all todos
// @Description List all todos
// @Security Bearer
// @Accept json
// @Produce json
// @Success 200 {array} model.Todo
// @Router /api/v1/todo [get]
func (a *TodoAPIV1) List() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("List")
		todos, err := a.todoSvc.List(c.Request.Context())
		if err != nil {
			api.ResponseError(c, http.StatusInternalServerError, errpkg.ERROR, err)
			return
		}

		api.ResponseSuccess(c, http.StatusOK, todos)
	}
}

// Read Todo
// @Tags Todo
// @Summary Read a todo
// @Description Read a todo
// @Security Bearer
// @Accept json
// @Produce json
// @Param id path string true "Todo ID"
// @Success 200 {object} model.Todo
// @Router /api/v1/todo/{id} [get]
func (a *TodoAPIV1) Read() gin.HandlerFunc {
	return func(c *gin.Context) {
		todoID := c.Param("id")
		todo, err := a.todoSvc.Read(c.Request.Context(), todoID)
		if err != nil {
			api.ResponseError(c, http.StatusInternalServerError, errpkg.ERROR, err)
			return
		}

		api.ResponseSuccess(c, http.StatusOK, todo)
	}
}

// Create Todo
// @Tags Todo
// @Summary Create a new todo
// @Description Create a new todo
// @Security Bearer
// @Accept json
// @Produce json
// @Param todo body dto.CreateTodoRequest true "Todo object that needs to be created"
// @Success 200 {object} model.Todo
// @Router /api/v1/todo [post]
func (a *TodoAPIV1) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &dto.CreateTodoRequest{}
		if err := c.ShouldBindJSON(req); err != nil {
			api.ResponseError(c, http.StatusBadGateway, errpkg.INVALID_PARAMS, err)
			return
		}

		err := req.Validate()
		if err != nil {
			api.ResponseError(c, http.StatusBadGateway, errpkg.INVALID_PARAMS, err)
			return
		}

		newTodo, err := a.todoSvc.Create(c.Request.Context(), req.ToModel())

		if err != nil {
			api.ResponseError(c, http.StatusInternalServerError, errpkg.ERROR, err)
			return
		}

		api.ResponseSuccess(c, http.StatusOK, newTodo)
	}
}

// Update Todo
// @Tags Todo
// @Summary Update a todo
// @Description Update a todo
// @Security Bearer
// @Accept json
// @Produce json
// @Param id path string true "Todo ID"
// @Param todo body dto.UpdateTodoRequest true "Todo object that needs to be updated"
// @Success 200 {object} model.Todo
// @Router /api/v1/todo/{id} [put]
func (a *TodoAPIV1) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		todoID := c.Param("id")

		req := &dto.UpdateTodoRequest{}
		if err := c.ShouldBindJSON(req); err != nil {
			api.ResponseError(c, http.StatusBadGateway, errpkg.INVALID_PARAMS, err)
			return
		}

		err := req.Validate()
		if err != nil {
			api.ResponseError(c, http.StatusBadGateway, errpkg.INVALID_PARAMS, err)
			return
		}

		updatedTodo, err := a.todoSvc.Update(c.Request.Context(), todoID, req.ToMap())

		if err != nil {
			api.ResponseError(c, http.StatusInternalServerError, errpkg.ERROR, err)
			return
		}

		api.ResponseSuccess(c, http.StatusOK, updatedTodo)
	}
}

// Delete Todo
// @Tags Todo
// @Summary Delete a todo
// @Description Delete a todo
// @Security Bearer
// @Accept json
// @Produce json
// @Param id path string true "Todo ID"
// @Success 200
// @Router /api/v1/todo/{id} [delete]
func (a *TodoAPIV1) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		todoID := c.Param("id")

		err := a.todoSvc.Delete(c.Request.Context(), todoID)

		if err != nil {
			api.ResponseError(c, http.StatusInternalServerError, errpkg.ERROR, err)
			return
		}

		api.ResponseSuccess(c, http.StatusOK, nil)
	}
}
