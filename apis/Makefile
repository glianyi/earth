PROTO_FILES=$(sort $(wildcard ./*.proto))
#PROTOC_INC_PATH=/usr/local/include
#GOOGLEAPIS_PATH=./third_party/googleapis

#PROTOC_FLAGS:=-I. -I$(GOOGLEAPIS_PATH)  -I$(PROTOC_INC_PATH)
PROTOC_FLAGS:=-I.

.PHONY: generate
generate:
	mkdir -p ../pkg/pb
	rm -rf ../pkg/pb/*
	protoc $(PROTOC_FLAGS) --go_out=../pkg/pb ${PROTO_FILES}
	protoc $(PROTOC_FLAGS) --go-grpc_out=../pkg/pb ${PROTO_FILES}

