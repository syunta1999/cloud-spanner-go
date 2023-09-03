
tidy:
	cd api
	go mod tidy

###############################
# Docker containers
#
###############################

build:
	docker compose build --no-cache

up:
	docker compose up -d
	docker compose logs -f

up-backend:
	docker compose up api -d
	docker compose logs -f

stop:
	docker compose stop

down:
	docker compose down

reset-db:
	docker compose down db
	docker compose up db

restart:
	docker compose restart

logs: ## Tail docker compose logs
	docker compose logs -f

ps: ## Check container status
	docker compose ps
