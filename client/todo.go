// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "microtodo": todo Resource Client
//
// Command:
// $ goagen
// --design=github.com/Microkubes/microtodo/design
// --out=$(GOPATH)/src/github.com/Microkubes/microtodo
// --version=v1.3.1

package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// AddTodoPath computes a request path to the add action of todo.
func AddTodoPath() string {

	return fmt.Sprintf("/")
}

// AddTodo makes a request to the add action endpoint of the todo resource
func (c *Client) AddTodo(ctx context.Context, path string, payload *Todo, contentType string) (*http.Response, error) {
	req, err := c.NewAddTodoRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewAddTodoRequest create the request corresponding to the add action endpoint of the todo resource.
func (c *Client) NewAddTodoRequest(ctx context.Context, path string, payload *Todo, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}

// ListTodoPath computes a request path to the list action of todo.
func ListTodoPath() string {

	return fmt.Sprintf("/")
}

// ListTodo makes a request to the list action endpoint of the todo resource
func (c *Client) ListTodo(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListTodoRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListTodoRequest create the request corresponding to the list action endpoint of the todo resource.
func (c *Client) NewListTodoRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
