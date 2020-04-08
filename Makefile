.PHONY: up down stats migration migrate-up migrate-down migrate-status

include .env
export

check-envs:
ifndef LOCAL_DB_PORT
	@echo Warning: LOCAL_DB_PORT isn\'t defined\; continue? [Y/n]
	@read line; if [ $$line == "n" ]; then echo aborting; exit 1 ; fi
endif
ifndef LOCAL_REST_PORT
	@echo Warning: LOCAL_REST_PORT isn\'t defined\; continue? [Y/n]
	@read line; if [ $$line == "n" ]; then echo aborting; exit 1 ; fi
endif

postgres:
	docker-compose up postgres

rest:
	docker-compose up rest

up: check-envs
	docker-compose up

down:
	docker-compose down

stats:
	docker stats $$(docker ps --filter network=go_payments --format="{{.Names}}")

fixtures:
	# example: make fixtures name=insert_streams_fixtures
	docker-compose run sql-migration goose -dir ./database/fixtures create ${name} sql

fixtures-up:
	# example: make migrate-up
	docker-compose run sql-migration goose -dir ./database/fixtures postgres "host=${DB_HOST} user=${DB_USER} password=${DB_PASS} dbname=${DB_NAME} sslmode=disable port=${DB_PORT}" up

migration:
	# example: make migration name=crate_streams_table
	docker-compose run sql-migration goose -dir ./database/migrations create ${name} sql

migrate-up:
	# example: make migrate-up
	docker-compose run sql-migration goose -dir ./database/migrations postgres "host=${DB_HOST} user=${DB_USER} password=${DB_PASS} dbname=${DB_NAME} sslmode=disable port=${DB_PORT}" up

migrate-down:
	# example: make migrate-down
	docker-compose run sql-migration goose -dir ./database/migrations postgres "host=${DB_HOST} user=${DB_USER} password=${DB_PASS} dbname=${DB_NAME} sslmode=disable port=${DB_PORT}" down

migrate-status:
	# example: make migrate-status
	docker-compose run sql-migration goose -dir ./database/migrations postgres "host=${DB_HOST} user=${DB_USER} password=${DB_PASS} dbname=${DB_NAME} sslmode=disable port=${DB_PORT}" status
