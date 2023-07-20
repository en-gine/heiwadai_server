buf: #protoファイルからapiコードを自動生成
	rm -rf ./api/v1/* && \
	buf generate

run: 
	go run .

test-all:
	go test -v ./test/...

install-buf:
	go install github.com/bufbuild/buf/cmd/buf@latest & \
  go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest & \
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest & \
  go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@latest
	export PATH="$(go env GOBIN):${PATH}" && export PATH="$(go env GOPATH)/bin:${PATH}"

migrate-create: # ex: TABLE_NAME=users
ifndef TABLE_NAME
	$(error コマンド引数に TABLE_NAME=テーブル名 をセットしてください。)
endif
	migrate create -ext sql -dir db/migrations -seq create_${TABLE_NAME}_table

init-db:
	@make migrate-up
	@make sqlboiler

migrate-up:
	migrate -path=db/migrations -database postgres://${PSQL_USER}:${PSQL_PASS}@${PSQL_HOST}:${PSQL_PORT}/${PSQL_DBNAME}?sslmode=disable up

migrate-down:
	migrate -path=db/migrations -database postgres://${PSQL_USER}:${PSQL_PASS}@${PSQL_HOST}:${PSQL_PORT}/${PSQL_DBNAME}?sslmode=disable down

migrate-drop:
	migrate -database postgres://${PSQL_USER}:${PSQL_PASS}@${PSQL_HOST}:${PSQL_PORT}/${PSQL_DBNAME}?sslmode=disable drop

install-sqlboiler:
	go install github.com/volatiletech/sqlboiler/v4@latest
	go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest

sqlboiler:#sqlboilerでmodelを自動生成
	sqlboiler psql --wipe

install-wire:
	go install github.com/google/wire/cmd/wire@latest