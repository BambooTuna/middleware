#!/bin/bash

# 第一引数にpackage.jsonの実行場所からの相対パスをセット
docker build `dirname $0` -t vue-builder:latest

docker run --rm -v /var/run/docker.sock:/var/run/docker.sock -v "${GOPATH}:/${GOPATH}" -w="/${PWD}" vue-builder:latest \
npm run build --prefix $1
