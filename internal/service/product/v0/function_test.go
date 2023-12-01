package v0

import (
	"context"
	"errors"
	"reflect"
	"testing"
	mProduct "tlkm-api/internal/model/product"
	"tlkm-api/internal/repository/product"
)

func TestServiceProduct_CreateProduct(t *testing.T) {
	type fields struct {
		repo RepoAttribute
	}
	ctx := context.Background()
	type args struct {
		ctx context.Context
		req mProduct.Product
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult mProduct.Product
		wantErr    bool
	}{
		{
			name: "case-1: success, happy flow",
			args: args{
				ctx: ctx,
				req: mProduct.Product{
					SKU:         "123456790",
					Title:       "Aqua 250ml",
					Description: "Air Mineral",
					Category:    "Air Mineral Berkualitas",
					ImgURL:      "https://images.tokopedia.net/img/cache/700/product-1/2018/5/17/22673038/22673038_d74c5460-b815-4ed4-89a4-f0a92cfb2b82_701_701.jpg",
					Weight:      float64(0.70),
					Price:       float64(11000),
				},
			},
			fields: fields{
				repo: RepoAttribute{
					ProductPostgre: &product.RepositoryMock{
						CreateProductFunc: func(ctx context.Context, req mProduct.Product) (mProduct.Product, error) {
							return mProduct.Product{
								SKU:         "123456790",
								Title:       "Aqua 250ml",
								Description: "Air Mineral",
								Category:    "Air Mineral Berkualitas",
								ImgURL:      "https://images.tokopedia.net/img/cache/700/product-1/2018/5/17/22673038/22673038_d74c5460-b815-4ed4-89a4-f0a92cfb2b82_701_701.jpg",
								Weight:      float64(0.70),
								Price:       float64(11000),
							}, nil
						},
					},
				},
			},
			wantResult: mProduct.Product{
				SKU:         "123456790",
				Title:       "Aqua 250ml",
				Description: "Air Mineral",
				Category:    "Air Mineral Berkualitas",
				ImgURL:      "https://images.tokopedia.net/img/cache/700/product-1/2018/5/17/22673038/22673038_d74c5460-b815-4ed4-89a4-f0a92cfb2b82_701_701.jpg",
				Weight:      float64(0.70),
				Price:       float64(11000),
			},
		},
		{
			name: "case-2: error, postgresql got issue",
			args: args{
				ctx: ctx,
				req: mProduct.Product{
					SKU:         "123456790",
					Title:       "Aqua 250ml",
					Description: "Air Mineral",
					Category:    "Air Mineral Berkualitas",
					ImgURL:      "https://images.tokopedia.net/img/cache/700/product-1/2018/5/17/22673038/22673038_d74c5460-b815-4ed4-89a4-f0a92cfb2b82_701_701.jpg",
					Weight:      float64(0.70),
					Price:       float64(11000),
				},
			},
			fields: fields{
				repo: RepoAttribute{
					ProductPostgre: &product.RepositoryMock{
						CreateProductFunc: func(ctx context.Context, req mProduct.Product) (mProduct.Product, error) {
							return mProduct.Product{}, errors.New("something went wrong")
						},
					},
				},
			},
			wantErr:    true,
			wantResult: mProduct.Product{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &ServiceProduct{
				repo: tt.fields.repo,
			}
			gotResult, err := service.CreateProduct(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ServiceProduct.CreateProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("ServiceProduct.CreateProduct() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestServiceProduct_UpdateProduct(t *testing.T) {
	type fields struct {
		repo RepoAttribute
	}
	ctx := context.Background()
	type args struct {
		ctx context.Context
		sku string
		req mProduct.UpdateProductRequest
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult mProduct.Product
		wantErr    bool
	}{
		{
			name: "case-1: success, happy flow",
			args: args{
				ctx: ctx,
				sku: "123456789",
				req: mProduct.UpdateProductRequest{
					SKU:         "123456790",
					Title:       "Aqua 250ml",
					Description: "Air Mineral",
					Category:    "Air Mineral Berkualitas",
				},
			},
			fields: fields{
				repo: RepoAttribute{
					ProductPostgre: &product.RepositoryMock{
						UpdateProductFunc: func(ctx context.Context, sku string, req mProduct.UpdateProductRequest) (mProduct.Product, error) {
							return mProduct.Product{
								SKU:         "123456790",
								Title:       "Aqua 250ml",
								Description: "Air Mineral",
								Category:    "Air Mineral Berkualitas",
								ImgURL:      "https://images.tokopedia.net/img/cache/700/product-1/2018/5/17/22673038/22673038_d74c5460-b815-4ed4-89a4-f0a92cfb2b82_701_701.jpg",
								Weight:      float64(0.70),
								Price:       float64(11000),
							}, nil
						},
					},
				},
			},
			wantResult: mProduct.Product{
				SKU:         "123456790",
				Title:       "Aqua 250ml",
				Description: "Air Mineral",
				Category:    "Air Mineral Berkualitas",
				ImgURL:      "https://images.tokopedia.net/img/cache/700/product-1/2018/5/17/22673038/22673038_d74c5460-b815-4ed4-89a4-f0a92cfb2b82_701_701.jpg",
				Weight:      float64(0.70),
				Price:       float64(11000),
			},
		},
		{
			name: "case-2: error, postgresql got issue",
			args: args{
				ctx: ctx,
				sku: "123456789",
				req: mProduct.UpdateProductRequest{
					SKU:         "123456790",
					Title:       "Aqua 250ml",
					Description: "Air Mineral",
					Category:    "Air Mineral Berkualitas",
				},
			},
			fields: fields{
				repo: RepoAttribute{
					ProductPostgre: &product.RepositoryMock{
						UpdateProductFunc: func(ctx context.Context, sku string, req mProduct.UpdateProductRequest) (mProduct.Product, error) {
							return mProduct.Product{}, errors.New("something went wrong")
						},
					},
				},
			},
			wantErr:    true,
			wantResult: mProduct.Product{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &ServiceProduct{
				repo: tt.fields.repo,
			}
			gotResult, err := service.UpdateProduct(tt.args.ctx, tt.args.sku, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ServiceProduct.UpdateProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("ServiceProduct.UpdateProduct() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestServiceProduct_GetProductBySKU(t *testing.T) {
	type fields struct {
		repo RepoAttribute
	}
	ctx := context.Background()
	type args struct {
		ctx context.Context
		sku string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult mProduct.Product
		wantErr    bool
	}{
		{
			name: "case-1: success, happy flow",
			args: args{
				ctx: ctx,
				sku: "123456790",
			},
			fields: fields{
				repo: RepoAttribute{
					ProductPostgre: &product.RepositoryMock{
						GetProductBySKUFunc: func(ctx context.Context, sku string) (mProduct.Product, error) {
							return mProduct.Product{
								SKU:         "123456790",
								Title:       "Aqua 250ml",
								Description: "Air Mineral",
								Category:    "Air Mineral Berkualitas",
								ImgURL:      "https://images.tokopedia.net/img/cache/700/product-1/2018/5/17/22673038/22673038_d74c5460-b815-4ed4-89a4-f0a92cfb2b82_701_701.jpg",
								Weight:      float64(0.70),
								Price:       float64(11000),
							}, nil
						},
					},
				},
			},
			wantResult: mProduct.Product{
				SKU:         "123456790",
				Title:       "Aqua 250ml",
				Description: "Air Mineral",
				Category:    "Air Mineral Berkualitas",
				ImgURL:      "https://images.tokopedia.net/img/cache/700/product-1/2018/5/17/22673038/22673038_d74c5460-b815-4ed4-89a4-f0a92cfb2b82_701_701.jpg",
				Weight:      float64(0.70),
				Price:       float64(11000),
			},
		},
		{
			name: "case-2: error, postgresql got issue",
			args: args{
				ctx: ctx,
				sku: "123456790",
			},
			fields: fields{
				repo: RepoAttribute{
					ProductPostgre: &product.RepositoryMock{
						GetProductBySKUFunc: func(ctx context.Context, sku string) (mProduct.Product, error) {
							return mProduct.Product{}, errors.New("something went wrong")
						},
					},
				},
			},
			wantErr:    true,
			wantResult: mProduct.Product{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &ServiceProduct{
				repo: tt.fields.repo,
			}
			gotResult, err := service.GetProductBySKU(tt.args.ctx, tt.args.sku)
			if (err != nil) != tt.wantErr {
				t.Errorf("ServiceProduct.GetProductBySKU() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("ServiceProduct.GetProductBySKU() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestServiceProduct_GetProducts(t *testing.T) {
	type fields struct {
		repo RepoAttribute
	}
	ctx := context.Background()
	type args struct {
		ctx   context.Context
		param mProduct.ParamSearch
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult mProduct.SearchProductResponse
		wantErr    bool
	}{
		{
			name: "case-1: success, exact-search 1 product",
			args: args{
				ctx: ctx,
				param: mProduct.ParamSearch{
					SKU:      "123456790",
					Title:    "Aqua 250ml",
					Category: "Air Mineral Berkualitas",
					PageSize: 20,
				},
			},
			fields: fields{
				repo: RepoAttribute{
					ProductPostgre: &product.RepositoryMock{
						GetProductsByParamFunc: func(ctx context.Context, param mProduct.ParamSearch) ([]mProduct.SearchProductData, error) {
							return []mProduct.SearchProductData{
								{
									SKU:         "123456790",
									Title:       "Aqua 250ml",
									Description: "Air Mineral",
									Category:    "Air Mineral Berkualitas",
									ImgURL:      "https://images.tokopedia.net/img/cache/700/product-1/2018/5/17/22673038/22673038_d74c5460-b815-4ed4-89a4-f0a92cfb2b82_701_701.jpg",
									Weight:      float64(0.70),
									Price:       float64(11000),
									AvgRating:   float64(3.5),
								},
							}, nil
						},
					},
				},
			},
			wantResult: mProduct.SearchProductResponse{
				Data: []mProduct.SearchProductData{
					{
						SKU:         "123456790",
						Title:       "Aqua 250ml",
						Description: "Air Mineral",
						Category:    "Air Mineral Berkualitas",
						ImgURL:      "https://images.tokopedia.net/img/cache/700/product-1/2018/5/17/22673038/22673038_d74c5460-b815-4ed4-89a4-f0a92cfb2b82_701_701.jpg",
						Weight:      float64(0.70),
						Price:       float64(11000),
						AvgRating:   float64(3.5),
					},
				},
				TotalPage: 1,
				TotalData: 1,
			},
		},
		{
			name: "case-2: success, sort rating active",
			args: args{
				ctx: ctx,
				param: mProduct.ParamSearch{
					Category:  "Air Mineral Berkualitas",
					SortBy:    "rating",
					Direction: "desc",
					PageSize:  20,
				},
			},
			fields: fields{
				repo: RepoAttribute{
					ProductPostgre: &product.RepositoryMock{
						GetProductsByParamFunc: func(ctx context.Context, param mProduct.ParamSearch) ([]mProduct.SearchProductData, error) {
							return []mProduct.SearchProductData{
								{
									SKU:         "123456790",
									Title:       "Aqua 250ml",
									Description: "Air Mineral Berkualitas",
									Category:    "Air Mineral",
									ImgURL:      "https://images.tokopedia.net/img/cache/700/product-1/2018/5/17/22673038/22673038_d74c5460-b815-4ed4-89a4-f0a92cfb2b82_701_701.jpg",
									Weight:      float64(0.70),
									Price:       float64(11000),
									AvgRating:   float64(3.5),
								},
								{
									SKU:         "123456789",
									Title:       "Pristine PH 8.6+",
									Description: "Air Mineral Berkualitas",
									Category:    "Air Mineral",
									ImgURL:      "https://www.static-src.com/wcsstore/Indraprastha/images/catalog/full//125/MTA-49969979/pristine_pristine-water-600-ml_full01.jpg",
									Weight:      float64(0.60),
									Price:       float64(10000),
									AvgRating:   float64(1.5),
								},
							}, nil
						},
					},
				},
			},
			wantResult: mProduct.SearchProductResponse{
				Data: []mProduct.SearchProductData{
					{
						SKU:         "123456790",
						Title:       "Aqua 250ml",
						Description: "Air Mineral Berkualitas",
						Category:    "Air Mineral",
						ImgURL:      "https://images.tokopedia.net/img/cache/700/product-1/2018/5/17/22673038/22673038_d74c5460-b815-4ed4-89a4-f0a92cfb2b82_701_701.jpg",
						Weight:      float64(0.70),
						Price:       float64(11000),
						AvgRating:   float64(3.5),
					},
					{
						SKU:         "123456789",
						Title:       "Pristine PH 8.6+",
						Description: "Air Mineral Berkualitas",
						Category:    "Air Mineral",
						ImgURL:      "https://www.static-src.com/wcsstore/Indraprastha/images/catalog/full//125/MTA-49969979/pristine_pristine-water-600-ml_full01.jpg",
						Weight:      float64(0.60),
						Price:       float64(10000),
						AvgRating:   float64(1.5),
					},
				},
				TotalPage: 1,
				TotalData: 2,
			},
		},
		{
			name: "case-3: success, no record found",
			args: args{
				ctx: ctx,
				param: mProduct.ParamSearch{
					Category:  "Air Mineral Berkualitas",
					SortBy:    "rating",
					Direction: "desc",
					PageSize:  20,
				},
			},
			fields: fields{
				repo: RepoAttribute{
					ProductPostgre: &product.RepositoryMock{
						GetProductsByParamFunc: func(ctx context.Context, param mProduct.ParamSearch) ([]mProduct.SearchProductData, error) {
							return []mProduct.SearchProductData{}, nil
						},
					},
				},
			},
			wantResult: mProduct.SearchProductResponse{
				Data:      []mProduct.SearchProductData{},
				TotalPage: 0,
				TotalData: 0,
			},
		},
		{
			name: "case-4: error, postgres query issue",
			args: args{
				ctx: ctx,
				param: mProduct.ParamSearch{
					SKU:      "123456790",
					Title:    "Aqua 250ml",
					Category: "Air Mineral Berkualitas",
				},
			},
			fields: fields{
				repo: RepoAttribute{
					ProductPostgre: &product.RepositoryMock{
						GetProductsByParamFunc: func(ctx context.Context, param mProduct.ParamSearch) ([]mProduct.SearchProductData, error) {
							return []mProduct.SearchProductData{}, errors.New("something went wrong")
						},
					},
				},
			},
			wantResult: mProduct.SearchProductResponse{},
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &ServiceProduct{
				repo: tt.fields.repo,
			}
			gotResult, err := service.GetProducts(tt.args.ctx, tt.args.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("ServiceProduct.GetProducts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("ServiceProduct.GetProducts() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
