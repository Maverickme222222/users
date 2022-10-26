package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Maverickme222222/users/usermgmt"
	"google.golang.org/grpc"
)

const (
	address = "localhost:9090"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Did not connect")
	}
	defer conn.Close()

	c := pb.NewUserManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.CreateNewUser(ctx, &pb.NewUser{
		Name: "test",
	})
	if err != nil {
		log.Fatalf("could not create user: %v", err)
	}

	log.Printf(`User details
	Name: %s
	`, r.GetName())

}
