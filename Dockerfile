FROM golang:1.9

COPY . /go/src/bitbucket.org/wthiti/coolbrobkk-qor
WORKDIR /go/src/bitbucket.org/wthiti/coolbrobkk-qor
RUN go get ./...
EXPOSE 9000

CMD go run main.go
