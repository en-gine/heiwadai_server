include .env

up:
	docker compose up -d

bash:
	docker compose exec server bash

stop:
	docker compose stop

restart:
	docker compose restart

gen: #protoファイルからapiコードを自動生成
	docker compose exec server buf generate

run: 
	docker compose exec server go run main.go
