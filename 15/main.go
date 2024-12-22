//Это может привести к утеки памяти. justSring указывает на 
//v не будут использоваться в дальнейшем, 
// GC не сможет удалить эту строку так как на неё указывает justSring

package main

import (
	"strings"
)

var justString string

func someFunc() {
  v := createHugeString(1 << 10)
  justString = strings.Clone(v[:100])
}
func main() {
  someFunc()
}