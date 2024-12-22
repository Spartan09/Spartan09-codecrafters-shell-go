package builtin

import (
	"fmt"
	"os"
)

type CdCommand struct{}

func (c CdCommand) Execute(args []string) error {
	if len(args) == 0 {
		return nil
	}
	if len(args) > 1 {
		return fmt.Errorf("cd: too many arguments")
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
