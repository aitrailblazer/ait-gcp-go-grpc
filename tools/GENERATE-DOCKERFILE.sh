#!/bin/bash
#
# Copyright 2023 AITrailblazer, LLC. All Rights Reserved.

set -e

. tools/PROTOS.sh
# clone_common_protos

cd cmd
cd srv

go clean -cache
go clean -testcache
go mod vendor #-v

docker build --pull --no-cache -f Dockerfile -t $GCP_SVC_NAME .
# docker scan $GCP_SVC_NAME

cd ..
cd ..


# gcloud builds submit --tag gcr.io/${GCP_PROJECT}/${GCP_SVC_NAME}
# gcloud alpha run deploy ${GCP_SVC_NAME} --image gcr.io/${GCP_PROJECT}/${GCP_SVC_NAME}   --project ${GCP_PROJECT} --platform managed --allow-unauthenticated  --use-http2 --sandbox minivm

# gcloud builds submit --tag gcr.io/${GCP_PROJECT}/${GCP_SVC_NAME}
# gcloud beta run deploy ${GCP_SVC_NAME} --image gcr.io/${GCP_PROJECT}/${GCP_SVC_NAME}   --project ${GCP_PROJECT} --platform managed --allow-unauthenticated --use-http2

# https://grpc-a.run.app




#!/usr/bin/env bash

# docker run -d \
#     --name caching-service \
#     caching-service

# docker run -d \
#     -p 8081:8080 \
#     --name caching-service \
#     caching-service

# cd client
# go mod init digicert.com/caching-client
# go mod download google.golang.org/grpc
# go mod tidy
# CGO_ENABLED=0 go build -v
# cd ..
############################
# local testing
############################
# Server supports reflection
# grpcurl -plaintext localhost:8081 list
# --Search
# --grpc.reflection.v1alpha.ServerReflection

# grpcurl -plaintext localhost:8081 list Search
# --Search.SayHello
# --Search.SearchAccountId

# grpcurl -plaintext \
#     -d '{"name": "hello"}' \
#     localhost:8081 \
#     Search.SayHello

# client/caching-client \
#      -grpcRemoteServer=0

# cd client
# go test -bench=. -benchtime 100x
# cd ..

# grpcui -plaintext \
#   localhost:8081


# echo We can programmatically determine the gRPC service\'s endpoint:

# ENDPOINT=$(\
#   gcloud run services list \
#   --project=${GCP_PROJECT} \
#   --platform=managed \
#   --format="value(status.address.url)" \
#   --filter="metadata.name=${GCP_SVC_NAME}") 
# ENDPOINT=${ENDPOINT#https://} && echo ${ENDPOINT}
# 

# # Server supports reflection
# echo grpcurl ${ENDPOINT}:443 \
#         list
# grpcurl ${ENDPOINT}:443 \
#         list

# echo grpcurl ${ENDPOINT}:443 \
#         list Search
# grpcurl ${ENDPOINT}:443 \
#         list Search


# grpcurl \
#     -d '{"name": "hello"}' \
#     ${ENDPOINT}:443 \
#     Search.SayHello

# echo client/caching-client \
#     -grpcRemoteServer=1

# client/caching-client \
#     -grpcRemoteServer=1

# cd client
# go test -bench=. -benchtime 100x
# cd ..

# grpcui \
#   ${ENDPOINT}:443 