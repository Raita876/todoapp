export VERSION ?= $(shell cat ./VERSION)
export PGPASSWORD = postgres


.PHONY: build
build:
	goreleaser release --snapshot --clean

.PHONY: release
release:
	goreleaser release --clean

.PHONY: run
run: build
	./dist/todoapp_linux_amd64_v1/todoappserver

.PHONY: help
help: build
	./dist/todoapp_linux_amd64_v1/todoappserver --help

.PHONY: up
up:
	docker compose -f ./deployments/compose.yaml up -d

.PHONY: down
down:
	docker compose -f ./deployments/compose.yaml down 

.PHONY: psql
psql:
	psql -h localhost -p 5432 -U postgres
