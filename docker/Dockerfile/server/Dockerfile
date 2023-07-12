FROM golang:1.20.5-bullseye
RUN mkdir server
WORKDIR /server

RUN apt-get update && apt-get install -y git
COPY ./server .

### connnect buf https://github.com/bufbuild/connect-go
RUN go mod tidy & \
  go install github.com/bufbuild/buf/cmd/buf@latest & \
  go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest & \
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest & \
  go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@latest

RUN export PATH="$(go env GOBIN):${PATH}" && export PATH="$(go env GOPATH)/bin:${PATH}"

### migrate https://github.com/golang-migrate/migrate
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

### sqlboiler
RUN	go install github.com/volatiletech/sqlboiler/v4@latest & \
  go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest


### For development
RUN  go mod tidy & \
  go install -v github.com/ramya-rao-a/go-outline@latest & \
  go install -v github.com/rogpeppe/godef@latest & \
  go install -v github.com/stamblerre/gocode@latest & \
  go install -v golang.org/x/tools/cmd/goimports@latest & \
  go install -v golang.org/x/tools/gopls@latest
