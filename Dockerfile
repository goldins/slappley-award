FROM golang:1.9

RUN go get github.com/tools/godep
RUN go get github.com/kelseyhightower/envconfig

ADD . /go/src/github.com/goldins/slappley-award

RUN cd /go/src/github.com/goldins/slappley-award && RUN godep go install

CMD ["/go/bin/slappley"]

EXPOSE 8000
