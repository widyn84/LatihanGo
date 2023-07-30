package main

import (
	"fmt"
	"math"
)

func calculate(d float64) (float64, float64) {
	var area = math.Pi * math.Pow(d/2, 2) // Hitung keliling
	var circumference = math.Pi * d
	return area, circumference
}

func main() {
	area, circum := calculate(3.5)
	fmt.Println(area)
	fmt.Println(circum)
}
