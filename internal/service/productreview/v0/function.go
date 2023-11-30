package v0

import (
	"context"
	"tlkm-api/internal/model/productreview"
)

func (service *ServiceProductReview) CreateProductReview(ctx context.Context, req productreview.ProductReview) (productreview.ProductReviewResponse, error) {
	resp, err := service.repo.ProductReviewPostgre.CreateProductReview(ctx, req)
	if err != nil {
		return productreview.ProductReviewResponse{}, err
	}

	return resp, nil
}
