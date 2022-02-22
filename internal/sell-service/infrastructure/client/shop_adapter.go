package client

import (
	"math/rand"
	"shop/internal/sell-service/application/command_workers"
	"shop/internal/sell-service/domain/entity"
	"shop/pkg/database"
)

type ShopAdapter struct {
	Repository *database.Repository
}

func (reader *ShopAdapter) CreateProduct(request command_workers.ProductRequest) error {
	err := reader.Repository.DB.Exec("insert into product (id, name, number) values (%s, %s, %s)",
		rand.Uint32(),
		request.Name,
		request.Number,
	).Error

	return err
}

func (reader *ShopAdapter) GetProductsList() (*[]entity.Product, error) {
	var products []entity.Product
	err := reader.Repository.DB.Find(&products).Scan(&products).Error

	if err != nil {
		return nil, err
	}

	return &products, nil
}

func (reader *ShopAdapter) GetProductsCount() (int64, error) {
	db := reader.Repository.DB.Table("product")
	var result int64
	err := db.Count(&result).Error

	if err != nil {
		return 0, err
	}

	return result, nil
}
