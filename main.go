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
	var jsonString = `[
		{
			"full_name" : "Febrianto",
			"email" : "febri@gmail.com",
			"age": 23
		},
		{
			"full_name" : "leonardo fransisco",
			"email" : "leo@gmail.com",
			"age": 24
		}
	]
		
	`

	var employees []Employee

	var err = json.Unmarshal([]byte(jsonString), &employees)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for index, employee := range employees {

		fmt.Printf("index : %v, employee object : %v \n", index, employee)
	}
}
