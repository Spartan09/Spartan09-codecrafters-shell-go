package builtin

import (
	"fmt"
	"github.com/codecrafters-io/shell-starter-go/internal/command"
	"github.com/codecrafters-io/shell-starter-go/internal/parser"
	"os"
	"path/filepath"
	"strings"
)

type TypeCommand struct {
	Checker command.BuiltinChecker
}

func (c *TypeCommand) Name() string { return "type" }
func (c *TypeCommand) Execute(args []string, redirect *parser.RedirectInfo) error {
	if len(args) != 1 {
		return fmt.Errorf("type: incorrect number of arguments")
	}

	var output string
	if c.Checker.IsBuiltin(args[0]) {
		output = fmt.Sprintf("%s is a shell builtin", args[0])
	} else {
		result, err := c.searchInPath(args[0])
		if err != nil {
			return err
		}
		output = result
	}

	if redirect != nil && redirect.StdoutFile != "" {
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
		fmt.Fprintln(f, output)
	} else {
		fmt.Println(output)
	}
	return nil
}

func (c *TypeCommand) searchInPath(cmd string) (string, error) {
	envPath := os.Getenv("PATH")
	paths := strings.Split(envPath, ":")
	for _, path := range paths {
		fp := filepath.Join(path, cmd)
		if _, err := os.Stat(fp); err == nil {
			return fp, nil
		}
	}
	return "", fmt.Errorf("%s: not found", cmd)
}
