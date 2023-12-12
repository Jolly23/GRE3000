FROM golang:alpine

MAINTAINER Jolly <zhaolei@pm.me>

RUN apk add -U tzdata git
ENV TZ=Asia/Harbin

RUN mkdir -p /go/src/GRE3000/logs
COPY . /go/src/GRE3000
WORKDIR /go/src/GRE3000

RUN go build

CMD ["./GRE3000"]
