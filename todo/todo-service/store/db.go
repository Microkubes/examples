package store

import (
	"fmt"
	"os"
	"time"

	"github.com/Microkubes/examples/todo/todo-service/app"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type TodoCollection struct {
	*mgo.Collection
}

type ITodoCollection interface {
	CreateTodo(payload *app.Todo) (string, error)
	ListTodos() (app.TodoMediaCollection, error)
}

func NewSession(host string, username string, password string, database string) *mgo.Collection {
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{host},
		Username: username,
		Password: password,
		Database: database,
		Source:   "admin",
		Timeout:  30 * time.Second,
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	session.SetMode(mgo.Monotonic, true)

	collection := session.DB("tododb").C("todos")
	return collection
}

func (c *TodoCollection) CreateTodo(payload *app.Todo) (string, error) {
	id := bson.NewObjectId()
	err := c.Insert(bson.M{
		"_id":         id,
		"title":       payload.Title,
		"description": payload.Description,
		"status":      "not_done",
		"created_at":  time.Now().UTC(),
	})
	if err != nil {
		return "", err
	}
	return id.Hex(), nil
}

func (c *TodoCollection) ListTodos() (app.TodoMediaCollection, error) {
	var result app.TodoMediaCollection
	err := c.Collection.Find(bson.M{}).All(&result)
	for _, todo := range result {
		bsonId := bson.ObjectId(*todo.ID).Hex()
		todo.ID = &bsonId
	}
	if err != nil {
		return nil, err
	}
	return result, nil
}
