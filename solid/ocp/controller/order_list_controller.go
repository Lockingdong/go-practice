package controller

import (
	"github.com/go-practice/solid/ocp/interactor"
	"github.com/jinzhu/copier"
)

type OrderListControllerQuery struct {
	ID          string
	ProductName string
	Status      string
}

type OrderListControllerResponse struct {
	ID          string
	ProductName string
	Status      string
	Amount      int
}

type IOrderListController interface {
	Handle(*OrderListControllerQuery) []*OrderListControllerResponse
}

// ^^^^^^

type OrderListController struct {
	orderListGenerator interactor.IOrderListGenerator
}

func NewOrderListController(generator interactor.IOrderListGenerator) IOrderListController {
	return &OrderListController{
		orderListGenerator: generator,
	}
}

func (c *OrderListController) Handle(q *OrderListControllerQuery) []*OrderListControllerResponse {
	responseList := make([]*OrderListControllerResponse, 0)

	resultList := c.orderListGenerator.List(&interactor.GeneratorQuery{
		ID:          q.ID,
		ProductName: q.ProductName,
		Status:      q.Status,
	})

	_ = copier.Copy(&responseList, &resultList)

	return responseList
}
