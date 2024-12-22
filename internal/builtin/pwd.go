package builtin

import (
	"fmt"
	"github.com/codecrafters-io/shell-starter-go/internal/parser"
	"os"
)

type PwdCommand struct{}

func (c *PwdCommand) Name() string { return "pwd" }
func (c *PwdCommand) Execute(args []string, redirect *parser.RedirectInfo) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
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
		fmt.Fprintln(f, dir)
	} else {
		fmt.Println(dir)
	}
	return nil
}
