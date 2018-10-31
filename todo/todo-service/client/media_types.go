// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "microtodo": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/Microkubes/examples/todo/todo-service/design
// --out=$(GOPATH)/src/github.com/Microkubes/examples/todo/todo-service
// --version=v1.3.1

package client

import (
	"github.com/goadesign/goa"
	"net/http"
)

// TodoMedia media type (default view)
//
// Identifier: application/json; view=default
type TodoMedia struct {
	CompletedAt *int    `bson:"completedAt,omitempty" form:"completedAt,omitempty" json:"completedAt,omitempty" yaml:"completedAt,omitempty"`
	CreatedAt   int     `bson:"createdAt,omitempty" form:"createdAt,omitempty" json:"createdAt,omitempty" yaml:"createdAt,omitempty"`
	CreatedBy   *string `bson:"createdBy,omitempty" form:"createdBy,omitempty" json:"createdBy,omitempty" yaml:"createdBy,omitempty"`
	Description *string `bson:"description,omitempty" form:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	Done        *bool   `bson:"done,omitempty" form:"done,omitempty" json:"done,omitempty" yaml:"done,omitempty"`
	ID          string  `bson:"_id,omitempty" form:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Title       *string `bson:"title,omitempty" form:"title,omitempty" json:"title,omitempty" yaml:"title,omitempty"`
}

// Validate validates the TodoMedia media type instance.
func (mt *TodoMedia) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}

	return
}

// DecodeTodoMedia decodes the TodoMedia instance encoded in resp body.
func (c *Client) DecodeTodoMedia(resp *http.Response) (*TodoMedia, error) {
	var decoded TodoMedia
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeErrorResponse decodes the ErrorResponse instance encoded in resp body.
func (c *Client) DecodeErrorResponse(resp *http.Response) (*goa.ErrorResponse, error) {
	var decoded goa.ErrorResponse
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// PaginatedTodosMedia media type (default view)
//
// Identifier: paginatedtodosmedia; view=default
type PaginatedTodosMedia struct {
	// List of todos
	Items []*TodoMedia `form:"items,omitempty" json:"items,omitempty" yaml:"items,omitempty" xml:"items,omitempty"`
	// Current page number
	Page *int `form:"page,omitempty" json:"page,omitempty" yaml:"page,omitempty" xml:"page,omitempty"`
	// Number of items per page
	PageSize *int `form:"pageSize,omitempty" json:"pageSize,omitempty" yaml:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Total number of items
	Total *int `form:"total,omitempty" json:"total,omitempty" yaml:"total,omitempty" xml:"total,omitempty"`
}

// Validate validates the PaginatedTodosMedia media type instance.
func (mt *PaginatedTodosMedia) Validate() (err error) {
	for _, e := range mt.Items {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodePaginatedTodosMedia decodes the PaginatedTodosMedia instance encoded in resp body.
func (c *Client) DecodePaginatedTodosMedia(resp *http.Response) (*PaginatedTodosMedia, error) {
	var decoded PaginatedTodosMedia
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
