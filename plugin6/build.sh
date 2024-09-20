#!/usr/bin/env bash

rm -rf vendor
rm -rf output


if [ ! -f go.sum ]; then
    touch go.sum
fi

go mod vendor

# 保存原来文件
cp go.mod go.tmp_mod
cp go.sum go.tmp_sum
cp main.go main.tmp_go


# 替换包名，实际上最好也把当前包名替换了也，防止插件包被循环引用
tmp_prefix="t_$(date +%s)"

replace_dir=()
# shellcheck disable=SC2045
for elem in $(ls vendor) ; do
  if ! [ "$elem" = "modules.txt" ] ; then  replace_dir+=("$elem") ; fi
done

echo "${replace_dir[*]}"

function move_dir() {
    mkdir -p "vendor/$tmp_prefix"
    for elem in "${replace_dir[@]}"; do
       mv vendor/"${elem}" vendor/"${tmp_prefix}"
    done
}

move_dir

function replace_file() {
    for elem in "${replace_dir[@]}"; do
       sed -i "s/${elem}/${tmp_prefix}\/${elem}/g" "$1"
    done
}

# shellcheck disable=SC2044
for file in $(find . -name '*.go'); do
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
