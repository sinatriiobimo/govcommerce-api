package v0

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"
	mProductReview "tlkm-api/internal/model/productreview"
	"tlkm-api/internal/repository/productreview"
)

func TestServiceProductReview_CreateProduct(t *testing.T) {
	type fields struct {
		repo RepoAttribute
	}
	ctx := context.Background()
	type args struct {
		ctx context.Context
		req mProductReview.ProductReview
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult mProductReview.ProductReviewResponse
		wantErr    bool
	}{
		{
			name: "case-1: success, happy flow",
			args: args{
				ctx: ctx,
				req: mProductReview.ProductReview{
					ProductSKU:    "123456790",
					Rating:        5,
					ReviewComment: "Sempurna",
				},
			},
			fields: fields{
				repo: RepoAttribute{
					ProductReviewPostgre: &productreview.RepositoryMock{
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
			},
			wantResult: mProductReview.ProductReviewResponse{
				ID:            1,
				ProductSKU:    "123456790",
				Rating:        5,
				ReviewComment: "Sempurna",
				ReviewDate:    time.Date(2023, time.January, 1, 12, 0, 0, 0, time.UTC),
			},
		},
		{
			name: "case-2: error, postgresql got issue",
			args: args{
				ctx: ctx,
				req: mProductReview.ProductReview{
					ProductSKU:    "123456790",
					Rating:        5,
					ReviewComment: "Sempurna",
				},
			},
			fields: fields{
				repo: RepoAttribute{
					ProductReviewPostgre: &productreview.RepositoryMock{
						CreateProductReviewFunc: func(ctx context.Context, req mProductReview.ProductReview) (mProductReview.ProductReviewResponse, error) {
							return mProductReview.ProductReviewResponse{}, errors.New("something went wrong")
						},
					},
				},
			},
			wantErr:    true,
			wantResult: mProductReview.ProductReviewResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &ServiceProductReview{
				repo: tt.fields.repo,
			}
			gotResult, err := service.CreateProductReview(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ServiceProductReview.CreateProductReview(() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("ServiceProductReview.CreateProductReview(() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
