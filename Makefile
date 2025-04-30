.PHONY: default
default: run

.PHONY: install
install:
	@go mod download

.PHONY: update
update:
	@go get -u ./... && go mod tidy

.PHONY: run
run:
	@go run cmd/run/main.go

.PHONY: generate
generate:
	@go run cmd/gen/main.go

.PHONY: migrate
generate:
	@go run cmd/migrate/main.go
