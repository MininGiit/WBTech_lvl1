package main

import "fmt"

func swap(a, b *int) {
	*a += *b
	*b = *a - *b
	*a -= *b
}

func main() {
	a, b := 12, 54
	swap(&a, &b)
	fmt.Println(a, b)
}