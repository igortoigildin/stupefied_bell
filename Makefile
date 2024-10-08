PROJECT_DIR = $(shell pwd)
PROJECT_BIN = $(PROJECT_DIR)/bin
$(shell [ -f bin] || mkdir -p $(PROJECT_BIN))
PATH := $(PROJECT_BIN):$(PATH)

GOLANGCI_LINT = $(PROJECT_BIN)/golanci-lint

.PHONY: .install-linter
.install-linter:
	### INSTALL GOLANGCI_LINT ###
	[ -f $(PROJECT_BIN)/golanci-lint ] || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(PROJECT_BIN) v1.59.1

.PHONY: lint
lint: .install-linter
	### RUN GOALNGCI-LINT ###
	$(GOLANGCI_LINT) run ./... --congif=./.golangci.yml

.PHONY: docker-up
docker-up:
	### START DOCKER COMPOSE ###
	docker compose -f ./docker/docker-compose.yml up -d

.PHONY: docker-down
docker-down:
	### STOP DOCKER COMPOSE ###
	docker compose -f ./docker/docker-compose.yml down

.PHONY: docker-restart
docker-restart: docker-down docker-up

.PHONY: docker-logs
docker-logs:
	### VIEW DOCKER COMPOSE LOGS ###
	docker compose logs -f

install-deps:
	GOBIN=$(PROJECT_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.2 
	GOBIN=$(PROJECT_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

.PHONY: make generate-delivery-api
generate-delivery-api:
	mkdir -p pkg/delivery_v1
	protoc --proto_path proto/delivery_v1 \
	--go_out=pkg/delivery_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/delivery_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	proto/delivery_v1/delivery.proto