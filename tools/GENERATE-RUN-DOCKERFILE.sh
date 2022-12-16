#!/bin/bash
#
# Copyright 2023 AITrailblazer, LLC. All Rights Reserved.

set -e

. tools/PROTOS.sh
# clone_common_protos
docker run -d \
    -p 8081:8080 \
    --name $GCP_SVC_NAME \
    $GCP_SVC_NAME
