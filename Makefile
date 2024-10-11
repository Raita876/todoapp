export VERSION ?= $(shell cat ./VERSION)
export PGPASSWORD = postgres
export GIN_MODE = release

TODOAPP_BIN ?= ./dist/todoapp_darwin_arm64/todoappserver

.PHONY: build
build:
	goreleaser release --snapshot --clean

.PHONY: release
release:
	goreleaser release --clean

.PHONY: dockerbuild
dockerbuild:
	docker build \
		-t todoappserver:$(VERSION) \
		--build-arg VERSION=$(VERSION) \
		.

.PHONY: swag
swag:
	swag init -g cmd/todoappserver/main.go

.PHONY: gotests
gotests:
	gotests -all -w \
		./internal/application/services/task_service.go \
		./internal/domain/entities/task.go \
		./internal/infrastructure/db/postgres/task_repository.go \
		./internal/interface/api/rest/task_controller.go

.PHONY: run
run: swag build
	$(TODOAPP_BIN) &

.PHONY: stop
stop:
	kill -9 $(shell pgrep todoappserver)

.PHONY: help
help: swag build
	./dist/todoapp_linux_amd64_v1/todoappserver --help

.PHONY: test
test:
	go test -v -cover -coverprofile=index.out ./internal/...

.PHONY: cover
cover: test
	go tool cover -html=index.out -o index.html
	python3 -m http.server 8765

.PHONY: golangci
golangci:
	golangci-lint run -v ./...

.PHONY: scenariotest
scenariotest:
	runn run --verbose ./test/runn/runn.yaml

.PHONY: up
up:
	docker compose -f ./deployments/compose.yaml up -d

.PHONY: down
down:
	docker compose -f ./deployments/compose.yaml down 

.PHONY: psql
psql:
	psql -h localhost -p 5432 -U postgres
