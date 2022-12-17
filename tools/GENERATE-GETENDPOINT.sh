#!/bin/bash
#
# Copyright 2023 AITrailblazer, LLC. All Rights Reserved.

set -e

. tools/PROTOS.sh

cd cmd
cd srv

ENDPOINT=$(\
  gcloud run services list \
  --project=${GCP_PROJECT} \
  --platform=managed \
  --format="value(status.address.url)" \
  --filter="metadata.name=${GCP_SVC_NAME}") 
ENDPOINT=${ENDPOINT#https://} && echo ${ENDPOINT}

cd ..
cd ..