package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	user := User{
		ID:    7423424,
		Name:  "Gopher",
		Email: "gopher@gmail.com",
	}

	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = os.WriteFile(`user.json`, data, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	content, err := os.ReadFile(`user.json`)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(content))

	var recoveredUser User
	err = json.Unmarshal(content, &recoveredUser)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(recoveredUser.Name)
}
