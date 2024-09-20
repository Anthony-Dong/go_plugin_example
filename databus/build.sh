#!/usr/bin/env bash

#  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

protoc -I . --plugin=protoc-gen-go="$(go env GOPATH)/bin/protoc-gen-go" --go_out=. data.proto