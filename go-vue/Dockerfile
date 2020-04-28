FROM golang:1.13 as builder

ARG PORT=${PORT:-9090}

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

ARG GO_FILE_PATH=${GO_FILE_PATH:-.}
ARG JS_FILE_PATH=${JS_FILE_PATH:-.}

WORKDIR /go/src/github.com/BambooTuna/middleware/go-vue
COPY ${GO_FILE_PATH} .
RUN go get ./... && \
    go build -o app


FROM node:12.12.0-alpine as vueBuilder
WORKDIR /src
COPY ${JS_FILE_PATH} .
RUN npm install && \
    npm run build

FROM scratch
# 軽量のalpineには必要ファイルがないため注意
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /go/src/github.com/BambooTuna/middleware/go-vue/app /app
COPY --from=vueBuilder /src/dist /dist
EXPOSE ${PORT}
ENTRYPOINT ["/app"]
