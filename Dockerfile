FROM golang:1.14

WORKDIR /go/src/app
COPY . .

RUN go get . && \
    go build . \
    sudo apt update \
    sudo apt install vim -y

EXPOSE 8081
CMD ["./app"]