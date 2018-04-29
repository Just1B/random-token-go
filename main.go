package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func main() {
	fmt.Println("\n\x1b[33;1m=============================================================")
	fmt.Println("         WELCOME PLEASE CONFIGURE YOUR TOKENS  :  				   ")
	fmt.Println("=============================================================\x1b[37;1m")

	var iteration int

	print("\nNumber of tokens to generate :  ")

	fmt.Scanln(&iteration)

	var bytes int

	print("\nNumber of bytes per token :  ")

	fmt.Scanln(&bytes)

	fmt.Println("\n")
	fmt.Println("=============================================================\n\x1b[37;1m")

	for i := 1; i <= iteration; i++ {
		token := GenerateRandomString(bytes)

		fmt.Println("Token", i, " -", token, "\n")
	}
}

// GenerateRandomBytes func
func GenerateRandomBytes(n int) ([]byte, error) {

	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

//GenerateRandomString func
func GenerateRandomString(s int) string {
	b, _ := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b)
}
