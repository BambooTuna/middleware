FROM envoyproxy/envoy:v1.14.1

ARG PORT=${PORT:-18080}
ARG API_SERVER_HOST=${API_SERVER_HOST:-host.docker.internal}
ARG API_SERVER_PORT=${API_SERVER_PORT:-8080}

WORKDIR /etc/envoy

COPY envoy.tpl.yaml envoy.tpl.yaml
RUN sed \
-e "s/<%PORT%>/${PORT}/" \
-e "s/<%API_SERVER_HOST%>/${API_SERVER_HOST}/" \
-e "s/<%API_SERVER_PORT%>/${API_SERVER_PORT}/" \
envoy.tpl.yaml > envoy.yaml

CMD /usr/local/bin/envoy -c /etc/envoy/envoy.yaml
EXPOSE ${PORT}
