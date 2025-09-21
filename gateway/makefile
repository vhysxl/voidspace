include .env
export

GENERATED_PROTO := ./proto/generated
PROTO_PATH := ../proto
PROTO_FILE := posts/post.proto

.PHONY: proto_generate

proto_generate:
	protoc --proto_path=$(PROTO_PATH) --go_out=$(GENERATED_PROTO) --go-grpc_out=$(GENERATED_PROTO) $(PROTO_FILE)
	


	