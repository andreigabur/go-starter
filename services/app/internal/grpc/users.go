package grpc

import (
	"context"

	"go-starter-app/internal/controllers"
	"go-starter-app/pkg/pb"
)

type UsersService struct {
	pb.UnimplementedUsersServiceServer
}

func NewUsersService() *UsersService {
	return &UsersService{}
}

func (s *UsersService) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	users := controllers.ListUsers()

	// Convert domain models to protobuf messages
	pbUsers := make([]*pb.User, len(users))
	for i, user := range users {
		pbUsers[i] = &pb.User{
			Id:    int32(user.ID),
			Name:  user.Name,
			Email: user.Email,
		}
	}

	return &pb.ListUsersResponse{
		Users: pbUsers,
	}, nil
}
