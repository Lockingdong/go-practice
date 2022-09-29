package bad_employee

type BadEmployee struct {
	wage         int
	regularHours int
}

func NewBadEmployee(wage, regularHours int) *BadEmployee {
	return &BadEmployee{
		wage:         wage,
		regularHours: regularHours,
	}
}

// for CFO
func (e *BadEmployee) CalculatePay() int {
	return e.ReportHours() * e.wage
}

// for COO
func (e *BadEmployee) ReportHours() int {
	return e.GetRegularHours() * 2
}

// for CTO
func (e *BadEmployee) Save() {
	// ... save to DB
}

func (e *BadEmployee) GetRegularHours() int {
	return e.regularHours
}
