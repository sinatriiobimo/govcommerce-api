package v0

import (
	"database/sql"
	rProductReview "tlkm-api/internal/repository/productreview"
)

type (
	ServiceProductReview struct {
		repo     RepoAttribute
		dbDriver *sql.DB
	}

	RepoAttribute struct {
		ProductReviewPostgre rProductReview.Repository
	}

	InitAttribute struct {
		Repo RepoAttribute
	}
)
