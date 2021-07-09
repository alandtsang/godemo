Gateway := gateway
RPCServer := rpcserver

gen:
	protoc -I . -I ${GOPATH}/src/github.com/googleapis/googleapis \
      --go_out=plugins=grpc:. proto/your_service.proto
	protoc -I . -I ${GOPATH}/src/github.com/googleapis/googleapis \
      --grpc-gateway_out . proto/your_service.proto

build: gen
	GO111MODULE=on go build -mod vendor -o $(Gateway) cmd/gateway/gateway.go
	GO111MODULE=on go build -mod vendor -o $(RPCServer) cmd/rpcserver/rpcserver.go

clean:
	@rm -f $(Gateway) $(RPCServer)