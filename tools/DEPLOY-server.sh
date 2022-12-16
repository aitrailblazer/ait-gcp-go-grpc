#!/bin/bash
#
# Copyright 2023 AITrailblazer, LLC. All Rights Reserved.

export PROJECT_ID=ait-gcp-go-grpc
VER=0.1
NAME=ait-gcp-go-grpc

cd cmd
cd api
cd srv

gcloud builds submit \
  --tag us-docker.pkg.dev/aitrailblazer-registry/gcr.io/$PROJECT_ID/$NAME:$VER

gcloud run deploy $NAME \
  --image us-docker.pkg.dev/aitrailblazer-registry/gcr.io/$PROJECT_ID/$NAME:$VER \
  --platform managed \
  --region us-central1 \
  --allow-unauthenticated \
  --max-instances=2

cd ..
cd ..
cd ..