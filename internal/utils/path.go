package utils

import (
	"log"
	"path"

	"github.com/kardianos/osext"
)

func GetExecutablePath() string {
	folderPath, err := osext.ExecutableFolder()
	
	if err != nil {
			log.Fatal(err)
	}

	return folderPath
}

func ConfigFilePath() string {
	basePath := GetExecutablePath();

	return path.Join(basePath, "config.json")
}