FROM golang:1.23.2-bullseye
RUN mkdir server
WORKDIR /server

RUN apt-get update && apt-get install -y git
COPY ./server .

## brew 
RUN apt-get install -y build-essential curl file git && \
  mkdir -p /home/linuxbrew/.linuxbrew && \
  curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh | bash && \
  echo 'eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)"' >> /root/.profile && \
  eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)" && \
  brew install gcc

### connnect buf https://github.com/bufbuild/connect-go
RUN eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)" && \
  go mod tidy && \
  go install github.com/bufbuild/buf/cmd/buf@latest && \
  go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest && \
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
  go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@latest && \
  go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

### migrate https://github.com/golang-migrate/migrate
RUN eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)" && \
  go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

### sqlboiler
RUN eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)" && \
  go install github.com/volatiletech/sqlboiler/v4@latest && \
  go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest

### For development
RUN eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)" && \
  go mod tidy && \
  go install -v github.com/ramya-rao-a/go-outline@latest && \
  go install -v github.com/rogpeppe/godef@latest && \
  go install -v github.com/stamblerre/gocode@latest && \
  go install -v golang.org/x/tools/cmd/goimports@latest && \
  go install -v golang.org/x/tools/gopls@latest && \
  go install -v github.com/air-verse/air@latest && \
  go get github.com/rs/cors

### For Lint
RUN eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)" && \
  go get github.com/nishanths/exhaustive && \
  curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.54.2