
package main

import (
	"context"
	"log"
	"net"

	pb "go-server/pb"

	"google.golang.org/grpc"
)

type userServer struct {
	pb.UnimplementedUserServiceServer
}

func (s *userServer) GetUserProfile(
	ctx context.Context,
	req *pb.UserRequest,
) (*pb.UserResponse, error) {
	log.Printf("Request untuk user ID: %d", req.GetUserId())

	if req.GetUserId() == 1 {
		return &pb.UserResponse{
			Name:  "Budi Santoso",
			Email: "budi@universitas.ac.id",
			Phone: "081234567890",
		}, nil
	}

	return &pb.UserResponse{
		Name:  "User Tidak Ditemukan",
		Email: "-",
		Phone: "-",
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Gagal membuka port: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, &userServer{})

	log.Println("gRPC Server berjalan pada port 50051")

	if err := server.Serve(listener); err != nil {
		log.Fatalf("Server gagal: %v", err)
	}
}
