# syntax=docker/dockerfile:1
FROM --platform=$TARGETPLATFORM ubuntu:20.04

ENV PATH=/usr/local/go/bin:/go/bin:$PATH

ENV GOPATH=/go
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io,direct

RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    curl \
    git \
    unzip \
    rsync \
    pkg-config \
    git-lfs \
    libncurses5 `#for clang-format` \
    make \
    xz-utils \
    && rm -rf /var/lib/apt/lists/*

RUN git config --global --add safe.directory /go/src/xiaobao.com/practice

WORKDIR /go/src/xiaobao.com/practice
