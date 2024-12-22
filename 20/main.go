package main

import (
	"fmt"
	"strings"
)

func Reverse(s string) string {
    words := strings.Split(s, " ")
    for i := 0; i < len(words)/2; i++ {
        j := len(words) - 1 - i
        words[i], words[j] = words[j], words[i]
    }
    return strings.Join(words, " ")
}

func main() {
	str := "one two tree"
	fmt.Println(Reverse(str))
}