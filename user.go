package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/RominGujarati/gRPC/user"
)

// Define User struct
type User struct {
	ID      int32
	Fname   string
	City    string
	Phone   int64
	Height  float32
	Married bool
}

// Simulated database (slice of users)
var users = []*User{
	{ID: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
	{ID: 2, Fname: "Alice", City: "NYC", Phone: 9876543210, Height: 5.5, Married: false},
	// Add more users as needed
}

// UserServiceServer is the server API for User service
type UserServiceServer struct{}

// GetUserById returns a single user by ID
func (s *UserServiceServer) GetUserById(ctx context.Context, req *pb.UserRequest) (*pb.User, error) {
	for _, u := range users {
		if u.ID == req.Id {
			return &pb.User{
				Id:      u.ID,
				Fname:   u.Fname,
				City:    u.City,
				Phone:   u.Phone,
				Height:  u.Height,
				Married: u.Married,
			}, nil
		}
	}
	return nil, fmt.Errorf("user not found")
}

// GetUsersByIds returns multiple users by IDs
func (s *UserServiceServer) GetUsersByIds(req *pb.UserIdsRequest, stream pb.UserService_GetUsersByIdsServer) error {
	for _, id := range req.Ids {
		for _, u := range users {
			if u.ID == id {
				if err := stream.Send(&pb.User{
					Id:      u.ID,
					Fname:   u.Fname,
					City:    u.City,
					Phone:   u.Phone,
					Height:  u.Height,
					Married: u.Married,
				}); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// SearchUsers returns users based on search criteria
func (s *UserServiceServer) SearchUsers(req *pb.SearchRequest, stream pb.UserService_SearchUsersServer) error {
	for _, u := range users {
		if (req.City == "" || u.City == req.City) &&
			(req.Phone == 0 || u.Phone == req.Phone) &&
			(req.Married == false || u.Married == req.Married) &&
			(req.Fname == "" || u.Fname == req.Fname) &&
			(req.Height == 0 || u.Height == req.Height) &&
			(req.Id == 0 || u.ID == req.Id) {
			if err := stream.Send(&pb.User{
				Id:      u.ID,
				Fname:   u.Fname,
				City:    u.City,
				Phone:   u.Phone,
				Height:  u.Height,
				Married: u.Married,
			}); err != nil {
				return err
			}
		}
	}
	return nil
}

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, &UserServiceServer{})

	log.Println("Starting server on port :50051...")
	if err := server.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
