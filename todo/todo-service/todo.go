package main

import (
	"github.com/Microkubes/examples/todo/todo-service/app"
	"github.com/Microkubes/examples/todo/todo-service/store"
	"github.com/goadesign/goa"
)

// TodoController implements the todo resource.
type TodoController struct {
	*goa.Controller
	collection store.ITodoCollection
}

// NewTodoController creates a todo controller.
func NewTodoController(service *goa.Service, collection store.ITodoCollection) *TodoController {
	return &TodoController{Controller: service.NewController("TodoController"), collection: collection}
}

// Add runs the add action.
func (c *TodoController) Add(ctx *app.AddTodoContext) error {
	// TodoController_Add: start_implement
	// authObj := auth.GetAuth(ctx.Context)

	// ctx.Payload.Owner = &authObj.UserID

	id, err := c.collection.CreateTodo(ctx.Payload)
	if err != nil {
		return ctx.InternalServerError(err)
	}
	return ctx.OK([]byte(id))
	// TodoController_Add: end_implement
}

// List runs the list action.
func (c *TodoController) List(ctx *app.ListTodoContext) error {
	// authObj := auth.GetAuth(ctx.Context)
	// TodoController_List: start_implement
	res, err := c.collection.ListTodos()
	if err != nil {
		return ctx.InternalServerError(err)
	}
	return ctx.OK(res)
	// TodoController_List: end_implement
}
