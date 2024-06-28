#!/usr/bin/env bash


rm -rf output

CGO_ENABLED=1 go build -ldflags "-s -w" -o output/plugin.so -buildmode=plugin