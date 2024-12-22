package main

import "fmt"

func quickSort(arr []int){
	if len(arr) < 2 {
	 	return 
	}
	left, right := 0, len(arr)-1
	pivotIndex := len(arr) / 2
	pivot := arr[pivotIndex]
	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]
	for i := 0; i < right; i++ {
		if arr[i] < pivot {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}
	arr[left], arr[right] = arr[right], arr[left] 
	quickSort(arr[:left])
	quickSort(arr[left+1:])
	return
}

func main(){
	arr := []int{23, 4, 3, 2 , 7, 12, 9, 23, 15}
	quickSort(arr)
	fmt.Println(arr)
}