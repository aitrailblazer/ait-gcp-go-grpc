#!/bin/bash
#
# Copyright 2023 AITrailblazer, LLC. All Rights Reserved.

set -e

. tools/PROTOS.sh
clone_common_protos

GCP_SVC_NAME=ait-gcp-go-grpc

docker build -f Dockerfile -t $GCP_SVC_NAME .
