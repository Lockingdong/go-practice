package json_presenter

import (
	"encoding/json"

	"github.com/go-practice/solid/ocp/controller"
	"github.com/go-practice/solid/ocp/interactor"
	"github.com/go-practice/solid/ocp/port/repo"
)

type JsonPresenter struct {
	controller controller.IOrderListController
}

func NewJsonPresenter() *JsonPresenter {
	repo := repo.NewOrderInmemRepo()
	generator := interactor.NewOrderListGenerator(repo)
	controller := controller.NewOrderListController(generator)
	return &JsonPresenter{
		controller: controller,
	}
}

func (p *JsonPresenter) Execute(q *controller.OrderListControllerQuery) string {

	result := p.controller.Handle(q)

	bytes, _ := json.Marshal(result)

	return string(bytes)
}
