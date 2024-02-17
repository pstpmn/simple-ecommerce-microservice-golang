package productHandler

import (
	"context"
	productCore "simple-ecomerce-microservice/services/product/core"
	prodPb "simple-ecomerce-microservice/services/product/productPb"
)

type (
	handler struct {
		usecase productCore.IProductUseCase
		prodPb.UnimplementedProductServiceServer
	}
)

func (ob *handler) GetProductDetails(pctx context.Context, req *prodPb.ProductId) (*prodPb.ProductProfile, error) {
	prod, err := ob.usecase.GetProductDetail(pctx, req.ProductId)
	if err != nil {
		return nil, err
	}
	return &prodPb.ProductProfile{
		Product: &prodPb.Product{
			Id:          prod.Product.Id,
			Name:        prod.Product.Name,
			Description: prod.Product.Description,
			Price:       prod.Product.Price,
			Stock:       prod.Product.Stock,
			// CreatedAt:   prod.Product.CreatedAt.String(),
		},
		Category: &prodPb.Category{
			Id:   prod.Category.Id,
			Name: prod.Category.Name,
			// CreatedAt: prod.Category.CreatedAt.String(),
		},
	}, err
}

func (ob *handler) StockManager(pctx context.Context, req *prodPb.StockManagerReq) (*prodPb.ProductProfile, error) {
	prod, err := ob.usecase.StockManager(pctx, req.ProductId, req.Topic, req.Amount)
	if err != nil {
		return nil, err
	}
	return &prodPb.ProductProfile{
		Product: &prodPb.Product{
			Id:          prod.Product.Id,
			Name:        prod.Product.Name,
			Description: prod.Product.Description,
			Price:       prod.Product.Price,
			Stock:       prod.Product.Stock,
			// CreatedAt:   prod.Product.CreatedAt.String(),
		},
		Category: &prodPb.Category{
			Id:   prod.Category.Id,
			Name: prod.Category.Name,
			// CreatedAt: prod.Category.CreatedAt.String(),
		},
	}, err
}

// func VerifyCustomer(customerId int) *pb.Ver
func NewProductGrpcHandler(usecase productCore.IProductUseCase) prodPb.ProductServiceServer {
	return &handler{
		usecase: usecase,
	}
}
