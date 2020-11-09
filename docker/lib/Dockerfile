FROM ubuntu:latest

ARG DEBIAN_FRONTEND=noninteractive

RUN apt-get update -y && apt-get install -y --no-install-recommends \
    build-essential \
    ca-certificates \
    cmake \
    curl \
    git \
    libavcodec-dev \
    libavformat-dev \
    libdc1394-22-dev \
    libgtk2.0-dev \
    libjpeg-dev \
    libpng-dev \
    libswscale-dev \
    libtbb-dev \
    libtbb2 \
    libtiff-dev \
    make \
    pkg-config \
    unzip \
    wget \
    && rm -rf /var/lib/apt/lists/*

COPY ./files/Makefile /root/Makefile

RUN cd /root && make install && rm -f /root/Makefile

LABEL maintainer="mateusz.urbanek.98@gmail.com" \
    type="lib" \
    version="0.0.1-init"