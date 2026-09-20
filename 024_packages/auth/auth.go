package auth

import "fmt"

func LoginWithCredentials(username string, password string) {
	fmt.Println("User credentials:\n"+
		"username:", username, "password:", password)
}
