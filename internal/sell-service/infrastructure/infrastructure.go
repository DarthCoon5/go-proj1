package infrastructure

import (
	"shop/internal/sell-service/application"
	"shop/internal/sell-service/infrastructure/client"
	"shop/pkg/database"
)

type GormInfrastructure struct {
	Repository *database.Repository
}

func (inf *GormInfrastructure) ApplicationInit() (app *application.Application, err error) {
	app = new(application.Application)
	app.QueryWorkers.ShopReceiver.ShopReceiverInterface = &client.ShopAdapter{
		Repository: inf.Repository,
	}
	return app, nil
}
