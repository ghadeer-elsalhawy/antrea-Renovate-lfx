package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
    password := []byte("mySecret123")
    hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
    if err != nil {
        panic(err)
    }
    fmt.Println("Hashed password:", string(hash))
}