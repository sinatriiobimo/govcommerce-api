package postgre

var (
	queryCreateProduct = `
	INSERT INTO products(
	sku, title, description, category, imgurl, weight, price, created_at, updated_at) 
	VALUES($1,$2,$3,$4,$5,$6,$7,NOW(),NOW())
	RETURNING sku, title, description, category, imgurl, weight, price`

	queryGetProductBySKU = `
	SELECT sku, title, description, category, imgurl, weight, price FROM products 
	WHERE sku=$1`

	queryGetProducts = ``
)
