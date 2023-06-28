FROM golang:1.19-alpine AS golang

# build 
RUN mkdir -p /build/simpleJsonApp
COPY main.go /build/simpleJsonApp/main.go
COPY internal/ /build/simpleJsonApp/internal/
COPY docs/ /build/simpleJsonApp/docs/
COPY go.* /build/simpleJsonApp/

WORKDIR /build/simpleJsonApp
RUN go build -o simpleJsonApp .

FROM alpine:3.16.0

RUN apk update && apk add --no-cache openssl \
    && rm -rf /var/cache/apk/*

COPY --from=golang /build/simpleJsonApp/simpleJsonApp /usr/local/bin/

ENTRYPOINT [ "/bin/sh", "-c", "simpleJsonApp" ]
