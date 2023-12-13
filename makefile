SERVICE_BINARY=go-crud-food

## up: starts containers in the background without forcing build
up:
    @echo Starting Docker images...
    docker-compose up -d
    @echo Docker images started!

## up_build: stops docker-compose (if running), builds projects and starts docker compose
up_build: build_broker
    @echo Stopping docker images (if running...)
    docker-compose down
    @echo Building (when required) and starting docker images...
    docker-compose up --build -d
    @echo Docker images built and started!

## down: stop docker compose
down:
    @echo Stopping docker compose...
    docker-compose down
    @echo Done!

## build_service: builds the service binary as a linux executable
build_service:
    @echo Building broker binary...
    set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0 && go build -o ${SERVICE_BINARY}
    @echo Done!
