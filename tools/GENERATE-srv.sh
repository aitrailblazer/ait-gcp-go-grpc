#!/bin/bash

PROJECT_ID=ait-gcp-go-grpc
# go mod init github.com/aitrailblazer/ait-gcp-go-grpc/rpc/v1
# go mod init github.com/aitrailblazer/ait-gcp-go-grpc/cmd/cli_grpc

# set -e

# . tools/PROTOS.sh
# clone_common_protos

# go install golang.org/x/vuln/cmd/govulncheck@latest

# go clean --cache

echo "==> Compile Go rpc:"${PROJECT_ID}

cd rpc
cd v1
cd pb

rm -rf go.mod
rm -rf go.sum
go mod init github.com/aitrailblazer/${PROJECT_ID}/rpc/v1/pb
go mod tidy
# govulncheck ./...
golangci-lint run ./...

# go build -v *.go

cd ..
cd ..
cd ..
CGO_ENABLED=0 go build -v rpc/v1/pb/*.go 

echo "==> Compile Go srv_grpc:"${PROJECT_ID}
cd cmd
cd srv_grpc

rm -rf go.mod
rm -rf go.sum
go mod init github.com/aitrailblazer/${PROJECT_ID}/cmd/srv_grpc

# go mod tidy 

# govulncheck ./...
golangci-lint run ./...
go mod vendor #-v

cd ..
cd ..
CGO_ENABLED=0 go build -v cmd/srv_grpc/*.go 
# # go run cmd/srv_grpc/main.go 
# # go run cmd/srv_grpc/main.go 
# # 2022/12/13 11:19:45 server listening at [::]:50051
# # 2022/12/13 11:20:04 AitrailblazerServiceSend VERSION 1.02 
# # 2022/12/13 11:20:04 Received message: world

# echo "==> Compile Go cli_grpc:"${PROJECT_ID}
# cd cmd
# cd cli_grpc

# rm -rf go.mod
# rm -rf go.sum
# go mod init github.com/aitrailblazer/${PROJECT_ID}/cmd/cli_grpc

# # go mod tidy 

# # govulncheck ./...
# golangci-lint run ./...
# # CGO_ENABLED=0 go build -v *.go

# cd ..
# cd ..
# CGO_ENABLED=0 go build -v cmd/cli_grpc/*.go 
# # go run cmd/cli_grpc/main.go 
# # go run cmd/cli_grpc/main.go 
# # 2022/12/13 11:20:04 Pong: index:369  message:"world"  ver:"1.02"  received_on:{seconds:1670959204  nanos:163510000}


# echo "==> Compile Go srv:"${PROJECT_ID}

# echo "==> Compile Go OpenAPI models"

# cd api
# cd v1
# cd models

# rm -rf go.mod
# rm -rf go.sum

# go mod init github.com/aitrailblazer/${PROJECT_ID}/api/v1/models
# go mod tidy 

# # govulncheck ./...
# golangci-lint run ./...

# cd ..
# cd ..
# cd ..
# CGO_ENABLED=0 go build -v api/v1/models/*.go

# echo "==> Compile Go OpenAPI"

# cd api
# cd v1
# cd api

# rm -rf go.mod
# rm -rf go.sum
# go mod init github.com/aitrailblazer/${PROJECT_ID}/api/v1/api

# go mod tidy

# # govulncheck ./...
# golangci-lint run ./...

# cd ..
# cd ..
# cd ..

# CGO_ENABLED=0 go build -v api/v1/api/*.go

# echo "==> Generating Go OpenAPI srv"

# cd cmd
# cd srv

# rm -rf go.mod
# rm -rf go.sum
# go mod init github.com/aitrailblazer/${PROJECT_ID}/cmd/srv

# go mod tidy

# # govulncheck ./...
# golangci-lint run ./...
# # rm -rf main
# go mod vendor #-v

# cd ..
# cd ..

# CGO_ENABLED=0 CGO_ENABLED=0 go build -v cmd/srv/srv.go
