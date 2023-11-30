package postgre

import (
	"database/sql"
	"log"
)

func (repo *ProductRepo) prepareStatements() {
	repo.statement = make(map[string]*sql.Stmt)
	repo.prepareProduct()
}

func (repo *ProductRepo) prepareProduct() {
	var (
		err             error
		createProduct   *sql.Stmt
		getProductBySKU *sql.Stmt
		getProducts     *sql.Stmt
		dbRead          = repo.db.TelkomRead
		dbWrite         = repo.db.TelkomWrite
	)

	if getProductBySKU, err = dbRead.Prepare(queryGetProductBySKU); err != nil {
		log.Panic("[prepareGetProductBySKU] err:", err)
	}
	if getProducts, err = dbRead.Prepare(queryGetProducts); err != nil {
		log.Panic("[prepareGetProducts] err:", err)
	}
	if createProduct, err = dbWrite.Prepare(queryCreateProduct); err != nil {
		log.Panic("[prepareCreateProduct] err:", err)
	}

	repo.statement[useStatementGetProductBySKU] = getProductBySKU
	repo.statement[useStatementGetProducts] = getProducts
	repo.statement[useStatementCreateProduct] = createProduct
}
