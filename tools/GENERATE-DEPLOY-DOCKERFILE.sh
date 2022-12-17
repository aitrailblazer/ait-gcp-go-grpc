#!/bin/bash
#
# Copyright 2023 AITrailblazer, LLC. All Rights Reserved.

set -e

. tools/PROTOS.sh

cd cmd
cd srv

# --use-http2  
echo gcloud run deploy  ${GCP_SVC_NAME} --image gcr.io/${GCP_PROJECT}/${GCP_SVC_NAME}:latest   --project ${GCP_PROJECT} --platform managed --allow-unauthenticated 
gcloud run deploy  ${GCP_SVC_NAME} --image gcr.io/${GCP_PROJECT}/${GCP_SVC_NAME}:latest   --project ${GCP_PROJECT} --platform managed --allow-unauthenticated  

cd ..
cd ..