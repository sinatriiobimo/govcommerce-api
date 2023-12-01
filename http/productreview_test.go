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
	"time"
	mProductReview "tlkm-api/internal/model/productreview"
	sProductReview "tlkm-api/internal/service/productreview"
	"tlkm-api/pkg"
)

func TestDeliveryProduct_CreateProductReview(t *testing.T) {
	type fields struct {
		service sProductReview.Service
	}

	type args struct {
		payload mProductReview.ProductReview
		rr      *httptest.ResponseRecorder
	}
	tests := []struct {
		name       string
		args       args
		fields     fields
		wantResult mProductReview.ProductReviewResponse
		wantErr    bool
	}{
		{
			name: "success: create success, happy flow",
			args: args{
				payload: mProductReview.ProductReview{
					ProductSKU:    "123456790",
					Rating:        5,
					ReviewComment: "Sempurna",
				},
				rr: httptest.NewRecorder(),
			},
			fields: fields{
				service: &sProductReview.ServiceMock{
					CreateProductReviewFunc: func(ctx context.Context, req mProductReview.ProductReview) (mProductReview.ProductReviewResponse, error) {
						return mProductReview.ProductReviewResponse{
							ID:            1,
							ProductSKU:    "123456790",
							Rating:        5,
							ReviewComment: "Sempurna",
							ReviewDate:    time.Date(2023, time.January, 1, 12, 0, 0, 0, time.UTC),
						}, nil
					},
				},
			},
			wantResult: mProductReview.ProductReviewResponse{
				ID:            1,
				ProductSKU:    "123456790",
				Rating:        5,
				ReviewComment: "Sempurna",
				ReviewDate:    time.Date(2023, time.January, 1, 12, 0, 0, 0, time.UTC),
			},
			wantErr: false,
		},
		{
			name: "error: if one of param empty",
			args: args{
				payload: mProductReview.ProductReview{
					Rating:        5,
					ReviewComment: "Sempurna",
				},
				rr: httptest.NewRecorder(),
			},
			fields: fields{
				service: &sProductReview.ServiceMock{
					CreateProductReviewFunc: func(ctx context.Context, req mProductReview.ProductReview) (mProductReview.ProductReviewResponse, error) {
						return mProductReview.ProductReviewResponse{}, errors.New("something went wrong")
					},
				},
			},
			wantResult: mProductReview.ProductReviewResponse{},
			wantErr:    true,
		},
		{
			name: "error: service got issue",
			args: args{
				payload: mProductReview.ProductReview{
					ProductSKU:    "123456790",
					Rating:        5,
					ReviewComment: "Sempurna",
				},
				rr: httptest.NewRecorder(),
			},
			fields: fields{
				service: &sProductReview.ServiceMock{
					CreateProductReviewFunc: func(ctx context.Context, req mProductReview.ProductReview) (mProductReview.ProductReviewResponse, error) {
						return mProductReview.ProductReviewResponse{}, errors.New("something went wrong")
					},
				},
			},
			wantResult: mProductReview.ProductReviewResponse{},
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			delivery := DeliveryProductReview{
				productReviewService: tt.fields.service,
			}

			reqBody, err := json.Marshal(tt.args.payload)
			if err != nil {
				fmt.Printf("Failed to marshal CreateProductReview: %v", err)
				return
			}

			req, err := http.NewRequest("POST", "/api/v1/productreview", bytes.NewBuffer(reqBody))
			assert.NoError(t, err)

			delivery.CreateProductReview(tt.args.rr, req)
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

			var responses mProductReview.ProductReviewResponse
			err = json.Unmarshal(dataBytes, &responses)
			assert.NoError(t, err)

			if !reflect.DeepEqual(responses, tt.wantResult) {
				t.Errorf("DeliveryProductReview.CreateProductReview() = %+v, want %+v", responses, tt.wantResult)
			}
		})
	}
}
