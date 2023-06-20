gen: #protoファイルからapiコードを自動生成
	rm -rf ./api/v1/* && \
	buf generate

run: 
	go run main.go

test-all:
	go test -v ./test/...

install-buf:
	go install github.com/bufbuild/buf/cmd/buf@latest & \
  go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest & \
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest & \
  go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@latest
	export PATH="$(go env GOBIN):${PATH}" && export PATH="$(go env GOPATH)/bin:${PATH}"