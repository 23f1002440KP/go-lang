package auth

import "fmt"

func  LoginWithCredentials(username string, password string) bool {  // This function is exported because it starts with a capital letter
	fmt.Println("Logging in with credentials")
	return true
}

