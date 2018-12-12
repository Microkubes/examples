### Multi-stage build
FROM golang:1.10-alpine3.7 as build

RUN apk add --no-cache curl git openssh

RUN go get -u -v github.com/goadesign/goa/...
RUN go get -u -v github.com/globalsign/mgo
RUN go get -u -v github.com/Microkubes/microservice-tools/...
RUN go get -u -v github.com/Microkubes/backends
RUN go get -u -v github.com/Microkubes/microservice-security/...

COPY . /go/src/github.com/Microkubes/examples/todo/todo-service

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install github.com/Microkubes/examples/todo/todo-service

### Main
FROM scratch

ENV API_GATEWAY_URL="http://localhost:8001"

COPY --from=build /go/bin/todo-service /todo-service

EXPOSE 8080
CMD ["/todo-service"]
