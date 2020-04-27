# envoyproxy
Vue <-> gRPC間のリバースプロキシ
Json <-> Protocol Buffersの変換をやってくれます

## 環境変数
- PORT
デフォルト: 9090


## docker-composeでの使用例
gRPCを使用したサーバーのホストとポート(8080)を環境変数にセット。
localhost:18080へアクセスするとうまいことサーバーに繋いでくれます。

※gRPCサーバーがローカルにある場合はホストには`host.docker.internal`を使用してください。`localhost`では接続できません。

```yaml
# docker-compose.yaml
version: "3"
services:
  envoy:
    restart: always
    build: middleware/envoyproxy
    ports:
      - 18080:18080
    environment:
      - PORT=18080
      - API_SERVER_HOST=apiServer
      - API_SERVER_PORT=8080
  apiServer:
      restart: always
      build: ~~~
      ~~~
```

## 参考
[公式のチュートリアル](https://github.com/grpc/grpc-web/tree/master/net/grpc/gateway/examples/helloworld#generate-protobuf-messages-and-)
