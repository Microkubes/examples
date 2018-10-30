package main

import (
	"fmt"
	"time"

	errors "github.com/JormungandrK/backends"
	"github.com/Microkubes/examples/todo/todo-service/app"
	"github.com/Microkubes/examples/todo/todo-service/config"
	"github.com/Microkubes/examples/todo/todo-service/db"
	"github.com/goadesign/goa"
)

type TodosController struct {
	*goa.Controller
	TodoStore db.TodoStore
	Config    *config.ServiceConfig
}

// NewTodosController creates a todos controller.
func NewTodosController(service *goa.Service, todoStore db.TodoStore, cfg *config.ServiceConfig) *TodosController {

	return &TodosController{
		Controller: service.NewController("TodosController"),
		TodoStore:  todoStore,
		Config:     cfg,
	}
}

// GetByID runs the getById action.
func (c *TodosController) GetByID(ctx *app.GetByIDTodoContext) error {
	res, err := c.TodoStore.DBGetByID(ctx.TodoID)

	if err != nil {
		if errors.IsErrInvalidInput(err) {
			return ctx.BadRequest(goa.ErrBadRequest(err.Error()))
		}

		if errors.IsErrNotFound(err) {
			return ctx.NotFound(goa.ErrNotFound(err.Error()))
		}

		return ctx.InternalServerError(goa.ErrBadRequest(err.Error()))
	}

	appTodo := &app.TodoMedia{
		ID:          res.ID,
		Title:       &res.Title,
		Description: &res.Description,
		Done:        res.Done,
		CreatedAt:   res.CreatedAt,
		CompletedAt: &res.CompletedAt,
	}

	return ctx.OK(appTodo)
}

// AddTodo runs the addTodo action.
func (c *TodosController) AddTodo(ctx *app.AddTodoTodoContext) error {
	// if !auth.HasAuth(ctx.Context) {
	// 	return ctx.InternalServerError(goa.ErrBadRequest("no-auth"))
	// }

	dbTodo := &db.Todo{
		Title:       ctx.Payload.Title,
		Description: ctx.Payload.Description,
		CreatedAt:   int(time.Now().Unix()),
		Done:        false,
	}

	//Add the todo
	todo, err := c.TodoStore.DBAddTodo(dbTodo)
	if err != nil {
		if errors.IsErrInvalidInput(err) || errors.IsErrAlreadyExists(err) {
			return ctx.BadRequest(goa.ErrBadRequest(err.Error()))
		}

		return ctx.InternalServerError(goa.ErrBadRequest(err.Error()))
	}

	appTodo := &app.TodoMedia{
		ID:          todo.ID,
		Title:       &todo.Title,
		Description: &todo.Description,
		CreatedAt:   todo.CreatedAt,
		CompletedAt: &todo.CompletedAt,
		Done:        todo.Done,
	}

	return ctx.Created(appTodo)
}

// DeleteTodo runs the deleteTodo action.
func (c *TodosController) DeleteTodo(ctx *app.DeleteTodoTodoContext) error {

	_, err := c.TodoStore.DBGetByID(ctx.TodoID)
	if err != nil {
		if errors.IsErrInvalidInput(err) {
			return ctx.BadRequest(goa.ErrBadRequest(err.Error()))
		}

		if errors.IsErrNotFound(err) {
			return ctx.NotFound(goa.ErrNotFound(err.Error()))
		}

		return ctx.InternalServerError(goa.ErrBadRequest(err.Error()))
	}

	err = c.TodoStore.DBDeleteTodo(ctx.TodoID)
	if err != nil {
		if errors.IsErrInvalidInput(err) {
			return ctx.BadRequest(goa.ErrBadRequest(err.Error()))
		}

		if errors.IsErrNotFound(err) {
			return ctx.NotFound(goa.ErrNotFound(err.Error()))
		}

		return ctx.InternalServerError(goa.ErrInternal(err.Error()))
	}

	return nil
}

// GetAllTodos runs the getAllTodos action.
func (c *TodosController) GetAllTodos(ctx *app.GetAllTodosTodoContext) error {
	order := ""
	sorting := ""
	limit := 0
	offset := 0

	if ctx.Order != nil {
		order = *ctx.Order
	}

	if ctx.Offset != nil {
		offset = *ctx.Offset
	}

	if ctx.Limit != nil {
		limit = *ctx.Limit
	}

	if ctx.Sorting != nil {
		sorting = *ctx.Sorting
	}

	todos, err := c.TodoStore.DBGetAllTodos(order, sorting, limit, offset)
	if err != nil {
		if errors.IsErrInvalidInput(err) {
			return ctx.BadRequest(goa.ErrBadRequest(err.Error()))
		}

		if errors.IsErrNotFound(err) {
			return ctx.NotFound(goa.ErrNotFound(err.Error()))
		}

		return ctx.InternalServerError(goa.ErrInternal(err.Error()))
	}

	return ctx.OK(todos)
}

// UpdateTodo runs the updateTodo action.
func (c *TodosController) UpdateTodo(ctx *app.UpdateTodoTodoContext) error {
	// if !auth.HasAuth(ctx.Context) {
	// 	return ctx.InternalServerError(goa.ErrBadRequest("no-auth"))
	// }

	if ctx.Payload.Title == nil && ctx.Payload.Description == nil && ctx.Payload.Done == nil {
		return ctx.BadRequest(goa.ErrBadRequest("title, description or done flag must be set"))
	}

	dbTodo := &db.Todo{
		ID: ctx.TodoID,
		// ModifiedBy: auth.GetAuth(ctx.Context).UserID,
		CompletedAt: int(time.Now().Unix()),
		// Done:        *ctx.Payload.Done,
		// Description: *ctx.Payload.Description,
	}

	existingTodo, err := c.TodoStore.DBGetByID(ctx.TodoID)
	if err != nil {
		if errors.IsErrInvalidInput(err) {
			return ctx.BadRequest(goa.ErrBadRequest(err.Error()))
		}

		if errors.IsErrNotFound(err) {
			return ctx.NotFound(goa.ErrNotFound(err.Error()))
		}

		return ctx.InternalServerError(goa.ErrInternal(err.Error()))
	}

	if ctx.Payload.Description != nil {
		dbTodo.Description = *ctx.Payload.Description
	} else {
		dbTodo.Description = existingTodo.Description
	}

	if ctx.Payload.Title != nil {
		dbTodo.Title = *ctx.Payload.Title
	} else {
		dbTodo.Title = existingTodo.Title
	}

	if ctx.Payload.Done != nil {
		dbTodo.Done = *ctx.Payload.Done
	} else {
		dbTodo.Done = existingTodo.Done
	}

	// dbTodo.CreatedBy = existingTodo.CreatedBy
	dbTodo.CreatedAt = existingTodo.CreatedAt

	//Update the todo
	res, err := c.TodoStore.DBUpdateTodo(dbTodo)
	if err != nil {
		if errors.IsErrInvalidInput(err) {
			return ctx.BadRequest(goa.ErrBadRequest(err.Error()))
		}

		if errors.IsErrNotFound(err) {
			return ctx.NotFound(goa.ErrNotFound(fmt.Sprintf("%s %s", "todo", err.Error())))
		}

		return ctx.InternalServerError(goa.ErrInternal(err.Error()))
	}

	appTodo := &app.TodoMedia{
		ID:          res.ID,
		Title:       &res.Title,
		Description: &res.Description,
		CreatedAt:   res.CreatedAt,
		CompletedAt: &res.CompletedAt,
		Done:        res.Done,
	}

	return ctx.OK(appTodo)
}
