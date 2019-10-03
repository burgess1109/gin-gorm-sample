ARG GO_VERSION=1.13

FROM golang:${GO_VERSION}-alpine

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

#ARG GO_PATH=/go
#ENV GOPATH ${GO_PATH}
#ENV PATH ${GO_PATH}/bin:$PATH
ENV GO111MODULE on

WORKDIR "/go/src/go-sample"

COPY . .
RUN make download

CMD ["make", "run-server"]
