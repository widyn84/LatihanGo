package main

import (
	"fmt"
	"strings"
)

func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}

func main() {
	var names = []string{"Widy", "Nugraha"}
	printMessage("Hallo", names)
}
