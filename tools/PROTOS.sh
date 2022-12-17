#!/bin/bash
#
# Copyright 2023 AITrailblazer, LLC. All Rights Reserved.
PROJECT_ID=ait-gcp-go-grpc
GCP_PROJECT=smartapi-3696
GCP_SVC_NAME=$GCP_PROJECT"-svc"

AIT_PROTOS=(
	api/v1/service.proto
)

COMMON_PROTOS_PATH='third_party/api-common-protos'
COMMON_PROTOS_GEN_VALIDATE_PATH='third_party/protoc-gen-validate'

GOOGLE_APIS_PATH='google'

# https://github.com/googleapis/googleapis.git
function clone_googleapis_protos {
	if [ ! -d $GOOGLE_APIS_PATH ]
	then
		git clone https://github.com/googleapis/googleapis $GOOGLE_APIS_PATH
	fi
}
function clone_common_protos {
	if [ ! -d $COMMON_PROTOS_PATH ]
	then
		git clone https://github.com/googleapis/api-common-protos $COMMON_PROTOS_PATH
	fi
}
function clone_validate_protos {
	if [ ! -d $COMMON_PROTOS_GEN_VALIDATE_PATH ]
	then
		git clone https://github.com/bufbuild/protoc-gen-validate.git $COMMON_PROTOS_GEN_VALIDATE_PATH
	fi
}


# Require a specific version of protoc for generating files.
# This stabilizes the generated file output, which includes the protoc version.
. tools/PROTOC-VERSION.sh
if [ "$(protoc --version)" != "libprotoc 3.$PROTOC_VERSION" ]; then
    echo "Please update your protoc to version 3.$PROTOC_VERSION, the current required version as specified in tools/PROTOC_VERSION.sh"
    exit
fi
