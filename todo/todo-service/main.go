//go:generate goagen bootstrap -d github.com/Microkubes/microtodo/design

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/JormungandrK/backends"
	"github.com/Microkubes/examples/todo/todo-service/app"
	"github.com/Microkubes/examples/todo/todo-service/config"
	"github.com/Microkubes/examples/todo/todo-service/db"
	toolscfg "github.com/Microkubes/microservice-tools/config"
	"github.com/Microkubes/microservice-tools/gateway"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	// Create service
	service := goa.New("todo-service")

	gatewayAdminURL, configFile := loadGatewaySettings()

	cfg, err := loadConfig(configFile)
	if err != nil {
		log.Fatal("Failed to load config: ", err)
	}

	unregisterService := registerMicroservice(cfg.ToStandardConfig(), gatewayAdminURL)
	defer unregisterService()
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
	todosService, err := setupTodosService(cfg.ToStandardConfig())
	if err != nil {
		log.Fatal("Failed to create a service: ", err)
	}

	// Mount "todo" controller
	c := NewTodosController(service, todosService, cfg)
	app.MountTodoController(service, c)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}

func registerMicroservice(cfg *toolscfg.ServiceConfig, gatewayAdminURL string) func() {
	registration := gateway.NewKongGateway(gatewayAdminURL, &http.Client{}, cfg.Service)

	err := registration.SelfRegister()
	if err != nil {
		panic(err)
	}

	return func() {
		registration.Unregister()
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

func loadConfig(file string) (*config.ServiceConfig, error) {
	cfg := &config.ServiceConfig{}
	err := toolscfg.LoadConfigAs(file, cfg)

	return cfg, err
}

func setupBackend(dbConfig toolscfg.DBConfig) (backends.Backend, backends.BackendManager, error) {
	dbinfoMap := map[string]*toolscfg.DBInfo{}
	dbinfoMap[dbConfig.DBName] = &dbConfig.DBInfo
	backendManager := backends.NewBackendSupport(dbinfoMap)
	backend, err := backendManager.GetBackend(dbConfig.DBName)
	return backend, backendManager, err
}

func setupRepository(backend backends.Backend) (backends.Repository, error) {
	return backend.DefineRepository("todos", backends.RepositoryDefinitionMap{
		"name": "todos",
		"indexes": []backends.Index{
			backends.NewUniqueIndex("id"),
		},
		"hashKey":       "id",
		"readCapacity":  int64(5),
		"writeCapacity": int64(5),
		"GSI": map[string]interface{}{
			"name": map[string]interface{}{
				"readCapacity":  1,
				"writeCapacity": 1,
			},
		},
	})
}

func setupTodosService(serviceConfig *toolscfg.ServiceConfig) (db.TodoStore, error) {
	backend, _, err := setupBackend(serviceConfig.DBConfig)
	if err != nil {
		return nil, err
	}
	todosRepo, err := setupRepository(backend)
	if err != nil {
		return nil, err
	}

	return db.NewTodosService(todosRepo), nil
}

func loadGatewaySettings() (string, string) {
	gatewayURL := os.Getenv("API_GATEWAY_URL")
	serviceConfigFile := os.Getenv("SERVICE_CONFIG_FILE")

	if gatewayURL == "" {
		gatewayURL = "http://localhost:8001"
	}
	if serviceConfigFile == "" {
		serviceConfigFile = "/run/secrets/microservice_todo_config.json"
	}

	return gatewayURL, serviceConfigFile
}
