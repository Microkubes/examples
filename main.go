//go:generate goagen bootstrap -d github.com/Microkubes/microtodo/design

package main

import (
	"github.com/Microkubes/microtodo/app"
	"github.com/Microkubes/microtodo/store"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"os"
)

func main() {
	// Create service
	service := goa.New("microtodo")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	collection := store.NewSession(os.Getenv("MICROTODO_HOST"), os.Getenv("MICROTODO_USERNAME"), os.Getenv("MICROTODO_PASSWORD"), "tododb")

	// Mount "todo" controller
	c := NewTodoController(service, &store.TodoCollection{collection})
	app.MountTodoController(service, c)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
