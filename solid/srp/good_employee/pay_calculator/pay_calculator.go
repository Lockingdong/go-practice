package pay_calculator

import (
	"github.com/go-practice/solid/srp/good_employee"
)

type PayCalculator struct {
	employee *good_employee.GoodEmployee
}

func NewPayCalculator(employee *good_employee.GoodEmployee) *PayCalculator {
	return &PayCalculator{
		employee: employee,
	}
}

func (s *PayCalculator) Calculate() int {
	return s.employee.GetRegularHours() * s.employee.GetWage()
}
