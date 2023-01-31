# ait-gcp-go-grpc

# How to Build and Deploy Production-Grade Web Services  with gRPC/OpenAPI with Go, API Gateway, Load Balancing and CloudRun Vol II.


Source code for the book.

[How to Build and Deploy Production-Grade Web Services  with gRPC/OpenAPI: Go, HTTP,Transcoding,API Gateway,Load Balancing, CloudRun Vol II.](https://www.amazon.com/dp/B0BPXZ4C97)

```sh
.
├── LICENSE
├── Makefile
├── README.md
├── api
│   ├── server
│   │   ├── Dockerfile
│   │   ├── go.mod
│   │   ├── go.sum
│   │   ├── main.go
│   │   ├── main_test.go
│   │   └── server
│   └── v1
│       └── service.proto
├── docs
│   └── openapi.yaml
├── rpc
│   ├── service.pb.go
│   └── service_grpc.pb.go
├── tools
│   ├── DEPLOY-server.sh
│   ├── GENERATE-GRPC.sh
│   ├── GENERATE-OPENAPI.sh
│   ├── GENERATE-server.sh
│   ├── PROJECT-activate.sh
│   ├── PROTOC-VERSION.sh
│   ├── PROTOS.sh
│   └── TEST-server.sh
```

Usage: 

```sh
    make protos 
```
Generating Go gRPC client/server:

GENERATE-GRPC.sh
```sh
├── rpc
│   ├── service.pb.go
│   └── service_grpc.pb.go
```

generate grpc YAML file 

├── docs
│   └── openapi.yaml

from .proto
```sh
│   └── v1
│       └── service.proto
```