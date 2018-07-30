FROM golang:1.10-alpine

WORKDIR /go/src/app
COPY . .

RUN apk add --update git

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]
