package main

import (
	"fmt"
	"github.com/MihailChapenko/lets-go-chat/pkg/hasher"
)

func main() {
	password := "secret"
	hash, err := hasher.HashPassword(password)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	match := hasher.CheckPassword(password, hash)
	fmt.Println("Password is valid:", match)
}
