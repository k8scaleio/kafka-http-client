FROM golang:alpine

RUN apk add --update --no-cache alpine-sdk bash ca-certificates \
      libressl \
      tar \
      git openssh openssl yajl-dev zlib-dev cyrus-sasl-dev openssl-dev build-base coreutils
WORKDIR /root
RUN git clone https://github.com/edenhill/librdkafka.git
WORKDIR /root/librdkafka
RUN /root/librdkafka/configure
RUN make
RUN make install
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
RUN go get -u gopkg.in/confluentinc/confluent-kafka-go.v1/kafka

RUN mkdir -p /go/src
COPY ./src /go/src/
# Set the Current Working Directory inside the container
WORKDIR /go/
ENV GOPATH=$GOPATH:/go/src/
RUN echo $GOPATH
RUN go build -o kafkaclient src/main/client.go

# Command to run the executable
CMD ["./kafkaclient"]
