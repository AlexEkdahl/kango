setup:
	@echo "Setting up the environment"
	@./scripts/setup.sh
	@make proto

cibuild:
	./scripts/cibuild.sh

#####################################

CLI=kango
SERVER=kango-server
SSH_SERVER=kango-ssh-server

CLI_SRC=./cmd/cli/main.go
SERVER_SRC=./cmd/server/main.go
SSH_SERVER_SRC=./cmd/ssh/main.go

BIN_DIR=./bin
.DEFAULT_GOAL := build
BUILD_CMD=go build -mod=readonly -ldflags="-s -w" -gcflags=all=-l -trimpath=true
BUILD_CMD_DOCKER=CGO_ENABLED=0 go build -mod=readonly -ldflags="-s -w -extldflags '-static'" -gcflags=all=-l -trimpath=true

build:
	@$(BUILD_CMD) -o $(BIN_DIR)/$(CLI) $(CLI_SRC)

docker-build-server:
	@$(BUILD_CMD_DOCKER) -o $(BIN_DIR)/$(SERVER) $(SERVER_SRC)

build-ssh-server:
	@$(BUILD_CMD) -o $(BIN_DIR)/$(SSH_SERVER) $(SSH_SERVER_SRC)

build-server:
	@$(BUILD_CMD) -o $(BIN_DIR)/$(SERVER) $(SERVER_SRC)

run: build
	@$(BIN_DIR)/$(CLI)

serve: build-server
	@$(BIN_DIR)/$(SERVER)

ssh: build-ssh-server
	@$(BIN_DIR)/$(SSH_SERVER)

proto:
	@mkdir -p pkg
	@protoc --go_out=pkg --go-grpc_out=pkg contract/kango.proto

test:
	go test ./... -v

clean:
	go clean
	rm -rf $(BIN_DIR)

