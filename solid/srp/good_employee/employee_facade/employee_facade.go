package employee_facade

import (
	"github.com/go-practice/solid/srp/good_employee"
	"github.com/go-practice/solid/srp/good_employee/employee_saver"
	"github.com/go-practice/solid/srp/good_employee/hour_reporter"
	"github.com/go-practice/solid/srp/good_employee/pay_calculator"
)

type EmployeeFacade struct {
	EmployeeSaver *employee_saver.EmployeeSaver
	HourReporter  *hour_reporter.HourReporter
	PayCalculator *pay_calculator.PayCalculator
}

func NewEmployeeFacade(employee *good_employee.GoodEmployee) *EmployeeFacade {
	return &EmployeeFacade{
		EmployeeSaver: employee_saver.NewEmployeeSaver(employee),
		HourReporter:  hour_reporter.NewHourReporter(employee),
		PayCalculator: pay_calculator.NewPayCalculator(employee),
	}
}
