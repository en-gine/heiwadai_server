FROM golang:1.24-bullseye
RUN mkdir server
WORKDIR /server

RUN apt-get update && apt-get install -y git
COPY ./server .

## Install build tools
RUN apt-get install -y build-essential curl file git gcc

### connect buf https://github.com/bufbuild/connect-go
RUN go mod tidy && \
  go install github.com/bufbuild/buf/cmd/buf@latest && \
  go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest && \
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
  go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@latest && \
  go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

### migrate https://github.com/golang-migrate/migrate
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

### sqlboiler
RUN go install github.com/aarondl/sqlboiler/v4@latest && \
  go install github.com/aarondl/sqlboiler/v4/drivers/sqlboiler-psql@latest

### For development
RUN go mod tidy && \
  go install -v github.com/ramya-rao-a/go-outline@latest && \
  go install -v github.com/rogpeppe/godef@latest && \
  go install -v github.com/stamblerre/gocode@latest && \
  go install -v golang.org/x/tools/cmd/goimports@latest && \
  go install -v golang.org/x/tools/gopls@v0.14.2 && \
  go install -v github.com/cosmtrek/air@v1.44.0 && \
  go get github.com/rs/cors

### For Lint
RUN go get github.com/nishanths/exhaustive && \
  curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.54.2