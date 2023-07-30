package main

import "fmt"

func main() {
	var fruits = [4]string{"apple", "grape", "banana", "melon"}
	fmt.Println("Jumlah Elemen \t \t", len(fruits))
	fmt.Println("Isi semua elemen \t", fruits)
	fmt.Println(fruits[0])
}
