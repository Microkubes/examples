package db

import (
	"encoding/json"
	"sync"

	errors "github.com/Microkubes/backends"
	"github.com/goadesign/goa"
)

//MapStore - map store
type MapStore map[string]interface{}

//DB database structure
type DB struct {
	sync.Mutex
	MapStore
}

// NewDB initializes a new "DB" with dummy data.
func NewDB() *DB {
	todos := &DB{
		MapStore: map[string]interface{}{
			"5ad14bb71da4d95a87710df1": map[string]interface{}{
				"ID":          "5ae4b86f1da4d9046c9113f4",
				"Title":       "Todo title",
				"Description": "Todo description",
				"CreatedAt":   1524938863,
				"CreatedBy":   "Creator",
				"CompletedAt": 123321456654,
				"Done":        false,
			},
		},
	}

	return todos

}

var (
	IDOK            = "5ad14bb71da4d95a87710df1"
	IDBadRequest    = "5ad14bb71da4d95a87710666"
	IDInternalError = "5ad14bb71da4d95a87710999"
	IDNotFound      = "5ad14bb71da4d95a87710555"

	TitleOK                  = "OK"
	TitleBadRequest          = "BadRequest"
	TitleInternalError       = "InternalError"
	TitleNotFound            = "NotFound"
	TitleExisting            = "Existing"
	DescriptionBadRequest    = "Description Bad Request"
	DescriptionInternalError = "Description Internal Server Error"
	DescriptionNotFound      = "Description Not Found"
	PageOK                   = 1
	PageSizeOK               = 2
	PageBadRequest           = -1
	PageSizeBadRequest       = -3
	PageInternalError        = 100
	PageSizeInternalError    = 101

	Title       = "Todo Title"
	Description = "Todo description"
)

var dummyTodo = &Todo{
	CreatedAt:   1524938863,
	CreatedBy:   "Creator",
	Description: "Todo description",
	ID:          IDOK,
	Title:       "Todo title",
}

// DBGetByID DBGetByID - returns todo by ID
func (r *DB) DBGetByID(todoID string) (*Todo, error) {

	//bad Request
	if todoID == IDBadRequest {
		return nil, errors.ErrInvalidInput()
	}

	//Internal Server Error
	if todoID == IDInternalError {
		return nil, errors.ErrBackendError()
	}

	//notFound error
	if todoID == IDNotFound {
		return nil, errors.ErrNotFound()
	}

	//OK
	return dummyTodo, nil
}

//DBAddTodo adds the Todo
func (r *DB) DBAddTodo(todo *Todo) (*Todo, error) {

	//bad Request
	if todo.Title == TitleBadRequest {
		return nil, errors.ErrInvalidInput()
	}

	//Internal Server Error
	if todo.Title == TitleInternalError {
		return nil, errors.ErrBackendError()
	}

	//OK
	if todo.Title == TitleOK {
		return dummyTodo, nil
	}

	//Not found
	if todo.Title == TitleNotFound {
		return nil, errors.ErrNotFound()
	}

	//otherwise
	return nil, nil
}

//DBUpdateTodo updates todo in DB
func (r *DB) DBUpdateTodo(todo *Todo) (*Todo, error) {

	//bad Request
	if todo.Title == TitleBadRequest {
		return nil, goa.ErrInvalidRequest("bad request")
	}

	//Internal Server Error
	if todo.Title == TitleInternalError {
		return nil, goa.ErrInternal("Internal server error")
	}

	//OK
	if todo.ID == IDOK {
		return dummyTodo, nil
	}

	//not found
	return nil, goa.ErrNotFound("There is no such todo")
}

//DBDeleteTodo deletes todo
func (r *DB) DBDeleteTodo(todoID string) error {

	//Error Bad Request
	if todoID == IDBadRequest {
		return errors.ErrInvalidInput("bad request")
	}

	//Error Internal Server Error
	if todoID == IDInternalError {
		//return goa.ErrInternal("internal server error")
		return errors.ErrBackendError("internal error")
	}

	//Error not found
	if todoID == IDNotFound {
		//return goa.ErrNotFound("not found")
		return errors.ErrNotFound("todo was not found")
	}

	return nil
}

//DBGetAllTodos lists all todos
func (r *DB) DBGetAllTodos(order string, sorting string, limit int, offset int) ([]byte, error) {
	return json.Marshal(dummyTodo)
}

// DBFindTodos finds todos based on search criteria
func (r *DB) DBFindTodos(filter *Filter) (*Todos, error) {

	if filter.Page == PageBadRequest {
		return nil, errors.ErrInvalidInput()
	}

	if filter.Page == PageInternalError {
		return nil, errors.ErrBackendError()
	}

	var items []*map[string]interface{}

	item := &map[string]interface{}{
		"title": 1,
		"id":    "123",
	}

	foo := append(items, item)

	var dummyTodos = &Todos{
		Page:     1,
		PageSize: 2,
		Total:    3,
		Items:    foo,
	}
	return dummyTodos, nil
}
