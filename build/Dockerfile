FROM golang:1.19.3-alpine3.17 as builder

RUN apk add --no-cache git curl openssl


RUN go install google.golang.org/protobuf/cmd/protoc-gen-go
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
RUN go get golang.org/x/tools/cmd/goimports

RUN wget https://github.com/grpc-ecosystem/grpc-gateway/releases/download/v1.16.0/protoc-gen-grpc-gateway-v1.16.0-linux-x86_64 \
    && chmod -x protoc-gen-grpc-gateway-v1.16.0-linux-x86_64 && mv protoc-gen-grpc-gateway-v1.16.0-linux-x86_64 /go/bin/protoc-gen-grpc-gateway

# the protoc can't run on alpine,
# we only need the protobuf's stdarnd library in the `/protoc/include`.
RUN mkdir -p /protoc && cd /protoc \
	&& wget https://github.com/google/protobuf/releases/download/v3.5.0/protoc-3.5.0-linux-x86_64.zip \
	&& unzip protoc-3.5.0-linux-x86_64.zip

FROM golang:1.19.3-alpine3.17

RUN apk add --no-cache git protobuf make curl openssl jq rsync build-base

COPY --from=builder /protoc/include /usr/local/include
COPY --from=builder /go/bin /go/bin
