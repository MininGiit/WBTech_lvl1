package main

import "fmt"

func initSet(arr []string) []string {
	set := make(map[string] bool, len(arr))
	for _, elem := range arr {
		set[elem] = true
	}	
	result := make([]string, 0, len(set))
	for key := range set {
		result = append(result, key)
	}
	return result
}

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	set := initSet(arr)
	fmt.Println(set)
}