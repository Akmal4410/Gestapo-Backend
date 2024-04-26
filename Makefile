postgres:
	@echo Creating a new container for postgres
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine

createdb:
	@echo Creating gestapo db
	docker exec -it postgres16 createdb --username=root --owner=root gestapo

dropdb:
	docker exec -it postgres16 dropdb gestapo

redis:
	@echo Creating a new container for postgres
	docker run --name redis7.2 -p 6379:6379 -d redis:7.2-alpine

prune:
	@echo Removing unused images starting with deploy-
	docker images | grep '^deploy-' | awk '{print $3}' | xargs -I {} docker rmi {}


authentication_server:
	@echo Running authentication service
	go run cmd/authentication_service/main.go 

grpc_gateway:
	@echo Running grpc gateway
	go run cmd/grpc_gateway/main.go

proto:
	@echo deleting generated files if exist..
	rm -f pkg/api/proto/*.go
	@echo Generating all proto pb files..
	protoc -I . \
	--go_out pkg/ --go_opt=paths=source_relative \
	--go-grpc_out pkg/ --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out pkg/ --grpc-gateway_opt=paths=source_relative \
	api/proto/*.proto
	@echo done..

evans:
	@echo Starting evans gRPC client..
	evans --host localhost --port 9001 -r repl      



compose_down: 
	@echo Stopping docker containers
	cd deploy && sudo docker compose down --remove-orphans
	@echo done

compose_up:
	@echo Start docker compose
	cd deploy && sudo docker compose up --build -d 
	@echo done

prune_images:
	@echo prune all images 
	cd deploy && sudo docker image prune -a
	@echo done 

AUTH_BINARY=authenticationServiceApp
GATEWAY_BINARY=gatewayApp

build_authentication:
	@echo Building authentication binary...
	cd cmd/authentication_service && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ${AUTH_BINARY} .
	@echo Moving file..
	mv cmd/authentication_service/${AUTH_BINARY} deploy/build
	@echo Done!

build_gateway:
	@echo Building gateway binary...
	cd cmd/grpc_gateway && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ${GATEWAY_BINARY} .
	@echo Moving file..
	mv cmd/grpc_gateway/${GATEWAY_BINARY} deploy/build
	@echo Done!


run: build_gateway build_authentication
	@echo Stopping docker images if running...
	cd deploy && docker compose down --remove-orphans
	@echo Building when required and starting docker images...
	cd deploy && sudo docker compose up --build -d
	@echo Docker images built and started!


 


.PHONY: postgres createdb dropdb server proto build_authentication run
