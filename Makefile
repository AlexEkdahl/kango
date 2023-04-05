setup:
	@echo "Setting up the environment"
	@./scripts/setup.sh
	@make proto

cibuild:
	./scripts/cibuild.sh

#####################################

BINARY=kango
SERVER=kango-server
SRC=./cmd/cli/main.go
SERVER_SRC=./cmd/server/main.go
BIN_DIR=./bin
.DEFAULT_GOAL := build
BUILD_CMD=go build -mod=readonly -ldflags="-s -w" -gcflags=all=-l -trimpath=true
BUILD_CMD_DOCKER=CGO_ENABLED=0 go build -mod=readonly -ldflags="-s -w -extldflags '-static'" -gcflags=all=-l -trimpath=true

build:
	@$(BUILD_CMD) -o $(BIN_DIR)/$(BINARY) $(SRC)

docker-build-server:
	@$(BUILD_CMD_DOCKER) -o $(BIN_DIR)/$(SERVER) $(SERVER_SRC)

build-server:
	@$(BUILD_CMD) -o $(BIN_DIR)/$(SERVER) $(SERVER_SRC)

run: build
	@$(BIN_DIR)/$(BINARY)

serve: build-server
	@$(BIN_DIR)/$(SERVER)

proto:
	@mkdir -p pkg
	@protoc --go_out=pkg --go-grpc_out=pkg contract/kango.proto

test:
	go test ./... -v

clean:
	go clean
	rm -rf $(BIN_DIR)

