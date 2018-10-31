// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "microtodo": Application User Types
//
// Command:
// $ goagen
// --design=github.com/Microkubes/examples/todo/todo-service/design
// --out=$(GOPATH)/src/github.com/Microkubes/examples/todo/todo-service
// --version=v1.3.1

package app

import (
	"github.com/goadesign/goa"
	"time"
)

// filterTodoPayload user type.
type filterTodoPayload struct {
	// Filter by fields key=>value
	Filter interface{} `form:"filter,omitempty" json:"filter,omitempty" yaml:"filter,omitempty" xml:"filter,omitempty"`
	// Sort specifications.
	Order []*orderSpecs `form:"order,omitempty" json:"order,omitempty" yaml:"order,omitempty" xml:"order,omitempty"`
	// Page number to fetch
	Page *int `form:"page,omitempty" json:"page,omitempty" yaml:"page,omitempty" xml:"page,omitempty"`
	// Number of items per page
	PageSize *int `form:"pageSize,omitempty" json:"pageSize,omitempty" yaml:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

// Validate validates the filterTodoPayload type instance.
func (ut *filterTodoPayload) Validate() (err error) {
	if ut.Page == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "page"))
	}
	if ut.PageSize == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "pageSize"))
	}
	return
}

// Publicize creates FilterTodoPayload from filterTodoPayload
func (ut *filterTodoPayload) Publicize() *FilterTodoPayload {
	var pub FilterTodoPayload
	if ut.Filter != nil {
		pub.Filter = ut.Filter
	}
	if ut.Order != nil {
		pub.Order = make([]*OrderSpecs, len(ut.Order))
		for i2, elem2 := range ut.Order {
			pub.Order[i2] = elem2.Publicize()
		}
	}
	if ut.Page != nil {
		pub.Page = *ut.Page
	}
	if ut.PageSize != nil {
		pub.PageSize = *ut.PageSize
	}
	return &pub
}

// FilterTodoPayload user type.
type FilterTodoPayload struct {
	// Filter by fields key=>value
	Filter interface{} `form:"filter,omitempty" json:"filter,omitempty" yaml:"filter,omitempty" xml:"filter,omitempty"`
	// Sort specifications.
	Order []*OrderSpecs `form:"order,omitempty" json:"order,omitempty" yaml:"order,omitempty" xml:"order,omitempty"`
	// Page number to fetch
	Page int `form:"page" json:"page" yaml:"page" xml:"page"`
	// Number of items per page
	PageSize int `form:"pageSize" json:"pageSize" yaml:"pageSize" xml:"pageSize"`
}

// Validate validates the FilterTodoPayload type instance.
func (ut *FilterTodoPayload) Validate() (err error) {

	return
}

// orderSpecs user type.
type orderSpecs struct {
	// Sort direction. One of 'asc' (ascending) or 'desc' (descenting).
	Direction *string `form:"direction,omitempty" json:"direction,omitempty" yaml:"direction,omitempty" xml:"direction,omitempty"`
	// Order by property
	Property *string `form:"property,omitempty" json:"property,omitempty" yaml:"property,omitempty" xml:"property,omitempty"`
}

// Publicize creates OrderSpecs from orderSpecs
func (ut *orderSpecs) Publicize() *OrderSpecs {
	var pub OrderSpecs
	if ut.Direction != nil {
		pub.Direction = ut.Direction
	}
	if ut.Property != nil {
		pub.Property = ut.Property
	}
	return &pub
}

// OrderSpecs user type.
type OrderSpecs struct {
	// Sort direction. One of 'asc' (ascending) or 'desc' (descenting).
	Direction *string `form:"direction,omitempty" json:"direction,omitempty" yaml:"direction,omitempty" xml:"direction,omitempty"`
	// Order by property
	Property *string `form:"property,omitempty" json:"property,omitempty" yaml:"property,omitempty" xml:"property,omitempty"`
}

// todo user type.
type todo struct {
	// Timestamp (milliseconds) when this todo item was completed.
	CompletedAt *time.Time `form:"completedAt,omitempty" json:"completedAt,omitempty" yaml:"completedAt,omitempty" xml:"completedAt,omitempty"`
	// Timestamp (milliseconds) when this todo item was created.
	CreatedAt *time.Time `form:"createdAt,omitempty" json:"createdAt,omitempty" yaml:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// Todo item text.
	Description *string `form:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty" xml:"description,omitempty"`
	// Is this todo item completed.
	Done *bool `form:"done,omitempty" json:"done,omitempty" yaml:"done,omitempty" xml:"done,omitempty"`
	// The item's unique identifier.
	ID *string `form:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty" xml:"id,omitempty"`
	// Todo item title.
	Title *string `form:"title,omitempty" json:"title,omitempty" yaml:"title,omitempty" xml:"title,omitempty"`
}

