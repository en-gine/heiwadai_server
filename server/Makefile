PSQL_PASS_ENCODED := $(shell echo -n ${PSQL_PASS} | python3 -c "import urllib.parse; import sys; print(urllib.parse.quote(sys.stdin.read(), safe=''))")

buf: #protoファイルからapiコードを自動生成
	buf generate

run: 
	@make lint
	go run .

lint:
	golangci-lint run

test-all:
	go test -v ./test/...

install-devtools:
	go install -v github.com/ramya-rao-a/go-outline@latest & \
	go install -v github.com/rogpeppe/godef@latest & \
	go install -v github.com/stamblerre/gocode@latest & \
	go install -v golang.org/x/tools/cmd/goimports@latest & \
	go install -v golang.org/x/tools/gopls@latest & \
	go install -v github.com/nishanths/exhaustive & \
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.54.2

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

install-migrate:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

migrate-up:
	migrate -path=db/migrations -database postgres://${PSQL_USER}:${PSQL_PASS_ENCODED}@${PSQL_HOST}:${PSQL_PORT}/${PSQL_DBNAME}?sslmode=disable up

migrate-down:
	migrate -path=db/migrations -database postgres://${PSQL_USER}:${PSQL_PASS_ENCODED}@${PSQL_HOST}:${PSQL_PORT}/${PSQL_DBNAME}?sslmode=disable down

migrate-drop:
	migrate -database postgres://${PSQL_USER}:${PSQL_PASS_ENCODED}@${PSQL_HOST}:${PSQL_PORT}/${PSQL_DBNAME}?sslmode=disable drop

install-sqlboiler:
	go install github.com/volatiletech/sqlboiler/v4@latest
	go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest

sqlboiler:#sqlboilerでmodelを自動生成
	sqlboiler psql --wipe

seeder-store:
	cd /server/db/seeder && go run . -StoreSeeder

seeder-user:
	cd /server/db/seeder && go run . -UserSeeder

seeder-admin:
	cd /server/db/seeder && go run . -AdminSeeder

seeder-server:
	cd /server/db/seeder && go run . -OnlyServer

book:
	go run /server/bookingTest/.