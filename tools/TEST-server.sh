#!/bin/bash
#
# Copyright 2023 AITrailblazer, LLC. All Rights Reserved.

export PROJECT_ID=ait-gcp-go-grpc

VER=1.0
	cd api
	cd server

	go clean -testcache
	go test -v  ./...

	cd ..
	cd ..

	cd api
	cd server

	# go clean -testcache
    # go test -v
	go test -v  -fuzz=. -fuzztime=12s
# /...

	cd ..
	cd ..