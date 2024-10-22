package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	users := []User{
		{Name: "Alice", Age: 20},
		{Name: "Bob", Age: 25},
	}

	jsonData, err := json.Marshal(users)
	if err != nil {
		fmt.Println("error marshaling JSON", err)
		return
	}
	fmt.Println("JSON data", string(jsonData))

	var userList []User
	if err := json.Unmarshal(jsonData, &userList); err != nil {
		fmt.Println("error unmarshaling JSON", err)
		return
	}

	for _, user := range userList {
		fmt.Printf("Name: %s, Age: %d", user.Name, user.Age)
	}

}
