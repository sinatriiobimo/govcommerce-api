package productreview

import "time"

type (
	ProductReview struct {
		ProductSKU    string `json:"productSKU"`
		Rating        int64  `json:"rating"`
		ReviewComment string `json:"reviewComment"`
	}

	ProductReviewResponse struct {
		ID            int64     `json:"id"`
		ProductSKU    string    `json:"productSKU"`
		Rating        int64     `json:"rating"`
		ReviewComment string    `json:"reviewComment"`
		ReviewDate    time.Time `json:"reviewDate"`
	}
)
