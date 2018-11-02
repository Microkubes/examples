package main

import (
	"fmt"
	"time"

	"github.com/JormungandrK/backends"
	errors "github.com/JormungandrK/backends"
	"github.com/Microkubes/examples/todo/todo-service/app"
	"github.com/Microkubes/examples/todo/todo-service/config"
	"github.com/Microkubes/examples/todo/todo-service/db"
	"github.com/Microkubes/microservice-security/auth"
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
		Done:        &res.Done,
		CreatedAt:   res.CreatedAt,
		CompletedAt: &res.CompletedAt,
		CreatedBy:   &res.CreatedBy,
	}

	return ctx.OK(appTodo)
}

// AddTodo runs the addTodo action.
func (c *TodosController) AddTodo(ctx *app.AddTodoTodoContext) error {
	if !auth.HasAuth(ctx.Context) {
		return ctx.InternalServerError(goa.ErrBadRequest("no-auth"))
	}

	dbTodo := &db.Todo{
		Title:       ctx.Payload.Title,
		Description: ctx.Payload.Description,
		CreatedAt:   int(time.Now().Unix()),
		CreatedBy:   auth.GetAuth(ctx.Context).UserID,
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
		Done:        &todo.Done,
		CreatedBy:   &todo.CreatedBy,
	}

	return ctx.Created(appTodo)
}

// DeleteTodo runs the deleteTodo action.
func (c *TodosController) DeleteTodo(ctx *app.DeleteTodoTodoContext) error {
	if !auth.HasAuth(ctx.Context) {
		return ctx.InternalServerError(goa.ErrBadRequest("no-auth"))
	}
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
	if !auth.HasAuth(ctx.Context) {
		return ctx.InternalServerError(goa.ErrBadRequest("no-auth"))
	}

	if ctx.Payload.Title == nil && ctx.Payload.Description == nil && ctx.Payload.Done == nil {
		return ctx.BadRequest(goa.ErrBadRequest("title, description or done flag must be set"))
	}

	dbTodo := &db.Todo{
		ID: ctx.TodoID,
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

	if ctx.Payload.Done != nil && *ctx.Payload.Done == true {
		dbTodo.CompletedAt = int(time.Now().Unix())
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

	dbTodo.CreatedBy = existingTodo.CreatedBy
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
		Done:        &res.Done,
		CreatedBy:   &res.CreatedBy,
	}

	return ctx.OK(appTodo)
}

// FilterTodos runs the filterTodos action.
func (c *TodosController) FilterTodos(ctx *app.FilterTodosTodoContext) error {
	// TodosController_FilterTodos: start_implement

	colFilter := map[string]interface{}{}

	if ctx.Payload.Filter != nil {
		var ok bool
		colFilter, ok = ctx.Payload.Filter.(map[string]interface{})
		if !ok {
			// try if string, to match by ID
			idval, ok := ctx.Payload.Filter.(string)
			if !ok {
				return ctx.BadRequest(goa.ErrBadRequest("invalid filter"))
			}
			colFilter = map[string]interface{}{"id": idval}
		}
	}

	sortSpecs := []db.SortSpec{}
	if ctx.Payload.Order != nil {
		for _, spec := range ctx.Payload.Order {
			sortSpecs = append(sortSpecs, db.SortSpec{
				Order:    *spec.Direction,
				Property: *spec.Property,
			})
		}
	}

	filter := db.Filter{
		After:    "",
		Page:     ctx.Payload.Page,
		PageSize: ctx.Payload.PageSize,
		Filter:   colFilter,
		Sort:     sortSpecs,
	}

	todos, err := c.TodoStore.DBFindTodos(&filter)
	if err != nil {
		if backends.IsErrInvalidInput(err) {
			return ctx.BadRequest(goa.ErrBadRequest(err))
		}
		return ctx.InternalServerError(goa.ErrInternal(err))
	}

	items := []*app.TodoMedia{}

	if todos.Items != nil {
		for _, rc := range todos.Items {
			todo := &db.Todo{}

			if err = backends.MapToInterface(rc, todo); err != nil {
				return ctx.InternalServerError(goa.ErrInternal(err))
			}
			items = append(items, toTodo(todo, true))
		}
	}

	res := &app.PaginatedTodosMedia{
		Page:     &todos.Page,
		PageSize: &todos.PageSize,
		Total:    &todos.Total,
		Items:    items,
	}
	return ctx.OK(res)
	// TodosController_FilterTodos: end_implement
}

func toTodo(td *db.Todo, maskCreds bool) *app.TodoMedia {
	createdAt := int(td.CreatedAt)
	return &app.TodoMedia{
		CreatedAt:   createdAt,
		ID:          td.ID,
		CompletedAt: &td.CompletedAt,
		Done:        &td.Done,
		Description: &td.Description,
		Title:       &td.Title,
		CreatedBy:   &td.CreatedBy,
	}
}
