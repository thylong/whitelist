FROM golang:alpine

WORKDIR /go/src/app
COPY ./main.go /go/src/app/main.go

RUN go build

EXPOSE 80
CMD ["./app"]