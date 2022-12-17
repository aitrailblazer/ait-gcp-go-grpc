#!/bin/bash
#
# Copyright 2023 AITrailblazer, LLC. All Rights Reserved.

set -e

. tools/PROTOS.sh

cd cmd
cd srv

gcloud builds submit --tag gcr.io/${GCP_PROJECT}/${GCP_SVC_NAME}

cd ..
cd ..