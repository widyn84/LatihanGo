package main

import "fmt"

func main() {
	var point = 2
	if point == 10 {
		fmt.Println("Lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("Lulus")
	} else if point == 4 {
		fmt.Println("Hampir Lulus")
	} else {
		fmt.Printf("Tidak Lulus. nilai anda %d\n", point)
	}
}
