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
	var jsonString = `{"Name": "John Wick", "Age": 27}`
	var jsonData = []byte(jsonString)
	var data user
	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("user :", data.FullName)
	fmt.Println("age :", data.Age)
}
