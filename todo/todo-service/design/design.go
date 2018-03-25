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
	Origin("*", func() {
		Methods("GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS")
	})
	Action("list", func() {
		Routing(GET("/todo"))
		Response(OK, func() {
			Media(CollectionOf(TodoMedia, func() {
				View("default")
			}))
		})
		Response(InternalServerError, ErrorMedia)
	})

	Action("add", func() {
		Routing(POST("/todo"))
		Payload(Todo)
		Response(OK)
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
	Attribute("owner", String, "Todo item owner's user ID")
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
		Attribute("createdAt", DateTime, func() {
			Metadata("struct:tag:json", "createdAt", "omitempty")
			Metadata("struct:tag:form", "createdAt", "omitempty")
			Metadata("struct:tag:yaml", "createdAt", "omitempty")
			Metadata("struct:tag:bson", "createdAt", "omitempty")
		})
		Attribute("completedAt", DateTime, func() {
			Metadata("struct:tag:json", "completedAt", "omitempty")
			Metadata("struct:tag:form", "completedAt", "omitempty")
			Metadata("struct:tag:yaml", "completedAt", "omitempty")
			Metadata("struct:tag:bson", "completedAt", "omitempty")
		})
		Attribute("owner", String, func() {
			Metadata("struct:tag:json", "owner", "omitempty")
			Metadata("struct:tag:form", "owner", "omitempty")
			Metadata("struct:tag:yaml", "owner", "omitempty")
			Metadata("struct:tag:bson", "owner", "omitempty")
		})

		Required("id", "createdAt", "done")
	})

	View("default", func() {
		Attribute("id")
		Attribute("title")
		Attribute("description")
		Attribute("done")
		Attribute("createdAt")
		Attribute("completedAt")
		Attribute("owner")
	})
})
