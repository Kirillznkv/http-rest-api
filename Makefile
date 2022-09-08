.PHONY: build
build: db_up
	go build -v ./cmd/apiserver

.PHONY: test
test:
	go test -v -rece -timeout 30s ./...

.PHONY: db_up
db_up:
	docker run --rm --name my_postgres \
	-e POSTGRES_PASSWORD=qwerty1234 \
	-e POSTGRES_USER=kshanti \
	-e POSTGRES_DB=mydb \
	-p 5432:5432 -d postgres

.PHONY: db_down
db_down:
	docker stop my_postgres

.DEFAULT_GOAL := build