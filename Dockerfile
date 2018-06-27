FROM golang:alpine

MAINTAINER Jolly <zhaolei@pm.me>

RUN apk add -U tzdata git
ENV TZ=Asia/Harbin

RUN mkdir -p /go/src/GRE3000/logs
COPY . /go/src/GRE3000
WORKDIR /go/src/GRE3000

RUN go get -v github.com/beego/bee \
        && apk del git \
        && rm -rf /var/cache/apk/*

CMD ["bee", "run"]
