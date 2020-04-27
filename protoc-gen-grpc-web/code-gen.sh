#!/bin/bash

#docker build . -t protoc-gen-grpc-web:latest
alias grpc='docker run --rm -v /var/run/docker.sock:/var/run/docker.sock -v "${GOPATH}:/${GOPATH}" -w="/${PWD}" protoc-gen-grpc-web'

for file in `\find $1 -name '*.proto'`; do
  mkdir -p $2
  docker run --rm -v /var/run/docker.sock:/var/run/docker.sock -v "${GOPATH}:/${GOPATH}" -w="/${PWD}" protoc-gen-grpc-web \
    protoc -I$1 -I${GOPATH}/src/github.com/googleapis/googleapis \
    --js_out=import_style=commonjs:$2 \
    --grpc-web_out=import_style=commonjs+dts,mode=grpcwebtext:$2 \
    $file
done
