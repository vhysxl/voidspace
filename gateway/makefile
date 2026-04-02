include .env
export

GENERATED_PROTO := ./proto/generated
PROTO_PATH := ../proto

.PHONY: proto_generate

proto_generate:
	protoc --proto_path=$(PROTO_PATH) --go_out=$(GENERATED_PROTO) --go-grpc_out=$(GENERATED_PROTO) $(PROTO_PATH)/users/v1/users.proto
	protoc --proto_path=$(PROTO_PATH) --go_out=$(GENERATED_PROTO) --go-grpc_out=$(GENERATED_PROTO) $(PROTO_PATH)/posts/v1/posts.proto
	protoc --proto_path=$(PROTO_PATH) --go_out=$(GENERATED_PROTO) --go-grpc_out=$(GENERATED_PROTO) $(PROTO_PATH)/comments/v1/comments.proto