// Publicize creates Todo from todo
func (ut *todo) Publicize() *Todo {
	var pub Todo
	if ut.CompletedAt != nil {
		pub.CompletedAt = ut.CompletedAt
	}
	if ut.CreatedAt != nil {
		pub.CreatedAt = ut.CreatedAt
	}
	if ut.Description != nil {
		pub.Description = ut.Description
	}
	if ut.Done != nil {
		pub.Done = ut.Done
	}
	if ut.ID != nil {
		pub.ID = ut.ID
	}
	if ut.Title != nil {
		pub.Title = ut.Title
	}
	return &pub
}

// Todo user type.
type Todo struct {
	// Timestamp (milliseconds) when this todo item was completed.
	CompletedAt *time.Time `form:"completedAt,omitempty" json:"completedAt,omitempty" yaml:"completedAt,omitempty" xml:"completedAt,omitempty"`
	// Timestamp (milliseconds) when this todo item was created.
	CreatedAt *time.Time `form:"createdAt,omitempty" json:"createdAt,omitempty" yaml:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// Todo item text.
	Description *string `form:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty" xml:"description,omitempty"`
	// Is this todo item completed.
	Done *bool `form:"done,omitempty" json:"done,omitempty" yaml:"done,omitempty" xml:"done,omitempty"`
	// The item's unique identifier.
	ID *string `form:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty" xml:"id,omitempty"`
	// Todo item title.
	Title *string `form:"title,omitempty" json:"title,omitempty" yaml:"title,omitempty" xml:"title,omitempty"`
}

// Todo payload
type todoPayload struct {
	// Todo description
	Description *string `form:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty" xml:"description,omitempty"`
	// Todo title
	Title *string `form:"title,omitempty" json:"title,omitempty" yaml:"title,omitempty" xml:"title,omitempty"`
}

// Validate validates the todoPayload type instance.
func (ut *todoPayload) Validate() (err error) {
	if ut.Title == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "title"))
	}
	if ut.Description == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "description"))
	}
	return
}

// Publicize creates TodoPayload from todoPayload
func (ut *todoPayload) Publicize() *TodoPayload {
	var pub TodoPayload
	if ut.Description != nil {
		pub.Description = *ut.Description
	}
	if ut.Title != nil {
		pub.Title = *ut.Title
	}
	return &pub
}

// Todo payload
type TodoPayload struct {
	// Todo description
	Description string `form:"description" json:"description" yaml:"description" xml:"description"`
	// Todo title
	Title string `form:"title" json:"title" yaml:"title" xml:"title"`
}

// Validate validates the TodoPayload type instance.
func (ut *TodoPayload) Validate() (err error) {
	if ut.Title == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "title"))
	}
	if ut.Description == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "description"))
	}
	return
}

// Todo update payload
type todoUpdatePayload struct {
	// Todo description
	Description *string `form:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty" xml:"description,omitempty"`
	// Todo status
	Done *bool `form:"done,omitempty" json:"done,omitempty" yaml:"done,omitempty" xml:"done,omitempty"`
	// Todo title
	Title *string `form:"title,omitempty" json:"title,omitempty" yaml:"title,omitempty" xml:"title,omitempty"`
}

// Publicize creates TodoUpdatePayload from todoUpdatePayload
func (ut *todoUpdatePayload) Publicize() *TodoUpdatePayload {
	var pub TodoUpdatePayload
	if ut.Description != nil {
		pub.Description = ut.Description
	}
	if ut.Done != nil {
		pub.Done = ut.Done
	}
	if ut.Title != nil {
		pub.Title = ut.Title
	}
	return &pub
}

// Todo update payload
type TodoUpdatePayload struct {
	// Todo description
	Description *string `form:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty" xml:"description,omitempty"`
	// Todo status
	Done *bool `form:"done,omitempty" json:"done,omitempty" yaml:"done,omitempty" xml:"done,omitempty"`
	// Todo title
	Title *string `form:"title,omitempty" json:"title,omitempty" yaml:"title,omitempty" xml:"title,omitempty"`
}
