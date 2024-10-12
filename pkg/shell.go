package pkg

import (
	"os"
)

func GetShell() string {
	shell := os.Getenv("SHELL")
	return shell
}
