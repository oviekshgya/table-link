package config

import (
	"github.com/spf13/viper"
	"log"
)

var Config *Configuration

type Configuration struct {
	Server   ServerConfiguration
	Database DatabaseConfiguration
}

type ServerConfiguration struct {
	AppName        string
	Port           string
	Secret         string
	Mode           string
	Env            string
	PORTGRPCSERVER string
}

type DatabaseConfiguration struct {
	Driver       string
	Dbname       string
	Username     string
	Password     string
	Host         string
	Port         int
	MaxLifetime  int
	MaxOpenConns int
	MaxIdleConns int
}

func Setup() {
	configuration := &Configuration{}

	viper.SetConfigFile("./.env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error saat membaca file .env: %v", err)
	}

	configuration.Database.Host = viper.GetString("DB_HOST")
	configuration.Database.Port = viper.GetInt("DB_PORT")
	configuration.Database.Username = viper.GetString("DB_USER")
	configuration.Database.Password = viper.GetString("DB_PASSWORD")
	configuration.Database.Dbname = viper.GetString("DB_NAME")
	configuration.Database.MaxIdleConns = viper.GetInt("DB_MAX_IDLE_CONNS")
	configuration.Database.MaxOpenConns = viper.GetInt("DB_MAX_OPEN_CONNS")
	configuration.Database.MaxLifetime = viper.GetInt("DB_MAX_LIFETIME")

	configuration.Server.AppName = viper.GetString("SERVICE_NAME")
	configuration.Server.Port = viper.GetString("SERVICE_PORT")
	configuration.Server.Secret = viper.GetString("SERVICE_SECRET")
	configuration.Server.Mode = viper.GetString("SERVICE_MODE")
	configuration.Server.Env = viper.GetString("SERVICE_ENV")

	viper.Set("database.host", configuration.Database.Host)
	viper.Set("database.port", configuration.Database.Port)
	viper.Set("database.user", configuration.Database.Username)
	viper.Set("database.password", configuration.Database.Password)
	viper.Set("database.name", configuration.Database.Dbname)

	Config = configuration

}

func GetConfig() *Configuration {
	return Config
}
