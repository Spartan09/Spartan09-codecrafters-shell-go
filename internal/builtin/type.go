package builtin

import (
	"fmt"
	"github.com/codecrafters-io/shell-starter-go/internal/command"
	"os"
	"path/filepath"
	"strings"
)

type TypeCommand struct {
	Checker command.BuiltinChecker
}

func (c *TypeCommand) Name() string { return "type" }
func (c *TypeCommand) Execute(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("type: incorrect number of arguments")
	}
	if c.Checker.IsBuiltin(args[0]) {
		fmt.Printf("%s is a shell builtin\n", args[0])
		return nil
	}
	return c.searchInPath(args[0])
}

func (c *TypeCommand) searchInPath(cmd string) error {
	envPath := os.Getenv("PATH")
	paths := strings.Split(envPath, ":")
	for _, path := range paths {
		fp := filepath.Join(path, cmd)
		if _, err := os.Stat(fp); err == nil {
			fmt.Println(fp)
			return nil
		}
	}
	return fmt.Errorf("%s: not found", cmd)
}
