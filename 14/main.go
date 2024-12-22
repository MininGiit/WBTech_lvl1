package main

import (
    "fmt"
    "reflect"
)

func determineType(val interface{}) {
    valType := reflect.TypeOf(val).Kind()
    switch reflect.TypeOf(val).Kind() {
    case reflect.Int:
        fmt.Println("int")
    case reflect.String:
        fmt.Println("string")
    case reflect.Bool:
        fmt.Println("bool")
    case reflect.Chan:
        fmt.Println("channel")
    default:
        fmt.Println("unknown type:", valType)
    }
}

func main() {
    num := 42
    str := "Привет!"
    flag := true
    ch := make(chan int)
	numf := 43.56
    determineType(num)
    determineType(str)
    determineType(flag)
    determineType(ch)
	determineType(numf)
}