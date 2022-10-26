package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"

	pb "github.com/Maverickme222222/users/usermgmt"
	"google.golang.org/grpc"
)

const (
	port = ":9090"
)

type UserManagementServer struct {
	pb.UnimplementedUserManagementServer
}

func (s *UserManagementServer) CreateNewUser(ctx context.Context, req *pb.NewUser) (*pb.NewUserResponse, error) {
	log.Printf("Received: %v", req.Name)
	var user_id int32 = int32(rand.Intn(1000))
	user := fmt.Sprintf("Added new user id %v for name %v", user_id, req.GetName())
	return &pb.NewUserResponse{
		Name: user,
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen at port %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterUserManagementServer(s, &UserManagementServer{})
	log.Printf("Server listening at %v", listen.Addr().String())
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
