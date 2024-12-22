package main

import (
	"fmt"
	"github.com/codecrafters-io/shell-starter-go/internal/shell"
	"os"
)

func main() {
	shell := shell.NewShell()
	if err := shell.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
