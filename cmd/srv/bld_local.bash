#!/usr/bin/env bash


# TODAY=$(date +%Y-%m-%d-%HH-%MM-%SS)

# VER="apttus-pricing-caching-service"
# VER+="_"v0.1
# VER+="_"$TODAY

# git tag -f $VER -m $VER"Release v0.1" 

# git push --tags
# go install \
#     github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
#     github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
#     google.golang.org/protobuf/cmd/protoc-gen-go \
#     google.golang.org/grpc/cmd/protoc-gen-go-grpc
# go install google.golang.org/protobuf/cmd/protoc-gen-go


# flatc -g --gen-object-api --grpc contracts.fbs
# flatc -g --gen-object-api models/digicert/pricing/contracts.fbs

# SRC_DIR=./models/digicert/pricing
# DST_DIR=./models/digicert/pricing

# go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@38aafd89f814f347db56a52efd055961651078ad
# go get google.golang.org/protobuf/cmd/protoc-gen-go 

# SRC_DIR=./protos
# DST_DIR=./protos

# protoc -I $SRC_DIR \
#     --go_out $DST_DIR --go_opt paths=source_relative \
#     --go-grpc_out $DST_DIR --go-grpc_opt=requireUnimplementedServers=false --go-grpc_opt=paths=source_relative \
#     calculator.proto

# echo protoc -I $SRC_DIR \
#     --go_out $DST_DIR --go_opt paths=source_relative \
#     --go-grpc_out $DST_DIR --go-grpc_opt=requireUnimplementedServers=false --go-grpc_opt=paths=source_relative \
#     contracts.proto

# protoc -I $SRC_DIR \
#     --go_out $DST_DIR --go_opt paths=source_relative \
#     --go-grpc_out $DST_DIR --go-grpc_opt=requireUnimplementedServers=false --go-grpc_opt=paths=source_relative \
#     contracts.proto

# cp -R  ../apttus-pricing-service/apttuspricing/ apttuspricing/

# cd models
# cd digicert
# cd pricing
# go mod init models.digicert.com/apttus-pricing-models
# # go mod vendor -v

# go mod tidy
# CGO_ENABLED=0 go build #-v
# cd ..
# cd ..
# cd ..

# cd caching
# go mod init digicert.com/apttus-pricing-caching
# go get digicert.com/apttus-pricing-caching
# go mod tidy
# go mod vendor -v

# CGO_ENABLED=0 go build #-v
# cd ..

 #####################################
# cd gen-test-data
# go mod init digicert.com/apttus-pricing-caching-gen-test-data
# go mod tidy
# CGO_ENABLED=0 go build -v
# cd ..

# cd gen
# go mod init digicert.com/apttus-pricing-caching-gen
# go mod tidy
# CGO_ENABLED=0 go build -v
# cd ..

# cd test-db
# ./bld.bash
# cd ..
 #####################################

cd client
./bld.bash
cd ..

echo -- Generate NewVersion
cd genversion
go mod init digicert.com/apttus-pricing-caching-service/genversion
go mod tidy
 
go generate
CGO_ENABLED=0 go build #-v
echo --- generated NewVersion:genversion.go
cat genversion.go
echo --- generated NewVersion:genversion.go

cd ..
echo --- generated NewVersion:gitupdate.bash
cat gitupdate.bash
echo --- generated NewVersion:gitupdate.bash

go mod init digicert.com/apttus-pricing-caching-server
go mod tidy
# CGO_ENABLED=0 go build -v
go clean -cache
go clean -testcache

CGO_ENABLED=0 go build -a -tags netgo -ldflags '-w' #-v

go mod vendor #-v

# docker build --pull --no-cache -f Dockerfile.local -t apttus-pricing-caching-service .
# docker scan apttus-pricing-caching-service
