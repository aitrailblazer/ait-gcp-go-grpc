#!/usr/bin/env bash


 #####################################

# cd client
# ./bld.bash
# cd ..

# echo -- Generate NewVersion
# cd genversion
# go mod init ...
# go mod tidy
 
# go generate
# CGO_ENABLED=0 go build #-v
# echo --- generated NewVersion:genversion.go
# cat genversion.go
# echo --- generated NewVersion:genversion.go

# cd ..
# echo --- generated NewVersion:gitupdate.bash
# cat gitupdate.bash
# echo --- generated NewVersion:gitupdate.bash

# go mod init ..
# go mod tidy
# CGO_ENABLED=0 go build -v
go clean -cache
go clean -testcache

CGO_ENABLED=0 go build -a -tags netgo -ldflags '-w' #-v

go mod vendor #-v

# docker build --pull --no-cache -f Dockerfile.local -t caching-service .
# docker scan caching-service
