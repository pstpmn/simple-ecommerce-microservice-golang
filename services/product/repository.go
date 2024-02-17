package product

import (
	"context"
	"errors"
	"fmt"
	productCore "simple-ecomerce-microservice/services/product/core"
	productModel "simple-ecomerce-microservice/services/product/models"

	"gorm.io/gorm"
)

type (
	repo struct {
		db *gorm.DB
	}
)

// UpdateOne implements productCore.IProductRepo.
func (ob *repo) UpdateOne(pctx context.Context, model productModel.ProductModel) error {
	if err := ob.
		db.
		Model(&model).
		Where(productModel.ProductModel{ProductId: model.ProductId}).
		Updates(&model).
		Error; err != nil {
		fmt.Printf("error func UpdateOne : %v", err.Error())
		return errors.New("something went wrong")
	}
	return nil
}

// FindProduct implements productCore.IProductRepo.
func (ob *repo) FindProduct(pctx context.Context, productId string) (productModel.ProductModel, error) {
	var model productModel.ProductModel
	err := ob.
		db.
		Where(&productModel.ProductModel{ProductId: productId}).
		Preload("Category").
		First(&model).
		Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Printf("error func FindProduct : %v", err.Error())
		return model, errors.New("something went wrong")
	}
	return model, nil
}

// FindProducts implements productCore.IProductRepo.
func (ob *repo) FindProducts(pctx context.Context) ([]productModel.ProductModel, error) {
	var model []productModel.ProductModel
	err := ob.
		db.
		Preload("Category").
		Find(&model).
		Error

	if err != nil {
		fmt.Printf("error func FindProducts : %v", err.Error())
		return model, errors.New("something went wrong")
	}
	return model, err
}

func NewRepository(conn *gorm.DB) productCore.IProductRepo {
	return &repo{
		db: conn,
	}
}
