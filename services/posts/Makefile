include .env
export

GENERATED_PROTO := ./proto/generated
PROTO_PATH := ../../proto
PROTO_FILE := posts/post.proto
DB_URL=postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)

.PHONY: migration_up migration_down migration_fix proto_generate

migration_up: #push migration
	migrate -path $(MIGRATION_DIR) -database "$(DB_URL)" up
 
migration_down: #rollback migrations
	migrate -path $(MIGRATION_DIR) -database "$(DB_URL)" down 1

migration_fix: #rollback to certain version
	migrate -path $(MIGRATION_DIR) -database "$(DB_URL)" force 3 

proto_generate:
	protoc --proto_path=${PROTO_PATH} --go_out=${GENERATED_PROTO} --go-grpc_out=${GENERATED_PROTO} ${PROTO_FILE}
	


	