package db

import (
	"encoding/json"

	"github.com/Microkubes/backends"
	"github.com/Microkubes/examples/todo/todo-service/app"
)

//Todo is the structure of the todo object
type Todo struct {
	ID          string `json:"id" bson:"_id"`
	Title       string `json:"title, omitempty" bson:"title"`
	Description string `json:"description, omitempty" bson:"description"`
	Done        bool   `json:"done" bson:"done"`
	CreatedAt   int    `json:"createdAt, omitempty" bson:"createdAt"`
	CompletedAt int    `json:"completedAt, omitempty" bson:"completedAt"`
	CreatedBy   string `json:"createdBy, omitempty" bson:"createdBy"`
}

// Todos represents a paginated results for multiple todos.
type Todos struct {
	Page     int                       `json:"page"`
	PageSize int                       `json:"pageSize"`
	Last     string                    `json:"last,omitempty"`
	Total    int                       `json:"total"`
	Items    []*map[string]interface{} `json:"items"`
}

// SortSpec represents a specification for a sort of todos list by a property
// and with a sorting direction.
type SortSpec struct {
	Property string `json:"property"`
	Order    string `json:"order"`
}

// Filter represents a request for looking up todos that match a certain criteria.
type Filter struct {
	Page     int                    `json:"page"`
	PageSize int                    `json:"pageSize"`
	After    string                 `json:"after,omitempty"`
	Filter   map[string]interface{} `json:"filter,omitempty"`
	Sort     []SortSpec             `json:"sort,omitempty"`
}

//TodoCollection is collection of todos
type TodoCollection []*app.TodoMedia

// TodoStore is an interface that defines the API for accessing and managing todos in the database.
type TodoStore interface {

	// Get performs a lookup for a todo by its ID.
	// Returns a pointer to a Todo and error.
	DBGetByID(todoID string) (*Todo, error)

	// AddTodo inserts Todo object in MongoDB.
	DBAddTodo(todo *Todo) (*Todo, error)

	//DBDeleteTodo deletes existing todo in MongoDB.
	DBDeleteTodo(todoID string) error

	//DBGetAllTodos lists all todos
	DBGetAllTodos(order string, sorting string, limit int, offset int) ([]byte, error)

	//DBUpdateTodo updates existing todo
	DBUpdateTodo(todo *Todo) (*Todo, error)

	// Find performs a lookup for todos that match certain criteria.
	DBFindTodos(filter *Filter) (*Todos, error)
}

//BackendTodosService holds data for implementation of the TodoStore interface.
type BackendTodosService struct {
	todosRepository backends.Repository
}

//DBGetByID returns only one Todo by its ID
func (r *BackendTodosService) DBGetByID(todoID string) (*Todo, error) {
	todo, err := r.todosRepository.GetOne(backends.NewFilter().Match("id", todoID), &Todo{})
	if err != nil {
		return nil, err
	}

	return todo.(*Todo), nil
}

//DBAddTodo adds the Todo in the DB
func (r *BackendTodosService) DBAddTodo(todo *Todo) (*Todo, error) {
	res, err := r.todosRepository.Save(todo, nil)
	if err != nil {
		return nil, err
	}

	return res.(*Todo), nil
}

//DBDeleteTodo deletes todo
func (r *BackendTodosService) DBDeleteTodo(todoID string) error {
	err := r.todosRepository.DeleteOne(backends.NewFilter().Match("id", todoID))
	if err != nil {
		return err
	}

	return nil
}

//DBGetAllTodos lists all todos
func (r *BackendTodosService) DBGetAllTodos(order string, sorting string, limit int, offset int) ([]byte, error) {
	var typeHint map[string]interface{}
	todos, err := r.todosRepository.GetAll(backends.NewFilter(), typeHint, order, sorting, limit, offset)
	if err != nil {
		return nil, err
	}

	return json.Marshal(todos)
}

//NewTodosService creates new TodoStore.
func NewTodosService(todosRepository backends.Repository) TodoStore {
	return &BackendTodosService{
		todosRepository: todosRepository,
	}
}

//DBUpdateTodo updates a todo in the DB.
func (r *BackendTodosService) DBUpdateTodo(todo *Todo) (*Todo, error) {
	td, err := r.todosRepository.Save(todo, backends.NewFilter().Match("id", todo.ID))
	if err != nil {
		return nil, err
	}

	dbTodo := &Todo{}

	if err = backends.MapToInterface(td, dbTodo); err != nil {
		return nil, err
	}

	return dbTodo, nil
}

// DBFindTodos performs a lookup for todos that match certain criteria.
func (r *BackendTodosService) DBFindTodos(filter *Filter) (*Todos, error) {
	if filter.Page <= 0 {
		return nil, backends.ErrInvalidInput("invalid page number")
	}
	limit := filter.PageSize
	if limit == 0 {
		limit = 10
	}
	offset := (filter.Page - 1) * limit

	selector := backends.NewFilter()
	if filter.Filter != nil {
		for prop, value := range filter.Filter {
			selector.MatchPattern(prop, value.(string))
		}
	}

	order := ""
	sort := ""
	if filter.Sort != nil {
		if len(filter.Sort) > 0 {
			order = filter.Sort[0].Order
			sort = filter.Sort[0].Property
		}
	}

	var typeHint map[string]interface{}

	result, err := r.todosRepository.GetAll(selector, typeHint, sort, order, limit, offset)
	if err != nil {
		return nil, err
	}

	totalTodos, err := r.todosRepository.GetAll(selector, typeHint, "", "", 0, 0)
	if err != nil {
		return nil, err
	}
	total := totalTodos.(*[]*map[string]interface{})

	todos := &Todos{
		Page:     filter.Page,
		PageSize: limit,
		Total:    len(*total),
	}

	rcArrPtr, ok := result.(*[]*map[string]interface{})
	if !ok {
		return nil, backends.ErrBackendError("Expected results to be of type []*Todo")
	}
	rcArr := *rcArrPtr
	todos.Items = rcArr

	return todos, nil
}
