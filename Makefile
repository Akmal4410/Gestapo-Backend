postgres:
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root gestapo

dropdb:
	docker exec -it postgres16 dropdb gestapo

server:
	go run cmd/main.go

proto:
	@echo deleting generated files if exist..
	rm -f pkg/api/proto/*.go
	@echo Generating all proto pb files..
	protoc -I . \
	--go_out pkg/ --go_opt=paths=source_relative \
	--go-grpc_out pkg/ --go-grpc_opt=paths=source_relative \
	api/proto/*.proto
	@echo done..

AUTH_BINARY=authenticationServiceApp

build_authentication:
	@echo Building authentication binary...
	cd cmd/authentication_service && go build -o ${AUTH_BINARY} .
	@echo Moving file..
	mv cmd/authentication_service/${AUTH_BINARY} deploy/build
	@echo Done!

run: build_authentication
	@echo Stopping docker images if running...
	sudo docker compose down --remove-orphans
	@echo Building when required and starting docker images...
	sudo docker compose up --build -d
	@echo Docker images built and started!

.PHONY: postgres createdb dropdb server proto build_authentication run
