package config

import (
	"context"
	"table-link/grpc/pb"
)

func (s *Server) CreateUser(ctx context.Context, req *pb.UserRequest) (*pb.Response, error) {
	result, err := UserService.CreateUser(req)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.ResponseLogin, error) {
	result, err := UserService.LoginUser(req)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *Server) GetAllUser(ctx context.Context, req *pb.GetAllUserRequest) (*pb.ResponseGetAllUser, error) {
	result, err := UserService.GetAllUser(req)
	if err != nil {
		return nil, err

	}
	return result, nil
}
