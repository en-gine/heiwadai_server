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
	docker compose exec server go run .

dev:
	docker compose exec server air

buf: 
	docker compose exec server buf generate

reload-env:
	docker-compose --env-file .env up -d

# build:
# 	docker build -t 228980314714.dkr.ecr.us-east-1.amazonaws.com/heiwadai-server:latest -f ./docker/Dockerfile/server/Dockerfile.prod .

# push:
# 	docker push 228980314714.dkr.ecr.us-east-1.amazonaws.com/heiwadai-server:latest

# login:
# 	aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin https://228980314714.dkr.ecr.us-east-1.amazonaws.com


build:
	docker build -t asia-northeast2-docker.pkg.dev/heiwadai/server/heiwadai-server -f ./docker/Dockerfile/server/Dockerfile.prod .

push:
	@make build
	docker push asia-northeast2-docker.pkg.dev/heiwadai/server/heiwadai-server

# deploy:
# 	gcloud beta run deploy heiwadai-server --project heiwadai --region asia-northeast2 --platform managed --source .

railway:
	railway link ${RAILWAY_PROJECT_ID}
	railway run -s heiwadai ls

init-proto:
	git submodule add git@github.com:en-gine/heiwadai_proto.git server/v1
	
update-proto:
	git submodule update --remote