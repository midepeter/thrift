FROM golang AS builder
ENV CGO_ENABLED 0
ENV GOOS=linux
ARG BUILD_REF

WORKDIR /home/midepeter/Desktop/Go

RUN mkdir /thrift
COPY go.* /thrift/
WORKDIR /thrift
RUN go mod download

COPY . /thrift
RUN go build -o thrift -ldflags "-X main.build=${BUILD_REF}"


FROM ubuntu:latest
RUN apt-get update && apt-get install -y
ENV LANG en_US.utf8
WORKDIR /root/
COPY --from=builder ./thrift ./

CMD [ "./thrift" ]

