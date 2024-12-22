package main

import(
	"fmt"
)

func WriteInChan(arr []int, ch chan <-int){
	for _, num := range arr {
		ch <- num
	}
	close(ch)
}

func Multiplication(ch1 <-chan int, ch2 chan <-int) {
	for num := range ch1 {
		ch2 <- num * 2
	}
	close(ch2)
}

func Out(ch <-chan int) {
	for num := range ch {
		fmt.Print(num, " ")
	}
}

func main() {
	arr := []int{3, 5, 34, 0, 2, 35, 19}
	chan1 := make(chan int)
	go WriteInChan(arr, chan1)
	chan2 := make(chan int)
	go Multiplication(chan1, chan2)
	Out(chan2)
}	