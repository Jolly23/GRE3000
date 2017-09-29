FROM golang

MAINTAINER Jolly <zhaolei@protonmail.com>


RUN echo "Asia/Shanghai" > /etc/timezone
RUN dpkg-reconfigure -f noninteractive tzdata

RUN mkdir -p /go/src/GRE3000/logs
COPY . /go/src/GRE3000
WORKDIR /go/src/GRE3000

RUN go get -v github.com/beego/bee

CMD ["bee", "run"]


# TODO
# 不在生产环境编译执行文件，直接本地编译好上传运行了
# CMD ["./GRE3000"]
