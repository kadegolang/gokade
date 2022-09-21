package tools

import (
	"golang.org/x/crypto/bcrypt"
)

// GeneratePassword 生成bcrypt hash
func Bcrypt(password string) string { //密文工具
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	if err != nil {
		panic(err.Error())
	}
	return string(hash)
}

// CheckPassword 检查密码正确性
func BcryptCheck(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

// package main

// import (
// 	"fmt"

// 	"golang.org/x/crypto/bcrypt"
// )

// func main() {
// 	password := "123abc"
// 	hash, err := bcrypt.GenerateFromPassword([]byte(password), 0)

// 	fmt.Println(string(hash), err)
// 	fmt.Println(bcrypt.CompareHashAndPassword(hash, []byte("123abcd")))
// 	fmt.Println(bcrypt.CompareHashAndPassword(hash, []byte("123abc")))
// }
