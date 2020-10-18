package handlers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kaisersuzaku/BE_A/handlers"
	"github.com/kaisersuzaku/BE_A/models"
	"github.com/kaisersuzaku/BE_A/services/mocks"
	"github.com/stretchr/testify/assert"
)

func TestOrderProduct(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		testName string
		expected func() (int, string)
		prepare  func() (*http.Request, *httptest.ResponseRecorder)
		prepMock func() *mocks.IOrderProductService
	}{
		{
			"TestOrderProduct : Status OK",
			func() (int, string) {
				resp := models.OrderProductResp{
					Status: "OK",
				}
				b, _ := json.Marshal(resp)
				return http.StatusOK, string(b)
			},
			func() (*http.Request, *httptest.ResponseRecorder) {
				req := models.OrderProductReq{
					ID:  1,
					Qty: 2,
				}
				b, _ := json.Marshal(req)
				br := ioutil.NopCloser(bytes.NewReader(b))
				r, _ := http.NewRequest(http.MethodPost, "/order-product", br)
				return r, httptest.NewRecorder()
			},
			func() *mocks.IOrderProductService {
				ops := mocks.IOrderProductService{}
				req := models.OrderProductReq{
					ID:  1,
					Qty: 2,
				}
				resp := models.OrderProductResp{
					Status: "OK",
				}
				ops.On("OrderProduct", context.Background(), req).Return(resp, models.RespError{})
				return &ops
			},
		},
		{
			"TestOrderProduct : Invalid Payload - Missing Field",
			func() (int, string) {
				resp := models.GetInvalidPayloadResp()
				b, _ := json.Marshal(resp)
				return http.StatusBadRequest, string(b)
			},
			func() (*http.Request, *httptest.ResponseRecorder) {
				req := models.OrderProductReq{
					ID: 1,
				}
				b, _ := json.Marshal(req)
				br := ioutil.NopCloser(bytes.NewReader(b))
				r, _ := http.NewRequest(http.MethodPost, "/order-product", br)
				return r, httptest.NewRecorder()
			},
			func() *mocks.IOrderProductService {
				ops := mocks.IOrderProductService{}
				return &ops
			},
		},
		{
			"TestOrderProduct : Invalid Payload - Invalid Content Type",
			func() (int, string) {
				resp := models.GetInvalidPayloadResp()
				b, _ := json.Marshal(resp)
				return http.StatusBadRequest, string(b)
			},
			func() (*http.Request, *httptest.ResponseRecorder) {
				b := []byte("Invalid Payload")
				br := ioutil.NopCloser(bytes.NewReader(b))

				r, _ := http.NewRequest(http.MethodPost, "/order-product", br)
				return r, httptest.NewRecorder()
			},
			func() *mocks.IOrderProductService {
				ops := mocks.IOrderProductService{}
				return &ops
			},
		},
		{
			"TestOrderProduct : Product Unavailable",
			func() (int, string) {
				resp := models.GetProductUnavailable()
				b, _ := json.Marshal(resp)
				return http.StatusBadRequest, string(b)
			},
			func() (*http.Request, *httptest.ResponseRecorder) {
				req := models.OrderProductReq{
					ID:  1,
					Qty: 2,
				}
				b, _ := json.Marshal(req)
				br := ioutil.NopCloser(bytes.NewReader(b))
				r, _ := http.NewRequest(http.MethodPost, "/order-product", br)
				return r, httptest.NewRecorder()
			},
			func() *mocks.IOrderProductService {
				ops := mocks.IOrderProductService{}
				req := models.OrderProductReq{
					ID:  1,
					Qty: 2,
				}
				resp := models.OrderProductResp{}
				ops.On("OrderProduct", context.Background(), req).Return(resp, models.GetProductUnavailable())
				return &ops
			},
		},
	}

	for _, test := range tests {

		ops := test.prepMock()
		oph := handlers.BuildOrderProductHandler(ops)

		req, rr := test.prepare()

		handler := http.HandlerFunc(oph.OrderProduct)
		handler.ServeHTTP(rr, req)

		httpStatus, bodyString := test.expected()
		assert.Equal(httpStatus, rr.Code, fmt.Sprintf("%s : expected http status %d, got %v", test.testName, httpStatus, rr.Code))
		assert.Equal(bodyString, rr.Body.String(), fmt.Sprintf("%s : expected body %s, got %v", test.testName, bodyString, rr.Body.String()))
	}
}
