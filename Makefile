.PHONY: tailwind-watch
tailwind-watch:
	./tailwindcss -i ./static/css/input.css -o ./static/css/style.css --watch

.PHONY: tailwind-build
tailwind-build:
	./tailwindcss -i ./static/css/input.css -o ./static/css/style.min.css --minify

.PHONY: templ-generate
templ-generate:
	templ generate

.PHONY: templ-watch
templ-watch:
	templ generate --watch
	
.PHONY: dev
dev:
	go build -o ./tmp/$(APP_NAME) ./cmd/$(APP_NAME)/main.go && air

.PHONY: build
build:
	make tailwind-build
	make templ-generate
	go build -ldflags "-X main.Environment=production" -o ./bin/$(APP_NAME) ./cmd/$(APP_NAME)/main.go

.PHONY: vet
vet:
	go vet ./...

.PHONY: staticcheck
staticcheck:
	staticcheck ./...

.PHONY: test
test:
	  go test -race -v -timeout 30s ./...

.PHONY: postgres-up
postgres-up:
	docker-compose up -d postgres

.PHONY: migrate-up migrate-down migrate-create

migrate-up:
	migrate -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" -path db/migrations up

migrate-down:
	migrate -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" -path db/migrations down

migrate-create:
	migrate create -ext sql -dir db/migrations -seq $(name)

.PHONY: sqlc-generate
sqlc-generate:
	sqlc generate