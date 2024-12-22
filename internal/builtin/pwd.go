package builtin

import (
	"fmt"
	"os"
)

type PwdCommand struct{}

func (c *PwdCommand) Name() string { return "pwd" }
func (c *PwdCommand) Execute(args []string, redirectFile string) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	if redirectFile != "" {
		f, err := os.OpenFile(redirectFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
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
