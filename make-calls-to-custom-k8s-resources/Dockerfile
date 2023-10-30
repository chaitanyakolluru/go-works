FROM golang:1.21 as build-stage

ARG read_repo_token

WORKDIR /fn

# configure token for pulling private repos
RUN git config --global url."https://ci:${read_repo_token}@gitlab.com/heb-engineering".insteadOf "https://gitlab.com/heb-engineering"
# configure private repo paths
RUN go env -w GOPRIVATE="gitlab.com/heb-engineering/*"

COPY go.mod go.sum ./
RUN go mod download

COPY main.go main.go
COPY *.go ./

RUN CGO_ENABLED=0 go build -o makecalls .

FROM alpine:3.16.0

RUN apk update && apk add --no-cache openssl \
    && rm -rf /var/cache/apk/*

COPY --from=build-stage /fn/makecalls /usr/local/bin/

ENTRYPOINT ["sleep", "3600"]
##ENTRYPOINT ["makecalls"]

