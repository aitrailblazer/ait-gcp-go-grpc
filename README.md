# ait-gcp-go-grpc

Hello! 

[ait-gcp-go-grpc](https://github.com/aitrailblazer/ait-gcp-go-grpc)

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