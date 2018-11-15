package main

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/Microkubes/examples/todo/todo-service/app"
	"github.com/Microkubes/examples/todo/todo-service/app/test"
	"github.com/Microkubes/examples/todo/todo-service/config"
	db "github.com/Microkubes/examples/todo/todo-service/db"
	"github.com/Microkubes/microservice-security/auth"
	"github.com/goadesign/goa"
)

var cfgBytes = []byte(`{
    "service":	{
		"name": "microservice-todo",
		"port": 8080,
		"paths": ["/todo"],
		"virtual_host": "microservice-todo.service.consul",
		"weight": 10,
		"slots": 100
	},
	"gatewayUrl": "http://kong:8000",
    "gatewayAdminUrl": "http://kong:8001",
    "database":{
        "dbName": "mongodb",
        "dbInfo": {
          "host": "mongo:27017",
          "database": "todos",
          "user": "restapi",
          "pass": "restapi"
        }
    },
    "security":{
        "keysDir": "/run/secrets",
        "ignorePatterns": ["/users/verify"],
        "jwt":{
          "name": "JWTSecurity",
          "description": "JWT security middleware",
          "tokenUrl": "http://kong:8000/jwt"
        },
        "saml":{
          "certFile": "/run/secrets/service.cert",
          "keyFile": "/run/secrets/service.key",
          "identityProviderUrl": "http://kong:8000/saml/idp",
          "userServiceUrl": "http://kong:8000/users",
          "registrationServiceUrl": "http://kong:8000/users/register",
          "rootURL": "http://localhost:8000/users"
        },
        "oauth2":{
          "description": "OAuth2 security middleware",
          "tokenUrl": "https://kong:8000/oauth2/token",
          "authorizeUrl": "https://kong:8000/oauth2/authorize"
        },
        "acl": {
          "policies":[{
              "id": "users-allow-admin-access",
              "description": "Allows access to everything to an admin user",
              "resources": ["/todo/<.+>"],
              "actions": ["api:read","api:write"],
              "effect": "allow",
              "subjects": ["<.+>"],
              "conditions": {
                "roles": {
                  "type": "RolesCondition",
                  "options": {
                    "values": ["admin", "system"]
                  }
                 }
              }
           },{
               "id": "users-allow-read-access",
               "description": "Allows users to create and read todos",
               "resources": ["/todo/<.+>", "/todo/add", "/todo/all", "/todo/<.+>/delete"],
               "actions": ["api:read", "api:write"],
               "effect": "allow",
               "subjects": ["<.+>"],
               "conditions": {
                "roles": {
                  "type": "RolesCondition",
                  "options": {
                    "values": ["user"]
                  }
                 }
              }
            },{
                "id": "read-swagger",
                "description": "Allows to service swagger.",
                "resources": ["/swagger<.+>"],
                "actions": ["api:read"],
                "effect": "allow",
                "subjects": ["<.+>"]
             }]
        }
    }
}
`)

var (
	service               = goa.New("todo-test")
	ctrl                  = NewTodosController(service, dbTest, cfg)
	ID                    = "5ad14bb71da4d95a87710df1"
	notFoundID            = "5ad14bb71da4d95a87710df8"
	badRequestID          = "5ad14bb71da4d95a87710666"
	internalServerErrorID = "5ad14bb71da4d95a87710999"
)

var payload = app.TodoPayload{
	// Todo description
	Description: db.TitleOK,
	// Todo title
	Title: db.TitleOK,
}

var payloadBadRequest = app.TodoPayload{
	// Todo description
	Description: db.TitleBadRequest,
	// Todo title
	Title: db.TitleBadRequest,
}

var payloadInternalServerError = app.TodoPayload{
	// Todo description
	Description: db.TitleInternalError,
	// Todo title
	Title: db.TitleInternalError,
}

var payloadNotFound = app.TodoPayload{
	// Todo description
	Description: db.TitleNotFound,
	// Todo title
	Title: db.TitleNotFound,
}

var payloadUpdateOK = app.TodoUpdatePayload{
	Description: &db.Description,
	Title:       &db.Title,
}

var payloadUpdateBadrequest = app.TodoUpdatePayload{
	// Todo description
	Description: &db.DescriptionBadRequest,
	// Todo title
	Title: &db.TitleBadRequest,
}

var payloadUpdateInternalServerError = app.TodoUpdatePayload{
	// Todo description
	Description: &db.DescriptionInternalError,
	// Todo title
	Title: &db.TitleInternalError,
}

var payloadUpdateNotFound = app.TodoUpdatePayload{
	// Todo description
	Description: &db.DescriptionNotFound,
	// Todo title
	Title: &db.TitleNotFound,
}

var payloadFilterOK = app.FilterTodoPayload{
	// Todos page
	Page: db.PageOK,
	// Todos pageSize
	PageSize: db.PageSizeOK,
}

var payloadFilterBadRequest = app.FilterTodoPayload{
	// Todos page
	Page: db.PageBadRequest,
	// Todos pageSize
	PageSize: db.PageSizeBadRequest,
}

