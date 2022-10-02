package pdf_presenter

import (
	"fmt"

	"github.com/go-practice/solid/ocp/controller"
	"github.com/go-practice/solid/ocp/interactor"
	"github.com/go-practice/solid/ocp/port/repo"
)

type PDFPresenter struct {
	controller controller.IOrderListController
}

func NewPDFPresenter() *PDFPresenter {
	repo := repo.NewOrderInmemRepo()
	generator := interactor.NewOrderListGenerator(repo)
	controller := controller.NewOrderListController(generator)
	return &PDFPresenter{
		controller: controller,
	}
}

func (p *PDFPresenter) Execute(q *controller.OrderListControllerQuery) {

	result := p.controller.Handle(q)
	for _, v := range result {
		fmt.Printf("ID: %s, ProductName: %s, Status: %s, Amount: %d\n", v.ID, v.ProductName, v.Status, v.Amount)
	}
}
