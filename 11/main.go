package main

import "fmt"

func intersection(arr1, arr2 []int) []int{
	set1 := make(map[int] bool, len(arr1))
	set2 := make(map[int] bool, len(arr2))
	for _, elem := range arr1 {
		set1[elem] = true
	}
	for _, elem := range arr2 {
		set2[elem] = true
	}
	var result []int 
	for key := range set1 {
		if set2[key] {
			result = append(result, key)
		}
	}
	return result
}

func main() {
	a := []int{1, 4, 5, 3, 4}
	b := []int{3, 2, 6, 8, 4}
	intersection := intersection(a, b)
	fmt.Println(intersection)
}