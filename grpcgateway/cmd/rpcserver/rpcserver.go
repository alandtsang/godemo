package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/alandtsang/grpc-gateway-demo/gen/go/your/service/v1"
)

type server struct {
	pb.UnimplementedYourServiceServer
}

func (s *server) Echo(ctx context.Context, in *pb.EchoReq) (*pb.EchoResp, error) {
	log.Printf("Received: %+v\n", in)
	greeting := fmt.Sprintf("Hello %s, your age is %d.", in.Name, in.Age)
	return &pb.EchoResp{Greeting: greeting}, nil
}

func main() {
	port := ":50001"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}
	fmt.Printf("listen on %s\n", port)

	s := grpc.NewServer()
	pb.RegisterYourServiceServer(s, &server{})
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}
}
