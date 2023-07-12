package main

import "fmt"

func main() {
	var firstName string = "Nama Depan"
	const lastName = " Nugraha"

	firstName = "Widy"

	var numberA int = 4
	var numberB *int = &numberA
	*numberB = 8

	fmt.Println("Halo ", firstName, lastName, numberA, "!\n")
}
