FROM golang

ADD . /go/src/github.com/demas/cowl-go
RUN go get github.com/demas/cowl-go
RUN go install github.com/demas/cowl-go

ENTRYPOINT /go/bin/cowl-go --operation=serve-rest-api
EXPOSE 4000