var payloadFilterInternalError = app.FilterTodoPayload{
	// Todos page
	Page: db.PageInternalError,
	// Todos pageSize
	PageSize: db.PageSizeInternalError,
}

var cfg = &config.ServiceConfig{}
var _ = json.Unmarshal(cfgBytes, cfg)
var dbTest = db.NewDB()

func Test_AddTodoTodoBadRequest(t *testing.T) {
	ctx := context.Background()
	authObj := &auth.Auth{UserID: ID}
	ctx = auth.SetAuth(ctx, authObj)

	test.AddTodoTodoBadRequest(t, ctx, service, ctrl, &payloadBadRequest)
}

func Test_AddTodoTodosInternalServerError(t *testing.T) {
	test.AddTodoTodoInternalServerError(t, context.Background(), service, ctrl, &payloadInternalServerError)
}

func Test_AddTodoTodoCreated(t *testing.T) {
	ctx := context.Background()
	authObj := &auth.Auth{UserID: ID}
	ctx = auth.SetAuth(ctx, authObj)

	test.AddTodoTodoCreated(t, ctx, service, ctrl, &payload)
}

func Test_DeleteTodoTodoBadRequest(t *testing.T) {
	ctx := context.Background()
	authObj := &auth.Auth{UserID: ID}
	ctx = auth.SetAuth(ctx, authObj)

	test.DeleteTodoTodoBadRequest(t, ctx, service, ctrl, db.IDBadRequest)
}

func Test_DeleteTodoTodoInternalServerError(t *testing.T) {
	test.DeleteTodoTodoInternalServerError(t, context.Background(), service, ctrl, db.IDInternalError)
}

func Test_DeleteTodoTodoNotFound(t *testing.T) {
	ctx := context.Background()
	authObj := &auth.Auth{UserID: ID}
	ctx = auth.SetAuth(ctx, authObj)

	test.DeleteTodoTodoNotFound(t, ctx, service, ctrl, db.IDNotFound)
}

func Test_DeleteTodoTodoOK(t *testing.T) {
	ctx := context.Background()
	authObj := &auth.Auth{UserID: ID}
	ctx = auth.SetAuth(ctx, authObj)

	test.DeleteTodoTodoOK(t, ctx, service, ctrl, db.IDOK)
}

func Test_GetAllTodosTodoBadRequest(t *testing.T) {
	order := "title"
	sorting := "none"
	limit := 1
	offset := 2
	test.GetAllTodosTodoBadRequest(t, context.Background(), service, ctrl, &limit, &offset, &order, &sorting)
}

func Test_GetAllTodosTodoOK(t *testing.T) {
	order := "title"
	sorting := "asc"
	limit := 1
	offset := 2
	test.GetAllTodosTodoOK(t, context.Background(), service, ctrl, &limit, &offset, &order, &sorting)
}

func Test_GetByIDTodoInternalServerError(t *testing.T) {
	test.GetByIDTodoInternalServerError(t, context.Background(), service, ctrl, db.IDInternalError)
}

func Test_GetByIDTodoNotFound(t *testing.T) {
	test.GetByIDTodoNotFound(t, context.Background(), service, ctrl, db.IDNotFound)
}

func Test_GetByIDTodoOK(t *testing.T) {
	test.GetByIDTodoOK(t, context.Background(), service, ctrl, db.IDOK)
}

func Test_UpdateTodoTodoBadRequest(t *testing.T) {
	ctx := context.Background()
	authObj := &auth.Auth{UserID: ID}
	ctx = auth.SetAuth(ctx, authObj)

	test.UpdateTodoTodoBadRequest(t, ctx, service, ctrl, db.IDBadRequest, &payloadUpdateBadrequest)
}

func Test_UpdateTodoTodoServerError(t *testing.T) {
	test.UpdateTodoTodoInternalServerError(t, context.Background(), service, ctrl, db.IDInternalError, &payloadUpdateInternalServerError)
}

func Test_UpdateTodoTodoNotFound(t *testing.T) {
	ctx := context.Background()
	authObj := &auth.Auth{UserID: ID}
	ctx = auth.SetAuth(ctx, authObj)

	test.UpdateTodoTodoNotFound(t, ctx, service, ctrl, db.IDNotFound, &payloadUpdateNotFound)
}

func Test_UpdateTodoTodoOK(t *testing.T) {
	ctx := context.Background()
	authObj := &auth.Auth{UserID: ID}
	ctx = auth.SetAuth(ctx, authObj)

	test.UpdateTodoTodoOK(t, ctx, service, ctrl, db.IDOK, &payloadUpdateOK)
}

func Test_FilterTodosTodoOK(t *testing.T) {
	test.FilterTodosTodoOK(t, context.Background(), service, ctrl, &payloadFilterOK)
}

func Test_FilterTodosTodoBadRequest(t *testing.T) {
	test.FilterTodosTodoBadRequest(t, context.Background(), service, ctrl, &payloadFilterBadRequest)
}

func Test_FilterTodosTodoInternalError(t *testing.T) {
	test.FilterTodosTodoInternalServerError(t, context.Background(), service, ctrl, &payloadFilterInternalError)
}
