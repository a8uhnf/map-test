SERVER_OUT := "bin/server"
CLIENT_OUT := "bin/client"
API_OUT := "api/api.pb.go"
API_REST_OUT := "api/api.pb.gw.go"
PKG := "github.com/a8uhnf/map-test"
SERVER_PKG_BUILD := "${PKG}/server"
CLIENT_PKG_BUILD := "${PKG}/client"
# PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)

.PHONY: all api server client

all: server client

api/api.pb.go: api/api.proto
	@protoc -I api/ \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I${HOME}/Downloads/protoc3/include \
		--go_out=plugins=grpc:api \
		api/api.proto

api/api.pb.gw.go: api/api.proto
	@protoc -I api/ \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I${HOME}/Downloads/protoc3/include \
		--grpc-gateway_out=logtostderr=true:api \
		api/api.proto

api: api/api.pb.go api/api.pb.gw.go ## Auto-generate grpc go sources

dep: ## Get the dependencies
	@go get -v -d ./...

server: api ## Build the binary file for server
	@go build -i -v -o $(SERVER_OUT) $(SERVER_PKG_BUILD)

client: api ## Build the binary file for client
	@go build -i -v -o $(CLIENT_OUT) $(CLIENT_PKG_BUILD)

clean: ## Remove previous builds
	@rm $(SERVER_OUT) $(CLIENT_OUT) $(API_OUT) $(API_REST_OUT)

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

