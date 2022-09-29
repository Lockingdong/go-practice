package main

import (
	"fmt"

	"github.com/go-practice/solid/srp/bad_employee"
	"github.com/go-practice/solid/srp/good_employee"
	"github.com/go-practice/solid/srp/good_employee/employee_facade"
)

func main() {

	// bad design
	badEmployee := bad_employee.NewBadEmployee(180, 8)
	pay := badEmployee.CalculatePay()
	hours := badEmployee.ReportHours()
	fmt.Printf("badEmployee CalculatePay: %d, ReportHours: %d\n", pay, hours)

	// good design
	goodEmployee := good_employee.NewGoodEmployee(180, 8)
	employeeFacade := employee_facade.NewEmployeeFacade(goodEmployee)
	pay2 := employeeFacade.PayCalculator.Calculate()
	hours2 := employeeFacade.HourReporter.Report()
	fmt.Printf("goodEmployee CalculatePay: %d, ReportHours: %d\n", pay2, hours2)

}
