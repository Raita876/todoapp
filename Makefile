export VERSION ?= $(shell cat ./VERSION)
export PGPASSWORD = postgres


.PHONY: build
build:
	goreleaser release --snapshot --clean

.PHONY: release
release:
	goreleaser release --clean

.PHONY: swag
swag:
	swag init -g cmd/todoappserver/main.go

.PHONY: run
run: swag build
	./dist/todoapp_linux_amd64_v1/todoappserver

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

.PHONY: up
up:
	docker compose -f ./deployments/compose.yaml up -d

.PHONY: down
down:
	docker compose -f ./deployments/compose.yaml down 

.PHONY: psql
psql:
	psql -h localhost -p 5432 -U postgres
