FROM golang:1.16 as build

WORKDIR /go/src/github.com/analogj/terraflow
COPY . /go/src/github.com/analogj/terraflow

RUN go mod vendor && \
    go build -ldflags '-w -extldflags "-static"' -o terraflow cmd/terraflow/terraflow.go

########
FROM ubuntu:bionic as runtime

COPY --from=build /go/src/github.com/analogj/terraflow/terraflow /usr/bin/
RUN chmod +x /usr/bin/terraflow && \
    terraflow --help
