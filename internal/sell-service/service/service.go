package service

import (
	"github.com/sirupsen/logrus"
	"shop/internal/sell-service/infrastructure"
	"shop/internal/sell-service/interfaces"
	"shop/pkg/database"
	"shop/pkg/server"
)

type Service struct {
	server *server.Server
	logger *logrus.Logger
}

type Config struct {
	ServerConfig     *ServerConfig
	RepositoryConfig *database.RepositoryConfig
}

type ServerConfig struct {
	Port string
}

func NewService(config *Config, logger *logrus.Logger) (*Service, error) {
	repoConfig := config.RepositoryConfig
	repo, err := database.RepositoryInit(repoConfig)
	if err != nil {
		return nil, err
	}

	gormInfrastructure := infrastructure.GormInfrastructure{
		Repository: repo,
	}

	app, err := infrastructure.GormInfrastructure.ApplicationInit(gormInfrastructure)
	if err != nil {
		return nil, err
	}

	return &Service{
		server: interfaces.NewHttpServer(config.ServerConfig.Port, app),
		logger: logger,
	}, nil
}

func (s *Service) Run() error {
	s.server.Run()
	return nil
}
