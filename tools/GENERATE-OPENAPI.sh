#!/bin/bash
#
# Copyright 2023 AITrailblazer, LLC. All Rights Reserved.

set -e

. tools/PROTOS.sh
clone_common_protos

go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest

echo "==> Generating OpenAPI spec for ${AIT_PROTOS[@]}"
# protoc ${HOSTED_PROTOS[*]} --proto_path='.' --proto_path=$COMMON_PROTOS_PATH --openapi_out='.'
protoc ${AIT_PROTOS[*]}  \
--proto_path='.'  \
--proto_path=$COMMON_PROTOS_PATH  \
--proto_path=$AIT_PROTOS  \
--openapi_out='docs/.'
