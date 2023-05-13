.PHONY: all clean gen-country gen-currency fmt test coverage help
.DEFAULT_GOAL := all

clean: ## clean
	find . -name '*_gen.go' -delete
	find . -name '*_gen.yaml' -delete
	rm swagger.yaml || true
	find . -name '*.coverage' -delete
	find . -type d -name 'node_modules' | xargs rm -rf

gen-country: ## generate country
	docker run --rm mikekonan/iso-3166-country-codes-scrapper:latest | docker run --rm -i -v $(PWD):/app node:12.19.0-alpine3.10 sh -c "cd /app/country/.generator && npm install && node app.js"

gen-currency: ## generate currency
	docker run --rm -i -v $(PWD):/app node:12.19.0-alpine3.10 sh -c "cd /app/currency/.generator && npm install && node generator.js"

gen-timezone: ## generate timezone
	docker run --rm -i -v $(PWD):/app node:12.19.0-alpine3.10 sh -c "cd /app/timezone/.generator && npm install && node generator.js"

gen-phone: ## generate currency
	docker run --rm -i -v $(PWD):/app node:12.19.0-alpine3.10 sh -c "cd /app/phone/.generator && npm install && node generator.js"

gen-language: ## generate currency
	docker run --rm -i -v $(PWD):/app node:12.19.0-alpine3.10 sh -c "cd /app/language/.generator && npm install && node generator.js"

concat-yaml: ## concat-yaml
	docker run --rm -i -v $(PWD):/app node:12.19.0-alpine3.10 sh -c "cd /app/.generator && npm install && node app.js"

fmt: ## format source code
	go fmt ./...

lint: ## lint
	golangci-lint run --path-prefix $(PWD)

test: ## run unit tests
	go test -race -cover -coverprofile=coverage.txt -covermode=atomic ./...

coverage: ## show code coverage
	go tool cover -html=coverage.txt

all: clean gen-country gen-currency gen-phone gen-language gen-timezone concat-yaml fmt lint test

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
