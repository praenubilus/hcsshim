FROM golang:1.17-bullseye

RUN mkdir /src
RUN cd /src;\
    git clone https://github.com/microsoft/hcsshim.git hcsshim;\
    mkdir -p /src/hcsshim/internal/tools/rootfs
    # git clone --depth 1 --branch v0.8.21 https://github.com/microsoft/hcsshim.git hcsshim

ADD rootfs_proxy.go /src/hcsshim/internal/tools/rootfs
ADD main.go /src/hcsshim/internal/tools/rootfs

WORKDIR /src/hcsshim/internal/tools/rootfs

RUN go mod tidy;\
    go build -o librootfs.so -buildmode=c-shared rootfs_proxy.go main.go

