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