package main

import (
	"fmt"
	"os/user"
)

func main() {
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("error to get user:", err)
	}
	fmt.Println(currentUser.Username)
}
