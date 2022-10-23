package main

import (
	"errors"
	"fmt"
)

var ErrOrderIDIsEmpty error = errors.New("Order ID is Empty")

type ErrInvalidOrderID struct {
	orderID string
}

func NewErrInvalidOrderID(orderID string) ErrInvalidOrderID {
	return ErrInvalidOrderID{
		orderID: orderID,
	}
}

func (e ErrInvalidOrderID) Error() string {
	return fmt.Sprintf("Order ID %s is invalid", e.orderID)
}

func (e ErrInvalidOrderID) OrderID() string {
	return e.orderID
}

func checkOrderExist(orderID string) (bool, error) {

	if orderID == "" {
		return false, ErrOrderIDIsEmpty
	}

	if len(orderID) < 3 {
		return false, NewErrInvalidOrderID(orderID)
	}

	if orderID == "123" {
		return true, errors.New(fmt.Sprintf("order id %s already exist", orderID))
	}

	if orderID == "456" {
		return true, fmt.Errorf("order id %s already exist", orderID)
	}

	return false, nil
}

func isErrInvalidOrderID(err error) bool {
	// _, ok := err.(ErrInvalidOrderID)
	// return ok

	var errInvalidOrderID ErrInvalidOrderID
	return errors.As(err, &errInvalidOrderID)
}

func isErrOrderIDIsEmpty(err error) bool {
	return errors.Is(err, ErrOrderIDIsEmpty)
}

func main() {

	if _, err := checkOrderExist("12"); err != nil {

		if isErrInvalidOrderID(err) {
			fmt.Printf("[ErrInvalidOrderID] %s \n", err.Error())
		} else if isErrOrderIDIsEmpty(err) {
			fmt.Printf("[ErrOrderIDIsEmpty] %s \n", err.Error())
		} else {
			fmt.Printf("[ErrOrderExist] %s \n", err.Error())
		}
	}

}
