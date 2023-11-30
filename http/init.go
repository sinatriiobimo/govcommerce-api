package main

import (
	"net/http"
	config "tlkm-api/configs"
)

func ProductDlv(s *Server) {
	cfg := config.Get()
	srv := InitHttp(cfg)

	dlv := NewDlvProduct(ProductInitAttribute{
		ProductService: srv.ServiceProduct,
	})
	s.AddRoute(http.MethodPost, "/api/v1/product", dlv.CreateProduct)
	s.AddRoute(http.MethodPatch, "/api/v1/product/:sku", dlv.UpdateProduct)
	s.AddRoute(http.MethodGet, "/api/v1/product/:sku", dlv.GetProductBySKU)
	s.AddRoute(http.MethodGet, "/api/v1/product", dlv.GetProducts)
}

func ProductReviewDlv(s *Server) {
	cfg := config.Get()
	srv := InitHttp(cfg)

	dlv := NewDlvProductReview(ProductReviewInitAttribute{
		ProductReviewService: srv.ServiceProductReview,
	})
	s.AddRoute(http.MethodPost, "/api/v1/productreview", dlv.CreateProductReview)
}

func NewDelivery(s *Server) {
	ProductDlv(s)
	ProductReviewDlv(s)
}
