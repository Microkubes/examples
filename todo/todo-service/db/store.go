package db

import (
	"encoding/json"

	"github.com/JormungandrK/backends"
	"github.com/Microkubes/examples/todo/todo-service/app"
)

//Todo is the structure of the organizaiton object
type Todo struct {
	ID          string `json:"id" bson:"_id"`
	Title       string `json:"title, omitempty" bson:"title"`
	Description string `json:"description, omitempty" bson:"description"`
	Done        bool   `json:"done" bson:"done"`
	CreatedAt   int    `json:"createdAt, omitempty" bson:"createdAt"`
	CompletedAt int    `json:"completedAt, omitempty" bson:"completedAt"`
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

	//DBDeleteTodo update existing todo in
	DBDeleteTodo(todoID string) error

	//DBGetAllTodos lists all todos
	DBGetAllTodos(order string, sorting string, limit int, offset int) ([]byte, error)

	//DBUpdateTodo update existing todo
	DBUpdateTodo(todo *Todo) (*Todo, error)
}

//BackendTodosService holds data for implementation of the MetadataService interface.
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

//DBAddTodo adds the Todo
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

//NewTodoService creates new MetadataService.
func NewTodosService(todosRepository backends.Repository) TodoStore {
	return &BackendTodosService{
		todosRepository: todosRepository,
	}
}

//DBUpdateTodo is
func (r *BackendTodosService) DBUpdateTodo(todo *Todo) (*Todo, error) {
	org, err := r.todosRepository.Save(todo, backends.NewFilter().Match("id", todo.ID))
	if err != nil {
		return nil, err
	}

	dbTodo := &Todo{}

	if err = backends.MapToInterface(org, dbTodo); err != nil {
		return nil, err
	}

	return dbTodo, nil
}
