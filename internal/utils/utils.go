package utils

import (
	"fmt"
	"log"
	"syscall"

	"golang.org/x/term"
)

func PromptSecret(promptText string) string {
	fmt.Print(promptText)

	bytepw, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatal(err)
	}

	pass := string(bytepw)
	fmt.Printf("You've entered: %q\n", pass)

	return pass
}
