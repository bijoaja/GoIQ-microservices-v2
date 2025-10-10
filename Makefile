.PHONY: build-all build-user build-auth build-gateway up down clean

build-all: build-user build-auth build-gateway

build-user:
	$(MAKE) -C user-service build

build-auth:
	$(MAKE) -C auth-service build

build-gateway:
	$(MAKE) -C gateway build

up:
	docker-compose up --build

down:
	docker-compose down

clean:
	rm -rf **/*.db
	rm -f user-service/user-service auth-service/auth-service gateway/gateway
