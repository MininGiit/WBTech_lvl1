//Это может привести к утеки памяти. justSring указывает на 
//v не будут использоваться в дальнейшем, 
// GC не сможет удалить эту строку так как на неё указывает justSring

package main

import (
  "strings"
  "math/rand"
)

var justString string

func createHugeString(n int) string{
  letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
  result := make([]byte, n)
 
  for i := 0; i < n; i++ {
   result[i] = letters[rand.Intn(len(letters))]
  }
  return string(result)
}

func someFunc() {
  v := createHugeString(1 << 10)
  justString = strings.Clone(v[:100])
}
func main() {
  someFunc()
}