FROM golang

ADD . /go/src/github.com/sniperkit/quzx-crawler

RUN go get github.com/sniperkit/quzx-crawler

RUN go install github.com/sniperkit/quzx-crawler

ENTRYPOINT ["/go/bin/cowl-go"]

CMD ["--operation=serve-rest-api"]

EXPOSE 4000
