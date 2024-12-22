package main

import "fmt"

func binarySearch(arr []int, target int, low, high int) (int, bool) {
 	if low > high {
  		return 0, false // Не найдено
	}
	mid := low + (high-low)/2
	if arr[mid] == target {
  		return mid, true
	} else if arr[mid] < target {
  		return binarySearch(arr, target, mid+1, high)
	} else {
  		return binarySearch(arr, target, low, mid-1)
	}
}

func main() {
	arr := []int{2, 3, 7, 8, 11, 12, 32}
	target := 8
	index, ok := binarySearch(arr, target, 0, len(arr)-1)
	if ok {
		fmt.Println(index)
	}
}

