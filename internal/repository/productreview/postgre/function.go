package postgre

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"tlkm-api/internal/model/productreview"
)

func (repo *ProductReviewRepo) CreateProductReview(ctx context.Context, req productreview.ProductReview) (result productreview.ProductReviewResponse, err error) {
	args := []interface{}{
		req.ProductSKU,
		req.Rating,
		req.ReviewComment,
	}

	rows, err := repo.getRowsWithArguments(ctx, useStatementCreateProductReview, args)
	if err != nil {
		fmt.Printf("[ProductReviewRepo.CreateProductReview] getRowsWithArguments: %v", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&result.ID, &result.ProductSKU, &result.Rating, &result.ReviewComment, &result.ReviewDate)
		if err != nil {
			fmt.Printf("[ProductReviewRepo.CreateProductReview] row scan: %v", err)
			continue
		}
	}

	return
}

func (repo *ProductReviewRepo) getRowsWithArguments(ctx context.Context, statementSelector string, args []interface{}) (rows *sql.Rows, err error) {
	var stmt *sql.Stmt

	if stmt = repo.getStatement(statementSelector); stmt == nil {
		err = errors.New("invalid statement")
		fmt.Printf("getStatement %v", err)
		return rows, err
	}

	if rows, err = stmt.Query(args...); err != nil {
		log.Println("[getRowsWithArguments] query:", err)
		fmt.Printf("query %v", err)
		return
	}

	return
}

func (repo *ProductReviewRepo) getStatement(statementSelector string) *sql.Stmt {
	if stmt, ok := repo.statement[statementSelector]; ok {
		return stmt
	}
	return nil
}
