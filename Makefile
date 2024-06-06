include .env
url="postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"
migration_dir="./migration"
up:
	GOOSE_MIGRATION_DIR=${migration_dir} goose postgres ${url} up

down:
	GOOSE_MIGRATION_DIR=${migration_dir} goose postgres ${url} down 

generate:
	make up
	sqlc generate
