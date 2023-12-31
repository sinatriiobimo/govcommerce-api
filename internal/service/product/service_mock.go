// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package product

import (
	"context"
	"sync"
	"tlkm-api/internal/model/product"
)

// Ensure, that ServiceMock does implement Service.
// If this is not the case, regenerate this file with moq.
var _ Service = &ServiceMock{}

// ServiceMock is a mock implementation of Service.
//
//	func TestSomethingThatUsesService(t *testing.T) {
//
//		// make and configure a mocked Service
//		mockedService := &ServiceMock{
//			CreateProductFunc: func(ctx context.Context, req product.Product) (product.Product, error) {
//				panic("mock out the CreateProduct method")
//			},
//			GetProductBySKUFunc: func(ctx context.Context, sku string) (product.Product, error) {
//				panic("mock out the GetProductBySKU method")
//			},
//			GetProductsFunc: func(ctx context.Context, param product.ParamSearch) (product.SearchProductResponse, error) {
//				panic("mock out the GetProducts method")
//			},
//			UpdateProductFunc: func(ctx context.Context, sku string, req product.UpdateProductRequest) (product.Product, error) {
//				panic("mock out the UpdateProduct method")
//			},
//		}
//
//		// use mockedService in code that requires Service
//		// and then make assertions.
//
//	}
type ServiceMock struct {
	// CreateProductFunc mocks the CreateProduct method.
	CreateProductFunc func(ctx context.Context, req product.Product) (product.Product, error)

	// GetProductBySKUFunc mocks the GetProductBySKU method.
	GetProductBySKUFunc func(ctx context.Context, sku string) (product.Product, error)

	// GetProductsFunc mocks the GetProducts method.
	GetProductsFunc func(ctx context.Context, param product.ParamSearch) (product.SearchProductResponse, error)

	// UpdateProductFunc mocks the UpdateProduct method.
	UpdateProductFunc func(ctx context.Context, sku string, req product.UpdateProductRequest) (product.Product, error)

	// calls tracks calls to the methods.
	calls struct {
		// CreateProduct holds details about calls to the CreateProduct method.
		CreateProduct []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Req is the req argument value.
			Req product.Product
		}
		// GetProductBySKU holds details about calls to the GetProductBySKU method.
		GetProductBySKU []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Sku is the sku argument value.
			Sku string
		}
		// GetProducts holds details about calls to the GetProducts method.
		GetProducts []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Param is the param argument value.
			Param product.ParamSearch
		}
		// UpdateProduct holds details about calls to the UpdateProduct method.
		UpdateProduct []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Sku is the sku argument value.
			Sku string
			// Req is the req argument value.
			Req product.UpdateProductRequest
		}
	}
	lockCreateProduct   sync.RWMutex
	lockGetProductBySKU sync.RWMutex
	lockGetProducts     sync.RWMutex
	lockUpdateProduct   sync.RWMutex
}

// CreateProduct calls CreateProductFunc.
func (mock *ServiceMock) CreateProduct(ctx context.Context, req product.Product) (product.Product, error) {
	if mock.CreateProductFunc == nil {
		panic("ServiceMock.CreateProductFunc: method is nil but Service.CreateProduct was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Req product.Product
	}{
		Ctx: ctx,
		Req: req,
	}
	mock.lockCreateProduct.Lock()
	mock.calls.CreateProduct = append(mock.calls.CreateProduct, callInfo)
	mock.lockCreateProduct.Unlock()
	return mock.CreateProductFunc(ctx, req)
}

// CreateProductCalls gets all the calls that were made to CreateProduct.
// Check the length with:
//
//	len(mockedService.CreateProductCalls())
func (mock *ServiceMock) CreateProductCalls() []struct {
	Ctx context.Context
	Req product.Product
} {
	var calls []struct {
		Ctx context.Context
		Req product.Product
	}
	mock.lockCreateProduct.RLock()
	calls = mock.calls.CreateProduct
	mock.lockCreateProduct.RUnlock()
	return calls
}

// GetProductBySKU calls GetProductBySKUFunc.
func (mock *ServiceMock) GetProductBySKU(ctx context.Context, sku string) (product.Product, error) {
	if mock.GetProductBySKUFunc == nil {
		panic("ServiceMock.GetProductBySKUFunc: method is nil but Service.GetProductBySKU was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Sku string
	}{
		Ctx: ctx,
		Sku: sku,
	}
	mock.lockGetProductBySKU.Lock()
	mock.calls.GetProductBySKU = append(mock.calls.GetProductBySKU, callInfo)
	mock.lockGetProductBySKU.Unlock()
	return mock.GetProductBySKUFunc(ctx, sku)
}

// GetProductBySKUCalls gets all the calls that were made to GetProductBySKU.
// Check the length with:
//
//	len(mockedService.GetProductBySKUCalls())
func (mock *ServiceMock) GetProductBySKUCalls() []struct {
	Ctx context.Context
	Sku string
} {
	var calls []struct {
		Ctx context.Context
		Sku string
	}
	mock.lockGetProductBySKU.RLock()
	calls = mock.calls.GetProductBySKU
	mock.lockGetProductBySKU.RUnlock()
	return calls
}

// GetProducts calls GetProductsFunc.
func (mock *ServiceMock) GetProducts(ctx context.Context, param product.ParamSearch) (product.SearchProductResponse, error) {
	if mock.GetProductsFunc == nil {
		panic("ServiceMock.GetProductsFunc: method is nil but Service.GetProducts was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Param product.ParamSearch
	}{
		Ctx:   ctx,
		Param: param,
	}
	mock.lockGetProducts.Lock()
	mock.calls.GetProducts = append(mock.calls.GetProducts, callInfo)
	mock.lockGetProducts.Unlock()
	return mock.GetProductsFunc(ctx, param)
}

// GetProductsCalls gets all the calls that were made to GetProducts.
// Check the length with:
//
//	len(mockedService.GetProductsCalls())
func (mock *ServiceMock) GetProductsCalls() []struct {
	Ctx   context.Context
	Param product.ParamSearch
} {
	var calls []struct {
		Ctx   context.Context
		Param product.ParamSearch
	}
	mock.lockGetProducts.RLock()
	calls = mock.calls.GetProducts
	mock.lockGetProducts.RUnlock()
	return calls
}

// UpdateProduct calls UpdateProductFunc.
func (mock *ServiceMock) UpdateProduct(ctx context.Context, sku string, req product.UpdateProductRequest) (product.Product, error) {
	if mock.UpdateProductFunc == nil {
		panic("ServiceMock.UpdateProductFunc: method is nil but Service.UpdateProduct was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Sku string
		Req product.UpdateProductRequest
	}{
		Ctx: ctx,
		Sku: sku,
		Req: req,
	}
	mock.lockUpdateProduct.Lock()
	mock.calls.UpdateProduct = append(mock.calls.UpdateProduct, callInfo)
	mock.lockUpdateProduct.Unlock()
	return mock.UpdateProductFunc(ctx, sku, req)
}

// UpdateProductCalls gets all the calls that were made to UpdateProduct.
// Check the length with:
//
//	len(mockedService.UpdateProductCalls())
func (mock *ServiceMock) UpdateProductCalls() []struct {
	Ctx context.Context
	Sku string
	Req product.UpdateProductRequest
} {
	var calls []struct {
		Ctx context.Context
		Sku string
		Req product.UpdateProductRequest
	}
	mock.lockUpdateProduct.RLock()
	calls = mock.calls.UpdateProduct
	mock.lockUpdateProduct.RUnlock()
	return calls
}
