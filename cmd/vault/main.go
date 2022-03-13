package main

import (
	"fmt"
	"log"
	"os"

	"github.com/shanmugharajk/vault/internal/database"
	"github.com/shanmugharajk/vault/pkg/cmd/root"
)

func main() {
	log.SetFlags(0)

	database.Connect()

	cmd := root.NewCmdRoot()
	if err := cmd.Execute(); err != nil {
		fmt.Printf("vault error: %s\n", err)
		os.Exit(1)
	}
}
