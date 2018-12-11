### Todo microservice

This directory contains the backend code for building a microservice using Goa. It contains documentation on how to setup and run the microservice, as well as explanation of the steps on how you can create your own microservice and deploy it into the Microkubes platform.

### Installation

First make sure that the Microkubes platform is up and running. Follow the instructions [here](https://github.com/Microkubes/microkubes) on how to setup Microkubes.

To install the microservice, first you have to build an image with Docker:

```
docker build -t microkubes/microservice-todo .
```

Then, in the [microkubes](https://github.com/Microkubes/microkubes) repo, in the       [docker-compose.fullstack.yml](https://github.com/Microkubes/microkubes/blob/master/docker/docker-compose.fullstack.yml) file,
add the `microservice-todo` service:

```
microservice-todo:
    image: microkubes/microservice-todo:latest
    environment:
      - API_GATEWAY_URL=http://kong:8001
      - MONGO_URL=mongo:27017
    deploy:
      restart_policy:
        condition: on-failure
    secrets:
      - public.pub
      - service.key
      - service.cert
      - system
      - default
      - system.pub
      - default.pub
      - microservice_todo_config.json
```

Also, in the same file, at the bottom where the `secrets` section is, add the file path for the config:

```
  microservice_todo_config.json:
    file: ./config/microservice_todo_config.json
```

Next, copy the contents of the config file located in this directory, and create a file called `microservice_todo_config.json` in the [config](https://github.com/Microkubes/microkubes/tree/master/docker/config) directory.

Finally redeploy the Microkubes stack and check if the service is running using `docker service ls`.

### API documentation

API documentation is available [here](./api.md).

### Libraries and tools

The microservice-todo is built using Go, and it uses the following libraries and tools:
- [Goa](https://github.com/goadesign/goa)
- [microservice-tools](https://github.com/Microkubes/microservice-tools)
- [microservice-security](https://github.com/Microkubes/microservice-security)
- [backends](https://github.com/Microkubes/backends)

#### Goa

Goa is a framework for building microservices and REST APIs in Go using a unique design-first approach. The API is defined using a DSL language in `design/design.go`. Once done, Goa provides a tool called `goagen` to bootstrap and generate code based on the definitions inside `design/design.go`.

After the `goagen` tool is run, it will autogenerate the directories `app`, `client`, `swagger` and `tool`.

You can learn more about Goa and how to define the API [here](https://goa.design/).

#### microservice-tools

`microservice-tools` is a microservice which provides shared tools used by the microservices in the Microkubes platform. In the case of the todo microservice, it uses the tools for loading configuration from a file and registering the microservice on the Kong API gateway.

Loading the config is done using the `loadConfig` function located in `main.go`. The configuration is needed to define things like metadata information for the service, database and security. Look at the `config.json` file as an example.

Registering the service on the Kong API gateway is done using the `registerMicroservice` function located in `main.go`. Kong is started automatically when the Microkubes platform is deployed.

#### microservice-security

`microservice-security` is a microservice which contains functions that are commonly used by all microservices in Microkubes for setting up the security. Also, it exposes functions to set up different security mechanisms for securing the microservices such as JWT, SAML and OAuth2.

In the context of the microservice todo, first thing to do is to create a security chain using the `NewSecurityFromConfig` function which is part of the `flow` package. When called, this function returns a security chain that should be mounted as a middleware to the Goa service. The code for this is in `main.go` inside the `main` function.

#### backends

`backends` is a package that supports multiple backends (databases). As of this writing, it supports MongoDB and DynamoDB. The purpose of this package is to act as an adapter so that there is only one unified API to use when manipulating different types of databases.

In the context of the microservice todo, first thing to do is to setup a backend. This is done inside the `setupBackend` function in `main.go`. Next is to setup a repository, which is basically the collection/table in the database. Usually there should be only one collection per database, but in special cases more collections can be added. Setting up a repository is done using the `setupRepository` function in `main.go`.

Once the database has been setup, a `TodosService` can be instatiated using the `setupTodosService` function in `main.go`. Inside `db/store.go`, there is an interface `TodoStore` that defines the API for accessing and managing todos in the database. All functions defined in the interface are implemented in the same file.

### Controller

After bootstraping the microservice using `goagen`, it will autogenerate a file, in this case `todo.go`, which acts as a controller for the API. For all actions defined in `design/design.go` it will create empty functions in `todo.go` ready to be implemented. For instance, for the `addTodo` action in `design/design.go`, it will create a function `AddTodo`. The functions defined in `db/store.go` should be called respectively inside all functions in `todo.go`.

### Tests

To run the tests for the backend type `go test -v`.