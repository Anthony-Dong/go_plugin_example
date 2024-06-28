#!/usr/bin/env bash


rm -rf output

CGO_ENABLED=1 go build -ldflags "-s -w" -o output/libplugin.so -buildmode=c-archive # 输出 libplugin.xx 方便链接
mv output/libplugin.h output/plugin.h