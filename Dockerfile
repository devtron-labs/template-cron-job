FROM golang:1.16.10-alpine3.13  AS build-env

RUN echo $GOPATH

RUN apk add --no-cache git gcc musl-dev
RUN apk add --update make
RUN go get github.com/google/wire/cmd/wire
WORKDIR /go/src/github.com/devtron-labs/template-cron-job
ADD . /go/src/github.com/devtron-labs/template-cron-job/
RUN GOOS=linux make build-all

# uncomment this post build arg
FROM alpine:3.15.0 as  devtron-all
RUN apk add --no-cache ca-certificates
RUN apk update
RUN apk add git
COPY --from=build-env  /go/src/github.com/devtron-labs/template-cron-job/template-cron-job .

COPY ./git-ask-pass.sh /git-ask-pass.sh
RUN chmod +x /git-ask-pass.sh

CMD ["./template-cron-job"]
