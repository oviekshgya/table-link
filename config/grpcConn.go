package config

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"table-link/grpc/pb"
	"table-link/src/interceptor"
)

type ServicesHandlers struct {
	PortGRPC string
}

type Server struct {
	pb.UnimplementedUserServiceServer
}

func (service ServicesHandlers) ConnGRPC() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", service.PortGRPC))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.NewAuthInterceptor("asdsada").Unary()),
		grpc.UnaryInterceptor(interceptor.NewAuthInterceptor("asdsada").Role()),
	)
	pb.RegisterUserServiceServer(s, &Server{})
	fmt.Println("CONNECTED: ", service.PortGRPC)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve grpc: %v", err)
	}
}
