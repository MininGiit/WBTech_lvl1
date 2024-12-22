package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(33554432 * 3)
	b := big.NewInt(33554432) 

	sum := new(big.Int).Add(a, b)
	fmt.Println("Sum: ", sum)

	diff := new(big.Int).Sub(a, b)
	fmt.Println("Diff: ", diff)

	prod := new(big.Int).Mul(a, b)
	fmt.Println("Mul: ", prod)
	
	quo := new(big.Int).Div(a, b)
	rem := new(big.Int).Rem(a, b)
	fmt.Printf("Div: %s, Rem: %s\n", quo, rem)
}