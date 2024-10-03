FROM golang:1.23-alpine AS base

ARG VERSION=0.0.0
WORKDIR /build

COPY . .

RUN CGO_ENABLE=0 go build \
    -ldflags "-s -w -X main.version=${VERSION} -X main.name=todoappserver" \
    -o /release/todoappserver \
    ./cmd/todoappserver

FROM alpine:3.20

COPY --from=base /release/todoappserver /usr/bin/

ENTRYPOINT ["/usr/bin/todoappserver"]
