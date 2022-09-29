package good_employee

type GoodEmployee struct {
	wage         int
	regularHours int
}

func NewGoodEmployee(wage, regularHours int) *GoodEmployee {
	return &GoodEmployee{
		wage:         wage,
		regularHours: regularHours,
	}
}

func (e *GoodEmployee) GetWage() int {
	return e.wage
}

func (e *GoodEmployee) GetRegularHours() int {
	return e.regularHours
}
