FROM debian:buster AS build

WORKDIR /work

RUN apt-get update && apt-get install -y \
    build-essential \
    freeglut3-dev \
    libfontconfig-dev \
    libfreetype6-dev \
    libgif-dev \
    libgl1-mesa-dev \
    libglu1-mesa-dev \
    libharfbuzz-dev \
    libicu-dev \
    libjpeg-dev \
    libpng-dev \
    libwebp-dev \
    git \
    python3 \
    cmake \
    ca-certificates \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# build libemoji
RUN git clone https://github.com/oyakodon/libemoji.git libemoji -b develop --recursive

RUN cd libemoji && \
    cmake . && \
    make

FROM golang:1.19.4-buster AS base

ARG AIR_VERSION=1.40.4

WORKDIR /work

RUN apt-get update && apt-get install -y \
    zlib1g-dev \
    libfontconfig-dev \
    libgl1-mesa-dev \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

COPY --from=build /work/libemoji/lib/ /work/third_party/libemoji/lib/

RUN go install github.com/cosmtrek/air@v${AIR_VERSION}

CMD ["air", "-c", ".air.toml"]
