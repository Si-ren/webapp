package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := []byte("password")
	hash, err := bcrypt.GenerateFromPassword(password, 0)
	fmt.Println(string(hash), err)
	fmt.Println(bcrypt.CompareHashAndPassword(hash, []byte("password")))
	hash1, err := bcrypt.GenerateFromPassword(password, 0)
	fmt.Println(string(hash1))
	fmt.Println(bcrypt.CompareHashAndPassword(hash1, []byte("password")))

}
