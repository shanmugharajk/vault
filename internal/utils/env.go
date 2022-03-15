package utils

import (
	"os"
	"syscall"
)

const (
	SHELL = "SHELL"
)

func Env(excludes func(env string) bool) []string {
	env := []string{}

	for _, v := range os.Environ() {
		if !excludes(v) {
			env = append(env, v)
		}
	}

	return env
}

func SetEnv(newEnv []string) error {
	return syscall.Exec(os.Getenv(SHELL), []string{os.Getenv(SHELL)}, newEnv)
}
