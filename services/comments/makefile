include .env
export

APP_NAME=voidspace
PROTO_PATH := ../../proto
PROTO_FILE := comments/comment.proto
MIGRATION_DIR=database/migrations/
GENERATED_PROTO=./proto/generated
DB_URL=mysql://$(DB_USER):$(DB_PASS)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)?multiStatements=true

migration_up: #push migration
	migrate -path ./database/migrations -database "$(DB_URL)" up
 
migration_down: #rollback migrations
	migrate -path $(MIGRATION_DIR) -database "$(DB_URL)" down 1

migration_fix: #rollback to certain version
	migrate -path $(MIGRATION_DIR) -database "$(DB_URL)" force 3 

proto_generate:
	protoc --proto_path=$(PROTO_PATH) --go_out=$(GENERATED_PROTO) --go-grpc_out=$(GENERATED_PROTO) $(PROTO_FILE)

migration_version: #check current migration version
	migrate -path $(MIGRAT	ION_DIR) -database "$(DB_URL)" version