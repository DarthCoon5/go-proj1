package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"math/rand"
	sell_service "shop/internal/sell-service/service"
	"shop/pkg/database"
	"time"
)

func main() {
	var logger = logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	rand.Seed(time.Now().UnixNano())

	if err := initConfig(); err != nil {
		logrus.Fatalf("Init configs error: %s", err.Error())
	}

	config := getServiceConfig()
	service, err := sell_service.NewService(config, logger)
	if err != nil {
		logger.Error(fmt.Sprintln(err))
		return
	}

	if err := service.Run(); err != nil {
		logger.Error(fmt.Sprintln(err))
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func getServiceConfig() *sell_service.Config {
	config := new(sell_service.Config)

	config.ServerConfig = &sell_service.ServerConfig{
		Port: viper.GetString("port"),
	}

	config.RepositoryConfig = &database.RepositoryConfig{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		User:     viper.GetString("db.user"),
		DbName:   viper.GetString("db.dbName"),
		Password: viper.GetString("db.password"),
		SslMode:  viper.GetString("db.sslMode"),
	}

	return config
}
