package productreview

import (
	"context"
	"tlkm-api/internal/model/productreview"
)

type Service interface {
	CreateProductReview(ctx context.Context, req productreview.ProductReview) (productreview.ProductReviewResponse, error)
}
