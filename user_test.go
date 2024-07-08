package main

import (
	"context"
	"log"
	"testing"

	pb "github.com/RominGujarati/gRPC/user"
	"google.golang.org/grpc"
)

func TestGetUserById(t *testing.T) {
	s := UserServiceServer{}
	req := &pb.UserRequest{Id: 1}

	user, err := s.GetUserById(context.Background(), req)
	if err != nil {
		t.Errorf("GetUserById returned error: %v", err)
	}

	expectedName := "Steve"
	if user.Fname != expectedName {
		t.Errorf("Expected user name: %s, got: %s", expectedName, user.Fname)
	}
}

func TestGetUsersByIds(t *testing.T) {
	// Establish a connection to the gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial server: %v", err)
	}
	defer conn.Close()

	// Create a gRPC client
	client := pb.NewUserServiceClient(conn)

	// Test case 1: Request with valid user IDs
	validIds := []int32{1, 2, 3}
	req := &pb.UsersRequest{Ids: validIds}

	resp, err := client.GetUsersByIds(context.Background(), req)
	if err != nil {
		t.Fatalf("failed to get users by IDs: %v", err)
	}

	// Validate response
	if len(resp.Users) != len(validIds) {
		t.Errorf("expected %d users, got %d", len(validIds), len(resp.Users))
	}

	log.Printf("GetUsersByIds Test Passed. Response: %v", resp)
}

// Test case for SearchUsers function
func TestSearchUsers(t *testing.T) {
	// Establish a connection to the gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial server: %v", err)
	}
	defer conn.Close()

	// Create a gRPC client
	client := pb.NewUserServiceClient(conn)

	// Test case 2: Search by city
	req := &pb.SearchRequest{City: "LA"}

	resp, err := client.SearchUsers(context.Background(), req)
	if err != nil {
		t.Fatalf("failed to search users: %v", err)
	}

	// Validate response
	for _, user := range resp.Users {
		if user.City != req.City {
			t.Errorf("expected city %s, got %s", req.City, user.City)
		}
	}

	log.Printf("SearchUsers Test Passed. Response: %v", resp)
}
