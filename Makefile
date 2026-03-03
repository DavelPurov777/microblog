DB_DRIVER = postgres
DB_STRING = "postgres://postgres:qwerty@localhost:5436/microblog?sslmode=disable"

MIGRATION_DIR = ./db/migrations

migrate-up: 
	goose -dir $(MIGRATION_DIR) $(DB_DRIVER) $(DB_STRING) up

migrate-down:
	goose -dir $(MIGRATION_DIR) $(DB_DRIVER) $(DB_STRING) down