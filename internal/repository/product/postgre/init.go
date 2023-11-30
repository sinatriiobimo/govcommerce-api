package postgre

import (
	"fmt"
	"log"
	repo "tlkm-api/internal/repository/product"
)

func New(attr InitAttribute) repo.Repository {
	if err := attr.validate(); err != nil {
		log.Panic(err)
	}

	new := &ProductRepo{
		db: attr.DB,
	}

	new.prepareStatements()
	return new
}

func (init InitAttribute) validate() error {
	if !init.DB.validate() {
		return fmt.Errorf("missing DB driver:%+v", init.DB)
	}

	return nil
}

func (DB DBList) validate() bool {
	if DB.TelkomRead == nil || DB.TelkomWrite == nil {
		return false
	}

	return true
}
