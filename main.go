package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FullName string `json:"Name"`
	Age      int
}

func main() {
	var object = []User{{"john wick", 23}, {"ethan hunt", 32}}

	var jsonData, err = json.Marshal(object)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(jsonData) // output berbentuk byte
	var jsonString = string(jsonData)
	fmt.Println(jsonString)
}
