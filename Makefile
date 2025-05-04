include .env

ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))

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

.PHONY: create_migration
create_migration:
	@atlas migrate diff $(ARGS) --env gorm --dir file://sql/migrations

.PHONY: migrate
migrate:
	@atlas migrate apply --env gorm --dir file://sql/migrations --url $(DATABASE_URL)

%::
	@true
