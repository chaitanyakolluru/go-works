FROM docker.nexus.heb.tools/golang:1.19-alpine AS golang

# build 
RUN mkdir -p /build/deploy-restarter
COPY cmd/ /build/deploy-restarter/cmd/
COPY pkg/ /build/deploy-restarter/pkg/
COPY util/ /build/deploy-restarter/util/
COPY go.* /build/deploy-restarter/

WORKDIR /build/deploy-restarter
RUN go build -o deploy-restarter ./cmd  

FROM docker.nexus.heb.tools/alpine:3.16.0

RUN apk update && apk add --no-cache openssl \
    && rm -rf /var/cache/apk/*

COPY --from=golang /build/deploy-restarter/deploy-restarter /usr/local/bin/

ENTRYPOINT [ "/bin/sh", "-c" ]
