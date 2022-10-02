package main

import (
	"fmt"

	"github.com/go-practice/solid/ocp/controller"
	"github.com/go-practice/solid/ocp/port/json_presenter"
	"github.com/go-practice/solid/ocp/port/pdf_presenter"
)

func main() {

	fmt.Println("=== From PDF Presenter: ===")

	printer := pdf_presenter.NewPDFPresenter()
	printer.Execute(&controller.OrderListControllerQuery{
		Status: "Paid",
	})

	fmt.Println("=== From Json Presenter: ===")
	presenter := json_presenter.NewJsonPresenter()
	jsonStr := presenter.Execute(&controller.OrderListControllerQuery{
		Status: "Paid",
	})
	fmt.Println(jsonStr)

}
