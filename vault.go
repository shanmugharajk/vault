package main

import (
	"fmt"

	"github.com/shanmugharajk/vault/internal/crypt"
	"github.com/shanmugharajk/vault/internal/database"
	"github.com/shanmugharajk/vault/internal/models"
)

func main() {
	fmt.Println("Vault - a simple password manager cli tool!")

	// == Encrypting / Decrypting
	hashKey := crypt.CreateHashKey("my password here", "salt here")
	fmt.Println("hashKey - ", hashKey)

	secretString := "secret to store"
	encryptedValue := crypt.Encrypt([]byte(secretString), hashKey)
	fmt.Println("encryptedValue - ", string(encryptedValue))

	decryptedValue := crypt.Decrypt([]byte(encryptedValue), hashKey)
	fmt.Println("decryptedValue - ", string(decryptedValue))

	// == Testing gorm ==
	database.Connect(&database.Config{
		Automigrate: true,
		Recreate:    true,
	})

	// Save data
	database.Db.Create(&models.Secret{
		Key:   "secret key",
		Value: "secret value",
	})

	// Read data
	var secret models.Secret
	database.Db.First(&secret)

	fmt.Println("Secret key - ", secret.Key)
	fmt.Println("Secret value - ", secret.Value)
}
