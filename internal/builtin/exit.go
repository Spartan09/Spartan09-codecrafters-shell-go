package builtin

import "os"

type ExitCommand struct{}

func (c *ExitCommand) Name() string { return "exit" }
func (c *ExitCommand) Execute(args []string, redirectFile string) error {
	os.Exit(0)
	return nil
}
