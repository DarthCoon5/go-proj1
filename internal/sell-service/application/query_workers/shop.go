package query_workers

import (
	"shop/internal/sell-service/domain/entity"
)

type ShopReceiver struct {
	ShopReceiverInterface ShopReceiverInterface
}

type ShopReceiverInterface interface {
	GetProductsList() (*[]entity.Product, error)
	GetProductsCount() (int64, error)
}

// get thing
func (receiver *ShopReceiver) GetProductsList() (*[]entity.Product, error) {
	return receiver.ShopReceiverInterface.GetProductsList()
}

func (receiver *ShopReceiver) GetProductsCount() (int64, error) {
	return receiver.ShopReceiverInterface.GetProductsCount()
}
