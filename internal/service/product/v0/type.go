package v0

import (
	"database/sql"
	rProduct "tlkm-api/internal/repository/product"
	rProductReview "tlkm-api/internal/repository/productreview"
)

type (
	ServiceProduct struct {
		repo     RepoAttribute
		dbDriver *sql.DB
	}

	RepoAttribute struct {
		ProductPostgre       rProduct.Repository
		ProductReviewPostgre rProductReview.Repository
	}

	InitAttribute struct {
		Repo RepoAttribute
	}
)
