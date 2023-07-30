package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	FullName string `json:"Name"`
	Age      int
}

func main() {
	var object = []user{{"John Wick", 27}, {"Ethan Hunt", 32}}
	var jsonData, err = json.Marshal(object)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var jsonString = string(jsonData)
	fmt.Println(jsonString)
}
