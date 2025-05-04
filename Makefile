include .env

CMD := $(firstword $(MAKECMDGOALS))
ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
migrations_dir := sql/migrations

define validate_args
  if [ -z "$(ARGS)" ]; then \
	  echo "✗ missing args, usage: make $(CMD) <ARGS>"; \
	  exit 1; \
	fi;
endef


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
	@$(validate_args)
	@atlas migrate diff $(ARGS) --env gorm --dir file://$(migrations_dir)

.PHONY: create_blank_migration
create_blank_migration:
	@$(validate_args)
	@atlas migrate new $(ARGS) --env gorm --dir file://$(migrations_dir)

.PHONY: migrate
migrate:
	@atlas migrate apply --env gorm --dir file://$(migrations_dir) --url $(DATABASE_URL)

.PHONY: migrate_down
migrate_down:
	@$(validate_args)
	@atlas migrate down --env gorm --dir file://$(migrations_dir) --url $(DATABASE_URL) --to-version $(ARGS)

%::
	@true
