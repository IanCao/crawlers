FROM golang:1.14 AS BUILD

WORKDIR $GOPATH/src/crawlers/crawler-core
ADD . $GOPATH/src/crawlers/crawler-core

RUN go build crawler.go
RUN rm -rf "$GOPATH/src/github.com/"
CMD ["./crawler"]
