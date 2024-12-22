package external

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func Execute(externalCommand []string) error {
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
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr
	err := execCmd.Run()
	if err != nil {
		return fmt.Errorf("failed to execute %s: %w", command, err)
	}
	return nil
}
