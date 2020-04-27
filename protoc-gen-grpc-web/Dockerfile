FROM node:12-buster

ARG APP_HOME=${APP_HOME:-/app}
RUN mkdir -p $APP_HOME

ARG GRPC_HOME=${GRPC_HOME:-/grpc}
RUN mkdir -p $GRPC_HOME
WORKDIR $GRPC_HOME

USER root

RUN npm i rimraf -g
RUN curl -L -O https://github.com/protocolbuffers/protobuf/releases/download/v3.11.2/protoc-3.11.2-linux-x86_64.zip
RUN curl -L -O https://github.com/grpc/grpc-web/releases/download/1.0.7/protoc-gen-grpc-web-1.0.7-linux-x86_64
RUN unzip protoc-3.11.2-linux-x86_64.zip && cp ./bin/protoc /usr/local/bin/. && chmod +x /usr/local/bin/protoc
RUN cp protoc-gen-grpc-web-1.0.7-linux-x86_64 /usr/local/bin/protoc-gen-grpc-web && chmod +x /usr/local/bin/protoc-gen-grpc-web

WORKDIR $APP_HOME
