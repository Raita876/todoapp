export PGPASSWORD = postgres

.PHONY: build
build:
	goreleaser release --snapshot --clean

.PHONY: release
release:
	goreleaser release --clean

.PHONY: run
run:
	./dist/todoapp_linux_amd64_v1/todoapp

.PHONY: up
up:
	docker compose -f ./deployments/docker-cmpose.yaml up -d

.PHONY: down
down:
	docker compose -f ./deployments/docker-cmpose.yaml down 

.PHONY: psql
psql:
	psql -h localhost -p 5432 -U postgres
