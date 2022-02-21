package interfaces

import (
	"github.com/gin-gonic/gin"
	"shop/internal/sell-service/application"
	"shop/internal/sell-service/interfaces/shop"
	"shop/pkg/server"
)

type HttpPort struct {
	app *application.Application
}

func NewHttpServer(host string, newApp *application.Application) *server.Server {
	return server.NewServer(host, &HttpPort{app: newApp})
}

func (port *HttpPort) RepositoryStatus() error {
	return nil
}

func (port *HttpPort) AddRoutes(router *gin.Engine) error {
	api := router.Group("/")

	workers := &shop.ShopWorkers{
		CommandReceiver: port.app.CommandWorkers.ShopPusher,
		QueryReceiver:   port.app.QueryWorkers.ShopReceiver,
	}

	shop.RegisterShopReceiverRoutes(workers, api)

	return nil
}
