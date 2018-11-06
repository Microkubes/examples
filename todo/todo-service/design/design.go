package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("microtodo", func() {
	Title("Micro Todo")
	Host("localhost:8080")
	Scheme("http")
})

var _ = Resource("todo", func() {
	BasePath("todo")
	Origin("*", func() {
		Methods("GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS")
	})

	Action("addTodo", func() {
		Description("Add new todo")
		Routing(POST("/add"))
		Payload(TodoPayload)
		Response(Created, TodoMedia)
		Response(NotFound, ErrorMedia)
		Response(BadRequest, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})

	Action("getAllTodos", func() {
		Description("Get all todos")
		// order string, sorting string, limit int, offset int
		Routing(GET("/all"))
		Params(func() {
			Param("order", String, "order by")
			Param("sorting", String, func() {
				Enum("asc", "desc")
			})
			Param("limit", Integer, "Limit todos per page")
			Param("offset", Integer, "number of todos to skip")
		})
		Response(OK)
		Response(NotFound, ErrorMedia)
		Response(BadRequest, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})

	Action("getById", func() {
		Description("Get todo by ID")
		Routing(GET("/:todoID"))
		Params(func() {
			Param("todoID", String, "Todo ID")
		})
		Response(OK, TodoMedia)
		Response(NotFound, ErrorMedia)
		Response(BadRequest, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})

	Action("deleteTodo", func() {
		Description("Delete todo")
		Routing(DELETE("/:todoID/delete"))
		Params(func() {
			Param("todoID", String, "Todo ID")
		})
		Response(OK)
		Response(NotFound, ErrorMedia)
		Response(BadRequest, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})

	Action("updateTodo", func() {
		Description("Update todo")
		Routing(PUT("/:todoID"), PATCH("/:todoID"))
		Params(func() {
			Param("todoID", String, "Todo ID")
		})
		Payload(TodoUpdatePayload)
		Response(OK, TodoMedia)
		Response(BadRequest, ErrorMedia)
		Response(NotFound, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})

	Action("filterTodos", func() {
		Description("Filter (lookup) todos")
		Routing(POST("/filter"))
		Payload(FilterTodoPayload)
		Response(OK, PaginatedTodosMedia)
		Response(BadRequest, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})
})

// Todo item
var Todo = Type("Todo", func() {
	Attribute("id", String, "The item's unique identifier.")
	Attribute("title", String, "Todo item title.")
	Attribute("description", String, "Todo item text.")
	Attribute("done", Boolean, "Is this todo item completed.")
	Attribute("createdAt", DateTime, "Timestamp (milliseconds) when this todo item was created.")
	Attribute("completedAt", DateTime, "Timestamp (milliseconds) when this todo item was completed.")
	Attribute("createdBy", String, "User who created the todo.")
})

// TodoMedia is todo item media type
var TodoMedia = MediaType("application/json", func() {
	TypeName("TodoMedia")

	Attributes(func() {
		Attribute("id", String, func() {
			Metadata("struct:tag:json", "id", "omitempty")
			Metadata("struct:tag:form", "id", "omitempty")
			Metadata("struct:tag:yaml", "id", "omitempty")
			Metadata("struct:tag:bson", "_id", "omitempty")
		})

		Attribute("title", String, func() {
			Metadata("struct:tag:json", "title", "omitempty")
			Metadata("struct:tag:form", "title", "omitempty")
			Metadata("struct:tag:yaml", "title", "omitempty")
			Metadata("struct:tag:bson", "title", "omitempty")
		})
		Attribute("description", String, func() {
			Metadata("struct:tag:json", "description", "omitempty")
			Metadata("struct:tag:form", "description", "omitempty")
			Metadata("struct:tag:yaml", "description", "omitempty")
			Metadata("struct:tag:bson", "description", "omitempty")
		})
		Attribute("done", Boolean, func() {
			Metadata("struct:tag:json", "done", "omitempty")
			Metadata("struct:tag:form", "done", "omitempty")
			Metadata("struct:tag:yaml", "done", "omitempty")
			Metadata("struct:tag:bson", "done", "omitempty")
		})
		Attribute("createdAt", Integer, func() {
			Metadata("struct:tag:json", "createdAt", "omitempty")
			Metadata("struct:tag:form", "createdAt", "omitempty")
			Metadata("struct:tag:yaml", "createdAt", "omitempty")
			Metadata("struct:tag:bson", "createdAt", "omitempty")
		})
		Attribute("completedAt", Integer, func() {
			Metadata("struct:tag:json", "completedAt", "omitempty")
			Metadata("struct:tag:form", "completedAt", "omitempty")
			Metadata("struct:tag:yaml", "completedAt", "omitempty")
			Metadata("struct:tag:bson", "completedAt", "omitempty")
		})
		Attribute("createdBy", String, func() {
			Metadata("struct:tag:json", "createdBy", "omitempty")
			Metadata("struct:tag:form", "createdBy", "omitempty")
			Metadata("struct:tag:yaml", "createdBy", "omitempty")
			Metadata("struct:tag:bson", "createdBy", "omitempty")
		})

		Required("id", "createdAt")
	})

	View("default", func() {
		Attribute("id")
		Attribute("title")
		Attribute("description")
		Attribute("done")
		Attribute("createdAt")
		Attribute("completedAt")
		Attribute("createdBy")
	})
})

// TodoPayload defines the payload for Todo objects.
var TodoPayload = Type("TodoPayload", func() {
	Description("Todo payload")
	Attribute("title", String, "Todo title")
	Attribute("description", String, "Todo description")
	Required("title", "description")
})

// TodoUpdatePayload defines the payload for Todo objects.
var TodoUpdatePayload = Type("TodoUpdatePayload", func() {
	Description("Todo update payload")
	Attribute("title", String, "Todo title")
	Attribute("description", String, "Todo description")
	Attribute("done", Boolean, "Todo status")
})

// FilterTodoPayload defines the filter object used in queries for filtering todos.
// This defines the pagination attributes, the filter and the sort/order specification.
var FilterTodoPayload = Type("FilterTodoPayload", func() {
	Attribute("page", Integer, "Page number to fetch")
	Attribute("pageSize", Integer, "Number of items per page")
	Attribute("filter", Any, "Filter by fields key=>value")
	Attribute("order", ArrayOf(OrderSpecs), "Sort specifications.")

	Required("page", "pageSize")
})

// OrderSpecs specification for the ordering of sorted results.
// Holds the property to sort by and the direction of sort - 'asc' ascending or 'desc' descending order.
var OrderSpecs = Type("OrderSpecs", func() {
	Attribute("property", String, "Order by property")
	Attribute("direction", String, "Sort direction. One of 'asc' (ascending) or 'desc' (descenting).")
})

// PaginatedTodosMedia defines the paginated result of multiple todos.
var PaginatedTodosMedia = MediaType("PaginatedTodosMedia", func() {
	TypeName("PaginatedTodosMedia")
	Attribute("page", Integer, "Current page number")
	Attribute("pageSize", Integer, "Number of items per page")
	Attribute("total", Integer, "Total number of items")
	Attribute("items", ArrayOf(TodoMedia), "List of todos")

	View("default", func() {
		Attribute("page")
		Attribute("pageSize")
		Attribute("total")
		Attribute("items")
	})

})
