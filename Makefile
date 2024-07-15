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
install:
	docker-compose exec app go mod tidy
vendor:
	docker-compose exec app go mod vendor
prod:
	docker-compose exec app go run src/main.go
dev:
	docker-compose exec app air
