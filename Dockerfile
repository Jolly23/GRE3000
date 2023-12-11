FROM golang:alpine

MAINTAINER Jolly <zhaolei@pm.me>

RUN apk add -U tzdata git
ENV TZ=Asia/Harbin
ENV GO111MODULE=auto

RUN mkdir -p /go/src/GRE3000/logs
COPY . /go/src/GRE3000
WORKDIR /go/src/GRE3000

RUN go install github.com/beego/bee/v2@latest \
        && apk del git \
        && rm -rf /var/cache/apk/*

CMD ["bee", "run", "-vendor=true"]
