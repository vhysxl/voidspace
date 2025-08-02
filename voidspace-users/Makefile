include .env
export

APP_NAME=voidspace
MIGRATION_DIR=database/migrations/
DB_URL=mysql://$(DB_USER):$(DB_PASS)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)?multiStatements=true

.PHONY: migration_up migration_down migration_fix

migration_up: #push migration
	migrate -path $(MIGRATION_DIR) -database "$(DB_URL)" up
 
migration_down: #rollback migrations
	migrate -path $(MIGRATION_DIR) -database "$(DB_URL)" down

migration_fix: #rollback to certain version
	migrate -path $(MIGRATION_DIR) -database "$(DB_URL)" force $(VERSION)