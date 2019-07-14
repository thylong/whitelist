FROM golang:alpine

RUN apk add curl

ADD . /go/src/github.com/thylong/tester
WORKDIR /go/src/github.com/thylong/tester

# Compile into executable binary
RUN go install .

CMD ["/go/bin/tester"]