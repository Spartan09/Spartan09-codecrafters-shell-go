package builtin

import (
	"github.com/codecrafters-io/shell-starter-go/internal/parser"
	"os"
)

type ExitCommand struct{}

func (c *ExitCommand) Name() string { return "exit" }
func (c *ExitCommand) Execute(args []string, redirect *parser.RedirectInfo) error {
	os.Exit(0)
	return nil
}
