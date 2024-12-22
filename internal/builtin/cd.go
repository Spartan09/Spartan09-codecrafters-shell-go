package builtin

import (
	"fmt"
	"os"
	"strings"
)

type CdCommand struct{}

func (c CdCommand) Execute(args []string, redirectFile string) error {
	if len(args) == 0 {
		return nil
	}
	if len(args) > 1 {
		return fmt.Errorf("cd: too many arguments")
	}

	if strings.HasPrefix(args[0], "~") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		args[0] = strings.Replace(args[0], "~", homeDir, 1)
	}

	err := os.Chdir(args[0])
	if err != nil {
		return fmt.Errorf("cd: %s: No such file or directory", args[0])
	}
	return nil
}

func (c CdCommand) Name() string {
	return "cd"
}
