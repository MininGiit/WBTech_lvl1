package main

import (
	"fmt"
	"math"
)

func group(temps []float64) map[int][]float64 {
	groupedTemps := make(map[int][]float64)
	for _, temp := range temps {
		var groupKey int
		if temp >= 0 {
			groupKey = int(math.Floor(temp/10)) * 10 
		} else {
			groupKey = int(math.Trunc(temp/10)) * 10 
		}
		
		groupedTemps[groupKey] = append(groupedTemps[groupKey], temp)
	}
	return groupedTemps
}

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	result := group(temperatures)
	for group, temperatures := range result{
		fmt.Printf("%d: ", group)
		for _, temperature := range temperatures {
			fmt.Print(temperature, " ")
		}
		fmt.Println()
	} 
}
