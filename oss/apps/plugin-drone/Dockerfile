FROM golang:1.13.4 as builder
COPY ../.. /build/
WORKDIR /build
RUN ls
RUN ls apps/plugin-drone
RUN GOPROXY=https://goproxy.cn,direct CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o plugin-drone-oss ./apps/plugin-drone

FROM alpine
RUN apk update \
    && apk add --no-cache ca-certificates \
    && rm -rf /var/cache/apk/*
MAINTAINER zhaoweiguo<docker@zhaoweiguo.com>
COPY --from=builder /build/plugin-drone-oss /
ENTRYPOINT ["/plugin-drone-oss"]




