package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "cys000522"
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 0)

	fmt.Println(string(hash), err)
	fmt.Println(bcrypt.CompareHashAndPassword(hash, []byte("cys000522")))
	fmt.Println(bcrypt.CompareHashAndPassword(hash, []byte("cys000522")))
}