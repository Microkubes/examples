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
	Action("list", func() {
		Routing(GET(""))
		Response(OK, func() {
			Media(CollectionOf(TodoMedia, func() {
				View("default")
			}))
		})
		Response(InternalServerError)
	})

	Action("add", func() {
		Routing(POST(""))
		Payload(Todo)
		Response(OK)
		Response(BadRequest)
		Response(InternalServerError)
	})
})

var Todo = Type("Todo", func() {
	Attribute("title", String)
	Attribute("description", String)
})

var TodoMedia = MediaType("application/json", func() {
	TypeName("TodoMedia")

	Attributes(func() {
		Attribute("id", String, func() {
			Metadata("struct:tag:json", "id", "omitempty")
			Metadata("struct:tag:form", "id", "omitempty")
			Metadata("struct:tag:yaml", "id", "omitempty")
			Metadata("struct:tag:bson", "_id", "omitempty")
		})
		Attribute("title", String)
		Attribute("description", String)
		Attribute("status", String)
		Attribute("created_at", DateTime)
	})

	View("default", func() {
		Attribute("id")
		Attribute("title")
		Attribute("description")
		Attribute("status")
		Attribute("created_at")
	})
})
