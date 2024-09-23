
.PHONY: build
build:
	goreleaser release --snapshot --clean

.PHONY: release
release:
	goreleaser release --clean

.PHONY: run
run:
	./dist/todoapp_linux_amd64_v1/todoapp
