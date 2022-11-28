#!/bin/bash
#
# Copyright 2023 AITrailblazer, LLC. All Rights Reserved.

export PROJECT_ID=ait-gcp-go-grpc
# gcloud init
echo "Activate your project:"
# gcloud auth list

# ACTIVE  ACCOUNT
# *       constantine@aitrailblazer.com
#         constantine@topinvestor.app
#         topinvestor-app-dev-srv@topinvestor-app-dev.iam.gserviceaccount.com

echo "Set the active account"
# gcloud config set account constantine@aitrailblazer.com

# gcloud config set project $(gcloud projects list --format='value(PROJECT_ID)')
gcloud config list project
# set -e

# . tools/PROTOS.sh
# clone_common_protos

# go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest

# echo "Generating OpenAPI spec for ${HOSTED_PROTOS[@]}"
# protoc ${HOSTED_PROTOS[*]} --proto_path='.' --proto_path=$COMMON_PROTOS_PATH --openapi_out='.'
