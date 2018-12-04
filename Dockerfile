FROM golang:1.11 as golang

ADD . $GOPATH/github.com/klnchu/logstash_pusher/
ENV GO111MODULE=auto
ENV GOPROXY=https://goproxy.io
RUN cd $GOPATH/github.com/klnchu/logstash_pusher && make

FROM busybox:1.27.2-glibc
COPY --from=golang /go/github.com/klnchu/logstash_pusher/logstash_pusher /
LABEL maintainer kollinchu@gmail.com
EXPOSE 9198
ENTRYPOINT ["/logstash_pusher"]  
