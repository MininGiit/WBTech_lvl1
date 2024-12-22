package main

import "fmt"

func isUniqueSym(str string) bool{
	uniqueSym := make(map[rune] bool, len(str))
	for _, sym := range str {
		if _, ok := uniqueSym[sym]; ok {
			return false
		} else {
			uniqueSym[sym] = true
		}
	}
	return true
}

func main() {
	str1 := "abcd"
	str2 := "abCdefAaf"
	str3 := "aabcd"
	fmt.Println(isUniqueSym(str1))
	fmt.Println(isUniqueSym(str2))
	fmt.Println(isUniqueSym(str3))
}