# Define the compose files to be used
COMPOSE_FILES := -f deployments/docker-compose.yaml -f deployments/docker-compose.kong.yaml

.PHONY: up down logs ps

# Start all services in detached mode
compose-up:
	@echo "Starting all services..."
	docker-compose $(COMPOSE_FILES) up -d

# Stop all services
compose-down:
	@echo "Stopping all services..."
	docker-compose $(COMPOSE_FILES) down

# View logs for all services
compose-logs:
	@echo "Tailing logs for all services..."
	docker-compose $(COMPOSE_FILES) logs -f

# List running services
compose-ps:
	@echo "Listing running services..."
	docker-compose $(COMPOSE_FILES) ps

# Rebuild and start all services
compose-rebuild:
	@echo "Rebuilding and starting all services..."
	docker-compose $(COMPOSE_FILES) up -d --build

# Reload Kong by restarting the container to apply declarative config changes
kong-sync:
	@echo "Reloading Kong to apply declarative config changes..."
	docker-compose $(COMPOSE_FILES) restart kong-gw