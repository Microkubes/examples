// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "microtodo": Application Contexts
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
	"net/http"
)

// AddTodoContext provides the todo add action context.
type AddTodoContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *Todo
}

// NewAddTodoContext parses the incoming request URL and body, performs validations and creates the
// context used by the todo controller add action.
func NewAddTodoContext(ctx context.Context, r *http.Request, service *goa.Service) (*AddTodoContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := AddTodoContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *AddTodoContext) OK(resp []byte) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	}
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *AddTodoContext) BadRequest() error {
	ctx.ResponseData.WriteHeader(400)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *AddTodoContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// ListTodoContext provides the todo list action context.
type ListTodoContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewListTodoContext parses the incoming request URL and body, performs validations and creates the
// context used by the todo controller list action.
func NewListTodoContext(ctx context.Context, r *http.Request, service *goa.Service) (*ListTodoContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ListTodoContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListTodoContext) OK(r TodoMediaCollection) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/json; type=collection")
	}
	if r == nil {
		r = TodoMediaCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ListTodoContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}
