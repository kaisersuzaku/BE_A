package services

import (
	"context"
	"log"

	"github.com/kaisersuzaku/BE_A/models"
	"github.com/kaisersuzaku/BE_A/repo"
)

type OrderProductService struct {
	pr repo.IProductRepo
}

func BuildOrderProductService(pr repo.IProductRepo) OrderProductService {
	return OrderProductService{
		pr: pr,
	}
}

type IOrderProductService interface {
	OrderProduct(ctx context.Context, req models.OrderProductReq) (resp models.OrderProductResp, err models.RespError)
	GetProductByID(ctx context.Context, id uint) (resp models.Product)
}

func (ops OrderProductService) GetProductByID(ctx context.Context, id uint) (resp models.Product) {
	ops.pr.Read(context.Background(), id, &resp)
	return
}

func (ops OrderProductService) OrderProduct(ctx context.Context, req models.OrderProductReq) (resp models.OrderProductResp, err models.RespError) {
	var product models.Product
	tx := ops.pr.Begin(context.Background())
	e := ops.pr.ReadForUpdateByID(context.Background(), tx, uint(req.ID), &product)
	if e != nil {
		ops.pr.Rollback(context.Background(), tx)
		log.Println(e)
		err = models.GetRequstFailed()
		return
	}

	if product.Stock == 0 {
		ops.pr.Rollback(context.Background(), tx)
		log.Println("Product stock is zero")
		err = models.GetProductUnavailable()
		return
	}

	if product.Stock < req.Qty {
		ops.pr.Rollback(context.Background(), tx)
		err = models.GetStockLessThanRequest()
		log.Println(err.ErrorMessage)
		return
	}
	product.Stock = product.Stock - req.Qty
	e = ops.pr.Update(context.Background(), tx, product)
	if e != nil {
		ops.pr.Rollback(context.Background(), tx)
		err = models.GetRequstFailed()
		log.Println(e)
		return
	}
	e = ops.pr.Commit(context.Background(), tx)
	if e != nil {
		err = models.GetRequstFailed()
		log.Println(e)
		return
	}

	resp.Status = "OK"
	return
}
