package main

import "fmt"

func Reverse(str string) string {
	runes := []rune(str)
	result := make([]rune, 0, len(runes))
	for i := len(runes) - 1; i >= 0; i-- {
		result = append(result, runes[i])
	}
	return string(result)
}

func main() {
	str := "🐘Привет 🐘 мир 🐘"
	fmt.Println(Reverse(str))
}