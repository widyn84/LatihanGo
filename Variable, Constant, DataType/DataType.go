package main

import "fmt"

func main() {
	var firstName string = "Nama Depan"
	const lastName = "Nugraha"

	firstName = "Widy"

	var bilanganBulat uint8 = 25
	var bilanganDesimal = 2.5
	var boolean = true

	fmt.Println("Halo", firstName, lastName, bilanganBulat, bilanganDesimal, boolean, "!\n")
}
