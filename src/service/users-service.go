package service

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"strconv"
	"table-link/grpc/pb"
	"table-link/pkg/helper"
	"table-link/src/model/users"
	"time"
)

func (service *UserServiceImpl) CreateUser(req *pb.UserRequest) (*pb.Response, error) {

	result, err := helper.WithTransaction(service.DB, func(tz *gorm.DB) (interface{}, error) {
		modelUser := users.UserModel{}
		roleid, _ := strconv.Atoi(req.RoleId)
		fmt.Println(roleid)
		hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, fmt.Errorf("encrypt password error: %s", err.Error())
		}

		if created := modelUser.Create(tz, users.Users{
			RoleID:   uint(roleid),
			Username: req.Email,
			Password: string(hash),
		}); created != nil {
			return &pb.Response{
				Status:  false,
				Message: fmt.Sprint(created),
			}, created
		}
		return &pb.Response{
			Status:  true,
			Message: "success",
		}, nil

	})
	if err != nil {
		return nil, err
	}
	return result.(*pb.Response), nil
}

func (service *UserServiceImpl) LoginUser(req *pb.LoginRequest) (*pb.ResponseLogin, error) {
	result, err := helper.WithTransaction(service.DB, func(tz *gorm.DB) (interface{}, error) {
		initialRedis := helper.InitializeRedis()

		modelUser := users.UserModel{}
		data, err := modelUser.GetByUsername(tz, req.Email)
		if err != nil {
			return nil, err
		}
		if err := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(req.Password)); err != nil {
			return nil, err
		}
		generate, err := helper.GenerateToken(data.ID, data.Username)
		if err != nil {
			return nil, err
		}
		if set := initialRedis.SetKey(fmt.Sprintf("login-%s", req.Email), generate, time.Duration(30*time.Second)); set != nil {
			log.Println("set reddis error")
		}

		return &pb.ResponseLogin{
			BaseResponse: &pb.Response{
				Status:  true,
				Message: "success",
			},
			Data: &pb.Data{
				AccessToken: generate,
			},
		}, nil
	})
	if err != nil {
		return nil, err
	}
	return result.(*pb.ResponseLogin), nil
}

func (service *UserServiceImpl) GetAllUser(req *pb.GetAllUserRequest) (*pb.ResponseGetAllUser, error) {
	result, err := helper.WithTransaction(service.DB, func(tz *gorm.DB) (interface{}, error) {
		modelUser := users.UserModel{}
		initialRedis := helper.InitializeRedis()

		var response pb.ResponseGetAllUser
		if get := initialRedis.GetKey(fmt.Sprintf("getall"), &response); get == nil {
			return &response, nil
		}

		data, err := modelUser.GetAll(tz)
		if err != nil {
			return nil, err
		}

		if len(data) > 0 {
			for i := 0; i < len(data); i++ {
				response.Data = &pb.User{
					Email:      data[i].Username,
					RoleId:     fmt.Sprint(data[i].RoleID),
					Name:       data[i].Username,
					LastAccess: fmt.Sprint("%v", time.Now()),
					RoleName:   fmt.Sprint(data[i].Role.Name),
				}
			}

		}

		if set := initialRedis.SetKey(fmt.Sprintf("getall"), &response, time.Duration(30*time.Second)); set != nil {
			log.Println("set reddis error")
		}
		return &response, nil
	})

	if err != nil {
		return nil, err
	}
	return result.(*pb.ResponseGetAllUser), nil
}
