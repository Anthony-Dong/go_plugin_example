#!/usr/bin/env bash

rm -rf vendor
rm -rf output


go mod vendor

# 保存原来文件
cp go.mod go.tmp_mod
cp go.sum go.tmp_sum
cp main.go main.tmp_go


# 替换包名，实际上最好也把当前包名替换了也，防止插件包被循环引用
tmp_prefix="t_1702492315"
mkdir -p "vendor/$tmp_prefix"
mv vendor/github.com vendor/"$tmp_prefix"

function replace_file() {
    sed -i "s/github.com/${tmp_prefix}\/github.com/g" "$1"
}

for file in `find . -name '*.go'`; do
   replace_file "$file"
done
replace_file vendor/modules.txt
replace_file go.mod
replace_file go.sum

# ls -al

CGO_ENABLED=1 go build -mod=vendor -ldflags "-s -w" -o output/plugin.so -buildmode=plugin

mv go.tmp_mod go.mod
mv go.tmp_sum go.sum
mv main.tmp_go main.go
rm -rf vendor
