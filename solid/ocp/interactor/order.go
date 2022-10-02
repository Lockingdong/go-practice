package interactor

type Order struct {
	ID          string
	ProductName string
	Status      string
	Amount      int
}

func NewOrder(id, productName, status string, amount int) *Order {
	return &Order{
		ID:          id,
		ProductName: productName,
		Status:      status,
		Amount:      amount,
	}
}
