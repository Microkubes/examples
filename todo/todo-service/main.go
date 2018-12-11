//go:generate goagen bootstrap -d github.com/Microkubes/microtodo/design

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Microkubes/backends"
	"github.com/Microkubes/examples/todo/todo-service/app"
	"github.com/Microkubes/examples/todo/todo-service/config"
	"github.com/Microkubes/examples/todo/todo-service/db"
	"github.com/Microkubes/microservice-security/chain"
	"github.com/Microkubes/microservice-security/flow"
	toolscfg "github.com/Microkubes/microservice-tools/config"
	"github.com/Microkubes/microservice-tools/gateway"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	// Create service
	service := goa.New("todo-service")

	// Load the Gateway URL and the config file path
	gatewayAdminURL, configFile := loadGatewaySettings()

	// Load config from file
	cfg, err := loadConfig(configFile)
	if err != nil {
		log.Fatal("Failed to load config: ", err)
	}

	unregisterService := registerMicroservice(cfg.ToStandardConfig(), gatewayAdminURL)

	// The unregistration is deferred for after main() executes. If we shut down
	// the service, it is nice to unregister, although it is not required.
	defer unregisterService()

	// Create security chain for the microservice
	securityChain, securityCleanup, err := flow.NewSecurityFromConfig(cfg.ToStandardConfig())
	if err != nil {
		log.Fatal("Failed to create security chain: ", err)
	}
	defer securityCleanup()

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount security chain as Goa Middleware
	service.Use(chain.AsGoaMiddleware(securityChain))

	// Instantiate TodosService
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

// Register the microservice on Kong Gateway
func registerMicroservice(cfg *toolscfg.ServiceConfig, gatewayAdminURL string) func() {
	// Creates new Kong gateway.Registration with the config settings. We pass the default http client here.
	registration := gateway.NewKongGateway(gatewayAdminURL, &http.Client{}, cfg.Service)

	// At this point we do a self-registration by calling SelfRegister
	err := registration.SelfRegister()
	if err != nil {
		// if there is an error it means we failed to self-register so we panic with error
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

// loadConfig loads a configuration from a file.
func loadConfig(file string) (*config.ServiceConfig, error) {
	cfg := &config.ServiceConfig{}
	err := toolscfg.LoadConfigAs(file, cfg)

	return cfg, err
}

// setupBackend creates a backend and returns it.
func setupBackend(dbConfig toolscfg.DBConfig) (backends.Backend, backends.BackendManager, error) {
	dbinfoMap := map[string]*toolscfg.DBInfo{}
	dbinfoMap[dbConfig.DBName] = &dbConfig.DBInfo

	// Define the supported backend (MongoDB/DynamoDB)
	backendManager := backends.NewBackendSupport(dbinfoMap)

	// Get the desired backend
	backend, err := backendManager.GetBackend(dbConfig.DBName)

	return backend, backendManager, err
}

// setupRepository defines the repository (collection/table) used in this microservice
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

// setupTodosService sets up new todos service.
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

// loadGatewaySettings loads the API Gateway URL and the path to the JSON config file from ENV variables.
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
