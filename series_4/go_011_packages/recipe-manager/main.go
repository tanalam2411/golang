package main

import (
	"fmt"
	"golang/series_4/go_011_packages/recipe-manager/recipe"
	"golang/series_4/go_011_packages/security"
)

const username = "test user"

func main() {

	if !security.IsAuthenticated(username){
		fmt.Printf("%s is not authenticated, logging\n", username)
		security.Login(username)
	}

	recipe.Add("Pizza")
	recipe.Remove("Sushi")

	security.Logout(username)

	if security.IsAuthenticated(username) {
		fmt.Println("ERROR - User still authenticated")
	} else {
		fmt.Println("Great, user successfully logged out")
	}
}
