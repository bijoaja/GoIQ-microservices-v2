.PHONY: setup build-all build-user build-auth build-gateway up-prod up-dev down clean clean-docker

setup:
	@chmod +x setup.sh
	@./setup.sh

# Windows Build
build-all: build-user build-auth build-gateway

build-user:
	$(MAKE) -C user-service build

build-auth:
	$(MAKE) -C auth-service build

build-gateway:
	$(MAKE) -C gateway build

# Docker Build

#production
up-prod: 
	docker compose -f docker-compose.prod.yml up --build -d

#development
up-dev:
	docker compose -f docker-compose.dev.yml up -d --build $(SERVICE)

down:
	docker-compose down

clean:
	rm -rf **/*.db
	rm -f user-service/user-service auth-service/auth-service gateway/gateway

clean-docker:
	docker system prune -af
