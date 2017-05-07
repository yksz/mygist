package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

func ReadUsername() (string, error) {
	r := bufio.NewReader(os.Stdin)
	fmt.Print("Username: ")
	line, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(line), nil
}

func ReadPassword() (string, error) {
	fmt.Printf("Password: ")
	password, err := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	if err != nil {
		return "", err
	}
	return string(password), nil
}
