#!/bin/bash

set -e

. tools/PROTOS.sh
clone_common_protos

go install golang.org/x/vuln/cmd/govulncheck@latest


# module code.ssnc.dev/cloud/apis
# go mod init code.ssnc.dev/cloud/cloud-registry

# This directory contains the generated srv
GENERATED='cmd/srv'
# go env -w GOPRIVATE=code.ssnc.dev/\*

echo "Generating Go srv"

echo "Generating Go models"

cd api
cd v1
cd models
rm -rf go.mod
rm -rf go.sum

go mod init github.com/aitrailblazer/ait-gcp-go-grpc/api/v1/models
go mod tidy 

golangci-lint run ./...
govulncheck ./...
golangci-lint run ./...
go build -v *.go

cd ..
cd ..
cd ..

echo "Generating Go server.gen"

cd api
cd v1
cd api
rm -rf go.mod
rm -rf go.sum
go mod init github.com/aitrailblazer/ait-gcp-go-grpc/api/v1/api

go mod tidy

golangci-lint run ./...
govulncheck ./...
golangci-lint run ./...
go build -v *.go
cd ..
cd ..
cd ..
echo "Generating Go srv"

cd cmd
cd srv

rm -rf go.mod
rm -rf go.sum
go mod init github.com/aitrailblazer/ait-gcp-go-grpc/cmd/srv

go mod tidy

golangci-lint run ./...
govulncheck ./...
golangci-lint run ./...
rm -rf server
go build -v server.go

cd ..
cd ..