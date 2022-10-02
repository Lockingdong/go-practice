package repo

import "github.com/go-practice/solid/ocp/interactor"

type OrderInmemRepo struct {
	list []*interactor.Order
}

func NewOrderInmemRepo() interactor.IOrderRepo {
	return &OrderInmemRepo{
		list: []*interactor.Order{
			interactor.NewOrder("111", "商品A", "Paid", 100),
			interactor.NewOrder("222", "商品B", "Paid", 200),
			interactor.NewOrder("333", "商品C", "Pending", 200),
			interactor.NewOrder("444", "商品D", "Pending", 200),
			interactor.NewOrder("555", "商品E", "Pending", 100),
		},
	}
}

func (r *OrderInmemRepo) List(q *interactor.GeneratorQuery) []*interactor.Order {
	// 省略實作 query db ...
	resultList := make([]*interactor.Order, 0)
	for _, v := range r.list {
		if isSatisfied(q.ID, v.ID) && isSatisfied(q.ProductName, v.ProductName) && isSatisfied(q.Status, v.Status) {
			resultList = append(resultList, v)
		}
	}

	return resultList
}

func isSatisfied(query, compared string) bool {
	if query == "" {
		return true
	}

	return query == compared
}
