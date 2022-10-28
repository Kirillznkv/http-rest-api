MIGRATIONS_DIR = db/migrations
DB_URL = "host=localhost port=5432 user=kshanti password=qwerty1234 dbname=mydb sslmode=disable"

all: db_up db_for_test_up migrate_up build

.PHONY: build
build:
	go build -v ./cmd/apiserver

.PHONY: test
test:
	go test -v -race ./...

.PHONY: db_up
db_up:
	docker run --rm --name my_postgres \
	-e POSTGRES_PASSWORD=qwerty1234 \
	-e POSTGRES_USER=kshanti \
	-e POSTGRES_DB=mydb \
	-p 5432:5432 -d postgres
	@sleep 3

.PHONY: db_down
db_down:
	docker stop my_postgres

.PHONY: migrate_up
migrate_up:
	goose -dir $(MIGRATIONS_DIR) postgres $(DB_URL) up

.PHONY: migrate_down
migrate_down:
	goose -dir $(MIGRATIONS_DIR) postgres $(DB_URL) down

#.DEFAULT_GOAL := build