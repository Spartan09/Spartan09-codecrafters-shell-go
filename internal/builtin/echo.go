package builtin

import (
	"fmt"
	"github.com/codecrafters-io/shell-starter-go/internal/parser"
	"os"
	"path/filepath"
	"strings"
)

type EchoCommand struct{}

func (c *EchoCommand) Name() string { return "echo" }
func (c *EchoCommand) Execute(args []string, redirect *parser.RedirectInfo) error {
	output := strings.Join(args, " ")

	if redirect != nil {
		// Create directories if needed for redirections
		if redirect.StdoutFile != "" {
			dir := filepath.Dir(redirect.StdoutFile)
			if err := os.MkdirAll(dir, 0755); err != nil {
				return fmt.Errorf("failed to create directory: %w", err)
			}
		}
		if redirect.StderrFile != "" {
			dir := filepath.Dir(redirect.StderrFile)
			if err := os.MkdirAll(dir, 0755); err != nil {
				return fmt.Errorf("failed to create directory: %w", err)
			}
		}

		// Handle stderr redirection
		var stderr *os.File = os.Stderr
		if redirect.StderrFile != "" {
			flags := os.O_CREATE | os.O_WRONLY
			if redirect.IsAppend {
				flags |= os.O_APPEND
			} else {
				flags |= os.O_TRUNC
			}
			f, err := os.OpenFile(redirect.StderrFile, flags, 0644)
			if err != nil {
				return fmt.Errorf("failed to open error file: %w", err)
			}
			defer f.Close()
			stderr = f
		}

		// Handle stdout redirection
		if redirect.StdoutFile != "" {
			flags := os.O_CREATE | os.O_WRONLY
			if redirect.IsAppend {
				flags |= os.O_APPEND
			} else {
				flags |= os.O_TRUNC
			}
			f, err := os.OpenFile(redirect.StdoutFile, flags, 0644)
			if err != nil {
				fmt.Fprintln(stderr, err)
				return nil
			}
			defer f.Close()
			fmt.Fprintln(f, output)
			return nil
		}
	}

	// If no stdout redirection, print to stdout
	fmt.Println(output)
	return nil
}
