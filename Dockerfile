FROM golang:1.9

ADD . /go/src/github.com/goldins/slappley-award

RUN go install github.com/goldins/slappley-award

CMD ["/go/bin/slappley-award"]
