package main

import (
	"fmt"
)

type Human struct{
	Name	string
	Age		int
}

func (h *Human) Say() {
	fmt.Println("Hello! My name is", h.Name)
}

type Action struct {
	Human
}

func main() {
	human := &Human{Name: "Slavik", Age: 22}
	human.Say()
	action := &Action{Human: Human{Name: "Action", Age: 23}}
	action.Say()
}