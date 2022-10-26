package secret

import (
	"fmt"
	"log"
	"syscall"

	"golang.org/x/term"
)

func ReadSecrets() (string, string) {
	var passphrase string
	var saltkey string

	for len(passphrase) < 10 {
		passphrase = ReadPassword("\nplease enter passphrase with minimum of 10 characters\n", 0)
	}

	for len(saltkey) < 5 {
		saltkey = ReadPassword("\nplease enter saltkey with minimum of 5 characters\n", 0)
	}

	return passphrase, saltkey
}

func ReadPassword(promptText string, minLength int) string {
	fmt.Print(promptText)

	bytepw, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatal(err)
	}

	pass := string(bytepw)
	for len(pass) < minLength {
		ReadPassword(promptText, minLength)
	}

	fmt.Printf("You've entered: %q\n\n", pass)

	return pass
}
