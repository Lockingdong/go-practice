package main

type TypeCCable struct{}
type USBCable struct{}

type ITungCharger interface {
	TypeCPort(TypeCCable)
	USBPort(USBCable)
}

type TungCharger struct {
}

func (c *TungCharger) TypeCPort(TypeCCable) {
	panic("implement me")
}

func (c *TungCharger) USBPort(USBCable) {
	panic("implement me")
}
