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

build:
	docker build -t asia-northeast2-docker.pkg.dev/heiwadai/server/heiwadai-server -f ./docker/Dockerfile/server/Dockerfile.prod .

push:
	@make build
	docker push asia-northeast2-docker.pkg.dev/heiwadai/server/heiwadai-server