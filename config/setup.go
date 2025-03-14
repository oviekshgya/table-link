package config

import (
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"table-link/domain"
	"table-link/pkg/helper"
)

func setConfiguration() {
	Setup()
}

type SetDatabaseConfig struct {
	MainDatabase *gorm.DB
}

func SetDatabase(configuration Configuration) {
	var conf domain.DatabaseConfig
	conf.MaxOpenConns = configuration.Database.MaxOpenConns
	conf.MaxIdleConns = configuration.Database.MaxIdleConns
	conf.Dbname = configuration.Database.Dbname
	conf.Username = configuration.Database.Username
	conf.Password = configuration.Database.Password
	conf.Host = configuration.Database.Host
	conf.Port = configuration.Database.Port
	conf.Driver = configuration.Database.Driver
	conf.ConnectDatabase()
}

func Start() {
	setConfiguration()
	conf := GetConfig()
	SetDatabase(*conf)
	helper.GetRoleData(domain.DB)

	gin.SetMode(conf.Server.Mode)

	StartService()
	var services ServicesHandlers
	services.PortGRPC = conf.Server.Port
	ctx, _ := context.WithCancel(context.Background())
	go helper.StartAddClient(ctx, domain.DB)
	services.ConnGRPC()

}
