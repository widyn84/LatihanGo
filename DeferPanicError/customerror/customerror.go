package main

import (
	"errors"
	"fmt"
	"strings"
)

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("Cannot be empty")
	}
	return true, nil
}

func main() {
	var name string
	fmt.Println("Type your name:	")
	fmt.Scanln(&name)
	if valid, err := validate(name); valid {
		fmt.Println("Hello", name)
	} else {
		fmt.Println(err.Error())
	}
}
