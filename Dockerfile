FROM golang:1.15.6-alpine

LABEL maintainer="Vague Coder <vaguecoder0to.n@gmail.com>"

WORKDIR $GOPATH/src/github.com/VagueCoder/Share-Price-Scraper

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

RUN go build -o ../../../../bin/Share-Price-Scraper *.go

CMD ["./Share-Price-Scraper"]

