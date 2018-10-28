//go:generate goagen bootstrap -d github.com/Microkubes/microtodo/design

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Microkubes/examples/todo/todo-service/app"
	"github.com/Microkubes/examples/todo/todo-service/store"
	"github.com/Microkubes/microservice-tools/config"
	toolscfg "github.com/Microkubes/microservice-tools/config"
	"github.com/Microkubes/microservice-tools/gateway"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	// Create service
	service := goa.New("microtodo")
	cfg := loadConfig()
	httpClient := &http.Client{}
	registration := gateway.NewKongGateway(cfg.GatewayAdminURL, httpClient, cfg.Service)
	if err := registration.SelfRegister(); err != nil {
		log.Fatal("Self registration failed: ", err)
	}
	defer registration.Unregister()
	// securityChain, _, err := flow.NewSecurityFromConfig(&config.ServiceConfig{
	// 	SecurityConfig: config.SecurityConfig{
	// 		KeysDir: Getenv("TODO_KEYS_DIR", "keys"),
	// 		JWTConfig: &config.JWTConfig{
	// 			Name:        "todo-jwt",
	// 			Description: "TODO JWT Security",
	// 			TokenURL:    Getenv("TODO_JWT_TOKEN_ISSUER_URL", "http://kong:8000/jwt/signin"),
	// 		},
	// 	},
	// })

	// if err != nil {
	// 	log.Fatal("Failed to set up the service security. ", err)
	// }
	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// attach the security chain as Goa middleware
	// service.Use(chain.AsGoaMiddleware(securityChain))

	collection := store.NewSession(os.Getenv("TODO_HOST"), os.Getenv("TODO_USERNAME"), os.Getenv("TODO_PASSWORD"), "tododb")

	// Mount "todo" controller
	c := NewTodoController(service, &store.TodoCollection{collection})
	app.MountTodoController(service, c)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}

// Getenv returns the value of an environment variable or a default value if the variable is unset.
func Getenv(variable, defaultValue string) string {
	value := os.Getenv(variable)
	if value == "" {
		return defaultValue
	}
	return value
}

func loadConfig() *config.ServiceConfig {
	sc := config.ServiceConfig{}
	confFile := "/run/secrets/microservice_todo_config.json"
	if os.Getenv("SERVICE_CONFIG_FILE") != "" {
		confFile = os.Getenv("SERVICE_CONFIG_FILE")
	}
	if err := toolscfg.LoadConfigAs(confFile, &sc); err != nil {
		panic(err)
	}
	return &sc
}
