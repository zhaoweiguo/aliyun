FROM golang:1.13.4 as builder
COPY . /build/
WORKDIR /build
RUN GOPROXY=https://goproxy.cn,direct CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o plugin-drone-dingtalk .

FROM alpine
RUN apk update \
    && apk add --no-cache ca-certificates \
    && rm -rf /var/cache/apk/*
MAINTAINER zhaoweiguo<docker@zhaoweiguo.com>
COPY --from=builder /build/plugin-drone-dingtalk /
ENTRYPOINT ["/plugin-drone-dingtalk"]




