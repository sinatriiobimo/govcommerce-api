package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"tlkm-api/internal/model/product"
	service "tlkm-api/internal/service/product"
	"tlkm-api/pkg"
)

type (
	DeliveryProduct struct {
		productService service.Service
	}
	ProductInitAttribute struct {
		ProductService service.Service
	}
)

func NewDlvProduct(init ProductInitAttribute) DeliveryProduct {
	if err := init.validate(); err != nil {
		log.Panicf("[DeliveryProduct] invalid init %s, given attribute:%+v", err, init)
	}

	return DeliveryProduct{
		productService: init.ProductService,
	}
}

func (init ProductInitAttribute) validate() error {
	if init.ProductService == nil {
		return errors.New("missing ProductService")
	}

	return nil
}

func (dc DeliveryProduct) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var (
		resp pkg.Response
		data product.Product
		req  product.Product
		err  error
	)

	start := time.Now()
	body, err := ioutil.ReadAll(r.Body)
	ctx := r.Context()

	defer r.Body.Close()
	if err != nil {
		fmt.Printf("Read body createProduct request failed, err : %v\n", err)
		resp.Header.Status = http.StatusInternalServerError
	}
	err = json.Unmarshal(body, &req)

	resp = dc.validateProduct(req)
	if resp.Error.Status == true {
		resp.Render(w, r)
		return
	}

	if err != nil {
		fmt.Printf("Unmarshal body createProduct request failed, err : %v\n", err)
		resp.Header.Status = http.StatusInternalServerError
	}

	data, err = dc.productService.CreateProduct(ctx, req)
	if err != nil {
		resp.Error = pkg.ErrorAttribute{
			Status:  false,
			Message: "createProduct cannot input with similar SKU, SKU must unique",
			Code:    http.StatusConflict,
		}
		resp.Header.Status = http.StatusConflict
		resp.Header.ProcessTime = float64(time.Since(start)) / float64(time.Millisecond)
		resp.Render(w, r)
		return
	}

	resp.Header.Status = http.StatusOK
	resp.Header.ProcessTime = float64(time.Since(start)) / float64(time.Millisecond)
	resp.Data = data
	resp.Render(w, r)
}

