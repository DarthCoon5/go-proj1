package command_workers

type ShopPusher struct {
	ShopReceiverInterface ShopReceiverInterface
}

type ProductRequest struct {
	Name   string
	Number int
}

type ShopReceiverInterface interface {
	CreateProduct(request ProductRequest) error
}

func (receiver *ShopPusher) CreateProduct(request ProductRequest) error {
	return receiver.ShopReceiverInterface.CreateProduct(request)
}
