package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	gw "github.com/alandtsang/grpc-gateway-demo/gen/go/your/service/v1"
)

var (
	// command-line options:
	// gRPC rpcserver endpoint
	grpcServerEndpoint = flag.String("grpc-rpcserver-endpoint", "localhost:50001", "gRPC rpcserver endpoint")
)

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC rpcserver endpoint
	// Note: Make sure the gRPC rpcserver is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterYourServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	// Start HTTP rpcserver (and proxy calls to gRPC rpcserver endpoint)
	port := ":8081"
	fmt.Printf("Listen on %s\n", port)
	return http.ListenAndServe(port, mux)
}
