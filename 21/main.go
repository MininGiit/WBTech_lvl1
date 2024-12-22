package main

import "fmt"

type Payment interface {
	Pay()
}

type Card struct{}

func (c *Card) Pay(){
	fmt.Println("payment by card")
}

type Phone struct{}

func (p *Phone) Call() {
	fmt.Println("Call")
}

type AdapterPhone struct{
	phone *Phone
}

func NewAdapter(phone *Phone) *AdapterPhone {
	return &AdapterPhone{phone: phone}
}

func (a *AdapterPhone) Pay() {
	fmt.Println("payment by phone")
}

func Buy(p Payment) {
	p.Pay()
} 

func main() {
	card := Card{}
	Buy(&card)
	phone := Phone{}
	adapter := NewAdapter(&phone)
	Buy(adapter)
}