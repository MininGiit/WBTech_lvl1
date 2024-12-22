package main

import (
	"fmt"
)

func setBit(num int64, i int, flag bool) int64{ //flag = true 
	mask := int64(1) << i
	if flag {
		return num | mask
	} else {
		return num & ^mask
	}
}

func main() {
	var number int64 = 16758 
	fmt.Println(setBit(number, 12, true))
	fmt.Println(setBit(number, 14, false))
}