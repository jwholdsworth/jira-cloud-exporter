FROM golang:1.13-alpine
LABEL maintainer "James W Holdsworth"
LABEL source "https://github.com/jwholdsworth/jira-cloud-exporter"

WORKDIR /go/src/app
COPY . .

RUN apk add --no-cache --update --verbose git

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]
