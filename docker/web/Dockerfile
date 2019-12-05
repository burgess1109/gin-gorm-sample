ARG GO_VERSION=1.13

FROM golang:${GO_VERSION}-alpine

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

ENV GO111MODULE on

WORKDIR "/go/src/go-sample"

COPY . .
RUN make download

CMD ["make", "run-server"]
