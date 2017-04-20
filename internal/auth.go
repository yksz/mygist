package internal

import (
	"fmt"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

func ReadPassword() (string, error) {
	fmt.Printf("Password: ")
	password, err := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	if err != nil {
		return "", err
	}
	return string(password), nil
}
