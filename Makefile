include .env

.PHONY: cover
cover:
	go test -v ./... -coverprofile=./cover.html $@ && go tool cover -html=./cover.html

.PHONY: migrations/new
migrations/new:
	@echo 'Creating migration files for ${name}...'
	migrate create -seq -ext=.sql -dir=./migrations ${name}

.PHONY: migrations/up
migrations/up:
	@echo 'Running up migration...'
	migrate -path ./migrations -database=${DB_DSN} up 1

.PHONY: migrations/down
migrations/down:
	@echo 'Running down migration...'
	migrate -path ./migrations -database=${DB_DSN} down 1

.PHONY: migrations/force
migrations/force:
	@echo 'Forcing migration...'
	migrate -path ./migrations -database=${DB_DSN} force ${version}