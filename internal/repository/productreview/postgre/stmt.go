package postgre

import (
	"database/sql"
	"log"
)

func (repo *ProductReviewRepo) prepareStatements() {
	repo.statement = make(map[string]*sql.Stmt)
	repo.prepareProductReview()
}

func (repo *ProductReviewRepo) prepareProductReview() {
	var (
		err                 error
		createProductReview *sql.Stmt
		dbWrite             = repo.db.TelkomWrite
	)

	if createProductReview, err = dbWrite.Prepare(queryCreateProductReview); err != nil {
		log.Panic("[prepareCreateProductReview] err:", err)
	}

	repo.statement[useStatementCreateProductReview] = createProductReview
}
