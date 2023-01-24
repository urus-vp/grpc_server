package main

import (
	"context"
	"flag"
	"fmt"
	"grpc_server/internal/config"
	"net"

	log "github.com/sirupsen/logrus"
	grpc_proto "github.com/urus-vp/grpc_proto"
	"google.golang.org/grpc"
)

type server struct {
	grpc_proto.UnimplementedGreeterServer
}

func (s *server) WriteFirebasePayload(ctx context.Context, in *grpc_proto.FirebasePayload) (*grpc_proto.FirebaseReply, error) {
	log.Printf("WriteFirebasePayload: Received: %v", in)
	return &grpc_proto.FirebaseReply{Message: "From server.WriteFirebasePayload..."}, nil
}

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	conf := config.Get()

	log.Println("Starting Server:", conf)

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	grpc_proto.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
