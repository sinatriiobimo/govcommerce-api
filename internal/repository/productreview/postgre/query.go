package postgre

var (
	queryCreateProductReview = `
	INSERT INTO product_reviews (product_sku, rating, review_comment, review_date, created_at) 
	VALUES ($1, $2, $3, NOW(), NOW())
  	RETURNING id, product_sku, rating, review_comment, review_date;`
)
