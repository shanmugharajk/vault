package secret

import (
	"os"
	"strings"

	"github.com/shanmugharajk/vault/internal/utils"
)

const (
	PASSPHRASE = "__phrase"
	SALT       = "__salt"
)

func excludes(env string) bool {
	return strings.Contains(env, PASSPHRASE) || strings.Contains(env, SALT)
}

func SetSecrets(passphrase string, salt string) error {
	newEnv := utils.Env(excludes)
	newEnv = append(newEnv, PASSPHRASE+"="+passphrase)
	newEnv = append(newEnv, SALT+"="+salt)
	return utils.SetEnv(newEnv)
}

func GetSecrets() (string, string) {
	return os.Getenv(PASSPHRASE), os.Getenv(SALT)
}

func DelSecrets() error {
	newEnv := utils.Env(excludes)
	return utils.SetEnv(newEnv)
}