func (dc DeliveryProduct) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var (
		resp pkg.Response
		data product.Product
		req  product.UpdateProductRequest
		err  error
	)

	start := time.Now()
	body, err := ioutil.ReadAll(r.Body)
	ctx := r.Context()

	sku := extractPathVariable(r.URL.String())
	if sku == "" {
		resp.Error = pkg.ErrorAttribute{
			Status:  false,
			Message: fmt.Sprintln("updateProduct err: SKU path cannot be empty"),
			Code:    http.StatusInternalServerError,
		}
		resp.Header.Status = http.StatusInternalServerError
		resp.Header.ProcessTime = float64(time.Since(start)) / float64(time.Millisecond)
		resp.Render(w, r)
		return
	}

	defer r.Body.Close()
	if err != nil {
		fmt.Printf("Read body updateProduct request failed, err : %v\n", err)
		resp.Header.Status = http.StatusInternalServerError
	}
	err = json.Unmarshal(body, &req)
	if err != nil {
		fmt.Printf("Unmarshal body updateProduct request failed, err : %v\n", err)
		resp.Header.Status = http.StatusInternalServerError
	}

	data, err = dc.productService.UpdateProduct(ctx, sku, product.UpdateProductRequest{
		SKU:         req.SKU,
		Title:       req.Title,
		Description: req.Description,
		Category:    req.Category,
	})
	if err != nil {
		resp.Error = pkg.ErrorAttribute{
			Status:  false,
			Message: fmt.Sprintf("updateProduct err: %v", err),
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

func (dc DeliveryProduct) GetProductBySKU(w http.ResponseWriter, r *http.Request) {
	var (
		resp pkg.Response
		data product.Product
		req  string
		err  error
	)

	start := time.Now()
	ctx := r.Context()

	req = extractPathVariable(r.URL.String())
	if req == "" {
		resp.Error = pkg.ErrorAttribute{
			Status:  false,
			Message: fmt.Sprintln("updateProduct err: SKU path cannot be empty"),
			Code:    http.StatusInternalServerError,
		}
		resp.Header.Status = http.StatusInternalServerError
		resp.Header.ProcessTime = float64(time.Since(start)) / float64(time.Millisecond)
		resp.Render(w, r)
		return
	}

	data, err = dc.productService.GetProductBySKU(ctx, req)
	if err != nil {
		resp.Error = pkg.ErrorAttribute{
			Status:  false,
			Message: fmt.Sprintf("getProductBySKU err: %v", err),
			Code:    http.StatusInternalServerError,
		}
		resp.Header.Status = http.StatusInternalServerError
		resp.Header.ProcessTime = float64(time.Since(start)) / float64(time.Millisecond)
		resp.Render(w, r)
		return
	}

	if data.SKU == "" {
		resp.Error = pkg.ErrorAttribute{
			Status:  true,
			Message: fmt.Sprintf("SKU %v: record not found", req),
			Code:    http.StatusNoContent,
		}
		resp.Header.Status = http.StatusNoContent
		resp.Header.ProcessTime = float64(time.Since(start)) / float64(time.Millisecond)
		resp.Render(w, r)
		return
	}
	resp.Header.Status = http.StatusOK
	resp.Header.ProcessTime = float64(time.Since(start)) / float64(time.Millisecond)
	resp.Data = data
	resp.Render(w, r)
}

func (dc DeliveryProduct) GetProducts(w http.ResponseWriter, r *http.Request) {
	var (
		resp              pkg.Response
		searchProductResp product.SearchProductResponse
		err               error
	)

	start := time.Now()
	ctx := r.Context()
	resp.Header.Status = http.StatusOK
	param := dc.setParamSearch(r)

	searchProductResp, err = dc.productService.GetProducts(ctx, param)
	if err != nil {
		resp.Header.Status = http.StatusInternalServerError
	}
	resp.Data = searchProductResp.Data
	lenData := len(searchProductResp.Data)
	if lenData == 0 {
		resp.Data = make([]product.Product, 0)
	}
	lastCursor := param.PageIndex * param.PageSize
	if lenData < param.PageSize {
		lastCursor = ((param.PageIndex - 1) * param.PageSize) + lenData
	}
	resp.Pagination.TotalData = searchProductResp.TotalData
	resp.Pagination.LastCursor = lastCursor
	resp.Pagination.Size = lenData
	resp.Pagination.TotalPage = searchProductResp.TotalPage

	if param.Direction != "" && param.SortBy != "" {
		if param.Direction == "" || "desc" == strings.ToLower(param.Direction) || strings.ToLower(param.Direction) != "asc" {
			param.Direction = "desc"
		}
		resp.Pagination.Sort = []pkg.Sort{
			{
				Type:  param.Direction,
				Field: param.SortBy,
			},
		}
	}

	resp.Header.Status = http.StatusOK
	resp.Header.ProcessTime = float64(time.Since(start)) / float64(time.Millisecond)
	resp.Render(w, r)
}

func (dc DeliveryProduct) setParamSearch(r *http.Request) (param product.ParamSearch) {
	var err error

	param.SKU = r.FormValue("sku")
	param.Title = r.FormValue("title")
	param.Category = r.FormValue("category")

	param.PageSize, err = strconv.Atoi(r.FormValue("page_size"))
	if param.PageSize <= 0 || err != nil {
		param.PageSize = 20
	}

	param.PageIndex, err = strconv.Atoi(r.FormValue("page_index"))
	if param.PageIndex <= 0 || err != nil {
		param.PageIndex = 0
	}

	tempSort := r.FormValue("sort")
	splitSort := strings.Split(tempSort, ":")
	var sortID int64
	if tempSort != "" {
		intSortId, err := strconv.ParseInt(splitSort[0], 10, 64)
		if err != nil {
			sortID = 0
		}
		sortID = intSortId
		if len(splitSort) > 1 {
			param.Direction = splitSort[1]
		}
		if param.Direction == "" || strings.ToLower(param.Direction) == "desc" {
			param.Direction = "desc"
		} else if strings.ToLower(param.Direction) == "asc" {
			param.Direction = "asc"
		}
		if sortID == 1 || sortID <= 0 {
			param.SortBy = "product_date"
		}
		if sortID == 2 {
			param.SortBy = "rating"
		}
	}

	return
}

func (dc DeliveryProduct) validateProduct(req product.Product) (resp pkg.Response) {
	if req.SKU == "" {
		resp.Header.Status = http.StatusInternalServerError
		resp.Error = pkg.ErrorAttribute{
			Status:  false,
			Message: "createProduct request failed: due to empty SKU",
			Code:    http.StatusInternalServerError,
		}
		return resp
	}

	if req.Title == "" {
		fmt.Println("createProduct request failed: due to empty title")
		resp.Header.Status = http.StatusInternalServerError
		resp.Error = pkg.ErrorAttribute{
			Status:  false,
			Message: "createProduct request failed: due to empty SKU",
			Code:    http.StatusInternalServerError,
		}
		return resp
	}

	if req.Category == "" {
		resp.Header.Status = http.StatusInternalServerError
		resp.Error = pkg.ErrorAttribute{
			Status:  false,
			Message: "createProduct request failed: due to empty category",
			Code:    http.StatusInternalServerError,
		}
		return resp
	}

	if req.Description == "" {
		resp.Header.Status = http.StatusInternalServerError
		resp.Error = pkg.ErrorAttribute{
			Status:  false,
			Message: "createProduct request failed: due to empty description",
			Code:    http.StatusInternalServerError,
		}
		return resp
	}

	if req.Weight <= 0 {
		resp.Header.Status = http.StatusInternalServerError
		resp.Error = pkg.ErrorAttribute{
			Status:  false,
			Message: "createProduct request failed: due to empty weight",
			Code:    http.StatusInternalServerError,
		}
		return resp
	}

	if req.Price <= 0 {
		resp.Header.Status = http.StatusInternalServerError
		resp.Error = pkg.ErrorAttribute{
			Status:  false,
			Message: "createProduct request failed: due to empty price",
			Code:    http.StatusInternalServerError,
		}
		return resp
	}

	if req.ImgURL == "" {
		resp.Header.Status = http.StatusInternalServerError
		resp.Error = pkg.ErrorAttribute{
			Status:  false,
			Message: "createProduct request failed: due to empty imageURL",
			Code:    http.StatusInternalServerError,
		}
		return resp
	}

	return resp
}

func extractPathVariable(path string) string {
	segments := strings.Split(path, "/")

	for i, segment := range segments {
		if segment == "product" && i+1 < len(segments) {
			return segments[i+1]
		}
	}

	return ""
}
