package v0

import (
	"context"
	"tlkm-api/internal/model/product"
)

func (service *ServiceProduct) CreateProduct(ctx context.Context, req product.Product) (product.Product, error) {
	resp, err := service.repo.ProductPostgre.CreateProduct(ctx, req)
	if err != nil {
		return product.Product{}, err
	}

	return resp, nil
}

func (service *ServiceProduct) UpdateProduct(ctx context.Context, sku string, req product.UpdateProductRequest) (product.Product, error) {
	resp, err := service.repo.ProductPostgre.UpdateProduct(ctx, sku, req)
	if err != nil {
		return product.Product{}, err
	}

	return resp, nil
}

func (service *ServiceProduct) GetProductBySKU(ctx context.Context, sku string) (product.Product, error) {
	resp, err := service.repo.ProductPostgre.GetProductBySKU(ctx, sku)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (service *ServiceProduct) GetProducts(ctx context.Context, param product.ParamSearch) (res product.SearchProductResponse, err error) {
	products := make([]product.SearchProductData, 0)

	products, err = service.repo.ProductPostgre.GetProductsByParam(ctx, param)
	if err != nil {
		return
	}
	res.TotalData = len(products)
	res.TotalPage = service.getTotalPage(int64(len(products)), param.PageSize)

	res.Data = products
	return
}

func (service *ServiceProduct) getTotalPage(total int64, pagesize int) (totalPage int) {
	if total < 1 || pagesize < 1 {
		return
	}

	totalPage = int(total) / pagesize
	if total%int64(pagesize) > 0 {
		totalPage += 1
	}
	return
}
