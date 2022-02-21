package query_workers

import (
	"shop/internal/sell-service/domain/entity"
)

type ShopReceiver struct {
	ShopReceiverInterface ShopReceiverInterface
}

type ShopReceiverInterface interface {
	ListProducts() (*[]entity.Product, error)
	CountProducts() (int64, error)
}

// get thing
func (receiver *ShopReceiver) ListProducts() (*[]entity.Product, error) {
	return receiver.ShopReceiverInterface.ListProducts()
}

func (receiver *ShopReceiver) CountProducts() (int64, error) {
	return receiver.ShopReceiverInterface.CountProducts()
}
