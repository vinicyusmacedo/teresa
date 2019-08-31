FROM golang:1.12 AS builder

ENV GO111MODULE=on
WORKDIR /go/src/github.com/luizalabs/teresa
COPY . /go/src/github.com/luizalabs/teresa

RUN make build-server

FROM debian:buster-slim
RUN apt-get update && \
apt-get install ca-certificates -y &&\
rm -rf /var/lib/apt/lists/* &&\
rm -rf /var/cache/apt/archives/*

WORKDIR /app
COPY --from=builder /go/src/github.com/luizalabs/teresa . 

ENTRYPOINT ["./teresa-server"]
CMD ["run"]
EXPOSE 50051
