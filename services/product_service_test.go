package services_test

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/kaisersuzaku/BE_A/repo/mocks"
	"gorm.io/gorm"

	"github.com/kaisersuzaku/BE_A/services"

	"github.com/kaisersuzaku/BE_A/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestOrderProduct(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		testName  string
		input1    func() context.Context
		input2    func() models.OrderProductReq
		expected1 func() models.OrderProductResp
		expected2 func() models.RespError
		prepare   func() *mocks.IProductRepo
	}{
		{
			"TestOrderProduct : Status OK",
			func() context.Context {
				return context.TODO()
			},
			func() models.OrderProductReq {
				return models.OrderProductReq{
					ID:  1,
					Qty: 2,
				}
			},
			func() models.OrderProductResp {
				return models.OrderProductResp{
					Status: "OK",
				}
			},
			func() models.RespError {
				return models.RespError{}
			},
			func() *mocks.IProductRepo {
				opr := mocks.IProductRepo{}
				var d gorm.DB
				var p1 models.Product
				var arg models.Product
				arg.ID = uint(1)
				arg.Name = "sarung"
				arg.Stock = uint(299)

				opr.On("Begin", context.Background()).Return(&d)
				opr.On("ReadForUpdateByID", context.Background(), &d, uint(1), &p1).Run(func(args mock.Arguments) {
					arg := args.Get(3).(*models.Product)
					arg.ID = uint(1)
					arg.Name = "sarung"
					arg.Stock = uint(301)
				}).Return(nil)
				opr.On("Update", context.Background(), &d, arg).Return(nil)
				opr.On("Commit", context.Background(), &d).Return(nil)
				return &opr
			},
		},
		{
			"TestOrderProduct : Transaction Error 01",
			func() context.Context {
				return context.TODO()
			},
			func() models.OrderProductReq {
				return models.OrderProductReq{
					ID:  1,
					Qty: 2,
				}
			},
			func() models.OrderProductResp {
				return models.OrderProductResp{}
			},
			func() models.RespError {
				return models.GetRequstFailed()
			},
			func() *mocks.IProductRepo {
				opr := mocks.IProductRepo{}
				var d gorm.DB
				var p1 models.Product
				var arg models.Product
				arg.ID = uint(1)
				arg.Name = "sarung"
				arg.Stock = uint(299)

				opr.On("Begin", context.Background()).Return(&d)
				opr.On("ReadForUpdateByID", context.Background(), &d, uint(1), &p1).Run(func(args mock.Arguments) {
					arg := args.Get(3).(*models.Product)
					arg.ID = uint(1)
					arg.Name = "sarung"
					arg.Stock = uint(301)
				}).Return(nil)
				opr.On("Update", context.Background(), &d, arg).Return(nil)
				opr.On("Commit", context.Background(), &d).Return(sql.ErrConnDone)
				return &opr
			},
		},
		{
			"TestOrderProduct : Transaction Error 02",
			func() context.Context {
				return context.TODO()
			},
			func() models.OrderProductReq {
				return models.OrderProductReq{
					ID:  1,
					Qty: 2,
				}
			},
			func() models.OrderProductResp {
				return models.OrderProductResp{}
			},
			func() models.RespError {
				return models.GetRequstFailed()
			},
			func() *mocks.IProductRepo {
				opr := mocks.IProductRepo{}
				var d gorm.DB
				var p1 models.Product
				var arg models.Product
				arg.ID = uint(1)
				arg.Name = "sarung"
				arg.Stock = uint(299)

				opr.On("Begin", context.Background()).Return(&d)
				opr.On("ReadForUpdateByID", context.Background(), &d, uint(1), &p1).Run(func(args mock.Arguments) {
					arg := args.Get(3).(*models.Product)
					arg.ID = uint(1)
					arg.Name = "sarung"
					arg.Stock = uint(301)
				}).Return(nil)
				opr.On("Update", context.Background(), &d, arg).Return(sql.ErrConnDone)
				opr.On("Rollback", context.Background(), &d).Return(nil)
				return &opr
			},
		},
		{
			"TestOrderProduct : Transaction Error 03",
			func() context.Context {
				return context.TODO()
			},
			func() models.OrderProductReq {
				return models.OrderProductReq{
					ID:  1,
					Qty: 2,
				}
			},
			func() models.OrderProductResp {
				return models.OrderProductResp{}
			},
			func() models.RespError {
				return models.GetRequstFailed()
			},
			func() *mocks.IProductRepo {
				opr := mocks.IProductRepo{}
				var d gorm.DB
				var p1 models.Product
				var arg models.Product
				arg.ID = uint(1)
				arg.Name = "sarung"
				arg.Stock = uint(299)

				opr.On("Begin", context.Background()).Return(&d)
				opr.On("ReadForUpdateByID", context.Background(), &d, uint(1), &p1).Return(sql.ErrConnDone)
				opr.On("Rollback", context.Background(), &d).Return(nil)
				return &opr
			},
		},
		{
			"TestOrderProduct : Transaction Error - Stock < Req Qty",
			func() context.Context {
				return context.TODO()
			},
			func() models.OrderProductReq {
				return models.OrderProductReq{
					ID:  1,
					Qty: 2,
				}
			},
			func() models.OrderProductResp {
				return models.OrderProductResp{}
			},
			func() models.RespError {
				return models.GetStockLessThanRequest()
			},
			func() *mocks.IProductRepo {
				opr := mocks.IProductRepo{}
				var d gorm.DB
				var p1 models.Product
				var arg models.Product
				arg.ID = uint(1)
				arg.Name = "sarung"
				arg.Stock = uint(1)

				opr.On("Begin", context.Background()).Return(&d)
				opr.On("ReadForUpdateByID", context.Background(), &d, uint(1), &p1).Run(func(args mock.Arguments) {
					arg := args.Get(3).(*models.Product)
					arg.ID = uint(1)
					arg.Name = "sarung"
					arg.Stock = uint(1)
				}).Return(nil)
				opr.On("Rollback", context.Background(), &d).Return(nil)
				return &opr
			},
		},
		{
			"TestOrderProduct : Transaction Error - Stock == 0",
			func() context.Context {
				return context.TODO()
			},
			func() models.OrderProductReq {
				return models.OrderProductReq{
					ID:  1,
					Qty: 2,
				}
			},
			func() models.OrderProductResp {
				return models.OrderProductResp{}
			},
			func() models.RespError {
				return models.GetProductUnavailable()
			},
			func() *mocks.IProductRepo {
				opr := mocks.IProductRepo{}
				var d gorm.DB
				var p1 models.Product
				var arg models.Product
				arg.ID = uint(1)
				arg.Name = "sarung"
				arg.Stock = uint(0)

				opr.On("Begin", context.Background()).Return(&d)
				opr.On("ReadForUpdateByID", context.Background(), &d, uint(1), &p1).Run(func(args mock.Arguments) {
					arg := args.Get(3).(*models.Product)
					arg.ID = uint(1)
					arg.Name = "sarung"
					arg.Stock = uint(0)
				}).Return(nil)
				opr.On("Rollback", context.Background(), &d).Return(nil)
				return &opr
			},
		},
	}
	for _, test := range tests {
		opr := test.prepare()
		ops := services.BuildOrderProductService(opr)
		resp, respErr := ops.OrderProduct(test.input1(), test.input2())
		assert.Equal(test.expected1(), resp, fmt.Sprintf("%s : Object not same, expected %v, got %v", test.testName, test.expected1(), resp))
		assert.Equal(test.expected2(), respErr, fmt.Sprintf("%s : Object not same, expected %v, got %v", test.testName, test.expected2(), respErr))
	}
}

