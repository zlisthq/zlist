# zlist
#
# VERSION 1.0

FROM golang
MAINTAINER Whiteworld <ljq258@gmail.com>

## for debug
# ADD . /go/src/github.com/zlisthq/zlist

RUN go get github.com/zlisthq/zlist

WORKDIR /go/src/github.com/zlisthq/zlist
ENTRYPOINT /go/bin/zlist
