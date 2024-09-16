include .env
export

.PHONY: clear
clear: ## Clear the working area and the project
	rm -rf bin/

.PHONY: fmt
fmt: ## Format code
	golangci-lint run --fix

.PHONY: run
run: ## Run app
	go mod tidy && go mod download && \
	GIN_MODE=debug CGO_ENABLED=0 go run -tags migrate ./cmd/app

.PHONY: compose-up
compose-up: ### Run docker-compose
	docker-compose --env-file ./.env up --build -d

.PHONY: compose-down
compose-down: ### Down docker-compose
	docker-compose down --remove-orphans

.PHONY: docker-rm-volume
docker-rm-volume: ### remove docker volume
	docker volume rm go-link-shortener_pg-data

.PHONY: migrate-create
migrate-create:  ### create new migration
	migrate create -ext sql -dir migrations $(name)

.PHONY: migrate-up
migrate-up: ### migration up
	migrate -path migrations -database '$(PG_URL)?sslmode=disable' up

.PHONY: help
.DEFAULT_GOAL := help
help:
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'