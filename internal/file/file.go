package file

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func IsExists(fileName string) bool {
	info, err := os.Stat(fileName)

	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func CreateIfNotExist(fileName string) bool {
	if IsExists(fileName) {
		return false
	}

	f, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	return true
}

func WriteFile(fileName string, content interface{}) {
	fmt.Println("content", content)

	fileContent, err := json.MarshalIndent(content, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(fileName, fileContent, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
