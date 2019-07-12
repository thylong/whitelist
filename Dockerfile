FROM golang:alpine AS builder

ADD . /go/src/github.com/thylong/whitelist
WORKDIR /go/src/github.com/thylong/whitelist

# Fetch dependencies
RUN apk add git dep make
RUN dep ensure

# Compile into executable binary
RUN go install .

FROM alpine:latest AS runner
WORKDIR /go/bin

COPY --from=builder /go/bin/whitelist /go/bin/whitelist
EXPOSE 8080 8081
CMD ["/go/bin/whitelist"]