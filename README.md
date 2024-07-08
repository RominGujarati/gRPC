# gRPC User Service

This repository contains a gRPC service for managing user details with search functionality.

## Prerequisites

- Go 1.11+ installed on your machine.
- Docker (optional, for containerization).

## Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/RominGujarati/gRPC.git
   cd gRPC

2. Initialize Go modules:
    go mod init github.com/RominGujarati/gRPC
    go mod tidy


Build and Run
Local Development
1. Build the application:
    go build -o grpc-user-service
2. Run the gRPC service:
    ./grpc-user-service
The gRPC service will start listening on port 50051.


Docker
1. Build the Docker image:
    docker build -t grpc-user-service .
2. Run the Docker container:
    docker run -p 50051:50051 grpc-user-service
The gRPC service will be accessible through Docker on port 50051.



Accessing gRPC Service Endpoints
You can interact with the gRPC service using tools like grpcurl or by writing a client in Go.

Example gRPCurl Commands:
1. Get user by ID:
    grpcurl -plaintext -d '{"id": 1}' localhost:50051 user.UserService/GetUserById

2. Get users by IDs (streaming):
    grpcurl -plaintext -d '{"ids": [1, 2]}' localhost:50051 user.UserService/GetUsersByIds

3. Search users:
    grpcurl -plaintext -d '{"city": "LA"}' localhost:50051 user.UserService/SearchUsers


Configuration
No additional configuration is required to run the gRPC service locally. Ensure port 50051 is accessible and not blocked by a firewall.