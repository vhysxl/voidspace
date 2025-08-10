include .env
export

APP_NAME=voidspace
MIGRATION_DIR=database/migrations/
GENERATED_PROTO=./proto/generated
DB_URL=mysql://$(DB_USER):$(DB_PASS)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)?multiStatements=true

.PHONY: migration_up migration_down migration_fix

migration_up: #push migration
	migrate -path $(MIGRATION_DIR) -database "$(DB_URL)" up
 
migration_down: #rollback migrations
	migrate -path $(MIGRATION_DIR) -database "$(DB_URL)" down 1

migration_fix: #rollback to certain version
	migrate -path $(MIGRATION_DIR) -database "$(DB_URL)" force 3 

proto_generate:
	protoc --go_out=${GENERATED_PROTO} --go-grpc_out=${GENERATED_PROTO} proto/users/v1/auth.proto