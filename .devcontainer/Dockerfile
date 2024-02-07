FROM ubuntu:22.04

RUN apt update && apt upgrade -y

RUN apt install -y curl

ARG GOLANG_VERSION=1.22.0
RUN curl -L https://go.dev/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz | tar -C /usr/local -zx

ENV PATH ${PATH}:/usr/local/go/bin
ENV PATH ${PATH}:/root/go/bin
