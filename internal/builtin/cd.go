package builtin

import (
	"fmt"
	"github.com/codecrafters-io/shell-starter-go/internal/parser"
	"os"
	"strings"
)

type CdCommand struct{}

func (c *CdCommand) Execute(args []string, redirect *parser.RedirectInfo) error {
	// CD implementation stays the same, just add redirect parameter
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

func (c *CdCommand) Name() string {
	return "cd"
}
