package main

import (
	"github.com/23f1002440KP/GO-Basics/auth"
	"github.com/23f1002440KP/GO-Basics/user"
	"github.com/fatih/color"
)

func main() {

	if auth.LoginWithCredentials("username", "password"){
		session := auth.GetSession()
		println("Logged in & ",session)
	} 

	user1 := user.UserNew{
		Username: "username",
		Email: "email",
	}

	// println(user1.Username)
	// println(user1.Email)
	color.Red(user1.Username)
	color.Green(user1.Email)

}