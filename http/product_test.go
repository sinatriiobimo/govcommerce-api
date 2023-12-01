package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	mProduct "tlkm-api/internal/model/product"
	sProduct "tlkm-api/internal/service/product"
	"tlkm-api/pkg"
)

func TestDeliveryProduct_CreateProduct(t *testing.T) {
	type fields struct {
		service sProduct.Service
	}

	type args struct {
		payload mProduct.Product
		rr      *httptest.ResponseRecorder
	}
	tests := []struct {
		name       string
		args       args
		fields     fields
		wantResult mProduct.Product
		wantErr    bool
	}{
		{
			name: "success: create success, happy flow",
			args: args{
				payload: mProduct.Product{
					SKU:         "123456790",
					Title:       "Aqua 250ml",
					Description: "Air Mineral",
					Category:    "Air Mineral Berkualitas",
					ImgURL:      "https://images.tokopedia.net/img/cache/700/product-1/2018/5/17/22673038/22673038_d74c5460-b815-4ed4-89a4-f0a92cfb2b82_701_701.jpg",
					Weight:      float64(0.70),
					Price:       float64(11000),
				},
				rr: httptest.NewRecorder(),
			},
			fields: fields{
				service: &sProduct.ServiceMock{
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
			wantResult: mProduct.Product{
				SKU:         "123456790",
				Title:       "Aqua 250ml",
				Description: "Air Mineral",
				Category:    "Air Mineral Berkualitas",
				ImgURL:      "https://images.tokopedia.net/img/cache/700/product-1/2018/5/17/22673038/22673038_d74c5460-b815-4ed4-89a4-f0a92cfb2b82_701_701.jpg",
				Weight:      float64(0.70),
				Price:       float64(11000),
			},
			wantErr: false,
		},
		{
			name: "error: if one of param empty",
			args: args{
				payload: mProduct.Product{
					Title:       "Aqua 250ml",
					Description: "Air Mineral",
					Category:    "Air Mineral Berkualitas",
					ImgURL:      "https://images.tokopedia.net/img/cache/700/product-1/2018/5/17/22673038/22673038_d74c5460-b815-4ed4-89a4-f0a92cfb2b82_701_701.jpg",
					Weight:      float64(0.70),
					Price:       float64(11000),
				},
				rr: httptest.NewRecorder(),
			},
			fields: fields{
				service: &sProduct.ServiceMock{
					CreateProductFunc: func(ctx context.Context, req mProduct.Product) (mProduct.Product, error) {
						return mProduct.Product{}, errors.New("something went wrong")
					},
				},
			},
			wantResult: mProduct.Product{},
			wantErr:    true,
		},
		{
			name: "error: service got issue",
			args: args{
				payload: mProduct.Product{
					SKU:         "123456790",
					Title:       "Aqua 250ml",
					Description: "Air Mineral",
					Category:    "Air Mineral Berkualitas",
					ImgURL:      "https://images.tokopedia.net/img/cache/700/product-1/2018/5/17/22673038/22673038_d74c5460-b815-4ed4-89a4-f0a92cfb2b82_701_701.jpg",
					Weight:      float64(0.70),
					Price:       float64(11000),
				},
				rr: httptest.NewRecorder(),
			},
			fields: fields{
				service: &sProduct.ServiceMock{
					CreateProductFunc: func(ctx context.Context, req mProduct.Product) (mProduct.Product, error) {
						return mProduct.Product{}, errors.New("something went wrong")
					},
				},
			},
			wantResult: mProduct.Product{},
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			delivery := DeliveryProduct{
				productService: tt.fields.service,
			}

			reqBody, err := json.Marshal(tt.args.payload)
			if err != nil {
				fmt.Printf("Failed to marshal CreateProduct: %v", err)
				return
			}

			req, err := http.NewRequest("POST", "/api/v1/product", bytes.NewBuffer(reqBody))
			assert.NoError(t, err)

			delivery.CreateProduct(tt.args.rr, req)
			if tt.wantErr == true {
				assert.Equal(t, http.StatusInternalServerError, tt.args.rr.Code)
			} else {
				assert.Equal(t, http.StatusOK, tt.args.rr.Code)
			}

			bodyResp, err := ioutil.ReadAll(tt.args.rr.Body)
			assert.NoError(t, err)

			var respData pkg.Response
			err = json.Unmarshal(bodyResp, &respData)
			assert.NoError(t, err)

			dataBytes, err := json.Marshal(respData.Data)
			assert.NoError(t, err)

			var responses mProduct.Product
			err = json.Unmarshal(dataBytes, &responses)
			assert.NoError(t, err)

			if !reflect.DeepEqual(responses, tt.wantResult) {
				t.Errorf("DeliveryProduct.CreateProduct() = %+v, want %+v", responses, tt.wantResult)
			}
		})
	}
}

func TestDeliveryProduct_UpdateProduct(t *testing.T) {
	type fields struct {
		service sProduct.Service
	}

	type args struct {
		sku     string
		payload mProduct.UpdateProductRequest
		rr      *httptest.ResponseRecorder
	}
	tests := []struct {
		name       string
		args       args
		fields     fields
		wantResult mProduct.Product
		wantErr    bool
	}{
		{
			name: "success: update success, happy flow",
			args: args{
				sku: "123456791",
				payload: mProduct.UpdateProductRequest{
					SKU:         "123456790",
					Title:       "Aqua 250ml",
					Description: "Air Mineral",
					Category:    "Air Mineral Berkualitas",
				},
				rr: httptest.NewRecorder(),
			},
			fields: fields{
				service: &sProduct.ServiceMock{
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
			wantResult: mProduct.Product{
				SKU:         "123456790",
				Title:       "Aqua 250ml",
				Description: "Air Mineral",
				Category:    "Air Mineral Berkualitas",
				ImgURL:      "https://images.tokopedia.net/img/cache/700/product-1/2018/5/17/22673038/22673038_d74c5460-b815-4ed4-89a4-f0a92cfb2b82_701_701.jpg",
				Weight:      float64(0.70),
				Price:       float64(11000),
			},
			wantErr: false,
		},
		{
			name: "error: service got issue",
			args: args{
				sku: "123456791",
				payload: mProduct.UpdateProductRequest{
					SKU:         "123456790",
					Title:       "Aqua 250ml",
					Description: "Air Mineral",
					Category:    "Air Mineral Berkualitas",
				},
				rr: httptest.NewRecorder(),
			},
			fields: fields{
				service: &sProduct.ServiceMock{
					UpdateProductFunc: func(ctx context.Context, sku string, req mProduct.UpdateProductRequest) (mProduct.Product, error) {
						return mProduct.Product{}, errors.New("something went wrong")
					},
				},
			},
			wantResult: mProduct.Product{},
			wantErr:    true,
		},
		{
			name: "error: if sku empty",
			args: args{
				sku: "",
				payload: mProduct.UpdateProductRequest{
					SKU:         "123456790",
					Title:       "Aqua 250ml",
					Description: "Air Mineral",
					Category:    "Air Mineral Berkualitas",
				},
				rr: httptest.NewRecorder(),
			},
			fields: fields{
				service: &sProduct.ServiceMock{
					UpdateProductFunc: func(ctx context.Context, sku string, req mProduct.UpdateProductRequest) (mProduct.Product, error) {
						return mProduct.Product{}, errors.New("something went wrong")
					},
				},
			},
			wantResult: mProduct.Product{},
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			delivery := DeliveryProduct{
				productService: tt.fields.service,
			}

			reqBody, err := json.Marshal(tt.args.payload)
			if err != nil {
				fmt.Printf("Failed to marshal UpdateProduct: %v", err)
				return
			}

			req, err := http.NewRequest("PATCH", fmt.Sprintf("/api/v1/product/%v", tt.args.sku), bytes.NewBuffer(reqBody))
			assert.NoError(t, err)

			delivery.UpdateProduct(tt.args.rr, req)
			if tt.wantErr == true {
				assert.Equal(t, http.StatusInternalServerError, tt.args.rr.Code)
			} else {
				assert.Equal(t, http.StatusOK, tt.args.rr.Code)
			}

			bodyResp, err := ioutil.ReadAll(tt.args.rr.Body)
			assert.NoError(t, err)

			var respData pkg.Response
			err = json.Unmarshal(bodyResp, &respData)
			assert.NoError(t, err)

			dataBytes, err := json.Marshal(respData.Data)
			assert.NoError(t, err)

			var responses mProduct.Product
			err = json.Unmarshal(dataBytes, &responses)
			assert.NoError(t, err)

			if !reflect.DeepEqual(responses, tt.wantResult) {
				t.Errorf("DeliveryProduct.UpdateProduct() = %+v, want %+v", responses, tt.wantResult)
			}
		})
	}
}

func TestDeliveryProduct_GetProductBySKU(t *testing.T) {
	type fields struct {
		service sProduct.Service
	}

	type args struct {
		sku string
		rr  *httptest.ResponseRecorder
	}
	tests := []struct {
		name       string
		args       args
		fields     fields
		wantResult mProduct.Product
		wantErr    bool
	}{
		{
			name: "success: get by SKU success, happy flow",
			args: args{
				sku: "123456791",
				rr:  httptest.NewRecorder(),
			},
			fields: fields{
				service: &sProduct.ServiceMock{
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
			wantResult: mProduct.Product{
				SKU:         "123456790",
				Title:       "Aqua 250ml",
				Description: "Air Mineral",
				Category:    "Air Mineral Berkualitas",
				ImgURL:      "https://images.tokopedia.net/img/cache/700/product-1/2018/5/17/22673038/22673038_d74c5460-b815-4ed4-89a4-f0a92cfb2b82_701_701.jpg",
				Weight:      float64(0.70),
				Price:       float64(11000),
			},
			wantErr: false,
		},
		{
			name: "error: service got issue",
			args: args{
				sku: "123456791",
				rr:  httptest.NewRecorder(),
			},
			fields: fields{
				service: &sProduct.ServiceMock{
					GetProductBySKUFunc: func(ctx context.Context, sku string) (mProduct.Product, error) {
						return mProduct.Product{}, errors.New("something went wrong")
					},
				},
			},
			wantResult: mProduct.Product{},
			wantErr:    true,
		},
		{
			name: "error: if sku empty",
			args: args{
				sku: "",
				rr:  httptest.NewRecorder(),
			},
			fields: fields{
				service: &sProduct.ServiceMock{
					GetProductBySKUFunc: func(ctx context.Context, sku string) (mProduct.Product, error) {
						return mProduct.Product{}, errors.New("something went wrong")
					},
				},
			},
			wantResult: mProduct.Product{},
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			delivery := DeliveryProduct{
				productService: tt.fields.service,
			}

			req, err := http.NewRequest("GET", fmt.Sprintf("/api/v1/product/%v", tt.args.sku), bytes.NewBuffer([]byte{}))
			assert.NoError(t, err)

			delivery.GetProductBySKU(tt.args.rr, req)
			if tt.wantErr == true {
				assert.Equal(t, http.StatusInternalServerError, tt.args.rr.Code)
			} else {
				assert.Equal(t, http.StatusOK, tt.args.rr.Code)
			}

			bodyResp, err := ioutil.ReadAll(tt.args.rr.Body)
			assert.NoError(t, err)

			var respData pkg.Response
			err = json.Unmarshal(bodyResp, &respData)
			assert.NoError(t, err)

			dataBytes, err := json.Marshal(respData.Data)
			assert.NoError(t, err)

			var responses mProduct.Product
			err = json.Unmarshal(dataBytes, &responses)
			assert.NoError(t, err)

			if !reflect.DeepEqual(responses, tt.wantResult) {
				t.Errorf("DeliveryProduct.GetProductBySKUt() = %+v, want %+v", responses, tt.wantResult)
			}
		})
	}
}
