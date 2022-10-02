package interactor

type IOrderRepo interface {
	List(*GeneratorQuery) []*Order
}
