package product

import (
	"context"
	"tlkm-api/internal/model/product"
)

type Service interface {
	CreateProduct(ctx context.Context, req product.Product) (product.Product, error)
	UpdateProduct(ctx context.Context, sku string, req product.UpdateProductRequest) (product.Product, error)
	GetProductBySKU(ctx context.Context, sku string) (product.Product, error)
	GetProducts(ctx context.Context, param product.ParamSearch) (product.SearchProductResponse, error)
}
