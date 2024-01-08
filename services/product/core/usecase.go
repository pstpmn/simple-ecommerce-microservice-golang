package productCore

import (
	"errors"
	productModel "simple-ecomerce-microservice/services/product/models"
)

type (
	usecase struct {
		repo   IProductRepo
		helper IHelper
	}
)

// StockManager implements IProductUseCase.
func (ob *usecase) StockManager(productId string, topic string, amount int64) (*ProductProfile, error) {
	if amount == 0 || amount < 0 {
		return nil, errors.New("amount stock don't less than or equal to zero value")
	}

	prod, err := ob.repo.FindProduct(productId)
	if err != nil {
		return nil, err
	}
	if prod.ProductId == "" {
		return nil, errors.New("not found product")
	}

	switch topic {
	case "add":
		prod.Stock += amount
	case "sub":
		prod.Stock -= amount
		if prod.Stock < 0 {
			return nil, errors.New("subtrack don't less then zero value")
		}
	case "tobe":
		prod.Stock = amount
	default:
		return nil, errors.New("invalid topic can't manage stock")
	}

	err = ob.repo.UpdateOne(prod)
	domain := ob.productToDomain(prod)
	return &domain, err
}

// DecreseStock implements IProductUseCase.
func (ob *usecase) DecreseStock(productId string, amountDecrese int64) (*ProductProfile, error) {
	prod, err := ob.repo.FindProduct(productId)
	if err != nil {
		return nil, err
	}
	if prod.ProductId == "" {
		return nil, errors.New("not found product")
	}

	prod.Stock -= amountDecrese
	err = ob.repo.UpdateOne(prod)
	domain := ob.productToDomain(prod)
	return &domain, err
}

func (ob *usecase) IncreseStock(productId string, amountIncrese int64) (*ProductProfile, error) {
	if amountIncrese == 0 {
		return nil, errors.New("amount stock don't zero value")
	}

	prod, err := ob.repo.FindProduct(productId)
	if err != nil {
		return nil, err
	}
	if prod.ProductId == "" {
		return nil, errors.New("not found product")
	}

	prod.Stock += amountIncrese
	err = ob.repo.UpdateOne(prod)
	domain := ob.productToDomain(prod)
	return &domain, err
}

func (*usecase) productToDomain(product productModel.ProductModel) ProductProfile {
	return ProductProfile{
		Product: &Product{
			Id:          product.ProductId,
			Name:        product.ProductName,
			Description: product.Description,
			Price:       product.Price,
			Stock:       product.Stock,
			CreatedAt:   product.CreatedAt,
		},
		Category: &Category{
			Id:        product.Category.Id,
			Name:      product.Category.Name,
			CreatedAt: product.Category.CreatedAt,
		},
	}
}
func (*usecase) productsToProductIntroductuinDomain(products []productModel.ProductModel) []ProductIntroduction {
	var domain []ProductIntroduction
	for _, v := range products {
		domain = append(domain, ProductIntroduction{
			Id:        v.ProductId,
			Name:      v.ProductName,
			CreatedAt: v.CreatedAt,
		})
	}
	return domain
}

// GetProductDetail implements IProductUseCase.
func (ob *usecase) GetProductDetail(productId string) (*ProductProfile, error) {
	prod, err := ob.repo.FindProduct(productId)
	if err != nil {
		return nil, err
	}
	if prod.ProductId == "" {
		return nil, errors.New("not found product")
	}

	domain := ob.productToDomain(prod)
	return &domain, nil
}

// GetProducts implements IProductUseCase.
func (ob *usecase) GetProducts() (*[]ProductIntroduction, error) {
	prod, err := ob.repo.FindProducts()
	domain := ob.productsToProductIntroductuinDomain(prod)
	return &domain, err
}

func NewUseCase(repo IProductRepo, helper IHelper) IProductUseCase {
	return &usecase{
		repo:   repo,
		helper: helper,
	}
}
