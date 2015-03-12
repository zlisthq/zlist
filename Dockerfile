# shadowsocks
#
# VERSION 0.0.1

FROM golang
MAINTAINER Whiteworld <ljq258@gmail.com>
RUN go get github.com/WhiteWorld/zlist

WORKDIR /go/src/github.com/WhiteWorld/zlist
ENTRYPOINT /go/bin/zlist

EXPOSE 8080
