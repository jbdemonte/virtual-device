FROM ubuntu:24.04

ENV GO_VERSION=1.22

RUN apt-get update
RUN apt-get install -y golang

WORKDIR /app

COPY * .