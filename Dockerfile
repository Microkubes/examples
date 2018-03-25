### Multi-stage build
FROM golang:1.10-alpine3.7 as build

RUN apk add --no-cache curl git openssh

RUN go get -u -v github.com/goadesign/goa/...
RUN go get -u -v github.com/globalsign/mgo

COPY . /go/src/github.com/Microkubes/microtodo

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install github.com/Microkubes/microtodo

### Main
FROM scratch

ENV API_GATEWAY_URL="http://localhost:8001"

COPY --from=build /go/bin/microtodo /microtodo

EXPOSE 8080
CMD ["/microtodo"]
