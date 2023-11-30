package v0

import (
	"fmt"
	"log"
	"tlkm-api/internal/service/product"
)

func New(attr InitAttribute) product.Service {
	if err := attr.validate(); err != nil {
		log.Panic(err)
	}

	new := &ServiceProduct{
		repo: attr.Repo,
	}
	return new
}

func (init InitAttribute) validate() error {
	if !init.Repo.validate() {
		return fmt.Errorf("missing repositories:%+v", init.Repo)
	}

	return nil
}

func (repos RepoAttribute) validate() bool {
	if repos.ProductPostgre == nil && repos.ProductReviewPostgre == nil {
		return false
	}
	return true
}
