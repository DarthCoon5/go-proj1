package application

import (
	"shop/internal/sell-service/application/command_workers"
	"shop/internal/sell-service/application/query_workers"
	"shop/internal/sell-service/domain/entity"
)

type Application struct {
	CommandWorkers CommandWorkers
	QueryWorkers   QueryWorkers
}

//CommandWorkers ...
type CommandWorkers struct {
	ShopPusher command_workers.ShopPusher
}

//QueryWorkers ...
type QueryWorkers struct {
	ShopReceiver query_workers.ShopReceiver
}

//TODO
type ShopWorker struct {
	ShopWorkerInterface ShopWorkerInterface
}
type ShopWorkerInterface interface {
	CreateProductAndList(request command_workers.ProductRequest) (*[]entity.Product, error)
}

func (worker *Application) CreateProductAndList(request command_workers.ProductRequest) (*[]entity.Product, error) {
	worker.CommandWorkers.ShopPusher.CreateProduct(request)
	//algo stuff
	return worker.QueryWorkers.ShopReceiver.ListProducts()
}
