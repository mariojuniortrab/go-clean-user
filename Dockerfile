FROM golang:latest

WORKDIR /go/app

CMD ["tail", "-f", "/dev/null"]
