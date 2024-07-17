ps:
	docker-compose ps
build:
	docker-compose up -d --build
up:
	docker-compose up -d
down:
	docker-compose down
restart:
	docker-compose restart
stop:
	docker-compose stop
app:
	docker-compose exec app sh
db:
	docker-compose exec db sh
installCLI:
	docker-compose exec app ./go_install.sh
install:
	docker-compose exec app go mod tidy
vendor:
	docker-compose exec app go mod vendor
prod:
	docker-compose exec app go run src/main.go
dev:
	docker-compose exec app air

include .env
DATABASE_URL := "postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_DATABASE}?sslmode=${DB_SSLMODE}&timezone=${DB_TIMEZONE}"
MIGRATIONS_PATH := src/database/migrations
step = 1

migrate-up:
	docker-compose exec app migrate -source "file://$(MIGRATIONS_PATH)" -database $(DATABASE_URL) up $(step)

migrate-down:
	docker-compose exec app migrate -source "file://$(MIGRATIONS_PATH)" -database $(DATABASE_URL) down $(step)

migrate-new:
	@docker-compose exec app migrate create -ext sql -dir "$(MIGRATIONS_PATH)" $(name)

migrate-force:
	@docker-compose exec app migrate -source "file://$(MIGRATIONS_PATH)" -database $(DATABASE_URL) force $(version)

migrate-status:
	echo "$(DATABASE_URL)"
	@docker-compose exec app migrate -source "file://$(MIGRATIONS_PATH)" -database $(DATABASE_URL) version
