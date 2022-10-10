package main

import (
	"errors"
	"fmt"
)

// vvvvv 抽象層 vvvvv
type ICar interface {
	Run()
}

type ICarFactory interface {
	MakeCar(brand string) (ICar, error)
}

// ^^^^^^^^^^^^^^^^^^^^

// vvvvv 以下為實作 vvvvv

// factory 實作
type CarFactory struct{}

func NewCarFactory() ICarFactory {
	return &CarFactory{}
}

func (*CarFactory) MakeCar(brand string) (ICar, error) {
	switch brand {
	case "BMW":
		return &CarBMW{}, nil
	case "MarchBuBu":
		return &CarMarchBuBu{}, nil
	default:
		return nil, errors.New("Car Brand Not Found")
	}
}

// Car 實作
type CarBMW struct{}

func (*CarBMW) Run() {
	fmt.Println("I'm a running BMW !!!")
}

type CarMarchBuBu struct{}

func (*CarMarchBuBu) Run() {
	fmt.Println("I'm a running MarchBuBu !!!")
}

// ^^^^^^^^^^^^^^^^^^^^

func main() {
	carFactory := NewCarFactory()

	marchBuBu, _ := carFactory.MakeCar("MarchBuBu")
	marchBuBu.Run()

	bmw, _ := carFactory.MakeCar("BMW")
	bmw.Run()
}