func TestGetProductByID(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		testName  string
		input1    func() context.Context
		input2    func() uint
		expected1 func() models.Product
		prepare   func() *mocks.IProductRepo
	}{
		{
			"TestGetProductByID : Found",
			func() context.Context {
				return context.TODO()
			},
			func() uint {
				return uint(1)
			},
			func() models.Product {
				var arg models.Product
				arg.ID = uint(1)
				arg.Name = "sarung"
				arg.Stock = uint(299)
				return arg
			},
			func() *mocks.IProductRepo {
				opr := mocks.IProductRepo{}
				var arg models.Product
				var p1 models.Product
				arg.ID = uint(1)
				arg.Name = "sarung"
				arg.Stock = uint(299)
				opr.On("Read", context.Background(), arg.ID, &p1).Run(func(args mock.Arguments) {
					arg := args.Get(2).(*models.Product)
					arg.ID = uint(1)
					arg.Name = "sarung"
					arg.Stock = uint(299)
				})
				return &opr
			},
		},
	}
	for _, test := range tests {
		opr := test.prepare()
		ops := services.BuildOrderProductService(opr)
		resp := ops.GetProductByID(test.input1(), test.input2())
		assert.Equal(test.expected1(), resp, fmt.Sprintf("%s : Object not same, expected %v, got %v", test.testName, test.expected1(), resp))
	}
}
