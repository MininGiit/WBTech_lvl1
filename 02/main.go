package main

import(
	"fmt"
	"sync"
)

func SquadArr(arr []int) []int {
	result := make([]int, len(arr))
	var wg sync.WaitGroup
	wg.Add (len(arr))
	for i, elem := range arr {
		go func ()  {
			result[i] = elem * elem
			wg.Done()
		} ()
	}
	wg.Wait()
	return result
}

func main(){
	arr := []int {2, 4, 6, 8, 10}
	result := SquadArr(arr)
	for _, elem := range result{
		fmt.Print(elem, " ")
	}
}