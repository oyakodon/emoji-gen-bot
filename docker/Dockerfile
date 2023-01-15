FROM debian:buster AS build

ARG TARGETPLATFORM

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

# for arm64 build
RUN if [ "$TARGETPLATFORM" = "linux/arm64" ]; then \
    echo "building for ARM64..." && \
    apt-get update && apt-get install -y ninja-build && apt-get clean && rm -rf /var/lib/apt/lists/* && \
    cd libemoji && \
    sed -i -E 's/target_cpu="x64"/target_cpu="arm64"/g' CMakeLists.txt && \
    sed -i -E 's/externals\/depot_tools\/ninja/\/usr\/bin\/ninja/g' CMakeLists.txt ;\
    fi

RUN cd libemoji && \
    cmake . && \
    make

FROM golang:1.19.5-alpine3.17 AS bin

WORKDIR /work

RUN apk update && \
    apk add --no-cache \
    libstdc++ \
    zlib-dev \
    fontconfig freetype-dev \
    mesa-dev \
    gcc \
    libc-dev

RUN ln -s /usr/lib/libfontconfig.so.1 /usr/lib/libfontconfig.so && \
    ln -s /usr/lib/libstdc++.so.6 /usr/lib/libstdc++.so

COPY --from=build /work/libemoji/lib/ /work/third_party/libemoji/lib/

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN mkdir -p /artifacts/ && \
    go build -o /artifacts/emoji-gen-bot main.go

FROM alpine:3.17

STOPSIGNAL SIGINT

WORKDIR /work

RUN apk update && \
    apk add --no-cache \
    libstdc++ \
    zlib-dev \
    fontconfig freetype-dev \
    mesa-dev \
    libc-dev

WORKDIR /work

COPY --from=bin /artifacts/emoji-gen-bot /work/

COPY ./config/config.yml /work/config/
COPY ./assets /work/assets

CMD [ "/work/emoji-gen-bot" ]
