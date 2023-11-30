package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"tlkm-api/internal/model/productreview"
	service "tlkm-api/internal/service/productreview"
	"tlkm-api/pkg"
)

type (
	DeliveryProductReview struct {
		productReviewService service.Service
	}
	ProductReviewInitAttribute struct {
		ProductReviewService service.Service
	}
)

func NewDlvProductReview(init ProductReviewInitAttribute) DeliveryProductReview {
	if err := init.validate(); err != nil {
		log.Panicf("[DeliveryProductReview] invalid init %s, given attribute:%+v", err, init)
	}

	return DeliveryProductReview{
		productReviewService: init.ProductReviewService,
	}
}

func (init ProductReviewInitAttribute) validate() error {
	if init.ProductReviewService == nil {
		return errors.New("missing ProductReviewService")
	}

	return nil
}

func (dc DeliveryProductReview) CreateProductReview(w http.ResponseWriter, r *http.Request) {
	var (
		resp pkg.Response
		data productreview.ProductReviewResponse
		req  productreview.ProductReview
		err  error
	)

	start := time.Now()
	body, err := ioutil.ReadAll(r.Body)
	ctx := r.Context()

	defer r.Body.Close()
	if err != nil {
		fmt.Printf("Read body createProductReview request failed, err : %v\n", err)
		resp.Header.Status = http.StatusInternalServerError
	}
	err = json.Unmarshal(body, &req)

	if err != nil {
		fmt.Printf("Unmarshal body createProductReview request failed, err : %v\n", err)
		resp.Header.Status = http.StatusInternalServerError
	}

	data, err = dc.productReviewService.CreateProductReview(ctx, req)
	if err != nil {
		resp.Error = pkg.ErrorAttribute{
			Status:  false,
			Message: "createProductReview: SKU invalid",
			Code:    http.StatusInternalServerError,
		}
		resp.Header.Status = http.StatusInternalServerError
		resp.Header.ProcessTime = float64(time.Since(start)) / float64(time.Millisecond)
		resp.Render(w, r)
		return
	}

	resp.Header.Status = http.StatusOK
	resp.Header.ProcessTime = float64(time.Since(start)) / float64(time.Millisecond)
	resp.Data = data
	resp.Render(w, r)
}
