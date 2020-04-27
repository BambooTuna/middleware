FROM envoyproxy/envoy:v1.14.1

ENV PORT 18080
ENV API_SERVER_HOST 127.0.0.1
ENV API_SERVER_PORT 8080

WORKDIR /etc/envoy

COPY . .
RUN chmod +x template-engine.sh && cat envoy.tpl.yaml | template-engine.sh > envoy.yaml

CMD /usr/local/bin/envoy -c /etc/envoy/envoy.yaml
EXPOSE ${PORT}
