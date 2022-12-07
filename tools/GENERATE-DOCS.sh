#!/bin/bash
#
set -e

. tools/PROTOS.sh
clone_common_protos

go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@latest

echo "Generating documentation for ${AIT_PROTOS[@]}"
mkdir -p ./docs/
protoc ${AIT_PROTOS[*]}  \
--proto_path='.'  \
--proto_path=$COMMON_PROTOS_PATH  \
--doc_opt='html,grpcapi.html'  \
--doc_out='./docs'

echo "Generating documentation for ${AIT_PROTOS[@]}"
mkdir -p ./docs/
protoc ${AIT_PROTOS[*]}  \
--proto_path='.'  \
--proto_path=$COMMON_PROTOS_PATH  \
--doc_opt='markdown,grpcapi.md'  \
--doc_out='./docs'
