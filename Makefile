GO_BIN?=$(shell pwd)/.bin/go

.PHONY: install-tools
install-tools: ## Install tooling defined in tools.go
	mkdir -p ${GO_BIN}
	@cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % sh -c 'GOBIN=${GO_BIN} go install %'

.PHONY: run
run:
	go run main.go

.PHONY: create-migration
create-migration:
	${GO_BIN}/goose -dir sql create new_migration sql
