FROM golang:alpine as builder

WORKDIR /go/src/raptor/server
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o server .

FROM alpine:latest

LABEL MAINTAINER=""

WORKDIR /go/src/raptor/server

COPY --from=0 /go/src/raptor/server/server ./
COPY --from=0 /go/src/raptor/server/resource ./resource/
COPY --from=0 /go/src/raptor/server/config.yaml ./

EXPOSE 8888
ENTRYPOINT ./server -c config.yaml
