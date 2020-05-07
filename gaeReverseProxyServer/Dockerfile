FROM golang:1.13 as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

ARG PACKAGE_REPOSITORY=${PACKAGE_REPOSITORY:-github.com/BambooTuna/middleware/gaeReverseProxyServer}
WORKDIR /go/src/${PACKAGE_REPOSITORY}
COPY . .
RUN go get ./... && \
    go build -o app && \
    mv app /app

FROM scratch
# 軽量のalpineには必要ファイルがないため注意
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /app /app
EXPOSE 8080
ENTRYPOINT ["/app"]
