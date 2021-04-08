FROM golang:1.16.3-buster AS builder
WORKDIR /src
COPY . .
ARG GOPROXY=https://goproxy.io,direct
RUN go build

FROM debian:buster
COPY --from=builder /src/webls .
RUN mkdir -p /data
ENV WEBLS_PATH /data
CMD ["./webls"]