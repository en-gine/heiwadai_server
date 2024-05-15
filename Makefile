include .env.makefile # Makefile用環境変数を読み込む

up:
	docker compose up -d

bash:
	docker compose exec server bash

stop:
	docker compose stop

restart:
	docker compose restart

run: 
	docker compose exec server go run .

dev:
	docker compose exec server air

buf: 
	docker compose exec server buf generate

reload-env:
	docker-compose --env-file .env up -d

build:
	docker build -t ${AWS_ID}.dkr.ecr.ap-northeast-1.amazonaws.com/heiwadai-server:latest -f ./docker/Dockerfile/server/Dockerfile.prod .

push:
	@make build
	docker push ${AWS_ID}.dkr.ecr.ap-northeast-1.amazonaws.com/heiwadai-server:latest

login:
	aws ecr get-login-password --region ap-northeast-1 | docker login --username AWS --password-stdin https://${AWS_ID}.dkr.ecr.ap-northeast-1.amazonaws.com


# build:
# 	docker build -t asia-northeast2-docker.pkg.dev/heiwadai/server/heiwadai-server -f ./docker/Dockerfile/server/Dockerfile.prod .

# push:
# 	@make build
# 	docker push asia-northeast2-docker.pkg.dev/heiwadai/server/heiwadai-server

# deploy:
# 	gcloud beta run deploy heiwadai-server --project heiwadai --region asia-northeast2 --platform managed --source .

init-proto:
	git submodule add git@github.com:en-gine/heiwadai_proto.git server/v1
	
update-proto:
	git submodule update --remote