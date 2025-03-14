package config

import (
	"table-link/domain"
	"table-link/src/model/role"
	"table-link/src/model/users"
	"table-link/src/service"
)

var UserService service.UserService

func StartService() {
	UserService = service.NewUserService(users.UserModel{DB: domain.DB}, role.RoleModel{DB: domain.DB}, domain.DB)
}
