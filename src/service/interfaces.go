package service

import (
	"gorm.io/gorm"
	"table-link/grpc/pb"
	"table-link/src/model/role"
	"table-link/src/model/users"
)

type UserService interface {
	CreateUser(req *pb.UserRequest) (*pb.Response, error)
	LoginUser(req *pb.LoginRequest) (*pb.ResponseLogin, error)
	GetAllUser(req *pb.GetAllUserRequest) (*pb.ResponseGetAllUser, error)
	Delete(req *pb.DeleteRequest) (*pb.Response, error)
}

type UserServiceImpl struct {
	UserModel users.UserModel
	DB        *gorm.DB
	RoleModel role.RoleModel
}

func NewUserService(usermodel users.UserModel, rolemodel role.RoleModel, db *gorm.DB) UserService {
	return &UserServiceImpl{
		UserModel: usermodel,
		DB:        db,
		RoleModel: rolemodel,
	}
}
