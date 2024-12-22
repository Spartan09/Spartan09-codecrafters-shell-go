package builtin

import (
	"fmt"
	"os"
)

type PwdCommand struct{}

func (p *PwdCommand) Execute(args []string) error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Println(pwd)
	return nil
}

func (p *PwdCommand) Name() string {
	return "pwd"
}
