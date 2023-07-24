include .env

up:
	docker compose up -d

bash:
	docker compose exec server bash

stop:
	docker compose stop

restart:
	docker compose restart

run: 
	docker compose exec server go run main.go

reload-env:
	docker-compose --env-file .env up -d