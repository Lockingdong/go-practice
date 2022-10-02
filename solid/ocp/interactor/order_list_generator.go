package interactor

import (
	"github.com/jinzhu/copier"
)

// generator query
type GeneratorQuery struct {
	ID          string
	ProductName string
	Status      string
}

// generator response
type GeneratorOrderResponse struct {
	ID          string
	ProductName string
	Status      string
	Amount      int
}

// generator interface
type IOrderListGenerator interface {
	List(req *GeneratorQuery) []*GeneratorOrderResponse
}

// ^^^^^^^

func NewOrderListGenerator(repo IOrderRepo) IOrderListGenerator {
	return &OrderListGenerator{
		repo: repo,
	}
}

type OrderListGenerator struct {
	repo IOrderRepo
}

func (l *OrderListGenerator) List(q *GeneratorQuery) []*GeneratorOrderResponse {

	responseList := make([]*GeneratorOrderResponse, 0)
	resultList := l.repo.List(q)

	_ = copier.Copy(&responseList, &resultList)

	return responseList
}
