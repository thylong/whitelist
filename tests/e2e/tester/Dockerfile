FROM golang:alpine

ADD . /go/src/github.com/thylong/tester
WORKDIR /go/src/github.com/thylong/tester

# Compile into executable binary
RUN go install .

CMD ["/go/bin/tester"]