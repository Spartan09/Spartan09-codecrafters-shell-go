package external

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// external/executor.go
func Execute(externalCommand []string, redirectFile string) error {
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

	// Handle stdout redirection
	if redirectFile != "" {
		f, err := os.OpenFile(redirectFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			return fmt.Errorf("failed to open output file: %w", err)
		}
		defer f.Close()
		execCmd.Stdout = f
		execCmd.Stderr = os.Stderr // Keep stderr going to terminal
	} else {
		execCmd.Stdout = os.Stdout
		execCmd.Stderr = os.Stderr
	}

	err := execCmd.Run()
	if err != nil && redirectFile == "" {
		// Only return error if we're not redirecting
		return fmt.Errorf("failed to execute %s: %w", command, err)
	}
	return nil
}
