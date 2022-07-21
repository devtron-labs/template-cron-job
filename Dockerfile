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
COPY --from=build-env  /go/src/github.com/devtron-labs/template-cron-job/auth_model.conf .
COPY --from=build-env  /go/src/github.com/devtron-labs/template-cron-job/vendor/github.com/argoproj/argo-cd/assets/ /go/src/github.com/devtron-labs/template-cron-job/vendor/github.com/argoproj/argo-cd/assets
COPY --from=build-env  /go/src/github.com/devtron-labs/template-cron-job/scripts/devtron-reference-helm-charts scripts/devtron-reference-helm-charts
COPY --from=build-env  /go/src/github.com/devtron-labs/template-cron-job/scripts/argo-assets/APPLICATION_TEMPLATE.JSON scripts/argo-assets/APPLICATION_TEMPLATE.JSON

COPY ./git-ask-pass.sh /git-ask-pass.sh
RUN chmod +x /git-ask-pass.sh

CMD ["./template-cron-job"]


#FROM alpine:3.15.0 as  template-cron-job-ea

#RUN apk add --no-cache ca-certificates
#COPY --from=build-env  /go/src/github.com/template-cron-job-labs/template-cron-job/auth_model.conf .
#COPY --from=build-env  /go/src/github.com/template-cron-job-labs/template-cron-job/cmd/external-app/template-cron-job-ea .

#COPY --from=build-env  /go/src/github.com/template-cron-job-labs/template-cron-job/vendor/github.com/argoproj/argo-cd/assets/ /go/src/github.com/template-cron-job-labs/template-cron-job/vendor/github.com/argoproj/argo-cd/assets
#COPY --from=build-env  /go/src/github.com/template-cron-job-labs/template-cron-job/scripts/template-cron-job-reference-helm-charts scripts/template-cron-job-reference-helm-charts
#COPY --from=build-env  /go/src/github.com/template-cron-job-labs/template-cron-job/scripts/argo-assets/APPLICATION_TEMPLATE.JSON scripts/argo-assets/APPLICATION_TEMPLATE.JSON

#CMD ["./template-cron-job-ea"]
