package main

import (
	"fmt"

	"github.com/shanmugharajk/vault/internal/crypt"
)

func main() {
	fmt.Println("Vault - a simple password manager cli tool!")

	hashKey := crypt.CreateHashKey("my password here", "salt here")
	fmt.Println("hashKey - ", hashKey)

	secretString := "secret to store"
	encryptedValue := crypt.Encrypt([]byte(secretString), hashKey)
	fmt.Println("encryptedValue - ", string(encryptedValue))

	decryptedValue := crypt.Decrypt([]byte(encryptedValue), hashKey)
	fmt.Println("decryptedValue - ", string(decryptedValue))
}
