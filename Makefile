.PHONY: docs
docs:
	swag init -d router -g swagger.go -o docs/swagger --parseDependency --parseInternal

.PHONY: compose-up
compose-up:
	docker compose -f docker/docker-compose.yml --env-file .env up -d

.PHONY: compose-up-b
compose-up-b:
	docker compose -f docker/docker-compose.yml --env-file .env up -d --build

.PHONY: compose-down
compose-down:
	docker compose -f docker/docker-compose.yml --env-file .env down

.PHONY: compose-down-v
compose-down-v:
	docker compose -f docker/docker-compose.yml --env-file .env down -v


.PHONY: compose-local-up
compose-local-up:
	docker compose -f docker/docker-compose.local.yml --env-file .env up -d

.PHONY: compose-local-down
compose-local-down:
	docker compose -f docker/docker-compose.local.yml --env-file .env down

.PHONY: compose-local-down-v
compose-local-down-v:
	docker compose -f docker/docker-compose.local.yml --env-file .env down -v

.PHONY: server
server:
	go run cmd/server/main.go