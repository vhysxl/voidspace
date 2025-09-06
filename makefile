.PHONY: run up down

up:
	docker compose -f services/users/docker-compose.yml up -d
	docker compose -f services/posts/docker-compose.yml up -d

dev: 
	start cmd /k "cd gateway && air" & \
	start cmd /k "cd services/users && air" & \
	start cmd /k "cd services/posts && air"

down:
	docker compose -f services/users/docker-compose.yml down
	docker compose -f services/posts/docker-compose.yml down
