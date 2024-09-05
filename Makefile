.PHONY: docs
docs:
	swag init -d router -g swagger.go -o docs/swagger --parseDependency --parseInternal

.PHONY: compose-up
compose-up:
	docker compose -f docker/docker-compose.yml --env-file .env up -d

.PHONY: compose-down
compose-down:
	docker compose -f docker/docker-compose.yml --env-file .env down

.PHONY: compose-down-v
compose-down-v:
	docker compose -f docker/docker-compose.yml --env-file .env down -v