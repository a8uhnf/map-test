SERVER_OUT := "bin/server"
CLIENT_OUT := "bin/client"
API_OUT := "api/api.pb.go"
API_REST_OUT := "api/api.pb.gw.go"
PKG := "github.com/a8uhnf/map-test"
SERVER_PKG_BUILD := "${PKG}/server"
CLIENT_PKG_BUILD := "${PKG}/client"

## RUN "make api" command to build and generate client.
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

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

