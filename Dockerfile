FROM ubuntu:24.04

ENV GO_VERSION=1.21

RUN apt-get update && apt-get install -y ca-certificates
RUN apt-get install -y golang
RUN apt-get install -y evtest udev

WORKDIR /virtual_device

COPY . /virtual_device

COPY docker-entrypoint.sh /docker-entrypoint.sh
RUN chmod +x /docker-entrypoint.sh

ENTRYPOINT ["/docker-entrypoint.sh"]
CMD ["go", "test", "-tags=integration", "-race", "-count=1", "-timeout", "60s", "./..."]
