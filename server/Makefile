PSQL_PASS_ENCODED := $(shell echo -n ${PSQL_PASS} | python3 -c "import urllib.parse; import sys; print(urllib.parse.quote(sys.stdin.read(), safe=''))")

buf: #protoファイルからapiコードを自動生成
	buf generate

run: 
	@make lint
	go run .

dev:
	@make lint
	air

lint:
	golangci-lint run

test-all:
	go test -v ./test/...

echo:
	echo ${PSQL_PASS_ENCODED}

install-devtools:
	go install -v github.com/ramya-rao-a/go-outline@latest & \
	go install -v github.com/rogpeppe/godef@latest & \
	go install -v github.com/stamblerre/gocode@latest & \
	go install -v golang.org/x/tools/cmd/goimports@latest & \
	go install -v golang.org/x/tools/gopls@latest & \
	go install -v github.com/nishanths/exhaustive & \
	go install -v github.com/cosmtrek/air@latest & \
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.54.2

install-buf:
	go install github.com/bufbuild/buf/cmd/buf@latest & \
  go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest & \
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest & \
  go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@latest & \
  go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest & \
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest & \
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest & \
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

swagger-merge:
	npx swagger-merge -v "1.0.0" -t "Heiwadai Api" -d "Heiwadai description" -p "./api/v1/" -o "./api/v1/merged-swagger.json" $(find ./api/v1/ -type f -name "*.json" | tr '\n' ' ')

sqlboiler:#sqlboilerでmodelを自動生成
	sqlboiler psql --wipe

grpcurl:
	grpcurl -plaintext localhost:3000 list

seeder-store:
	cd /server/db/seeder && go run . -StoreSeeder

seeder-user:
	cd /server/db/seeder && go run . -UserSeeder

seeder-admin:
	cd /server/db/seeder && go run . -AdminSeeder

seeder-coupon:
	cd /server/db/seeder && go run . -CouponSeeder

seeder-server:
	cd /server/db/seeder && go run . -OnlyServer

bookTest:
	go run /server/test/bookingTest/.

mailTest:
	go run /server/test/mailTest/.

uploadTest:
	go run /server/test/upload/.

checkinTest:
	go run /server/test/checkinTest/.

couponTest:
	go run /server/test/couponTest/.

postTest:
	go run /server/test/postTest/.

messageTest:
	go run /server/test/messageTest/.

birthCouponTest:
	go run /server/test/birthCouponTest/.

passTest:
	go run /server/test/passTest/.


kanaTest:
	# go run /server/test/kana/.
	go test -run . -v ./test/kana/.
