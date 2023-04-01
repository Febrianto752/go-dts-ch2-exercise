package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Age      int    `json:"Age"`
}

func main() {
	var jsonString = `
		{
			"full_name" : "Febrianto",
			"email" : "febri@gmail.com",
			"age": 23
		}
	`

	var tmp interface{}

	var err = json.Unmarshal([]byte(jsonString), &tmp)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var result = tmp.(map[string]interface{})

	fmt.Println("full_name :", result["full_name"])
	fmt.Println("email :", result["email"])
	fmt.Println("age :", result["age"])
}
