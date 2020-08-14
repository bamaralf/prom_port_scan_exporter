FROM golang:1.14

WORKDIR /go/src/app
COPY . .

RUN go get . && \
    go build .

EXPOSE 8081
CMD ["./app"]