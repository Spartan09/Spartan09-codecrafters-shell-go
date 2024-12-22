package external

import (
	"fmt"
	"github.com/codecrafters-io/shell-starter-go/internal/parser"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func Execute(externalCommand []string, redirect *parser.RedirectInfo) error {
	if len(externalCommand) == 0 {
		return fmt.Errorf("no command provided")
	}

	command := externalCommand[0]
	args := externalCommand[1:]

	var commandPath string
	envPath := os.Getenv("PATH")
	paths := strings.Split(envPath, string(os.PathListSeparator))
	for _, path := range paths {
		fp := filepath.Join(path, command)
		if _, err := os.Stat(fp); err == nil {
			commandPath = fp
			break
		}
	}
	if commandPath == "" {
		return fmt.Errorf("%s: command not found", command)
	}

	execCmd := exec.Command(commandPath, args...)
	execCmd.Stdin = os.Stdin

	// Handle directory creation and redirections
	if redirect != nil {
		// Create directories
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

		// Setup redirections
		if redirect.StdoutFile != "" {
			flags := os.O_CREATE | os.O_WRONLY
			if redirect.IsAppend {
				flags |= os.O_APPEND
			} else {
				flags |= os.O_TRUNC
			}
			f, err := os.OpenFile(redirect.StdoutFile, flags, 0644)
			if err != nil {
				return err
			}
			defer f.Close()
			execCmd.Stdout = f
		} else {
			execCmd.Stdout = os.Stdout
		}

		if redirect.StderrFile != "" {
			flags := os.O_CREATE | os.O_WRONLY
			if redirect.IsAppend {
				flags |= os.O_APPEND
			} else {
				flags |= os.O_TRUNC
			}
			f, err := os.OpenFile(redirect.StderrFile, flags, 0644)
			if err != nil {
				return err
			}
			defer f.Close()
			execCmd.Stderr = f
		} else {
			execCmd.Stderr = os.Stderr
		}
	} else {
		execCmd.Stdout = os.Stdout
		execCmd.Stderr = os.Stderr
	}

	err := execCmd.Run()
	if err != nil && redirect.StdoutFile == "" && redirect.StderrFile == "" {
		// Only return error if neither stdout nor stderr is redirected
		return err
	}
	return nil
}
