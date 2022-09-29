package hour_reporter

import (
	"github.com/go-practice/solid/srp/good_employee"
)

type HourReporter struct {
	employee     *good_employee.GoodEmployee
	regularHours int
}

func NewHourReporter(employee *good_employee.GoodEmployee) *HourReporter {
	return &HourReporter{
		employee: employee,
	}
}

func (s *HourReporter) Report() int {
	return s.employee.GetRegularHours() * 2
}
