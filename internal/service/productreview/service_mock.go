// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package productreview

import (
	"context"
	"sync"
	"tlkm-api/internal/model/productreview"
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
//			CreateProductReviewFunc: func(ctx context.Context, req productreview.ProductReview) (productreview.ProductReviewResponse, error) {
//				panic("mock out the CreateProductReview method")
//			},
//		}
//
//		// use mockedService in code that requires Service
//		// and then make assertions.
//
//	}
type ServiceMock struct {
	// CreateProductReviewFunc mocks the CreateProductReview method.
	CreateProductReviewFunc func(ctx context.Context, req productreview.ProductReview) (productreview.ProductReviewResponse, error)

	// calls tracks calls to the methods.
	calls struct {
		// CreateProductReview holds details about calls to the CreateProductReview method.
		CreateProductReview []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Req is the req argument value.
			Req productreview.ProductReview
		}
	}
	lockCreateProductReview sync.RWMutex
}

// CreateProductReview calls CreateProductReviewFunc.
func (mock *ServiceMock) CreateProductReview(ctx context.Context, req productreview.ProductReview) (productreview.ProductReviewResponse, error) {
	if mock.CreateProductReviewFunc == nil {
		panic("ServiceMock.CreateProductReviewFunc: method is nil but Service.CreateProductReview was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Req productreview.ProductReview
	}{
		Ctx: ctx,
		Req: req,
	}
	mock.lockCreateProductReview.Lock()
	mock.calls.CreateProductReview = append(mock.calls.CreateProductReview, callInfo)
	mock.lockCreateProductReview.Unlock()
	return mock.CreateProductReviewFunc(ctx, req)
}

// CreateProductReviewCalls gets all the calls that were made to CreateProductReview.
// Check the length with:
//
//	len(mockedService.CreateProductReviewCalls())
func (mock *ServiceMock) CreateProductReviewCalls() []struct {
	Ctx context.Context
	Req productreview.ProductReview
} {
	var calls []struct {
		Ctx context.Context
		Req productreview.ProductReview
	}
	mock.lockCreateProductReview.RLock()
	calls = mock.calls.CreateProductReview
	mock.lockCreateProductReview.RUnlock()
	return calls
}
