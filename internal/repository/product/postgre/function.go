package postgre

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
	"tlkm-api/internal/model/product"
)

func (repo *ProductRepo) CreateProduct(ctx context.Context, req product.Product) (result product.Product, err error) {
	args := []interface{}{
		req.SKU,
		req.Title,
		req.Description,
		req.Category,
		req.ImgURL,
		req.Weight,
		req.Price,
	}

	rows, err := repo.getRowsWithArguments(ctx, useStatementCreateProduct, args)
	if err != nil {
		fmt.Printf("[ProductRepo.CreateProduct] getRowsWithArguments: %v", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&result.SKU, &result.Title, &result.Category, &result.Description, &result.ImgURL, &result.Weight, &result.Price)
		if err != nil {
			fmt.Printf("[ProductRepo.CreateProduct] row scan: %v", err)
			continue
		}
	}

	return
}

func (repo *ProductRepo) UpdateProduct(ctx context.Context, sku string, req product.UpdateProductRequest) (result product.Product, err error) {
	query := "UPDATE products SET"
	setClauses := make([]string, 0)

	if req.SKU != "" {
		setClauses = append(setClauses, fmt.Sprintf("sku = '%s'", req.SKU))
	}
	if req.Title != "" {
		setClauses = append(setClauses, fmt.Sprintf("title = '%s'", req.Title))
	}

	if req.Description != "" {
		setClauses = append(setClauses, fmt.Sprintf("description = '%s'", req.Description))
	}

	if req.Category != "" {
		setClauses = append(setClauses, fmt.Sprintf("category = '%s'", req.Category))
	}

	if len(setClauses) == 0 {
		return product.Product{}, errors.New("[ProductRepo.UpdateProduct] no fields provided for update")
	}

	query += " " + strings.Join(setClauses, ",")
	query += fmt.Sprintf(", updated_at='NOW()' WHERE sku = '%s' RETURNING sku, title, description, category, imgurl, weight, price;", sku)

	row := repo.db.TelkomWrite.QueryRow(query)
	if err := row.Scan(&result.SKU, &result.Title, &result.Description, &result.Category, &result.ImgURL, &result.Weight, &result.Price); err != nil {
		return product.Product{}, err
	}

	return
}

func (repo *ProductRepo) GetProductBySKU(ctx context.Context, sku string) (result product.Product, err error) {
	rows, err := repo.getRowsWithArguments(ctx, useStatementGetProductBySKU, []interface{}{sku})
	if err != nil {
		return result, err
	}

	for rows.Next() {
		if err = rows.Scan(&result.SKU, &result.Title, &result.Category, &result.Description, &result.ImgURL, &result.Weight, &result.Price); err != nil {
			if err == sql.ErrNoRows {
				return result, errors.New("[ProductRepo.GetProductBySKU] record not found")
			}
			return result, err
		}
	}

	return result, nil
}

func (repo *ProductRepo) GetProducts(ctx context.Context, param product.ParamSearch) ([]product.Product, error) {

	return []product.Product{}, nil
}

func (repo *ProductRepo) GetProductsByParam(ctx context.Context, param product.ParamSearch) (products []product.SearchProductData, total int64, err error) {
	query := `
		SELECT
			p.sku, p.title, p.description, p.category, p.imgurl, p.weight, p.price,
			COALESCE(AVG(pr.rating), 0) AS avg_rating
		FROM
			products p
		LEFT JOIN
			product_reviews pr ON p.sku = pr.product_sku
		WHERE
			1 = 1
	`

	setClauses := make([]string, 0)
	if param.SKU != "" {
		setClauses = append(setClauses, fmt.Sprintf(" AND p.sku = '%s'", param.SKU))
	}
	if param.Title != "" {
		setClauses = append(setClauses, fmt.Sprintf(" AND p.title ILIKE '%s'", fmt.Sprintf("%%%s%%", param.Title)))
	}
	if param.Category != "" {
		setClauses = append(setClauses, fmt.Sprintf(" AND p.category = '%s'", param.Category))
	}

	query += " " + strings.Join(setClauses, "")

	switch strings.ToLower(param.SortBy) {
	case "rating":
		query += " GROUP BY p.sku, p.title, p.description, p.category, p.imgurl, p.weight, p.price"
		query += " ORDER BY avg_rating"
	default:
		query += " GROUP BY p.sku, p.title, p.description, p.category, p.imgurl, p.weight, p.price"
		query += " ORDER BY p.created_at"
	}

	direction := "ASC"
	if strings.ToUpper(param.Direction) == "DESC" {
		direction = "DESC"
	}
	query += fmt.Sprintf(" %s", direction)

	query += fmt.Sprintf(" LIMIT %d OFFSET %d", param.PageSize, param.PageIndex)

	rows, err := repo.db.TelkomWrite.QueryContext(ctx, query)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var p product.SearchProductData
		err := rows.Scan(&p.SKU, &p.Title, &p.Description, &p.Category, &p.ImgURL, &p.Weight, &p.Price, &p.AvgRating)
		if err != nil {
			return nil, 0, err
		}
		products = append(products, p)
	}

	total = int64(len(products))
	return products, total, nil
}

func (repo *ProductRepo) getRowsWithArguments(ctx context.Context, statementSelector string, args []interface{}) (rows *sql.Rows, err error) {
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

func (repo *ProductRepo) getStatement(statementSelector string) *sql.Stmt {
	if stmt, ok := repo.statement[statementSelector]; ok {
		return stmt
	}
	return nil
}
