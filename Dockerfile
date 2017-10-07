FROM golang

MAINTAINER Jolly <zhaolei@protonmail.com>


RUN echo "Asia/Harbin" > /etc/timezone
RUN dpkg-reconfigure -f noninteractive tzdata

RUN mkdir -p /go/src/GRE3000/logs
COPY . /go/src/GRE3000
WORKDIR /go/src/GRE3000

RUN go get -v github.com/beego/bee

CMD ["bee", "run"]
