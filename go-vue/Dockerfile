FROM golang:1.13-alpine as builder

ENV PORT 9090

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /go/src/github.com/BambooTuna/middleware/go-vue
COPY . .
RUN go get ./... && \
go build -o app

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /go/src/github.com/BambooTuna/middleware/go-vue/app /app
EXPOSE ${PORT}
ENTRYPOINT ["/app"]
