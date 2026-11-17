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

	user1 := User{
		ID:    1,
		Name:  "User1",
		Email: "userone@gmail.com",
	}

	user2 := User{
		ID:    2,
		Name:  "User2",
		Email: "usertwo@gmail.com",
	}

	user3 := User{
		ID:    3,
		Name:  "User3",
		Email: "userthree@gmail.com",
	}

	users := []User{user1, user2, user3}
	newUsers, err := json.Marshal(users)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = os.WriteFile(`newUsers.json`, newUsers, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	content, err = os.ReadFile(`newUsers.json`)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(content))

	newUsersSlice := []User{}
	err = json.Unmarshal(content, &newUsersSlice)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, value := range newUsersSlice {
		fmt.Println(value)
	}
}
