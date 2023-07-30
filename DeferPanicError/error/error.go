package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Printf("Type some number:	")
	fmt.Scanln(&input)
	var number int
	var err error
	number, err = strconv.Atoi(input)
	if err == nil {
		fmt.Println(number, "Is number")
	} else {
		fmt.Println(input, "Is not number")
		fmt.Println(err.Error())
	}
}
