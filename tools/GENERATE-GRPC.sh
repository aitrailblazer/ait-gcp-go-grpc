#!/bin/bash
#
# Copyright 2023 AITrailblazer, LLC. All Rights Reserved.

set -e

. tools/PROTOS.sh
clone_common_protos

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

echo "Generating Go gRPC client/server model for ${SERVICE_PROTOS[@]}"
protoc ${AIT_PROTOS[*]}  \
--proto_path='.'  \
--proto_path=$COMMON_PROTOS_PATH   \
--proto_path=$AIT_PROTOS_PATH  \
--go_opt='module=github.com/aitrailblazer/ait-gcp-go-grpc'  \
--go_out='.'

echo "Generating Go gRPC client/server for ${AIT_PROTOS[@]}"
protoc ${AIT_PROTOS[*]}  \
--proto_path='.'  \
--proto_path=$COMMON_PROTOS_PATH  \
--proto_path=$AIT_PROTOS_PATH  \
--go-grpc_opt='module=github.com/aitrailblazer/ait-gcp-go-grpc'  \
--go-grpc_out='.'

