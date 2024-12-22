package main


import "fmt"

func removeElem(s []int, i int) []int {
	if i < 0 || i >= len(s) {
  		return s
	}
 	copy(s[i:], s[i+1:])
	return s[:len(s)-1]
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	newS := removeElem(s, 2)
	fmt.Println(newS)
}