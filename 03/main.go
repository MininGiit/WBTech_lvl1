package main

import (
	"sync"
	"fmt"
)

func WriteInChan(arr []int, ch chan <-int) {
	var wg sync.WaitGroup
	wg.Add (len(arr))
	for _, elem := range arr{
		go func ()  {
			ch <- elem * elem
			wg.Done()
		} ()
	}
	wg.Wait()
	close(ch)
}

func SquadArr(arr []int) <-chan int {
	ch := make(chan int)
	go WriteInChan(arr, ch)
	return ch
}

func Sum(ch <-chan int) int {
	sum := 0
	for val := range ch {
		sum += val
	}
	return sum
}

func main() {
	arr := []int {2, 4, 6, 8, 10}
	ch := SquadArr(arr)
	sum := Sum(ch)
	fmt.Println(sum)
}