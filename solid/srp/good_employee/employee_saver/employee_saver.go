package employee_saver

import (
	"github.com/go-practice/solid/srp/good_employee"
)

type EmployeeSaver struct {
	employee *good_employee.GoodEmployee
}

func NewEmployeeSaver(employee *good_employee.GoodEmployee) *EmployeeSaver {
	return &EmployeeSaver{
		employee: employee,
	}
}

func (s *EmployeeSaver) Save() {
	// ... save to DB
}
