APP_NAME = server
GOOSE_DBSTRING = "root:root1234@tcp(127.0.0.1:33306)/book-shop"
GOOSE_MIGRATIONS_DIR ?= sql/schema

run:
	go run ./cmd/${APP_NAME}/

wire:
	cd internal/wire && wire

migrate-up:
	@GOOSE_DRIVER=mysql GOOSE_DBSTRING=${GOOSE_DBSTRING} goose -dir ${GOOSE_MIGRATIONS_DIR} up

migrate-down:	
	@GOOSE_DRIVER=mysql GOOSE_DBSTRING=${GOOSE_DBSTRING} goose -dir ${GOOSE_MIGRATIONS_DIR} down

migrate-reset:
	@GOOSE_DRIVER=mysql GOOSE_DBSTRING=${GOOSE_DBSTRING} goose -dir ${GOOSE_MIGRATIONS_DIR} reset

sqlc:
	sqlc generate
