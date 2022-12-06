#!/bin/bash
#

go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

echo oapi-codegen --config=api/v1/models.cfg.yaml docs/openapi.yaml
echo output: api/v1/models/models.gen.go
mkdir -p api/v1/models
oapi-codegen --config=api/v1/models.cfg.yaml docs/openapi.yaml

echo oapi-codegen --config=api/server.cfg.yaml docs/openapi.yaml
echo output: api/v1/server.gen.go

oapi-codegen --config=api/v1/server.cfg.yaml docs/openapi.yaml
