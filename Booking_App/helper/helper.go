package helper

import (
	"fmt"
	"strings"
)

func ValidateUserData(firstName string, lastName string, email string) bool {

	isValidName := len(lastName) >= 2 && len(firstName) >= 2
	isValidEmail := strings.Contains(email, "@")

	if !isValidName {
		fmt.Printf("first and last name should be alteast 2 chars or more\n")
	}

	if !isValidEmail {
		fmt.Printf("Email: %v is not valid!. Please try with valid email id \n", email)
	}

	return isValidName && isValidEmail
}
