#!/usr/bin/env bash

go mod vendor -v

GCP_PROJECT=smartapi-295619
GCP_SVC_NAME=grpc-cloud-run-service

docker build -f Dockerfile -t $GCP_SVC_NAME .

# gcloud builds submit --tag gcr.io/${GCP_PROJECT}/${GCP_SVC_NAME}
# gcloud alpha run deploy ${GCP_SVC_NAME} --image gcr.io/${GCP_PROJECT}/${GCP_SVC_NAME}   --project ${GCP_PROJECT} --platform managed --allow-unauthenticated  --use-http2 --sandbox minivm

# gcloud builds submit --tag gcr.io/${GCP_PROJECT}/${GCP_SVC_NAME}
# gcloud beta run deploy ${GCP_SVC_NAME} --image gcr.io/${GCP_PROJECT}/${GCP_SVC_NAME}   --project ${GCP_PROJECT} --platform managed --allow-unauthenticated --use-http2

# https://grpc-a.run.app

# # go test
# git add .

# echo $VER
# git commit -a -m $VER
# git push -f -u origin master


