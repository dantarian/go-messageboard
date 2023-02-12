GO_BIN?=$(shell pwd)/.bin/go
POSTGRES_DB?=${POSTGRES_DB}

export PATH := ${GO_BIN}:${PATH}

db_type?=postgres
db_host?=database
db_name?=${POSTGRES_DB}
db_user?=postgres
db_password?=${POSTGRES_PASSWORD}

generate: export PSQL_DBNAME=${db_name}
generate: export PSQL_HOST=${db_host}
generate: export PSQL_PORT=5432
generate: export PSQL_USER=${db_user}
generate: export PSQL_PASS=${db_password}
generate: export PSQL_SSLMODE=disable

.PHONY: install-tools
install-tools: ## Install tooling defined in tools.go
	mkdir -p ${GO_BIN}
	@cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % sh -c 'GOBIN=${GO_BIN} go install %'

.PHONY: generate
generate:
	go generate ./...

.PHONY: run
run:
	go run main.go

.PHONY: create-migration
create-migration:
	${GO_BIN}/goose -dir sql create new_migration sql

.PHONY: migrate-up
migrate-up: ## Apply migrations to database
	${GO_BIN}/goose -dir sql ${db_type} "host=${db_host} user=${db_user} password=${db_password} dbname=${db_name}" up
