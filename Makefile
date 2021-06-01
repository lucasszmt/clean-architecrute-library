#!make
include .env
export $(shell sed 's/=.*//' .env)

build:
	@echo "Building"
	go run main.go

MIGRATIONS_PATH="$(shell pwd)/database/migrations"
DB_DSN="$$DB_TYPE://$$DB_USER:$$DB_PASS@$$DB_HOST:$$DB_PORT/$$DB_NAME?sslmode=disable"

migrations:
	docker run --rm -v $(MIGRATIONS_PATH):/migrations --network "$$DOCKER_NETWORK" migrate/migrate -path=/migrations -database $(DB_DSN) up

clean:
	docker run --rm -it -v $(MIGRATIONS_PATH):/migrations  --network "$$DOCKER_NETWORK" migrate/migrate -path=/migrations -database $(DB_DSN) down

#Creates the migration file already named
create_migration:
	docker run --rm -v $(MIGRATIONS_PATH):/migrations --user $(shell id -u):$(shell id -g) \
					--network "$$DOCKER_NETWORK" migrate/migrate create -ext sql -dir ./migrations -seq $(MIG_NAME)