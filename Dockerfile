FROM ubuntu:24.04

ENV GO_VERSION=1.22

RUN apt-get update && apt-get install -y ca-certificates
RUN apt-get install -y golang
RUN apt-get install -y evtest

WORKDIR /virtual_device

COPY . /virtual_device