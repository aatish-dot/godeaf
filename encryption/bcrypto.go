package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	mypass := `wow&777`
	encryptpass, err := bcrypt.GenerateFromPassword([]byte(mypass), bcrypt.MinCost)
	if err != nil {
		fmt.Println("error reported by bcrypt", err)
	}
	fmt.Println(encryptpass)

	loginpass := `wow&777`
	err = bcrypt.CompareHashAndPassword(encryptpass, []byte(loginpass))
	if err != nil {
		fmt.Println("error reported by bcrypt on comparison of pass", err)
		return
	}

	fmt.Println("Login successful ")

}
