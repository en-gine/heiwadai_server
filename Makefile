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

migrate-create: # ex: TABLE_NAME=users
ifndef TABLE_NAME
	$(error コマンド引数に TABLE_NAME=テーブル名 をセットしてください。)
endif
	docker compose exec server migrate create -ext sql -dir ./migrations/ -seq create_${TABLE_NAME}_table

migrate-up:
	docker compose exec server migrate -path=./migrations/ -database postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable up