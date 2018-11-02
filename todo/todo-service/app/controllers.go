// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "microtodo": Application Controllers
//
// Command:
// $ goagen
// --design=github.com/Microkubes/examples/todo/todo-service/design
// --out=$(GOPATH)/src/github.com/Microkubes/examples/todo/todo-service
// --version=v1.3.1

package app

import (
	"context"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/cors"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// TodoController is the controller interface for the Todo actions.
type TodoController interface {
	goa.Muxer
	AddTodo(*AddTodoTodoContext) error
	DeleteTodo(*DeleteTodoTodoContext) error
	FilterTodos(*FilterTodosTodoContext) error
	GetAllTodos(*GetAllTodosTodoContext) error
	GetByID(*GetByIDTodoContext) error
	UpdateTodo(*UpdateTodoTodoContext) error
}

// MountTodoController "mounts" a Todo resource controller on the given service.
func MountTodoController(service *goa.Service, ctrl TodoController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/todo/add", ctrl.MuxHandler("preflight", handleTodoOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/todo/:todoID/delete", ctrl.MuxHandler("preflight", handleTodoOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/todo/filter", ctrl.MuxHandler("preflight", handleTodoOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/todo/all", ctrl.MuxHandler("preflight", handleTodoOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/todo/:todoID", ctrl.MuxHandler("preflight", handleTodoOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewAddTodoTodoContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*TodoPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.AddTodo(rctx)
	}
	h = handleTodoOrigin(h)
	service.Mux.Handle("POST", "/todo/add", ctrl.MuxHandler("addTodo", h, unmarshalAddTodoTodoPayload))
	service.LogInfo("mount", "ctrl", "Todo", "action", "AddTodo", "route", "POST /todo/add")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDeleteTodoTodoContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.DeleteTodo(rctx)
	}
	h = handleTodoOrigin(h)
	service.Mux.Handle("DELETE", "/todo/:todoID/delete", ctrl.MuxHandler("deleteTodo", h, nil))
	service.LogInfo("mount", "ctrl", "Todo", "action", "DeleteTodo", "route", "DELETE /todo/:todoID/delete")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewFilterTodosTodoContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*FilterTodoPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.FilterTodos(rctx)
	}
	h = handleTodoOrigin(h)
	service.Mux.Handle("POST", "/todo/filter", ctrl.MuxHandler("filterTodos", h, unmarshalFilterTodosTodoPayload))
	service.LogInfo("mount", "ctrl", "Todo", "action", "FilterTodos", "route", "POST /todo/filter")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetAllTodosTodoContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.GetAllTodos(rctx)
	}
	h = handleTodoOrigin(h)
	service.Mux.Handle("GET", "/todo/all", ctrl.MuxHandler("getAllTodos", h, nil))
	service.LogInfo("mount", "ctrl", "Todo", "action", "GetAllTodos", "route", "GET /todo/all")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetByIDTodoContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.GetByID(rctx)
	}
	h = handleTodoOrigin(h)
	service.Mux.Handle("GET", "/todo/:todoID", ctrl.MuxHandler("getById", h, nil))
	service.LogInfo("mount", "ctrl", "Todo", "action", "GetByID", "route", "GET /todo/:todoID")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdateTodoTodoContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*TodoUpdatePayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.UpdateTodo(rctx)
	}
	h = handleTodoOrigin(h)
	service.Mux.Handle("PUT", "/todo/:todoID", ctrl.MuxHandler("updateTodo", h, unmarshalUpdateTodoTodoPayload))
	service.LogInfo("mount", "ctrl", "Todo", "action", "UpdateTodo", "route", "PUT /todo/:todoID")
	service.Mux.Handle("PATCH", "/todo/:todoID", ctrl.MuxHandler("updateTodo", h, unmarshalUpdateTodoTodoPayload))
	service.LogInfo("mount", "ctrl", "Todo", "action", "UpdateTodo", "route", "PATCH /todo/:todoID")
}

// handleTodoOrigin applies the CORS response headers corresponding to the origin.
func handleTodoOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, HEAD, OPTIONS")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalAddTodoTodoPayload unmarshals the request body into the context request data Payload field.
func unmarshalAddTodoTodoPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &todoPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalFilterTodosTodoPayload unmarshals the request body into the context request data Payload field.
func unmarshalFilterTodosTodoPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &filterTodoPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalUpdateTodoTodoPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateTodoTodoPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &todoUpdatePayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}
