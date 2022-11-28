#!/bin/bash
#
# Copyright 2023 AITrailblazer, LLC. All Rights Reserved.

export PROJECT_ID=ait-gcp-go-grpc

go install golang.org/x/vuln/cmd/govulncheck@latest

cd rpc

govulncheck ./...
golangci-lint run ./...

go mod init github.com/aitrailblazer/${PROJECT_ID}/rpc
go mod tidy

go build -v *.go

cd ..

cd api
cd server

govulncheck ./...
golangci-lint run ./...
# go mod init github.com/aitrailblazer/${PROJECT_ID}/server

go mod tidy
go build -o server
 
cd ..
cd ..
